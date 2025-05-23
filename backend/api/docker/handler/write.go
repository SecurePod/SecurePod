package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/malsuke/SecurePod-backend/api/docker/container"
)

func Write(c echo.Context) error {
	cli, err := container.CreateDockerClient()
	if err != nil {
		return err
	}
	r := new(container.WriteRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	err = container.Write(c.Request().Context(), cli, *r)
	if err != nil {
		return err
	}
	return c.JSON(200, "OK")
}
