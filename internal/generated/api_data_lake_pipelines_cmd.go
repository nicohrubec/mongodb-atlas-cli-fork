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

type CreatePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *CreatePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreatePipelineOpts) Run(ctx context.Context) error {
	params := &admin.CreatePipelineApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.CreatePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreatePipelineBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreatePipelineOpts{}
	cmd := &cobra.Command{
		Use:     "createPipeline",
		// Aliases: []string{"?"},
		Short:   "Create One Data Lake Pipeline",
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
type DeletePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
}

func (opts *DeletePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeletePipelineOpts) Run(ctx context.Context) error {
	params := &admin.DeletePipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.DeletePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeletePipelineBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeletePipelineOpts{}
	cmd := &cobra.Command{
		Use:     "deletePipeline",
		// Aliases: []string{"?"},
		Short:   "Remove One Data Lake Pipeline",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type DeletePipelineRunDatasetOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
	pipelineRunId string
}

func (opts *DeletePipelineRunDatasetOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeletePipelineRunDatasetOpts) Run(ctx context.Context) error {
	params := &admin.DeletePipelineRunDatasetApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
		PipelineRunId: opts.pipelineRunId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.DeletePipelineRunDatasetWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeletePipelineRunDatasetBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeletePipelineRunDatasetOpts{}
	cmd := &cobra.Command{
		Use:     "deletePipelineRunDataset",
		// Aliases: []string{"?"},
		Short:   "Delete Pipeline Run Dataset",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")
	cmd.Flags().StringVar(&opts.pipelineRunId, "pipelineRunId", , "Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	_ = cmd.MarkFlagRequired("pipelineRunId")

	return cmd
}
type GetPipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
}

func (opts *GetPipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetPipelineOpts) Run(ctx context.Context) error {
	params := &admin.GetPipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.GetPipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetPipelineBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetPipelineOpts{}
	cmd := &cobra.Command{
		Use:     "getPipeline",
		// Aliases: []string{"?"},
		Short:   "Return One Data Lake Pipeline",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type GetPipelineRunOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
	pipelineRunId string
}

func (opts *GetPipelineRunOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetPipelineRunOpts) Run(ctx context.Context) error {
	params := &admin.GetPipelineRunApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
		PipelineRunId: opts.pipelineRunId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.GetPipelineRunWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetPipelineRunBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetPipelineRunOpts{}
	cmd := &cobra.Command{
		Use:     "getPipelineRun",
		// Aliases: []string{"?"},
		Short:   "Return One Data Lake Pipeline Run",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")
	cmd.Flags().StringVar(&opts.pipelineRunId, "pipelineRunId", , "Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")
	_ = cmd.MarkFlagRequired("pipelineRunId")

	return cmd
}
type ListPipelineRunsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	createdBefore time.Time
}

func (opts *ListPipelineRunsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPipelineRunsOpts) Run(ctx context.Context) error {
	params := &admin.ListPipelineRunsApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		CreatedBefore: opts.createdBefore,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineRunsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListPipelineRunsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListPipelineRunsOpts{}
	cmd := &cobra.Command{
		Use:     "listPipelineRuns",
		// Aliases: []string{"?"},
		Short:   "Return All Data Lake Pipeline Runs from One Project",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")
	cmd.Flags().StringVar(&opts.createdBefore, "createdBefore", , "If specified, Atlas returns only Data Lake Pipeline runs initiated before this time and date.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type ListPipelineSchedulesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
}

func (opts *ListPipelineSchedulesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPipelineSchedulesOpts) Run(ctx context.Context) error {
	params := &admin.ListPipelineSchedulesApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineSchedulesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListPipelineSchedulesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListPipelineSchedulesOpts{}
	cmd := &cobra.Command{
		Use:     "listPipelineSchedules",
		// Aliases: []string{"?"},
		Short:   "Return Available Ingestion Schedules for One Data Lake Pipeline",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type ListPipelineSnapshotsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	completedAfter time.Time
}

func (opts *ListPipelineSnapshotsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPipelineSnapshotsOpts) Run(ctx context.Context) error {
	params := &admin.ListPipelineSnapshotsApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		CompletedAfter: opts.completedAfter,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineSnapshotsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListPipelineSnapshotsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListPipelineSnapshotsOpts{}
	cmd := &cobra.Command{
		Use:     "listPipelineSnapshots",
		// Aliases: []string{"?"},
		Short:   "Return Available Backup Snapshots for One Data Lake Pipeline",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")
	cmd.Flags().StringVar(&opts.completedAfter, "completedAfter", , "Date and time after which MongoDB Cloud created the snapshot. If specified, MongoDB Cloud returns available backup snapshots created after this time and date only. This parameter expresses its value in the ISO 8601 timestamp format in UTC.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type ListPipelinesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *ListPipelinesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPipelinesOpts) Run(ctx context.Context) error {
	params := &admin.ListPipelinesApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelinesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListPipelinesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListPipelinesOpts{}
	cmd := &cobra.Command{
		Use:     "listPipelines",
		// Aliases: []string{"?"},
		Short:   "Return All Data Lake Pipelines from One Project",
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
type PausePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
}

func (opts *PausePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *PausePipelineOpts) Run(ctx context.Context) error {
	params := &admin.PausePipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.PausePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func PausePipelineBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := PausePipelineOpts{}
	cmd := &cobra.Command{
		Use:     "pausePipeline",
		// Aliases: []string{"?"},
		Short:   "Pause One Data Lake Pipeline",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type ResumePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
}

func (opts *ResumePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ResumePipelineOpts) Run(ctx context.Context) error {
	params := &admin.ResumePipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ResumePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ResumePipelineBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ResumePipelineOpts{}
	cmd := &cobra.Command{
		Use:     "resumePipeline",
		// Aliases: []string{"?"},
		Short:   "Resume One Data Lake Pipeline",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type TriggerSnapshotIngestionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
}

func (opts *TriggerSnapshotIngestionOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *TriggerSnapshotIngestionOpts) Run(ctx context.Context) error {
	params := &admin.TriggerSnapshotIngestionApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.TriggerSnapshotIngestionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func TriggerSnapshotIngestionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := TriggerSnapshotIngestionOpts{}
	cmd := &cobra.Command{
		Use:     "triggerSnapshotIngestion",
		// Aliases: []string{"?"},
		Short:   "Trigger on demand snapshot ingestion",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type UpdatePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	pipelineName string
}

func (opts *UpdatePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdatePipelineOpts) Run(ctx context.Context) error {
	params := &admin.UpdatePipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.UpdatePipelineWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdatePipelineBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdatePipelineOpts{}
	cmd := &cobra.Command{
		Use:     "updatePipeline",
		// Aliases: []string{"?"},
		Short:   "Update One Data Lake Pipeline",
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
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", , "Human-readable label that identifies the Data Lake Pipeline.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}

func DataLakePipelinesBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dataLakePipelines",
		Short:   "Returns, adds, edits, and removes Atlas Data Lake Pipelines and associated runs.",
	}
	cmd.AddCommand(
		CreatePipelineBuilder(),
		DeletePipelineBuilder(),
		DeletePipelineRunDatasetBuilder(),
		GetPipelineBuilder(),
		GetPipelineRunBuilder(),
		ListPipelineRunsBuilder(),
		ListPipelineSchedulesBuilder(),
		ListPipelineSnapshotsBuilder(),
		ListPipelinesBuilder(),
		PausePipelineBuilder(),
		ResumePipelineBuilder(),
		TriggerSnapshotIngestionBuilder(),
		UpdatePipelineBuilder(),
	)
	return cmd
}
