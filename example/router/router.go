// Code generated by go-astro. DO NOT EDIT.
package router

import (
	"github.com/labstack/echo/v4"
	"go.quinn.io/go-astro/example/pages"
)

// RegisterRoutes adds all page routes to the Echo instance
func RegisterRoutes(e *echo.Echo) {

	e.GET("/blog/:slug", HandleBlogPost)
	e.GET("/", HandleIndex)
}



func HandleBlogPost(c echo.Context) error {
	slug := c.Param("slug")
	return pages.BlogPost(slug).Render(c.Request().Context(), c.Response().Writer)
}


func HandleIndex(c echo.Context) error {
	return pages.Index().Render(c.Request().Context(), c.Response().Writer)
}


