package keeper

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter = otel.Meter("github.com/yaseen786000/cloudos-chain/x/counter")

	countMetric metric.Int64Counter
)

func init() {
	var err error
	countMetric, err = meter.Int64Counter("count")
	if err != nil {
		panic(err)
	}
}
