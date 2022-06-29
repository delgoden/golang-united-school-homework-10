package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET("/name/:PARAM", h.Hello)
	router.GET("/bad", h.Bad)
	router.POST("/data", h.Massage)
	router.POST("/headers", h.Sum)
	router.NoRoute(h.NoRoute)
}

func (h *handler) Hello(c *gin.Context) {
	name := c.Param("PARAM")
	c.String(http.StatusOK, "Hello, %s!", name)
}

func (h *handler) Bad(c *gin.Context) {
	c.Status(http.StatusInternalServerError)
}

func (h *handler) Massage(c *gin.Context) {
	param, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.String(http.StatusOK, "I got message:\n%s", param)
}

func (h *handler) Sum(c *gin.Context) {
	a := c.GetHeader("a")
	b := c.GetHeader("b")
	numA, err := strconv.Atoi(a)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	numB, err := strconv.Atoi(b)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.Header("a+b", strconv.Itoa(numA+numB))
}

func (h *handler) NoRoute(c *gin.Context) {
	c.Status(http.StatusOK)
}
