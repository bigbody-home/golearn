package main

import (
	"fmt"
	_ "golearn/gin/backmemory/api/docs"
	_interface "golearn/gin/backmemory/interface"
	"golearn/gin/backmemory/model"
	"golearn/gin/backmemory/service"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 添加注释以描述 server 信息
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
var swaggerHandler gin.HandlerFunc

func init() {
	swaggerHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}

func main() {
	g := gin.Default()
	g.Use(service.TaskMiddle())
	g.Any("/:resource", Handler)
	g.Any("/:resource/:id", Handler2)
	g.GET("swagger/*any", swaggerHandler)
	g.Group("/user", Handler)
	g.Run()
}

// @Summary      Delete one Content
// @Description  Delete One content
// @Tags         memory pad
// @Accept       json
// @Produce      json
// @Success      200  {object}	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /memopad/:id [post]
func Handler(c *gin.Context) {
	fmt.Println("func1")
	mm := make(map[string]_interface.Interfaces)
	mm["memopad"] = &model.Content{}
	t := c.Param("resource")
	method := c.Request.Method
	switch method {
	case "GET":
		mm[t].List(c)
	case "POST":
		mm[t].Create(c)
	}

}

// @Summary      Delete or update one Content
// @Description  Delete One content
// @Tags         memory pad
// @Accept       json
// @Produce      json
// @Param		 id
// @Success      200  {object}	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /memopad/:id [delete]

func Handler2(c *gin.Context) {
	fmt.Println("func1")
	mm := make(map[string]_interface.Interfaces)
	mm["memopad"] = &model.Content{}
	t := c.Param("resource")
	method := c.Request.Method
	switch method {
	case "PUT":
		mm[t].Update(c)
	case "DELETE":
		mm[t].Delete(c)
	}

}
