package follow

import (
	"context"
	"fmt"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/user/follow/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListFollowings(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.ListFollowingsInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.ListFollowingsInput{
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
			params: &types.ListFollowingsInput{
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
			res, err := ListFollowings(ctx, c.client, c.params)

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

func Test_ListFollowers(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.ListFollowersInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.ListFollowersInput{
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
			params: &types.ListFollowersInput{
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
			res, err := ListFollowers(ctx, c.client, c.params)

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

func Test_CreateFollowing(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.CreateFollowingInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.CreateFollowingInput{
				ID:       "1234567890",
				TargetID: "0987654321",
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
			params: &types.CreateFollowingInput{
				ID:       "1234567890",
				TargetID: "0987654321",
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
			res, err := CreateFollowing(ctx, c.client, c.params)

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

func Test_DeleteFollowing(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.DeleteFollowingInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.DeleteFollowingInput{
				SourceUserID: "1234567890",
				TargetID:     "0987654321",
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
			params: &types.DeleteFollowingInput{
				SourceUserID: "1234567890",
				TargetID:     "0987654321",
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
			res, err := DeleteFollowing(ctx, c.client, c.params)

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
