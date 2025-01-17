package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/influxdata/influx-cli/v2/pkg/stdio"
)

type Client struct {
	clients.CLI
	api.UsersApi
	api.OrganizationsApi
}

type CreateParams struct {
	clients.OrgParams
	Name     string
	Password string
}

var ErrMustSpecifyUser = errors.New("must specify user ID or user name")

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	if params.Password != "" && len(params.Password) < stdio.MinPasswordLen {
		return clients.ErrPasswordIsTooShort
	}

	orgID, err := c.GetOrgIdI(ctx, params.OrgID, params.OrgName, c.OrganizationsApi)
	if err != nil {
		return err
	}

	user, err := c.PostUsers(ctx).User(api.User{Name: params.Name}).Execute()
	if err != nil {
		return fmt.Errorf("failed to create user %q: %w", params.Name, err)
	}
	if err := c.printUsers(printUserOpts{user: &user}); err != nil {
		return err
	}

	memberBody := api.AddResourceMemberRequestBody{Id: *user.Id}
	if _, err := c.PostOrgsIDMembers(ctx, orgID).AddResourceMemberRequestBody(memberBody).Execute(); err != nil {
		_, _ = c.StdIO.WriteErr([]byte("WARN: initial password not set for user, use `influx user password` to set it\n"))
		return fmt.Errorf("failed setting org membership for user %q, use `influx org members add` to retry: %w", user.Name, err)
	}

	if params.Password == "" {
		_, _ = c.StdIO.WriteErr([]byte("WARN: initial password not set for user, use `influx user password` to set it\n"))
		return nil
	}

	passwordBody := api.PasswordResetBody{Password: params.Password}
	if err := c.PostUsersIDPassword(ctx, user.GetId()).PasswordResetBody(passwordBody).Execute(); err != nil {
		return fmt.Errorf("failed setting password for user %q, use `influx user password` to retry: %w", user.Name, err)
	}
	return nil
}

func (c Client) Delete(ctx context.Context, id influxid.ID) error {
	user, err := c.GetUsersID(ctx, id.String()).Execute()
	if err != nil {
		return fmt.Errorf("user %q not found: %w", id, err)
	}
	if err := c.DeleteUsersID(ctx, id.String()).Execute(); err != nil {
		return fmt.Errorf("failed to delete user %q: %w", id, err)
	}
	return c.printUsers(printUserOpts{user: &user, deleted: true})
}

type ListParams struct {
	Id   influxid.ID
	Name string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	req := c.GetUsers(ctx)
	if params.Id.Valid() {
		req = req.Id(params.Id.String())
	}
	if params.Name != "" {
		req = req.Name(params.Name)
	}
	users, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list users: %w", err)
	}
	printOpts := printUserOpts{}
	if users.Users != nil {
		printOpts.users = *users.Users
	}
	return c.printUsers(printOpts)
}

type UpdateParmas struct {
	Id   influxid.ID
	Name string
}

func (c Client) Update(ctx context.Context, params *UpdateParmas) error {
	update := api.User{}
	if params.Name != "" {
		update.SetName(params.Name)
	}
	user, err := c.PatchUsersID(ctx, params.Id.String()).User(update).Execute()
	if err != nil {
		return fmt.Errorf("failed to update user %q: %w", params.Id, err)
	}
	return c.printUsers(printUserOpts{user: &user})
}

type SetPasswordParams struct {
	Id       influxid.ID
	Name     string
	Password string
}

func (c Client) SetPassword(ctx context.Context, params *SetPasswordParams) (err error) {
	if !params.Id.Valid() && params.Name == "" {
		return ErrMustSpecifyUser
	}
	id := params.Id.String()
	displayName := id
	if !params.Id.Valid() {
		displayName = params.Name
		users, err := c.GetUsers(ctx).Name(params.Name).Execute()
		if err != nil {
			return fmt.Errorf("failed to find user %q: %w", params.Name, err)
		}
		if users.Users == nil || len(*users.Users) == 0 {
			return fmt.Errorf("no user found with name %q", params.Name)
		}
		id = (*users.Users)[0].GetId()
	}

	password := params.Password
	if password == "" {
		password, err = c.StdIO.GetPassword(fmt.Sprintf("Please type new password for %q", displayName))
		if err != nil {
			return err
		}
	}

	body := api.PasswordResetBody{Password: password}
	if err := c.PostUsersIDPassword(ctx, id).PasswordResetBody(body).Execute(); err != nil {
		return fmt.Errorf("failed to set password for user %q: %w", params.Id.String(), err)
	}
	_, err = c.StdIO.Write([]byte(fmt.Sprintf("Successfully updated password for user %q\n", displayName)))
	return err
}

type printUserOpts struct {
	user    *api.UserResponse
	users   []api.UserResponse
	deleted bool
}

func (c Client) printUsers(opts printUserOpts) error {
	if c.PrintAsJSON {
		var v interface{}
		if opts.user != nil {
			v = opts.user
		} else {
			v = opts.users
		}
		return c.PrintJSON(v)
	}

	headers := []string{"ID", "Name"}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	if opts.user != nil {
		opts.users = append(opts.users, *opts.user)
	}

	var rows []map[string]interface{}
	for _, u := range opts.users {
		row := map[string]interface{}{
			"ID":   u.GetId(),
			"Name": u.GetName(),
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}
