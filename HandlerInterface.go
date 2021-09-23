package go_flat_http

import (
	"context"
)

type HandlerInterface func(ctx context.Context, input []byte, response Response) Response
