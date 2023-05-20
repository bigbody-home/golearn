package _interface

import "github.com/gin-gonic/gin"

// Interfaces This is for resource CRUD interface, every resource can realize it
type Interfaces interface {
	List(c *gin.Context)
	Update(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}
