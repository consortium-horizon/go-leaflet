package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "html/template"
    "strings"
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
    Group string
}

func (m Marker) GetColor() template.JS {
    db := initDB()
    var err error
    var valueread []byte
    valueread, err = db.Read(m.Group)
    if err==nil{
        var g Group
        err = g.GobDecode(valueread)
        return template.JS(g.Color)
    }
    return "black"
}

func (m Marker) GetIcon() template.JS {
    db := initDB()
    var err error
    var valueread []byte
    valueread, err = db.Read(m.Group)
    if err==nil{
        var g Group
        err = g.GobDecode(valueread)
        return template.JS(g.Icon)
    }
    return "icon-flag"
}

func (m Marker) GetGroupName() template.JS {
    db := initDB()
    var err error
    var valueread []byte
    valueread, err = db.Read(m.Group)
    if err==nil{
        var g Group
        err = g.GobDecode(valueread)
        return template.JS(strings.Replace(g.Name, " ", "", -1))
    }
    return "Unknow Group"
}

func (m Marker) GetTitle() template.JS {
    return template.JS(strings.Replace(m.Title, " ", "", -1))
}

func (m Marker) GetGroup() Group {
    db := initDB()
    var err error
    var valueread []byte
    valueread, err = db.Read(m.Group)
    var g Group
    if err==nil{
        err = g.GobDecode(valueread)
        return g
    }
    return g
}

type Group struct {
    Name string
    Desc string
    Color string
    Icon string   
}

func (g Group) GetName() template.JS {
    return template.JS(strings.Replace(g.Name, " ", "", -1))
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
    err = encoder.Encode(d.Group)
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
    err = decoder.Decode(&d.Desc)
    if err!=nil {
        return err
    }
    return decoder.Decode(&d.Group)
}

func (d *Group) GobEncode() ([]byte, error) {
    w := new(bytes.Buffer)
    encoder := gob.NewEncoder(w)
    err := encoder.Encode(d.Name)
    if err!=nil {
        return nil, err
    }
    err = encoder.Encode(d.Desc)
    if err!=nil {
        return nil, err
    }
    err = encoder.Encode(d.Color)
    if err!=nil {
        return nil, err
    }
    err = encoder.Encode(d.Icon)
    if err!=nil {
        return nil, err
    }
    return w.Bytes(), nil
}

func (d *Group) GobDecode(buf []byte) error {
    r := bytes.NewBuffer(buf)
    decoder := gob.NewDecoder(r)
    err := decoder.Decode(&d.Name)
    if err!=nil {
        return err
    }
    err = decoder.Decode(&d.Desc)
    if err!=nil {
        return err
    }
    err = decoder.Decode(&d.Color)
    if err!=nil {
        return err
    }
    return decoder.Decode(&d.Icon)
}

func (m* Marker) WriteToDB() error {
    keybase := "marker"
    index := 0
    key := keybase + strconv.Itoa(index)
    db := initDB()
    for db.Has(key) {
        index++
        key = keybase + strconv.Itoa(index)
    }
    return m.WriteToDBWithKey(key)
}

func (m* Marker) WriteToDBWithKey(key string) error {
    db := initDB()
    valueasBytes, err := m.GobEncode()
    if err!=nil {
        fmt.Println("%v",err)
    }
    db.Write(key, valueasBytes )
    return err
}

func (m* Group) WriteToDB() error {
    keybase := "group"
    index := 0
    key := keybase + strconv.Itoa(index)
    db := initDB()
    for db.Has(key) {
        index++
        key = keybase + strconv.Itoa(index)
    }
    return m.WriteToDBWithKey(key)
}

func (m* Group) WriteToDBWithKey(key string) error {
    db := initDB()
    valueasBytes, err := m.GobEncode()
    if err!=nil {
        fmt.Println("%v",err)
    }
    db.Write(key, valueasBytes )
    return err
}

func markers() map[int]Marker {
    markers := make(map[int]Marker)
    db := initDB()
    keybase := "marker"
    var err error
    keys := db.KeysPrefix(keybase, nil)
    for key := range keys {
        var valueread []byte
        valueread, err = db.Read(key)
        index,_ := strconv.Atoi(strings.TrimPrefix(key, keybase))
        if err==nil{
            var m Marker
            err = m.GobDecode(valueread)
            markers[index] = m
        }
    }
    return markers
}

func getMarker(key string) *Marker {
    db := initDB()
    var err error
    var valueread []byte
    valueread, err = db.Read(key)
    if err==nil{
        var m Marker
        err = m.GobDecode(valueread)
        return &m
    }
    return nil
}

func getGroup(key string) *Group {
    db := initDB()
    var err error
    var valueread []byte
    valueread, err = db.Read(key)
    if err==nil{
        var g Group
        err = g.GobDecode(valueread)
        return &g
    }
    return nil
}

