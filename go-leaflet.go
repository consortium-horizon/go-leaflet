package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "html/template"
    "github.com/peterbourgon/diskv"
    "fmt"
    "encoding/gob"
    "bytes"
    "strconv"
)

type  Marker struct {
    X int
    Y int
    Title string
    Desc string
}

func (d *Marker) GobEncode() ([]byte, error) {
    w := new(bytes.Buffer)
    encoder := gob.NewEncoder(w)
    err := encoder.Encode(d.X)
    if err!=nil {
        return nil, err
    }
    err = encoder.Encode(d.Y)
    if err!=nil {
        return nil, err
    }
    err = encoder.Encode(d.Title)
    if err!=nil {
        return nil, err
    }
    err = encoder.Encode(d.Desc)
    if err!=nil {
        return nil, err
    }
    return w.Bytes(), nil
}

func (d *Marker) GobDecode(buf []byte) error {
    r := bytes.NewBuffer(buf)
    decoder := gob.NewDecoder(r)
    err := decoder.Decode(&d.X)
    if err!=nil {
        return err
    }
    err = decoder.Decode(&d.Y)
    if err!=nil {
        return err
    }
    err = decoder.Decode(&d.Title)
    if err!=nil {
        return err
    }
    return decoder.Decode(&d.Desc)
}

func (m* Marker) WriteToDB() error {
    db := initDB()
    keybase := "marker"
    index := 0
    key := keybase + strconv.Itoa(index)
    for db.Has(key) {
        index++
        key = keybase + strconv.Itoa(index)
    }
    valueasBytes, err := m.GobEncode()
    if err!=nil {
        fmt.Println("%v",err)
    }

    db.Write(key, valueasBytes )
    return err
}

func markers() []Marker {
    markers := make([]Marker, 0)
    db := initDB()
    keybase := "marker"
    index := 0
    key := keybase + strconv.Itoa(index)
    var err error
    for db.Has(key) {
        var valueread []byte
        valueread, err = db.Read(key)
        if err==nil{
            var m Marker
            err = m.GobDecode(valueread)
            markers = append(markers, m)
        }
        //next iter
        index++
        key = keybase + strconv.Itoa(index)
    }
    return markers
}

func markersJS() template.JS {
    markers := markers()
    output := ""
    for _, m := range markers{
        output += fmt.Sprintf("var marker = L.marker([%d, %d]).addTo(map);\n",m.X,m.Y)
    }
    return template.JS(output)
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



    router.Static("/css", "./css")
    router.Static("/js", "./js")
    router.Static("/images", "./images")
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "go-leaflet",
            "markers": markersJS,
        })
    })
    router.GET("/markers", func(c *gin.Context) {
        c.HTML(http.StatusOK, "markers.tmpl", gin.H{
            "title": "Markers",
            "markers": markers,
        })
    })

    router.POST("/markers/add", func(c *gin.Context) {
        title := c.PostForm("title")
        desc := c.PostForm("desc")
        x, err := strconv.Atoi(c.PostForm("x"))
        y, err := strconv.Atoi(c.PostForm("y"))
        if err!=nil {
            fmt.Printf("%v",err)
            return
        }
        //fmt.Printf("title: %v; desc: %v; x: %v; y: %v", title, desc, x, y)
        m := Marker{Title:title, Desc: desc, X: x, Y: y }
        m.WriteToDB()
        c.Redirect(http.StatusMovedPermanently, "/markers")
    })
    router.Run(":8080")
}