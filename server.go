package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "static/assets")
	e.File("/", "static/index.html")

	e.POST("/genbook", genbook)

	e.Logger.Info(e.Start(":80"))
}

func genbook(c echo.Context) error {
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowURLSchemes("http", "https")

	url := p.Sanitize(c.FormValue("url"))

	fmt.Println("Processing: " + url)

	out, err := exec.Command("fanficfare", "-j", url).Output()
	if err != nil {
		fmt.Printf("fanficfare failed with %s\n", err)
		return c.HTML(http.StatusInternalServerError, "Could not generate epub: FanFicFare internal error")
	}

	if string(out[0]) == "B" {
		return c.HTML(http.StatusNotFound, string(out))
	}

	var data map[string]interface{}
	err_m := json.Unmarshal([]byte(string(out)), &data)
	if err_m != nil {
		fmt.Println(err_m)
		panic(err)
		return c.HTML(http.StatusInternalServerError, "Could not extract epub metadata: Echo internal error")
	}

	fmt.Println("Created: " + data["output_filename"].(string))

	return c.Attachment(data["output_filename"].(string), data["output_filename"].(string))

}
