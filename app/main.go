package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

type Member struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var (
	members = map[int]*Member{}
	seq     = 1
)

func createMember(c echo.Context) error {
	m := &Member{
		Id: seq,
	}

	if err := c.Bind(m); err != nil {
		return err
	}

	members[m.Id] = m
	seq++
	return c.JSON(http.StatusCreated, m)
}

func getMember(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, members[id])
}

func updateMember(c echo.Context) error {
	m := new(Member)

	err := c.Bind(m)
	if err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))

	members[id].Name = m.Name

	return c.JSON(http.StatusAccepted, members[id])
}

func deleteMember(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(members, id)
	return c.NoContent(http.StatusNoContent)
}

func getAllMember(c echo.Context) error {
	return nil
}


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.HideBanner = true

	e.POST("/members", createMember)
	e.GET("/members/:id", getMember)
	e.PUT("/members/:id", updateMember)
	e.DELETE("/members/:id", deleteMember)
	e.GET("/members", getAllMember)

	e.Logger.Fatal(e.Start(":8080"))
}
