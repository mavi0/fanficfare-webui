package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "static/assets")
	e.File("/", "static/index.html")

	e.POST("/downloadAndGet", downloadAndGet)
	e.POST("/triggerDownload", triggerDownload)
	e.GET("/book/:filename", getBook)

	e.Logger.Info(e.Start(":80"))
}

func downloadbook(c echo.Context) string {
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowURLSchemes("http", "https")

	url := p.Sanitize(c.FormValue("url"))

	fmt.Println("Processing: " + url)

	cmd := exec.Command("fanficfare", "--non-interactive", "-o", "is_adult=true", "-j", url)
	cmd.Dir = "./books"
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("fanficfare failed with %s\n", err)
		panic("Could not generate epub: FanFicFare internal error")
	}

	if string(out[0]) == "B" {
		fmt.Printf("not found %s\n", string(out))
		panic("404 not found")
	}

	var data map[string]interface{}
	err_m := json.Unmarshal([]byte(string(out)), &data)
	if err_m != nil {
		fmt.Println(err_m)
		panic(err)
	}

	fmt.Println("Created: " + data["output_filename"].(string))
	return data["output_filename"].(string)
}

func triggerDownload(c echo.Context) error {
	file_name := downloadbook(c)
	return c.String(http.StatusOK, url.QueryEscape(file_name))
}

func downloadAndGet(c echo.Context) error {
	file_name := downloadbook(c)
	return c.Attachment("./books/" + file_name, file_name)
}

func getBook(c echo.Context) error {
	file_name_encoded := c.Param("filename")
	file_name, err := url.QueryUnescape(file_name_encoded)
	if err != nil {
		panic("file_name is invalid")
	}
	return c.Attachment("./books/" + file_name, file_name)
}
