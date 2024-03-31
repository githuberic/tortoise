package main

import (
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

/**
1: curl -l 'http://localhost:1323/users/Joe'
2: curl -l 'http://localhost:1323/show?team=x-men&member=wolverine'
3: curl -POST -d "name=Joe Smith" -d "email=joe@labstack.com" 'http://localhost:1323/save/www'
4: curl -POST -F "name=Joe Smith" -F "avatar=@/Users/guoqingliu/Downloads/20191113161112816.jpeg" http://localhost:1323/save/fd
*/
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save/www", saveWWW)
	e.POST("/save/fd", saveFD)
	/*
		e.POST("/users", saveUser)
		e.PUT("/users/:id", updateUser)
		e.DELETE("/users/:id", deleteUser)*/
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func saveWWW(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func saveFD(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
