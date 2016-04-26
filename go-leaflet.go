package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "html/template"
    "github.com/peterbourgon/diskv"
    "fmt"
)

type  Marker struct {
    x int
    y int
    title string
    desc string
}


func markers() template.JS {
    return "var marker = L.marker([1000, 1000]).addTo(map);"
}
func initDB() *diskv.Diskv {
    // Simplest transform function: put all the data files into the base dir.
    flatTransform := func(s string) []string { return []string{} }

    // Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
    d := diskv.New(diskv.Options{
        BasePath:     "data",
        Transform:    flatTransform,
        CacheSizeMax: 1024 * 1024,
    })
    return d;
}


func main() {
    router := gin.Default()

    db := initDB()
    key := "marker1"
    value := Marker{x: 1000, y: 1000, title: "yes man", desc: "trop swag ce marker"}
    db.Write(key, []byte(fmt.Sprintf("%+v", value)))
    valueread, _ := db.Read(key)

    router.Static("/css", "./css")
    router.Static("/js", "./js")
    router.Static("/images", "./images")
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "go-leaflet",
            "markers": markers,
            "marker1": valueread,
        })
    })
    router.Run(":8080")
}