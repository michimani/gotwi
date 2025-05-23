package quotetweet

import (
	"context"
	"fmt"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/tweet/quotetweet/types"
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
			name: "normal: valid parameters",
			client: gotwi.NewMockGotwiClientWithFunc(gotwi.MockFuncInput{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			}),
			params: &types.ListInput{
				ID: "1234567890",
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
			params: &types.ListInput{
				ID: "1234567890",
			},
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
