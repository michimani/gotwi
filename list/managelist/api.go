package managelist

import (
	"context"
	"errors"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/list/managelist/types"
)

const (
	createEndpoint = "https://api.twitter.com/2/lists"
	updateEndpoint = "https://api.twitter.com/2/lists/:id"
	deleteEndpoint = "https://api.twitter.com/2/lists/:id"
)

// Enables the authenticated user to create a List.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists
func Create(ctx context.Context, c gotwi.IClient, p *types.CreateInput) (*types.CreateOutput, error) {
	if p == nil {
		return nil, errors.New("CreateInput is nil")
	}
	res := &types.CreateOutput{}
	if err := c.CallAPI(ctx, createEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to update the meta data of a specified List that they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/put-lists-id
func Update(ctx context.Context, c gotwi.IClient, p *types.UpdateInput) (*types.UpdateOutput, error) {
	if p == nil {
		return nil, errors.New("UpdateInput is nil")
	}
	res := &types.UpdateOutput{}
	if err := c.CallAPI(ctx, updateEndpoint, "PUT", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Enables the authenticated user to delete a List that they own.
// https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id
func Delete(ctx context.Context, c gotwi.IClient, p *types.DeleteInput) (*types.DeleteOutput, error) {
	if p == nil {
		return nil, errors.New("DeleteInput is nil")
	}
	res := &types.DeleteOutput{}
	if err := c.CallAPI(ctx, deleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
