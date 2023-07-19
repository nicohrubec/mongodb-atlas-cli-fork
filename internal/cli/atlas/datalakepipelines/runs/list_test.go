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

//go:build unit

// This code was autogenerated at 2023-04-25T17:59:19+01:00. Note: Manual updates are allowed, but may be overwritten.

package runs

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201002/admin"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPipelineRunsLister(ctrl)

	expected := &atlasv2.PaginatedPipelineRun{
		Results: []atlasv2.IngestionPipelineRun{
			{
				Id:          pointer.Get("507f1f77bcf86cd799439011"),
				DatasetName: pointer.Get("dataset"),
				State:       pointer.Get("IDLE"),
			},
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store:        mockStore,
		pipelineName: "pipeline",
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		PipelineRuns(listOpts.ProjectID, listOpts.pipelineName).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID                         DATASET NAME   STATE
507f1f77bcf86cd799439011   dataset        IDLE

`, buf.String())
	t.Log(buf.String())
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID, flag.Pipeline, flag.Output},
	)
}
