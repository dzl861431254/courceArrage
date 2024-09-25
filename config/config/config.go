package config

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"gitlab.shoplazza.site/common/nemo/nemo/component/db"
	"gitlab.shoplazza.site/common/nemo/nemo/component/pubsub/kafka"
	"gitlab.shoplazza.site/common/nemo/nemo/component/pubsub/sqs"
	"gitlab.shoplazza.site/common/nemo/nemo/component/redis"
	nconfig "gitlab.shoplazza.site/common/nemo/nemo/config"
	"gitlab.shoplazza.site/common/nemo/nemo/i18n"
	"gitlab.shoplazza.site/common/nemo/nemo/metric"
	"gitlab.shoplazza.site/common/nemo/nemo/server/grpch"
	"gitlab.shoplazza.site/common/nemo/nemo/server/grpcx"
	"gitlab.shoplazza.site/common/nemo/nemo/zlog"
)

type config struct {
	Env                string `env:"ENV" envDefault:"dev"`
	BasePath           string `env:"BASE_PATH" envDefault:""`
	LogLevel           string `env:"LOG_LEVEL" envDefault:"warn"`
	SentryDSN          string `env:"SENTRY_DSN"`
	Cluster            string `env:"CLUSTER"`
	GracePeriodSeconds int64  `env:"GRACE_PERIOD_SECONDS" envDefault:"30"`

	metric.MetricConfig
	grpcx.GRPCServerConfig
	grpch.HttpConfig

	db.DBConfig
	redis.RedisConfig
	kafka.KafkaConfig
	sqs.SQSConfig

	i18n.I18NConfig
}

var Cfg *config

func init() {
	if len(os.Getenv("ENV")) == 0 {
		_, filename, _, _ := runtime.Caller(0)
		rootPath := path.Dir(path.Dir(path.Dir(filename)))
		err := godotenv.Load(rootPath + "/.env")
		if err != nil {
			panic(err)
		}
	}

	Cfg = &config{}
	// load all config
	// Environment variables are currently supported and compatible old env
	nconfig.MustLoad(Cfg, nconfig.UseEnv())
	if Cfg.SentryDSN != "" {
		zlog.EnableSentryLogger(Cfg.SentryDSN)
	}

	if Cfg.BasePath == "" {
		// get root dir
		_, filename, _, _ := runtime.Caller(0)
		Cfg.BasePath = path.Dir(path.Dir(filename))
	}

	if Cfg.FilesPath == "" {
		Cfg.FilesPath = path.Join(Cfg.BasePath, "locale")
	}

	// i18n.MustInit(Cfg.I18NConfig)
	// DB = db.InitGormV2DB(Cfg.DBConfig)

	fmt.Printf("%+v", Cfg)
}
