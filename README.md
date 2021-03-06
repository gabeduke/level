![](https://github.com/gabeduke/level/workflows/Test/badge.svg)
![](https://github.com/gabeduke/level/workflows/Lint/badge.svg)
![](https://github.com/gabeduke/level/workflows/Fmt/badge.svg)
[![codecov](https://codecov.io/gh/gabeduke/level/branch/master/graph/badge.svg)](https://codecov.io/gh/gabeduke/level)

# Level

Level is an API to query readings from the National Water Service. Readings can be found for most major water ways. 

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/gabeduke/level)

## API Documentation

More detailed docs can be found [HERE](https://gabeduke.github.io/level/)

<!-- markdown-swagger -->
 Endpoint    | Method | Auth? | Description            
 ----------- | ------ | ----- | -----------------------
 `/healthz`  | GET    | No    | get health             
 `/level`    | GET    | No    | get level by station   
 `/slack`    | POST   | No    | return a slack response
 `/stations` | GET    | No    | get stations           
<!-- /markdown-swagger -->

## Run

Try out the API for free on Google Cloud Run:

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/gabeduke/level.git)

### Docker

`make run` will serve the project in a local container

### Develop

`make dev` will run the project in Go dev mode

### Bootstrap (terraform)

```bash
git submodule update
make init
make import
make apply
```
