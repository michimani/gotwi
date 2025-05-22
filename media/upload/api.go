package upload

import (
	"context"
	"errors"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/media/upload/types"
)

const (
	initializeEndpoint = "https://api.x.com/2/media/upload/initialize"
	appendEndpoint     = "https://api.x.com/2/media/upload/:mediaID/append"
	finalizeEndpoint   = "https://api.x.com/2/media/upload/:mediaID/finalize"
)

func Initialize(ctx context.Context, c gotwi.IClient, p *types.InitializeInput) (*types.InitializeOutput, error) {
	if p == nil {
		return nil, errors.New("InitializeInput is nil")
	}
	res := &types.InitializeOutput{}
	if err := c.CallAPI(ctx, initializeEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func Append(ctx context.Context, c gotwi.IClient, p *types.AppendInput) (*types.AppendOutput, error) {
	if p == nil {
		return nil, errors.New("AppendInput is nil")
	}
	boundary := p.GenerateBoundary()
	contentType := fmt.Sprintf("multipart/form-data;charset=UTF-8;boundary=%s", boundary)

	res := &types.AppendOutput{}
	ctx = context.WithValue(ctx, "Content-Type", contentType)
	if err := c.CallAPI(ctx, appendEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func Finalize(ctx context.Context, c gotwi.IClient, p *types.FinalizeInput) (*types.FinalizeOutput, error) {
	if p == nil {
		return nil, errors.New("FinalizeInput is nil")
	}
	res := &types.FinalizeOutput{}
	if err := c.CallAPI(ctx, finalizeEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
