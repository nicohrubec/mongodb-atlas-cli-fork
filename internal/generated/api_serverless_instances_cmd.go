// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"context"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type CreateServerlessInstanceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *CreateServerlessInstanceOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateServerlessInstanceOpts) Run(ctx context.Context) error {
	params := &admin.CreateServerlessInstanceApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.CreateServerlessInstanceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateServerlessInstanceBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateServerlessInstanceOpts{}
	cmd := &cobra.Command{
		Use:     "createServerlessInstance",
		// Aliases: []string{"?"},
		Short:   "Create One Serverless Instance in One Project",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type DeleteServerlessInstanceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	name string
}

func (opts *DeleteServerlessInstanceOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteServerlessInstanceOpts) Run(ctx context.Context) error {
	params := &admin.DeleteServerlessInstanceApiParams{
		GroupId: opts.groupId,
		Name: opts.name,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.DeleteServerlessInstanceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteServerlessInstanceBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteServerlessInstanceOpts{}
	cmd := &cobra.Command{
		Use:     "deleteServerlessInstance",
		// Aliases: []string{"?"},
		Short:   "Remove One Serverless Instance from One Project",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.name, "name", , "Human-readable label that identifies the serverless instance.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}
type GetServerlessInstanceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	name string
}

func (opts *GetServerlessInstanceOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetServerlessInstanceOpts) Run(ctx context.Context) error {
	params := &admin.GetServerlessInstanceApiParams{
		GroupId: opts.groupId,
		Name: opts.name,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.GetServerlessInstanceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetServerlessInstanceBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetServerlessInstanceOpts{}
	cmd := &cobra.Command{
		Use:     "getServerlessInstance",
		// Aliases: []string{"?"},
		Short:   "Return One Serverless Instance from One Project",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.name, "name", , "Human-readable label that identifies the serverless instance.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}
type ListServerlessInstancesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListServerlessInstancesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListServerlessInstancesOpts) Run(ctx context.Context) error {
	params := &admin.ListServerlessInstancesApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.ListServerlessInstancesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListServerlessInstancesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListServerlessInstancesOpts{}
	cmd := &cobra.Command{
		Use:     "listServerlessInstances",
		// Aliases: []string{"?"},
		Short:   "Return All Serverless Instances from One Project",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")

	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type UpdateServerlessInstanceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	name string
}

func (opts *UpdateServerlessInstanceOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateServerlessInstanceOpts) Run(ctx context.Context) error {
	params := &admin.UpdateServerlessInstanceApiParams{
		GroupId: opts.groupId,
		Name: opts.name,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.UpdateServerlessInstanceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateServerlessInstanceBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateServerlessInstanceOpts{}
	cmd := &cobra.Command{
		Use:     "updateServerlessInstance",
		// Aliases: []string{"?"},
		Short:   "Update One Serverless Instance in One Project",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.name, "name", , "Human-readable label that identifies the serverless instance.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func ServerlessInstancesBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "serverlessInstances",
		Short:   "Returns, adds, edits, and removes serverless instances.",
	}
	cmd.AddCommand(
		CreateServerlessInstanceBuilder(),
		DeleteServerlessInstanceBuilder(),
		GetServerlessInstanceBuilder(),
		ListServerlessInstancesBuilder(),
		UpdateServerlessInstanceBuilder(),
	)
	return cmd
}
