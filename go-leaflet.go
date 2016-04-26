package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func main() {
    router := gin.Default()

    router.Static("/css", "./css")
    router.Static("/js", "./js")
    router.Static("/images", "./images")
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "go-leaflet",
        })
    })
    router.Run(":8080")
}