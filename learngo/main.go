package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/usrname/learngo/scrapper"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := scrapper.CleanString(c.FormValue("term"))
	scrapper.Scrape(term)
	return c.Attachment(fileName, "job.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
