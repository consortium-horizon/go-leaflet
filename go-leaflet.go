package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "html/template"
)

func markers() template.JS {
    return "var marker = L.marker([1000, 1000]).addTo(map);"
}


func main() {
    router := gin.Default()

    router.Static("/css", "./css")
    router.Static("/js", "./js")
    router.Static("/images", "./images")
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "go-leaflet",
            "markers": markers,
        })
    })
    router.Run(":8080")
}