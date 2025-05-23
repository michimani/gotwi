package spacelookup

import (
	"context"
	"errors"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/space/spacelookup/types"
	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.GetInput
		mockRes *types.GetOutput
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			params:  &types.GetInput{ID: "test-id"},
			mockRes: &types.GetOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.GetInput{ID: "test-id"},
			mockRes: nil,
			mockErr: errors.New("error"),
			wantErr: true,
		},
		{
			name:    "error: params is nil",
			params:  nil,
			mockRes: nil,
			mockErr: nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Parallel()

			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					return nil
				},
			})

			res, err := Get(context.Background(), mockClient, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.mockRes, res)
		})
	}
}

func Test_List(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.ListInput
		mockRes *types.ListOutput
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			params:  &types.ListInput{IDs: []string{"test-id"}},
			mockRes: &types.ListOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.ListInput{IDs: []string{"test-id"}},
			mockRes: nil,
			mockErr: errors.New("error"),
			wantErr: true,
		},
		{
			name:    "error: params is nil",
			params:  nil,
			mockRes: nil,
			mockErr: nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Parallel()

			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					return nil
				},
			})

			res, err := List(context.Background(), mockClient, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.mockRes, res)
		})
	}
}

func Test_ListByCreatorIDs(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.ListByCreatorIDsInput
		mockRes *types.ListByCreatorIDsOutput
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			params:  &types.ListByCreatorIDsInput{UserIDs: []string{"test-id"}},
			mockRes: &types.ListByCreatorIDsOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.ListByCreatorIDsInput{UserIDs: []string{"test-id"}},
			mockRes: nil,
			mockErr: errors.New("error"),
			wantErr: true,
		},
		{
			name:    "error: params is nil",
			params:  nil,
			mockRes: nil,
			mockErr: nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Parallel()

			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					return nil
				},
			})

			res, err := ListByCreatorIDs(context.Background(), mockClient, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.mockRes, res)
		})
	}
}

func Test_ListBuyers(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.ListBuyersInput
		mockRes *types.ListBuyersOutput
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			params:  &types.ListBuyersInput{ID: "test-id"},
			mockRes: &types.ListBuyersOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.ListBuyersInput{ID: "test-id"},
			mockRes: nil,
			mockErr: errors.New("error"),
			wantErr: true,
		},
		{
			name:    "error: params is nil",
			params:  nil,
			mockRes: nil,
			mockErr: nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Parallel()

			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					return nil
				},
			})

			res, err := ListBuyers(context.Background(), mockClient, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.mockRes, res)
		})
	}
}

func Test_ListTweets(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.ListTweetsInput
		mockRes *types.ListTweetsOutput
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			params:  &types.ListTweetsInput{ID: "test-id"},
			mockRes: &types.ListTweetsOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.ListTweetsInput{ID: "test-id"},
			mockRes: nil,
			mockErr: errors.New("error"),
			wantErr: true,
		},
		{
			name:    "error: params is nil",
			params:  nil,
			mockRes: nil,
			mockErr: nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Parallel()

			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					return nil
				},
			})

			res, err := ListTweets(context.Background(), mockClient, c.params)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, res)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.mockRes, res)
		})
	}
}
