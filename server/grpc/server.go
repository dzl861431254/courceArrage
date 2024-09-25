package grpc

import (
	"baymax/config/config"
	"baymax/handler"
	"context"
	"github.com/getsentry/sentry-go"
	"gitlab.shoplazza.site/common/nemo/nemo"
	"gitlab.shoplazza.site/common/nemo/nemo/middleware"
	"gitlab.shoplazza.site/common/nemo/nemo/server/grpcx"
	"gitlab.shoplazza.site/common/nemo/nemo/zlog"
	"gitlab.shoplazza.site/common/pbc/baymax/go_code/baymax/healthz"
)

func StartServer(ctx context.Context, cfg grpcx.GRPCServerConfig) error {
	s := NewServer(cfg, nemo.New("baymax"))
	select {
	case <-ctx.Done():
	}
	s.GracefulStop()
	return nil
}

func NewServer(cfg grpcx.GRPCServerConfig, s *nemo.App) *grpcx.Server {
	// init GRPC server with global config
	s.NewGrpcServer(cfg)

	// default middleware
	defaultInterceptors, err := middleware.Default(sentry.CurrentHub())
	if err != nil {
		zlog.Error(err.Error())
	}

	s.StartMetricServer(context.Background(), config.Cfg.MetricConfig)

	// Register default middleware
	s.GrpcServer.AddInterceptors(defaultInterceptors)

	// Register custom middleware
	// s.GrpcServer.AddInterceptor(custom middleware)

	// Register Service
	s.GrpcServer.RegisterService(&healthz.HealthzService_ServiceDesc, handler.NewHealthzHandler())

	go func() {
		s.GrpcServer.Serve()
	}()
	return s.GrpcServer
}
