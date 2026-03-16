package sample

import (
	"net/http"

	"github.com/gkatanacio/go-lambdalith-template/internal/handlerutil"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Hello(w http.ResponseWriter, r *http.Request) {
	msg := c.service.Hello(r.Context())

	handlerutil.DataResponse(w, http.StatusOK, msg)
}

func (c *Controller) Echo(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Message string `json:"message"`
	}
	if err := handlerutil.JsonRequestBody(r, &payload); err != nil {
		handlerutil.ErrorResponse(w, handlerutil.BadRequest("invalid request body"))
		return
	}

	handlerutil.DataResponse(w, http.StatusOK, payload)
}
