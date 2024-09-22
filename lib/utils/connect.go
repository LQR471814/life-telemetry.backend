package utils

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func SlogInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			slog.Info("request", "procedure", req.Spec().Procedure)

			res, err := next(ctx, req)
			if err != nil {
				slog.Error("request failed", "procedure", req.Spec().Procedure, "err", err)
			}

			return res, err
		}
	}
}

func ListenGRPC(mux *http.ServeMux, port int) {
	slog.Info("listening to gRPC...", "port", 8000)
	err := http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", 8000),
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		Fatal(fmt.Errorf("serve http: %w", err))
	}
}
