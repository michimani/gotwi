package managelist

import (
	"context"
	"fmt"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/list/managelist/types"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.CreateInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.CreateInput{
				Name:        "test list",
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			wantErr: false,
		},
		{
			name: "error",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return fmt.Errorf("CallAPI error")
				},
			}),
			params: &types.CreateInput{
				Name:        "test list",
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			wantErr: true,
		},
		{
			name: "error: params is nil",
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
			res, err := Create(ctx, c.client, c.params)

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

func Test_Update(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.UpdateInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.UpdateInput{
				ID:          "1234567890",
				Name:        gotwi.String("test list"),
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			wantErr: false,
		},
		{
			name: "error",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return fmt.Errorf("CallAPI error")
				},
			}),
			params: &types.UpdateInput{
				ID:          "1234567890",
				Name:        gotwi.String("test list"),
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			wantErr: true,
		},
		{
			name: "error: params is nil",
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
			res, err := Update(ctx, c.client, c.params)

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

func Test_Delete(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.DeleteInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.DeleteInput{
				ID: "1234567890",
			},
			wantErr: false,
		},
		{
			name: "error",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return fmt.Errorf("CallAPI error")
				},
			}),
			params: &types.DeleteInput{
				ID: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "error: params is nil",
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
			res, err := Delete(ctx, c.client, c.params)

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
