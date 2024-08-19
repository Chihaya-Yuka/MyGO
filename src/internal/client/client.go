package client

import (
	"context"
	"llm-client/models"
)

type Client interface {
	Ask(ctx context.Context, request models.Request) (*models.Response, error)
}
