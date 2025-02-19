package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/models"
	"github.com/kaolnwza/proj-blueprint/libs/response"
)

func (h handler) CreateUserHandler(c *gin.Context) {
	var req models.ReqUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.MakeError(c, response.NewBadRequestError("", "", nil))
		return
	}

	if err := h.userSvc.CreateUser(c.Request.Context(), req); err != nil {
		response.MakeError(c, err)
		return
	}

	response.MakeSuccess(c, nil)
}
