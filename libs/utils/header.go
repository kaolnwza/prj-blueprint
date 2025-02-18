package utils

import (
	"context"
	"net/http"

	"github.com/kaolnwza/proj-blueprint/libs/constants"
)

func NewHeader(ctx context.Context) http.Header {
	requestId := MustGetContext[string](ctx, constants.CtxRequestId)
	header := http.Header{}
	header.Set(constants.RequestIdKey, requestId)
	return header
}
