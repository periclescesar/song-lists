package pkg

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	lists := r.Group("lists")
	{
		lists.GET("", getLists())
		lists.POST("", createList())
		lists.GET("/:id", getSingleList())
		lists.PUT("/:id", updateList())
		lists.DELETE("/:id", deleteList())
	}

	return r
}
