package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/go-github/v39/github"
	"github.com/google/uuid"
	log "github.com/micro/micro/v3/service/logger"
	"golang.org/x/oauth2"

	"github.com/micro/micro/v3/service/runtime/source/git"
)

var repos = map[string]string{
	"github.com/m3o/m3o-js": "%v-api-js",
}

func Push(pat string) {
	for {
		for originRepo, newRepoTemplate := range repos {
			log.Infof("Processing repo %v", originRepo)

			gitter := git.NewGitter(map[string]string{})
			err := gitter.Checkout(originRepo, "main")
			if err != nil {
				log.Errorf("   Failed to check out repo %v: %v", originRepo, err)
				continue
			}

			files, err := ioutil.ReadDir(gitter.RepoDir())
			if err != nil {
				log.Error("   Failed to read repo dir %v: %v", gitter.RepoDir(), err)
			}

			for _, file := range files {
				if strings.HasPrefix(file.Name(), ".") {
					continue
				}
				if !file.IsDir() {
					continue
				}

				repoName := fmt.Sprintf(newRepoTemplate, file.Name())
				log.Infof("   Processing folder %v", file.Name())

				tmpDir := filepath.Join(os.TempDir(), uuid.Must(uuid.NewUUID()).String())

				os.MkdirAll(tmpDir, 0777)

				log.Infof("   Setting up git repo in tempdir %v", tmpDir)
				path := filepath.Join(gitter.RepoDir(), file.Name())
				log.Infof("   Copying from %v", path)

				outp, err := exec.Command("cp", "-r", path, tmpDir).CombinedOutput()
				if err != nil {
					log.Errorf("Error copying: %v, output: %v", err, string(outp))
					continue
				}

				targetDir := filepath.Join(tmpDir, file.Name())

				time.Sleep(3 * time.Second)
				log.Infof("   Preparing to push %v", targetDir)

				ctx := context.Background()
				ts := oauth2.StaticTokenSource(
					&oauth2.Token{AccessToken: pat},
				)
				tc := oauth2.NewClient(ctx, ts)
				client := github.NewClient(tc)
				// list all organizations for user "willnorris"
				_, _, err = client.Repositories.Create(ctx, "m3oapis", &github.Repository{
					Name: &repoName,
				})
				if err != nil {
					log.Errorf("   Failed to create repo %v: %v", repoName, err)
				}

				// git remote add origin https://[USERNAME]:[NEW TOKEN]@github.com/[USERNAME]/[REPO].git
				// see https://stackoverflow.com/questions/18935539/authenticate-with-github-using-a-token

				cmd := exec.Command("git", "init")
				cmd.Dir = targetDir
				outp, err = cmd.CombinedOutput()
				if err != nil {
					log.Errorf("   Failed to set origin %v: %v", err, string(outp))
					continue
				}

				cmd = exec.Command("git", "checkout", "-b", "main")
				cmd.Dir = targetDir
				outp, err = cmd.CombinedOutput()
				if err != nil {
					log.Errorf("   Failed to set origin %v: %v", err, string(outp))
					continue
				}

				cmd = exec.Command("git", "remote", "add", "origin", fmt.Sprintf("https://m3o-actions:%v@github.com/m3oapis/%v.git", pat, repoName))
				cmd.Dir = targetDir
				outp, err = cmd.CombinedOutput()
				if err != nil {
					log.Errorf("   Failed to set origin %v: %v", err, string(outp))
					continue
				}

				cmd = exec.Command("git", "add", ".")
				cmd.Dir = targetDir
				outp, err = cmd.CombinedOutput()
				if err != nil {
					log.Errorf("   Failed to add files %v: %v", err, string(outp))
					continue
				}

				cmd = exec.Command("git", "commit", "-am", "Update")
				cmd.Dir = targetDir
				outp, err = cmd.CombinedOutput()
				if err != nil {
					log.Errorf("   Failed to commit files %v: %v", err, string(outp))
					continue
				}

				cmd = exec.Command("git", "push", "origin", "-f", "main")
				cmd.Dir = targetDir
				outp, err = cmd.CombinedOutput()
				if err != nil {
					log.Errorf("   Failed to push %v: %v", err, string(outp))
					continue
				}
			}
		}
		time.Sleep(24 * time.Hour)
	}
}

func main() {
	Push(os.Getenv("GITHUB_CLIENT_PAT"))
}
