package main

import (
	"github.com/influxdata/influx-cli/v2/clients/user"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/urfave/cli"
)

func newUserCmd() cli.Command {
	return cli.Command{
		Name:  "user",
		Usage: "User management commands",
		Subcommands: []cli.Command{
			newUserCreateCmd(),
			newUserDeleteCmd(),
			newUserListCmd(),
			newUserUpdateCmd(),
			newUserSetPasswordCmd(),
		},
	}
}

func newUserCreateCmd() cli.Command {
	var params user.CreateParams
	return cli.Command{
		Name:  "create",
		Usage: "Create user",
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:   "org-id",
				Usage:  "The ID of the organization",
				EnvVar: "INFLUX_ORG_ID",
				Value:  &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "The name of the organization",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The user name",
				EnvVar:      "INFLUX_NAME",
				Required:    true,
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "password, p",
				Usage:       "The user password",
				Destination: &params.Password,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := user.Client{
				CLI:              getCLI(ctx),
				UsersApi:         api.UsersApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newUserDeleteCmd() cli.Command {
	var id influxid.ID
	return cli.Command{
		Name:  "delete",
		Usage: "Delete user",
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:     "id, i",
				Usage:    "The user ID",
				Required: true,
				Value:    &id,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			client := user.Client{CLI: getCLI(ctx), UsersApi: getAPI(ctx).UsersApi}
			return client.Delete(getContext(ctx), id)
		},
	}
}

func newUserListCmd() cli.Command {
	var params user.ListParams
	return cli.Command{
		Name:    "list",
		Aliases: []string{"find", "ls"},
		Usage:   "List users",
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:  "id, i",
				Usage: "The user ID",
				Value: &params.Id,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The user name",
				Destination: &params.Name,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			client := user.Client{CLI: getCLI(ctx), UsersApi: getAPI(ctx).UsersApi}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newUserUpdateCmd() cli.Command {
	var params user.UpdateParmas
	return cli.Command{
		Name: "update",
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:     "id, i",
				Usage:    "The user ID",
				Required: true,
				Value:    &params.Id,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The user name",
				Destination: &params.Name,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			client := user.Client{CLI: getCLI(ctx), UsersApi: getAPI(ctx).UsersApi}
			return client.Update(getContext(ctx), &params)
		},
	}
}

func newUserSetPasswordCmd() cli.Command {
	var params user.SetPasswordParams
	return cli.Command{
		Name: "password",
		Flags: append(
			commonFlagsNoPrint(),
			&cli.GenericFlag{
				Name:  "id, i",
				Usage: "The user ID",
				Value: &params.Id,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The user name",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "password, p",
				Usage:       "Password to set on the user",
				Destination: &params.Password,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			client := user.Client{CLI: getCLI(ctx), UsersApi: getAPI(ctx).UsersApi}
			return client.SetPassword(getContext(ctx), &params)
		},
	}
}
