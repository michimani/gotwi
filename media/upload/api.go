package upload

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/media/upload/types"
)

const (
	initializeEndpoint = "https://api.x.com/2/media/upload/initialize"
	appendEndpoint     = "https://api.x.com/2/media/upload/:mediaID/append"
	finalizeEndpoint   = "https://api.x.com/2/media/upload/:mediaID/finalize"
)

func Initialize(ctx context.Context, c *gotwi.Client, p *types.InitializeInput) (*types.InitializeOutput, error) {
	res := &types.InitializeOutput{}
	if err := c.CallAPI(ctx, initializeEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func Append(ctx context.Context, c *gotwi.Client, p *types.AppendInput) (*types.AppendOutput, error) {
	boundary := p.GenerateBoundary()
	contentType := fmt.Sprintf("multipart/form-data;charset=UTF-8;boundary=%s", boundary)

	res := &types.AppendOutput{}
	ctx = context.WithValue(ctx, "Content-Type", contentType)
	if err := c.CallAPI(ctx, appendEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func Finalize(ctx context.Context, c *gotwi.Client, p *types.FinalizeInput) (*types.FinalizeOutput, error) {
	res := &types.FinalizeOutput{}
	if err := c.CallAPI(ctx, finalizeEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
