# Juniper
Companion library for gin-gonic to provide common functionality across multiple projects

Balically im pulling out all the duplicate code from variaus web apps into here

## Installing
The first time you install this package you will need to net GOPRIVATE envar
```bash
GOPRIVATE="github.com/acevrialltd" go get github.com/aceviralltd/juniper-go
```
## Note for future matt
I have chosen a flat package structure (test excluded) as a purely asthetic choice on implementation  
Havinge everything under the juniper. package name will be nicer

### Cli
Command line helper components

### Controller
Route base controller

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

