# News Hub (microservices version) - Users API 
A simple users API for a hub of news developed for academic purposes.
<hr /> 

## Requirements
- Golang 1.16
<hr /> 

## Setup _(Suggested)_
- Install `golang` from https://go.dev/doc/install.
- Clone this repository and move into the folder.

<hr /> 

## Run Application
Start running the app locally with the following command ``make run``.
If the application setup is ok you should see the following message and the application listening on `8081` port.
```console
news-hub-microservices_users-api $make run
sh scripts/run.sh
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> news-hub-microservices_users-api/internal/controllers.HealthChecksController.Ping-fm (4 handlers)
[GIN-debug] GET    /v1/                      --> news-hub-microservices_users-api/internal/controllers.UsersController.Get-fm (4 handlers)
[GIN-debug] POST   /v1/                      --> news-hub-microservices_users-api/internal/controllers.UsersController.Create-fm (4 handlers)
[GIN-debug] POST   /v1/login                 --> news-hub-microservices_users-api/internal/controllers.UsersController.Authenticate-fm (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8081

```
<hr /> 

## Run Tests

Run all tests, with the following command:
```console
make tests
```

Run all tests with code coverage + report, with the following command:
```console
make tests-coverage
```