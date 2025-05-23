package tweetcount

import (
	"context"
	"fmt"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/tweet/tweetcount/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListRecent(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.ListRecentInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.ListRecentInput{
				Query: "test query",
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
			params: &types.ListRecentInput{
				Query: "test query",
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
			res, err := ListRecent(ctx, c.client, c.params)

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

func Test_ListAll(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.ListAllInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.ListAllInput{
				Query: "test query",
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
			params: &types.ListAllInput{
				Query: "test query",
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
			res, err := ListAll(ctx, c.client, c.params)

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
