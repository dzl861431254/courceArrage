package http

import (
	"context"
	"gitlab.shoplazza.site/common/nemo/nemo"
	"gitlab.shoplazza.site/common/nemo/nemo/server/grpch"
	"gitlab.shoplazza.site/common/pbc/baymax/go_code/baymax/healthz"
)

func StartServer(ctx context.Context, cfg grpch.HttpConfig) error {
	s := NewServer(cfg, nemo.New("baymax"))
	select {
	case <-ctx.Done():
	}
	s.GracefulStop()
	return nil
}

func NewServer(cfg grpch.HttpConfig, s *nemo.App) *grpch.Server {
	s.NewHttpServer(cfg)

	// Register Http Healthz
	s.HttpServer.Handler(healthz.RegisterHealthzServiceHandler)

	// Register Handler

	// Do things that grpc can't handle, such as file upload and download
	/**
	s.HttpServer.GWHandler(func(ctx context.Context, mux *httptrace.Router) error {
		mux.HandleFunc("/yourpath_ok", func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		})
		return nil
	})
	*/

	go func() {
		s.HttpServer.Serve()
	}()
	return s.HttpServer
}
