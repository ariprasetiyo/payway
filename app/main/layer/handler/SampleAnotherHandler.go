package handler

import (
	"net/http"
	"payway/app/main/layer/interfaces"
	"payway/app/main/layer/model"

	"github.com/gin-gonic/gin"
)

func NewSampleAnotherHandler() interfaces.Handler {
	return &sampleAnotherHandler{}
}

type sampleAnotherHandler struct {
}

func (handler sampleAnotherHandler) Execute(c *gin.Context) {
	response := model.SampleJsonStruct{Name: "name", Number: 1}
	model := model.Response{http.StatusOK, response}
	c.JSON(model.HttpStatus, model.Response)
}
