package pkg

import (
	"GoSOLID/pkg/domain"
	"GoSOLID/pkg/services"
	"GoSOLID/pkg/services/repositories"
	"github.com/gin-gonic/gin"
)

func deleteList() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func updateList() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "arroz",
		})
	}
}

/**
{
        "name": "lista 1",
        "type": "custom",
        "songs": [
          1,
          2,
          3
        ]
      }
*/

func createList() func(c *gin.Context) {
	return func(c *gin.Context) {
		var body domain.List

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(domain.BADREQUEST, gin.H{
				"error": "json parse error",
			})
			return
		}

		var repo repositories.ListRepository
		if c.GetHeader("Request-test") == "1" {
			//jd := drivers.NewJsonDriver("../data/list.json")
			//repo = repositories.NewJsonListRepository(*jd)
		} else {
			repo = repositories.NewSqliteListRepository()
		}

		listService := services.NewListService(repo)

		statusCode, err := listService.CreateList(&body)
		if err != nil {
			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, body)
	}
}

func getSingleList() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "batata",
		})
	}
}

func getLists() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "macarrão",
		})
	}
}
