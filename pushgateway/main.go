package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	// Prometheus 메트릭 레지스트리 생성
	//prometheusRegistry := prometheus.NewRegistry()

	// Node Exporter 메트릭을 정의
	metrics := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "An example metric for Node Exporter",
		},
		[]string{"server", "os"},
	)

	// 메트릭을 레지스트리에 등록
	//prometheusRegistry.MustRegister(metrics)

	// 예제 메트릭 업데이트
	for i := 0; ; i++ {
		metrics.WithLabelValues("server01", "windows").Set(float64(i))
		metrics.WithLabelValues("server02", "linux").Dec()
		//metrics.WithLabelValues("server03", "osx").Set(rand.Float64())

		// 메트릭을 Pushgateway에 푸시
		if err := push.New("http://localhost:9091", "push_gateway").
			Collector(metrics).Push(); err != nil {
			fmt.Printf("Could not push to Pushgateway: %v\n", err)
		} else {
			fmt.Printf("--> push set (%d) \n", i)
		}

		time.Sleep(1000 * time.Millisecond)
	}

}
