package handler

import (
	"log/slog"

	docker "github.com/malsuke/SecurePod-backend/api/docker/container"
	"github.com/malsuke/SecurePod-backend/utils"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	ssh = []*docker.ContainerService{
		docker.NewContainerWithConfig(
			&container.Config{
				Image: "password-crack-attack:latest",
				Tty:   true,
			},
			&container.HostConfig{
				AutoRemove: true,
				// PortBindings: nat.PortMap{
				// 	"22/tcp": []nat.PortBinding{
				// 		{
				// 			HostPort: "0",
				// 		},
				// 	},
				// },
				Resources: container.Resources{
					Memory: 1024 * 1024 * 1024,
				},
			},
			nil,
			nil,
		),
		docker.NewContainerWithConfig(
			&container.Config{
				Image: "password-crack-defense:latest",
				Tty:   true,
			},
			&container.HostConfig{
				AutoRemove: true,
				// PortBindings: nat.PortMap{
				// 	"22/tcp": []nat.PortBinding{
				// 		{
				// 			HostPort: "0",
				// 		},
				// 	},
				// },
				Resources: container.Resources{
					Memory: 1024 * 1024 * 1024,
				},
			},
			nil,
			nil,
		),
	}

	sqli = []*docker.ContainerService{
		docker.NewContainerWithConfig(
			&container.Config{
				Image: "sqli-app:latest",
			},
			&container.HostConfig{
				PortBindings: nat.PortMap{
					"80/tcp": []nat.PortBinding{
						{
							HostPort: "0",
						},
					},
				},
				AutoRemove: true,
				Resources: container.Resources{
					Memory: 1024 * 1024 * 1024,
				},
			},
			nil,
			nil,
		),
		docker.NewContainerWithConfig(
			&container.Config{
				Image: "sqli-db:latest",
			},
			&container.HostConfig{
				AutoRemove: true,
				Resources: container.Resources{
					Memory: 1024 * 1024 * 1024,
				},
			},
			nil,
			nil,
		),
	}

	cli = []*docker.ContainerService{
		docker.NewContainerWithConfig(
			&container.Config{
				Image: "cli:latest",
				Tty:   true,
			},
			&container.HostConfig{
				PortBindings: nat.PortMap{
					"6080/tcp": []nat.PortBinding{
						{
							HostPort: "0",
						},
					},
				},
				AutoRemove: true,
				Resources: container.Resources{
					Memory: 1024 * 1024 * 1024,
				},
			},
			nil,
			nil,
		),
	}

	ContainerList = map[string][]*docker.ContainerService{
		"sshBrute": ssh,
		"sqli":     sqli,
		"cli":      cli,
	}
)

func Create(c echo.Context) error {
	tag := c.Param("tag")
	slog.Info(tag)
	ctx := c.Request().Context()
	cli, err := docker.CreateDockerClient()
	if err != nil {
		return err
	}

	var ids []map[string]string

	nid := utils.GenerateUUID()
	nid, err = docker.CreateNetwork(ctx, cli, nid)
	if err != nil {
		return err
	}
	log.Debug().Str("network", nid).Msg("network created")

	for _, container := range ContainerList[tag] {
		if tag == "sqli" {
			container.SetNetworkEndpointConfigWithAlias(nid)
		} else {
			container.SetNetworkEndpointConfig(nid)
		}
		log.Debug().Str("network", nid).Msg("network attached")
		id, err := container.CreateContainer(ctx, cli)
		if err != nil {
			return err
		}
		ids = append(ids, map[string]string{
			"id": *id,
		})
		c.Logger().Debug(id)
	}
	return c.JSON(200, ids)
}
