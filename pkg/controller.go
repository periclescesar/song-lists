package pkg

import (
	"GoSOLID/pkg/domain"
	"GoSOLID/pkg/provider/dic"
	"GoSOLID/pkg/services"
	"GoSOLID/pkg/services/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
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

		ctn, _ := dic.NewContainer(di.Request)

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(domain.BADREQUEST, gin.H{
				"error": "json parse error",
			})
			return
		}

		var repo repositories.ListRepository
		if c.Query("draft") == "1" {
			repo = ctn.GetJsonListRepository()
		} else {
			repo = ctn.GetListRepository()
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
