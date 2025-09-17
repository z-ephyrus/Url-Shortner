package main

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Link struct{
	 Id string
	 Url string
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
var linkMap = map[string]*Link{"example": { Id: "exa", Url: "https://gobyexample.com/", },}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	e.GET("/:id",rediectHandler)
	e.GET("/",indexHandler)
	e.POST("/submit",submitHandler)

	e.Logger.Fatal(e.Start(":8080"))
	
}

func rediectHandler(c echo.Context) error {
	id := c.Param("id")
	link,f := linkMap[id]

	if !f{
		return c.String(http.StatusNotFound,"link not found")
	} 
	return c.Redirect(http.StatusMovedPermanently, link.Url)
}



func generateString(lenght int)string  {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		var result []byte

		for i:=0; i<lenght; i++{
			index := r.Intn(len(charset))
			result = append(result, charset[index])
		}
		return string(result)
}


func indexHandler(c echo.Context) error {
	html := `
		<h1>Submit a new website</h1>
		<form action="/submit" method="POST">
		<label for="url">Website URL:</label>
		<input type="text" id="url" name="url">
		<input type="submit" value="Submit">
		</form>
		<h2>Existing Links </h2>
		<ul>
	`

	for _, link := range linkMap {
		html += `<li><a href="/` + link.Id + `">` + link.Id + `</a></li>`
	}
	html += `</ul>`

	return c.HTML(http.StatusOK, html)
}

func submitHandler(c echo.Context) error  {
	url := c.FormValue("url")
	if url == ""{
		return c.String(http.StatusBadRequest, "Enter a URL")
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
	    url = "https://" + url
	}

	id := generateString(8)

	linkMap[id] = &Link{Id: id, Url: url} 

	return c.Redirect(http.StatusSeeOther, "/")
}