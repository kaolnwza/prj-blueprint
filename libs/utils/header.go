package utils

import (
	"context"
	"net/http"

	"github.com/kaolnwza/proj-blueprint/libs/consts"
)

func NewHeader(ctx context.Context) http.Header {
	requestId := MustGetContext[string](ctx, consts.CtxRequestId)
	header := http.Header{}
	header.Set(consts.RequestIdKey, requestId)
	return header
}
