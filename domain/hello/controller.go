package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) HandleRoot(c *gin.Context) {
	message := ctrl.service.GetMessage()
	c.JSON(http.StatusOK, gin.H{"message": message.Message})
}

func (ctrl *Controller) HandleGreet(c *gin.Context) {
	id := c.Param("id")
	greet := ctrl.service.GetGreet(id)
	c.JSON(http.StatusOK, greet)
}

func (ctrl *Controller) HandleAddGreet(c *gin.Context) {
	var greet Greet
	if err := c.ShouldBind(&greet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctrl.service.AddGreet(greet)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Greet added successfully",
	})
}

func (ctrl *Controller) HandleUpdateGreet(c *gin.Context) {
	id := c.Param("id")
	var greet Greet
	c.BindJSON(&greet)
	ctrl.service.UpdateGreet(id, greet)
	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl *Controller) HandleDeleteGreet(c *gin.Context) {
	id := c.Param("id")
	ctrl.service.DeleteGreet(id)
	c.JSON(http.StatusOK, gin.H{})
}
