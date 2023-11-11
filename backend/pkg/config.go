package pkg

import (
	"github.com/DenisPalnitsky/file-manager-task/backend/pkg/metrics"
	"github.com/DenisPalnitsky/file-manager-task/backend/pkg/rest"
)

type Config struct {
	IsDebugMode bool   `default:"false" envvar:"DEBUG_MODE"`
	LogLevel    string `default:"info" envvar:"LOG_LEVEL"`
	Metrics     metrics.MetricsConfig
	Service     rest.HttpConfig
}
