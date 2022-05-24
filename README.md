# Juniper
Companion library for gin-gonic to provide common functionality across multiple projects

Balically im pulling out all the duplicate code from variaus web apps into here

## Installing
In order for this repo to be pulled you need to have a GOPRIVATE env setup on the pulling machine
### Configure ~/.netrc
```
machine github.com
    login [github_username]
    password [github_personal_access_token]

machine api.github.com
    login [github_username]
    password [github_personal_access_token]
```

### Pull with direct GOPRIVATE env
GOPRIVATE will be set for this operation only
```bash
GOPRIVATE="github.com/acevrialltd/juniper-go" go get github.com/aceviralltd/juniper-go
```

### Setup go env for private repos
Set GOPRIVATE for all go get/install requests
`~/.config/go/env`
```
GOPRIVATE=github.com/aceviralltd
```


## Note for future matt
I have chosen a flat package structure (test excluded) as a purely asthetic choice on implementation  
Havinge everything under the juniper. package name will be nicer

### Cli
Command line helper components

### Controller
Route base controller

### Cron
Run cron tasks on a schedule, based on the cli register

### Env
Dotenv wrapper

### Gorm 
custom loggers (production/debug)

### JWT
helpers for managing jwt's and bearer tokens

### Log
Custom gin log formatters

### Response
Common response shapes for generic http errors

### Test
http test helpers

### Validation
Custom struct tag extraction for validator

