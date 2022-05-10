package juniper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseMutateFunc takes the ckrrent response data and mutates it somehow
type ResponseMutateFunc func(ctx *gin.Context, data *gin.H)

// Controller shared logic for interacting with the gin framework
//
// This does not make an assumptions about the environment or dependencies
// required by the final controllers that embed this struct
type Controller struct {
	InjectGlobalData ResponseMutateFunc
}

// JSON is just an alios of the gin context method that injects the hud data
func (con Controller) JSON(ctx *gin.Context, code int, data gin.H) {
	if con.InjectGlobalData != nil {
		con.InjectGlobalData(ctx, &data)
	}

	ctx.JSON(code, data)
}

// AbortWithStatusJSON is just an alios of the gin context method that injects the hud data
func (con Controller) AbortWithStatusJSON(ctx *gin.Context, code int, data gin.H) {
	if con.InjectGlobalData != nil {
		con.InjectGlobalData(ctx, &data)
	}

	ctx.AbortWithStatusJSON(code, data)
}

// bindForm attempt to bind form input to struct
//
// This will actually bind from any source not just form data
//
// if it fails then abort will be auto called
func (con Controller) BindForm(ctx *gin.Context, input interface{}) (err error) {
	if err = ctx.ShouldBind(input); err != nil {
		con.AbortWithStatusJSON(ctx, http.StatusBadRequest, gin.H{
			"outcome": false,
			"message": "Bad Input",
			"fields":  errorFIelds(err),
		})
	}

	return
}

// BindUri attempt to bind path input to struct
//
// if it fails then abort will be auto called
func (con Controller) BindUri(ctx *gin.Context, input interface{}) (err error) {
	if err = ctx.ShouldBindUri(input); err != nil {
		con.AbortWithStatusJSON(ctx, http.StatusNotFound, ResponseNotFound)
	}

	return
}

// bindForm attempt to bind query string input to struct
//
// if it fails then abort will be auto called
func (con Controller) BindQuery(ctx *gin.Context, input interface{}) (err error) {
	if err = ctx.ShouldBindQuery(input); err != nil {
		con.AbortWithStatusJSON(ctx, http.StatusBadRequest, gin.H{
			"outcome": false,
			"message": "Bad Input",
			"fields":  errorFIelds(err),
		})
	}

	return
}
