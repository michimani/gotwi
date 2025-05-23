package userlookup

import (
	"context"
	"fmt"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/user/userlookup/types"
	"github.com/stretchr/testify/assert"
)

func Test_List(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.ListInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.ListInput{
				IDs: []string{"1234567890"},
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
			params: &types.ListInput{
				IDs: []string{"1234567890"},
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
			res, err := List(ctx, c.client, c.params)

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

func Test_Get(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.GetInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.GetInput{
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
			params: &types.GetInput{
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
			res, err := Get(ctx, c.client, c.params)

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

func Test_ListByUsernames(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.ListByUsernamesInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.ListByUsernamesInput{
				Usernames: []string{"username1", "username2"},
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
			params: &types.ListByUsernamesInput{
				Usernames: []string{"username1", "username2"},
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
			res, err := ListByUsernames(ctx, c.client, c.params)

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

func Test_GetByUsername(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.GetByUsernameInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.GetByUsernameInput{
				Username: "username",
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
			params: &types.GetByUsernameInput{
				Username: "username",
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
			res, err := GetByUsername(ctx, c.client, c.params)

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

func Test_GetMe(t *testing.T) {
	cases := []struct {
		name    string
		client  gotwi.IClient
		params  *types.GetMeInput
		wantErr bool
	}{
		{
			name: "success",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params:  &types.GetMeInput{},
			wantErr: false,
		},
		{
			name: "error",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return fmt.Errorf("CallAPI error")
				},
			}),
			params:  &types.GetMeInput{},
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
			res, err := GetMe(ctx, c.client, c.params)

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
