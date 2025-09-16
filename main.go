package main

import (
	"math/rand"
	"net/http"
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
	//e.GET("/",indexHandler)
	//e.POST("/Submit",submitHandler)

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