package main

// 参考: https://cloud.tencent.com/document/product/1416/56033
import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	klog "k8s.io/klog/v2"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	recordMetrics()
	fmt.Println("server Metrics started")

	klog.Info("starting gaia schedule scheduler ...")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
