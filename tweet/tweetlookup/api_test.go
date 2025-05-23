package tweetlookup

import (
	"context"
	"errors"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/tweet/tweetlookup/types"
	"github.com/stretchr/testify/assert"
)

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
			params:  &types.ListInput{},
			mockRes: &types.ListOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.ListInput{},
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
			ctx := context.Background()
			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					if c.mockRes != nil {
						*i.(*types.ListOutput) = *c.mockRes
					}
					return nil
				},
			})

			res, err := List(ctx, mockClient, c.params)
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
			params:  &types.GetInput{},
			mockRes: &types.GetOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.GetInput{},
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
			ctx := context.Background()
			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					if c.mockRes != nil {
						*i.(*types.GetOutput) = *c.mockRes
					}
					return nil
				},
			})

			res, err := Get(ctx, mockClient, c.params)
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
