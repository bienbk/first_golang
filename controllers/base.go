package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type CommonErrorResponse struct {
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type baseController struct {
	web.Controller
}

func (c *baseController) OK(code int, data interface{}) {
	c.Ctx.Output.SetStatus(code)
	c.SetData(data)
	c.ServeJSON()
}

func (c *baseController) JSON(code int, data interface{}) {
	c.Ctx.Output.SetStatus(code)
	c.SetData(data)
	c.ServeJSON()
}

func (c *baseController) BadRequest(errName string, message string, details interface{}) {
	c.Ctx.Output.SetStatus(http.StatusBadRequest)
	c.SetData(CommonErrorResponse{
		Error:   errName,
		Message: message,
		Details: details,
	})
	c.ServeJSON()
}

func (c *baseController) InternalErr(errName string, message string, details interface{}) {
	c.Ctx.Output.SetStatus(http.StatusInternalServerError)
	c.SetData(CommonErrorResponse{
		Error:   errName,
		Message: message,
		Details: details,
	})
	c.ServeJSON()
}

func (c *baseController) UnprocessableEntity(errName string, message string, details interface{}) {
	c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
	c.SetData(CommonErrorResponse{
		Error:   errName,
		Message: message,
		Details: details,
	})
	c.ServeJSON()
}

func (c *baseController) Unauthorized(errName string, message string, details interface{}) {
	c.Ctx.Output.SetStatus(http.StatusUnauthorized)
	c.SetData(CommonErrorResponse{
		Error:   errName,
		Message: message,
		Details: details,
	})
	c.ServeJSON()
}
