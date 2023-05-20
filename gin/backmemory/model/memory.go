// Package model defied all kinds of resource model
// every resource model realize _interface.Interfaces
package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golearn/gin/backmemory/db"
	"gorm.io/gorm"
	"log"
)

// Content describe memory pad's content
type Content struct {
	*gorm.Model
	What     string `json:"what" binding:"required"`
	CId      int    `json:"c_id" binding:"required"`
	CreateBy string `json:"create_by" binding:"required"`
}

// @Summary      Show All Contents
// @Description  get string by ID
// @Tags         memory pad
// @Accept       json
// @Produce      json
// @Param
// @Success      200  {object}  []model.Content
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /memopad [get]
func (c2 *Content) List(c *gin.Context) {
	fmt.Println("iii")
	var res []Content
	dbr := db.Db

	tx := dbr.Find(&res)
	log.Printf("总共有%d条备忘录", tx.RowsAffected)
	c.JSON(200, res)
}

// @Summary      Show All Contents
// @Description  get string by ID
// @Tags         memory pad
// @Accept       json
// @Produce      json
// @Param
// @Success      200  {object}  []model.Content
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /memopad/{id} [get]
func (c2 *Content) Update(c *gin.Context) {
	//
	id := c.Param("id")
	var ne Content
	var old Content
	err := c.ShouldBindWith(&ne, binding.JSON)
	fmt.Println(ne)
	if err != nil {
		panic(err)
	}
	dbr := db.Db
	err = dbr.AutoMigrate(&ne)
	if err != nil {
		log.Fatal(err)
	}
	dbr.First(&old, id)
	old.CreateBy = ne.CreateBy
	old.CId = ne.CId
	old.What = ne.What
	dbr.Save(&old)
	log.Println("Update okay!")
	c.JSON(200, &ne)
}

// Create @Summary      Create a Contents
// @Description  post a record for memory pad
// @Tags         memory pad
// @Accept       json
// @Produce      json
//
//	@Param		 {
//	   "what":"have lunch with tim",
//	   "c_id": 222,
//	   "create_by":"luke"
//	}
//
// @Success      200  {object}  model.Content
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /memopad [post]
func (c2 *Content) Create(c *gin.Context) {
	var mc Content
	err := c.ShouldBindWith(&mc, binding.JSON)
	fmt.Println(mc)
	if err != nil {
		panic(err)
	}
	dbr := db.Db
	err = dbr.AutoMigrate(&mc)
	if err != nil {
		log.Fatal(err)
	}
	dbr.Create(&mc)
	c.JSON(200, &mc)
}

// Delete @Summary      Delete one Content
// @Description  Delete One content
// @Tags         memory pad
// @Accept       json
// @Produce      json
// @Param		 id
// @Success      200  {object}  []model.Content
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /memopad/{id} [delete]
func (c2 *Content) Delete(c *gin.Context) {
	dbr := db.Db
	id := c.Param("id")
	fmt.Println("id is", id)
	dbr.Delete(&Content{}, id)
	c.String(200, "Success")
}
