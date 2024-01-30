package overssh

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func downloadHandler(c echo.Context) error {
	params := c.Param("id")

	if params == "eba5349c34" {
		c.File("./public/example_file.txt")
		return nil
	}

	if params == "" {
		return notFound(c)
	}

	if pipe, ok := Pipes[params]; ok {
		defer pipe.Close()
		return c.Stream(200, "text", pipe.transfer.reader)
	}

	return notFound(c)
}

func notFound(c echo.Context) error {
	return c.File("./public/404.html")
}

func index(c echo.Context) error {
	return c.File("./public/index.html")
}

func StartDownloadServer() error {
	e := echo.New()
	e.HideBanner = true
	e.Static("/", "./public")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)

	e.GET("/d/:id", downloadHandler)

	echo.NotFoundHandler = func(c echo.Context) error {
		return notFound(c)
	}

	if os.Getenv("DEV") == "true" {
		e.Start(":3000")
	} else {
		err := e.StartTLS(":443", "/etc/letsencrypt/live/www.overs.sh/fullchain.pem", "/etc/letsencrypt/live/www.overs.sh/privkey.pem")
		if err != nil {
			return err
		}
	}

	return nil
}
