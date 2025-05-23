package listfollow_test

import (
	"context"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/list/listfollow"
	"github.com/michimani/gotwi/list/listfollow/types"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_ListFollowers(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.ListFollowersInput
		mock    func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error
		expect  *types.ListFollowersOutput
		wantErr bool
	}{
		{
			name: "success",
			params: &types.ListFollowersInput{
				ID: "test-list-id",
			},
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				output := i.(*types.ListFollowersOutput)
				output.Data = []resources.User{}
				return nil
			},
			expect: &types.ListFollowersOutput{
				Data: []resources.User{},
			},
			wantErr: false,
		},
		{
			name: "error",
			params: &types.ListFollowersInput{
				ID: "test-list-id",
			},
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				return assert.AnError
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:   "error: params is nil",
			params: nil,
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				return nil
			},
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			client := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI:              c.mock,
				MockIsReady:              func() bool { return true },
				MockAccessToken:          func() string { return "test-token" },
				MockAuthenticationMethod: func() gotwi.AuthenticationMethod { return gotwi.AuthenMethodOAuth2BearerToken },
			})

			res, err := listfollow.ListFollowers(context.Background(), client, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, res)
		})
	}
}

func Test_ListFollowed(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.ListFollowedInput
		mock    func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error
		expect  *types.ListFollowedOutput
		wantErr bool
	}{
		{
			name: "success",
			params: &types.ListFollowedInput{
				ID: "test-user-id",
			},
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				output := i.(*types.ListFollowedOutput)
				output.Data = []resources.List{}
				return nil
			},
			expect: &types.ListFollowedOutput{
				Data: []resources.List{},
			},
			wantErr: false,
		},
		{
			name: "error",
			params: &types.ListFollowedInput{
				ID: "test-user-id",
			},
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				return assert.AnError
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:   "error: params is nil",
			params: nil,
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				return nil
			},
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			client := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI:              c.mock,
				MockIsReady:              func() bool { return true },
				MockAccessToken:          func() string { return "test-token" },
				MockAuthenticationMethod: func() gotwi.AuthenticationMethod { return gotwi.AuthenMethodOAuth2BearerToken },
			})

			res, err := listfollow.ListFollowed(context.Background(), client, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, res)
		})
	}
}

func Test_Create(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.CreateInput
		mock    func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error
		expect  *types.CreateOutput
		wantErr bool
	}{
		{
			name: "success",
			params: &types.CreateInput{
				ID:     "test-user-id",
				ListID: "test-list-id",
			},
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				output := i.(*types.CreateOutput)
				output.Data.Following = true
				return nil
			},
			expect: &types.CreateOutput{
				Data: struct {
					Following bool `json:"following"`
				}{
					Following: true,
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			params: &types.CreateInput{
				ID:     "test-user-id",
				ListID: "test-list-id",
			},
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				return assert.AnError
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:   "error: params is nil",
			params: nil,
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				return nil
			},
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			client := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI:              c.mock,
				MockIsReady:              func() bool { return true },
				MockAccessToken:          func() string { return "test-token" },
				MockAuthenticationMethod: func() gotwi.AuthenticationMethod { return gotwi.AuthenMethodOAuth2BearerToken },
			})

			res, err := listfollow.Create(context.Background(), client, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, res)
		})
	}
}

func Test_Delete(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.DeleteInput
		mock    func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error
		expect  *types.DeleteOutput
		wantErr bool
	}{
		{
			name: "success",
			params: &types.DeleteInput{
				ID:     "test-user-id",
				ListID: "test-list-id",
			},
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				output := i.(*types.DeleteOutput)
				output.Data.Following = false
				return nil
			},
			expect: &types.DeleteOutput{
				Data: struct {
					Following bool `json:"following"`
				}{
					Following: false,
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			params: &types.DeleteInput{
				ID:     "test-user-id",
				ListID: "test-list-id",
			},
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				return assert.AnError
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:   "error: params is nil",
			params: nil,
			mock: func(ctx context.Context, endpoint string, method string, p util.Parameters, i util.Response) error {
				return nil
			},
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			client := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI:              c.mock,
				MockIsReady:              func() bool { return true },
				MockAccessToken:          func() string { return "test-token" },
				MockAuthenticationMethod: func() gotwi.AuthenticationMethod { return gotwi.AuthenMethodOAuth2BearerToken },
			})

			res, err := listfollow.Delete(context.Background(), client, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, res)
		})
	}
}
