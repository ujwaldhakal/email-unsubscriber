package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"html/template"
	"io"
	service "github.com/ujwaldhakal/email-unsubscriber/model"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var frontEnd = &cobra.Command{
	Use:   "serve-frontend",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		t := &Template{
			templates: template.Must(template.ParseGlob("public/*.html")),
		}
		e.Renderer = t

		e.GET("/", func(c echo.Context) error {
			return c.Render(http.StatusOK, "index.html", nil)
		})

		e.GET("/get-services", func(c echo.Context) error {
			data := service.Service{}
			return c.JSON(http.StatusOK, data.Get())
		})

		e.GET("/unsubscribe", func(c echo.Context) error {
			data := service.Service{}
			data.Unsubscribe(c.QueryParam("id"))
			return c.JSON(http.StatusOK, nil)
		})

		e.Logger.Fatal(e.Start(":1323"))
	},
	
}



func init()  {
	rootCmd.AddCommand(frontEnd)
}
