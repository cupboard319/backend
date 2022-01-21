package handler

import (
	"context"
	"encoding/json"
	"fmt"

	auth2 "github.com/m3o/services/pkg/auth"
	projects "github.com/m3o/services/projects/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
)

const (
	prefixProjectsByUser = "projectsByUser"
	prefixProjectsByID   = "projectsByID"

	roleAdmin = "admin"
)

// pass a blank projectID for listing projects by userID
func projectsByUserKey(userID, projectID string) string {
	return fmt.Sprintf("%s/%s/%s", prefixProjectsByUser, userID, projectID)
}

func projectsByID(projectID string) string {
	return fmt.Sprintf("%s/%s", prefixProjectsByID, projectID)
}

type Project struct {
	ID      string
	Name    string
	Members []Member
}

type Member struct {
	ID    string
	Roles []string
}

type Projects struct{}

func New(srv *service.Service) *Projects {
	return &Projects{}
}

func (p *Projects) Create(ctx context.Context, request *projects.CreateRequest, response *projects.CreateResponse) error {
	panic("implement me")
}

func (p *Projects) Read(ctx context.Context, request *projects.ReadRequest, response *projects.ReadResponse) error {
	method := "projects.Read"
	acc, err := auth2.VerifyMicroCustomer(ctx, method)
	if err != nil {
		return err
	}
	if len(request.Id) == 0 {
		return errors.BadRequest(method, "Missing ID param")
	}
	recs, err := store.Read(projectsByID(request.Id))
	if err != nil && err != store.ErrNotFound {
		log.Errorf("Error reading project %s", err)
		return errors.InternalServerError(method, "Error reading project")
	}
	if len(recs) == 0 {
		return errors.NotFound(method, "Project not found")
	}
	var project Project
	if err := json.Unmarshal(recs[0].Value, &project); err != nil {
		log.Errorf("Error reading project %s", err)
		return errors.InternalServerError(method, "Error reading project")
	}
	found := false
	for _, v := range project.Members {
		if v.ID == acc.ID {
			found = true
			break
		}
	}
	if !found {
		return errors.Unauthorized(method, "Unauthorized")
	}
	response.Project = projectToProto(&project)
	return nil
}

func createDefaultProject(userID string) (*Project, error) {
	project := Project{
		ID:   userID,
		Name: "default",
		Members: []Member{
			{
				ID:    userID,
				Roles: []string{roleAdmin},
			},
		},
	}

	err := writeProject(project, userID)
	return &project, err
}

func writeProject(project Project, userID string) error {
	if err := store.Write(store.NewRecord(projectsByUserKey(userID, project.ID), project)); err != nil {
		return err
	}
	if err := store.Write(store.NewRecord(projectsByUserKey(userID, project.ID), project)); err != nil {
		return err
	}
	return nil
}

// List the projects this user has access to
func (p *Projects) List(ctx context.Context, request *projects.ListRequest, response *projects.ListResponse) error {
	method := "projects.List"
	acc, err := auth2.VerifyMicroCustomer(ctx, method)
	if err != nil {
		return err
	}
	recs, err := store.Read(projectsByUserKey(acc.ID, ""), store.ReadPrefix())
	if err != nil && err != store.ErrNotFound {
		log.Errorf("Error listing projects %s", err)
		return errors.InternalServerError(method, "Error looking up projects")
	}
	if len(recs) == 0 {
		// lazy create
		proj, err := createDefaultProject(acc.ID)
		if err != nil {
			log.Errorf("Error creating default project %s", err)
			return errors.InternalServerError(method, "Error looking up projects")
		}
		response.Projects = []*projects.Project{projectToProto(proj)}
		return nil
	}
	response.Projects = make([]*projects.Project, len(recs))
	for i, v := range recs {
		var proj Project
		if err := json.Unmarshal(v.Value, &proj); err != nil {
			log.Errorf("Error marshalling project %s", err)
			return errors.InternalServerError(method, "Error looking up projects")
		}
		response.Projects[i] = projectToProto(&proj)
	}
	return nil
}

func projectToProto(proj *Project) *projects.Project {
	members := []*projects.Member{}
	for _, v := range proj.Members {
		members = append(members, &projects.Member{
			Id:    v.ID,
			Roles: v.Roles,
		})
	}
	return &projects.Project{
		Id:      proj.ID,
		Name:    proj.Name,
		Members: members,
	}
}