func groups() map[int]Group {
    groups := make(map[int]Group)
    db := initDB()
    keybase := "group"
    var err error
    keys := db.KeysPrefix(keybase, nil)
    for key := range keys {
        var valueread []byte
        valueread, err = db.Read(key)
        index,_ := strconv.Atoi(strings.TrimPrefix(key, keybase))
        if err==nil{
            var m Group
            err = m.GobDecode(valueread)
            groups[index] = m
        }
    }
    return groups
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

// simulate some private data
var secrets = gin.H{
    "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
    "austin": gin.H{"email": "austin@example.com", "phone": "666"},
    "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}
func main() {
    router := gin.Default()



    router.Static("/css", "./css")
    router.Static("/font", "./font")
    router.Static("/js", "./js")
    router.Static("/images", "./images")
    router.LoadHTMLGlob("templates/*")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "markers": markers,
            "onMap": true,
            "groups": groups,
        })
    })

    // Group using gin.BasicAuth() middleware
    // gin.Accounts is a shortcut for map[string]string
    authorized := router.Group("/markers", gin.BasicAuth(gin.Accounts{
        "admin":   "lch",
    }))

    authorized.GET("/new", func(c *gin.Context) {
        c.HTML(http.StatusOK, "markers-new.tmpl", gin.H{
            "title": "Ajouter un marqueur",
            "markers": markers,
            "groups": groups,
        })
    })
    authorized.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "markers-manage.tmpl", gin.H{
            "title": "Liste des marqueurs",
            "markers": markers,
            "groups": groups,
        })
    })
    authorized.GET("../groups", func(c *gin.Context) {
        c.HTML(http.StatusOK, "groups-manage.tmpl", gin.H{
            "title": "Liste des types de marqueurs",
            "markers": markers,
            "groups": groups,
        })
    })

    authorized.POST("/add", func(c *gin.Context) {
        title := c.PostForm("title")
        desc := c.PostForm("desc")
        x, err := strconv.Atoi(c.PostForm("x"))
        y, err := strconv.Atoi(c.PostForm("y"))
        group := c.PostForm("group")
        fmt.Printf("%v",group)
        if err!=nil {
            fmt.Printf("%v",err)
            return
        }
        //fmt.Printf("title: %v; desc: %v; x: %v; y: %v; color: %v; icon:%v", title, desc, x, y, color, icon)
        m := Marker{Title:title, Desc: desc, X: x, Y: y, Group: group }
        m.WriteToDB()
        c.Redirect(http.StatusMovedPermanently, c.Request.Header.Get("Referer"))
    })
    authorized.POST("/addgroup", func(c *gin.Context) {
        name := c.PostForm("name")
        desc := c.PostForm("desc")
        color := c.PostForm("color")
        icon := c.PostForm("icon")
        g := Group{Name:name, Desc: desc, Color:color, Icon:icon }
        g.WriteToDB()
        c.Redirect(http.StatusMovedPermanently, c.Request.Header.Get("Referer"))
    })

    authorized.POST("/delete", func(c *gin.Context) {
        key := c.PostForm("key")
        //fmt.Printf("title: %v; desc: %v; x: %v; y: %v", title, desc, x, y)
        db := initDB();
        db.Erase(key)
        c.Redirect(http.StatusMovedPermanently, c.Request.Header.Get("Referer"))
    })
    authorized.POST("/deletegroup", func(c *gin.Context) {
        key := c.PostForm("key")
        //fmt.Printf("title: %v; desc: %v; x: %v; y: %v", title, desc, x, y)
        db := initDB();
        db.Erase(key)
        c.Redirect(http.StatusMovedPermanently, c.Request.Header.Get("Referer"))
    })

    authorized.POST("/edit", func(c *gin.Context) {
        key := c.PostForm("key")
        m := getMarker(key)
        if m!=nil {
            c.HTML(http.StatusOK, "markers-new.tmpl", gin.H{
            "title": "Editer un marker",
            "marker": *m,
            "groups": groups,
            "key": key,
            "edit": true,
            })

        }
    })

    authorized.POST("/editgroup", func(c *gin.Context) {
        key := c.PostForm("key")
        g := getGroup(key)
        if g!=nil {
            c.HTML(http.StatusOK, "markers-new.tmpl", gin.H{
            "title": "Editer un groupe",
            "group": *g,
            "groups": groups,
            "key": key,
            "editgroup": true,
            })

        }
    })

    authorized.POST("/save", func(c *gin.Context) {
        title := c.PostForm("title")
        desc := c.PostForm("desc")
        x, err := strconv.Atoi(c.PostForm("x"))
        y, err := strconv.Atoi(c.PostForm("y"))
        group := c.PostForm("group")
        key := c.PostForm("key")
        if err!=nil {
            fmt.Printf("%v",err)
            return
        }
        //fmt.Printf("title: %v; desc: %v; x: %v; y: %v", title, desc, x, y)
        m := Marker{Title:title, Desc: desc, X: x, Y: y, Group: group }
        m.WriteToDBWithKey(key)
        c.Redirect(http.StatusMovedPermanently, "/markers")
    })

    authorized.POST("/savegroup", func(c *gin.Context) {
        name := c.PostForm("name")
        desc := c.PostForm("desc")
        color := c.PostForm("color")
        icon := c.PostForm("icon")
        key := c.PostForm("key")
        g := Group{Name:name, Desc: desc, Color:color, Icon:icon }
        g.WriteToDBWithKey(key)
        c.Redirect(http.StatusMovedPermanently, "/groups")
    })




    // /admin/secrets endpoint
    // hit "localhost:8080/admin/secrets
    // authorized.GET("/secrets", func(c *gin.Context) {
    //     // get user, it was set by the BasicAuth middleware
    //     user := c.MustGet(gin.AuthUserKey).(string)
    //     if secret, ok := secrets[user]; ok {
    //         c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
    //     } else {
    //         c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
    //     }
    // })
    
    router.Run(":8080")
}