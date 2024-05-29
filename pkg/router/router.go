package router

import (
	"directories/pkg/selects"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Router struct {
}

type Table struct {
	Name string `uri:"table" binding:"required"`
}

type EntityId struct {
	Id int `uri:"id"`
}

type Limit struct {
	Limit int `form:"limit"`
}

type Offset struct {
	Offset int `form:"offset"`
}

func NewRouter() *Router {
	return &Router{}
}

func (h *Router) GetRouter(dbConn *sqlx.DB) *gin.Engine {

	router := gin.New()
	api := router.Group("/api")
	{
		lists := api.Group("/data")
		{
			lists.GET("/:table", func(c *gin.Context) {

				var table Table
				if err := c.ShouldBindUri(&table); err != nil {
					c.JSON(400, gin.H{"msg": err})
					return
				}

				var limit Limit
				if err := c.Bind(&limit); err != nil {
					c.JSON(400, gin.H{"msg": err})
					return
				}

				var offset Offset
				if err := c.Bind(&offset); err != nil {
					c.JSON(400, gin.H{"msg": err})
					return
				}

				data, err := selects.GetAll(dbConn, table.Name, limit.Limit)
				if err != nil {
					c.JSON(400, gin.H{"msg": err})
					return
				}

				c.JSON(200, data)

			})

			lists.GET("/:table/:id", func(c *gin.Context) {

				var table Table
				if err := c.ShouldBindUri(&table); err != nil {
					c.JSON(400, gin.H{"msg": err})
					return
				}

				var id EntityId
				if err := c.BindUri(&id); err != nil {
					c.JSON(400, gin.H{"msg": err})
					return
				}

				data, err := selects.GetById(dbConn, table.Name, id.Id)
				if err != nil {
					c.JSON(400, gin.H{"msg": err})
					return
				}

				c.JSON(200, data)

			})

		}

	}

	return router
}
