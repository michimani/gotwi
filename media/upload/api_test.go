package upload

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/media/upload/types"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.InitializeInput
		wantErr bool
	}{
		{
			name: "normal: valid parameters",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.InitializeInput{
				TotalBytes: 1000,
				MediaType:  types.MediaTypeJPEG,
			},
			wantErr: false,
		},
		{
			name: "error: parameters is nil",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			ctx := context.Background()
			res, err := Initialize(ctx, c.client, c.params)

			if c.wantErr {
				asst.Error(err)
				asst.Nil(res)
				return
			}

			asst.NoError(err)
			asst.NotNil(res)
		})
	}
}

func TestAppend(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.AppendInput
		wantErr bool
	}{
		{
			name: "normal: valid parameters",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.AppendInput{
				MediaID:      "1234567890",
				Media:        bytes.NewReader([]byte("test data")),
				SegmentIndex: 0,
			},
			wantErr: false,
		},
		{
			name: "error: parameters is nil",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params:  nil,
			wantErr: true,
		},
		{
			name: "error: CallAPI returns error",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return fmt.Errorf("CallAPI error")
				},
			}),
			params: &types.AppendInput{
				MediaID:      "1234567890",
				Media:        bytes.NewReader([]byte("test data")),
				SegmentIndex: 0,
			},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			ctx := context.Background()
			res, err := Append(ctx, c.client, c.params)

			if c.wantErr {
				asst.Error(err)
				asst.Nil(res)
				return
			}

			asst.NoError(err)
			asst.NotNil(res)
		})
	}
}

func TestFinalize(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.FinalizeInput
		wantErr bool
	}{
		{
			name: "normal: valid parameters",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.FinalizeInput{
				MediaID: "1234567890",
			},
			wantErr: false,
		},
		{
			name: "error: parameters is nil",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			ctx := context.Background()
			res, err := Finalize(ctx, c.client, c.params)

			if c.wantErr {
				asst.Error(err)
				asst.Nil(res)
				return
			}

			asst.NoError(err)
			asst.NotNil(res)
		})
	}
}
