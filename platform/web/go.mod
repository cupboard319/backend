module github.com/micro/services/platform/web

go 1.13

require (
	github.com/micro/go-micro/v2 v2.9.1-0.20200618113919-8c7c27c573f5
	github.com/micro/micro/v2 v2.4.0
	github.com/micro/services/platform/service v0.0.0-20200313185528-4a795857eb73
)

replace github.com/micro/services/platform/service => ../service
