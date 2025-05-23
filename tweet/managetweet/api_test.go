package managetweet

import (
	"context"
	"errors"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/tweet/managetweet/types"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.CreateInput
		mockRes *types.CreateOutput
		mockErr error
		want    *types.CreateOutput
		wantErr bool
	}{
		{
			name: "success",
			params: &types.CreateInput{
				Text: gotwi.String("test tweet"),
			},
			mockRes: &types.CreateOutput{
				Data: struct {
					ID   *string `json:"id"`
					Text *string `json:"text"`
				}{
					ID:   gotwi.String("1234567890"),
					Text: gotwi.String("test tweet"),
				},
			},
			mockErr: nil,
			want: &types.CreateOutput{
				Data: struct {
					ID   *string `json:"id"`
					Text *string `json:"text"`
				}{
					ID:   gotwi.String("1234567890"),
					Text: gotwi.String("test tweet"),
				},
			},
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.CreateInput{},
			mockRes: nil,
			mockErr: errors.New("error occurred"),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "error: params is nil",
			params:  nil,
			mockRes: nil,
			mockErr: nil,
			want:    nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					if c.mockRes != nil {
						*i.(*types.CreateOutput) = *c.mockRes
					}
					return nil
				},
			})

			got, err := Create(context.Background(), mockClient, c.params)
			if c.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}

func Test_Delete(t *testing.T) {
	cases := []struct {
		name    string
		params  *types.DeleteInput
		mockRes *types.DeleteOutput
		mockErr error
		want    *types.DeleteOutput
		wantErr bool
	}{
		{
			name: "success",
			params: &types.DeleteInput{
				ID: "1234567890",
			},
			mockRes: &types.DeleteOutput{
				Data: struct {
					Deleted *bool `json:"deleted"`
				}{
					Deleted: gotwi.Bool(true),
				},
			},
			mockErr: nil,
			want: &types.DeleteOutput{
				Data: struct {
					Deleted *bool `json:"deleted"`
				}{
					Deleted: gotwi.Bool(true),
				},
			},
			wantErr: false,
		},
		{
			name:    "error",
			params:  &types.DeleteInput{},
			mockRes: nil,
			mockErr: errors.New("error occurred"),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "error: params is nil",
			params:  nil,
			mockRes: nil,
			mockErr: nil,
			want:    nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockClient := gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					if c.mockErr != nil {
						return c.mockErr
					}
					if c.mockRes != nil {
						*i.(*types.DeleteOutput) = *c.mockRes
					}
					return nil
				},
			})

			got, err := Delete(context.Background(), mockClient, c.params)
			if c.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}
