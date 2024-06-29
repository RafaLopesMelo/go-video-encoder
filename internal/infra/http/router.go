package http

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    v1 := r.Group("/api/v1")
    v1.GET("/", func (c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World!",
        })
    })

    return r
}
