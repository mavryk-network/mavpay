package extension

import (
	"context"
	"io"

	"github.com/alis-is/jsonrpc2/endpoints"
	"github.com/alis-is/jsonrpc2/rpc"
	"github.com/mavryk-network/mavpay/constants"
)

/*
All references to endpoints are redirected from here to the endpoints package.
This is done so we can adjust calls (inject call prefix) while mainting consistent behavior across the extension package.
"github.com/alis-is/jsonrpc2/endpoints" should not be used directly
*/

type EndpointClient = endpoints.EndpointClient

func Request[TParams rpc.ParamsType, TResult rpc.ResultType](ctx context.Context, c EndpointClient, method string, params TParams) (*rpc.Response[TResult], error) {
	return endpoints.Request[TParams, TResult](ctx, c, constants.EXTENSION_CALL_PREFIX+method, params)
}

func Notify[TParams rpc.ParamsType](ctx context.Context, c EndpointClient, method string, params TParams) error {
	return endpoints.Notify(ctx, c, constants.EXTENSION_CALL_PREFIX+method, params)
}

func RequestTo[TParams rpc.ParamsType, TResult rpc.ResultType](ctx context.Context, c EndpointClient, method string, params TParams, result *rpc.Response[TResult]) error {
	return endpoints.RequestTo(ctx, c, constants.EXTENSION_CALL_PREFIX+method, params, result)
}

func RegisterEndpointMethod[TParam rpc.ParamsType, TResult rpc.ResultType](c endpoints.EndpointServer, method string, handler endpoints.RpcMethod[TParam, TResult]) {
	endpoints.RegisterEndpointMethod(c, constants.EXTENSION_CALL_PREFIX+method, handler)
}

func NewPlainObjectStream(rw io.ReadWriteCloser) endpoints.ObjectStream {
	return endpoints.NewPlainObjectStream(rw)
}

func NewStreamEndpoint(ctx context.Context, stream endpoints.ObjectStream) *endpoints.StreamEndpoint {
	return endpoints.NewStreamEndpoint(ctx, stream)
}
