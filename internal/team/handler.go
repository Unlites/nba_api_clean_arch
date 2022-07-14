package team

import "github.com/gin-gonic/gin"

type Handler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
