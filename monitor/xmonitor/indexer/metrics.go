package indexer

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	latencyHist = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "latency_seconds",
		Help:      "Cross chain latency in seconds per stream per xdapp (submit-emit timestamp)",
		Buckets:   prometheus.ExponentialBucketsRange(time.Second.Seconds(), time.Hour.Seconds(), 10),
	}, []string{"stream", "xdapp"})

	successCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "success_total",
		Help:      "Total number of successful cross chain transactions per stream per xdapp",
	}, []string{"stream", "xdapp"})

	revertCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "revert_total",
		Help:      "Total number of reverted cross chain transactions per stream per xdapp",
	}, []string{"stream", "xdapp"})

	excessGasHist = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "excess_gas",
		Help:      "Excess gas per stream per xdapp (msg.GasLimit - receipt.GasUsed)",
		Buckets:   prometheus.ExponentialBucketsRange(1, 1e6, 10),
	}, []string{"stream", "xdapp"})
)

type sample struct {
	Stream    string
	XDApp     string
	Latency   time.Duration
	ExcessGas uint64
	Success   bool
}

func instrumentSample(s sample) {
	if s.Success {
		successCounter.WithLabelValues(s.Stream, s.XDApp).Inc()
	} else {
		revertCounter.WithLabelValues(s.Stream, s.XDApp).Inc()
	}
	latencyHist.WithLabelValues(s.Stream, s.XDApp).Observe(s.Latency.Seconds())
	excessGasHist.WithLabelValues(s.Stream, s.XDApp).Observe(float64(s.ExcessGas))
}
