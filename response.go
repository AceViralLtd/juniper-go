package juniper

import "github.com/gin-gonic/gin"

var (
	ResponseSuccess      = gin.H{"outcome": true, "message": nil}
	ResponseNotFound     = gin.H{"outcome": false, "message": "Not Found"}
	ResponseFailed       = gin.H{"outcome": false, "message": "Failed"}
	ResponseBadInput     = gin.H{"outcome": false, "message": "Bad Input"}
	ResponseUnauthorized = gin.H{"outcome": false, "message": "Unauthorized"}
)
