package filteredstream

import (
	"context"
	"errors"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/tweet/filteredstream/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListRules(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.ListRulesInput
		mockRes *types.ListRulesOutput
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			params:  &types.ListRulesInput{},
			mockRes: &types.ListRulesOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.ListRulesInput{},
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
			client := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					if c.mockRes != nil {
						*i.(*types.ListRulesOutput) = *c.mockRes
					}
					return nil
				},
			})

			res, err := ListRules(ctx, client, c.params)
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

func Test_CreateRules(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.CreateRulesInput
		mockRes *types.CreateRulesOutput
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			params:  &types.CreateRulesInput{},
			mockRes: &types.CreateRulesOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.CreateRulesInput{},
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
			client := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					if c.mockRes != nil {
						*i.(*types.CreateRulesOutput) = *c.mockRes
					}
					return nil
				},
			})

			res, err := CreateRules(ctx, client, c.params)
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

func Test_DeleteRules(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.DeleteRulesInput
		mockRes *types.DeleteRulesOutput
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			params:  &types.DeleteRulesInput{},
			mockRes: &types.DeleteRulesOutput{},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.DeleteRulesInput{},
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
			client := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					if c.mockRes != nil {
						*i.(*types.DeleteRulesOutput) = *c.mockRes
					}
					return nil
				},
			})

			res, err := DeleteRules(ctx, client, c.params)
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
