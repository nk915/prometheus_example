package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Gin 라우터 생성
	r := gin.Default()

	// Prometheus 레지스트리 생성
	prometheusRegistry := prometheus.NewRegistry()

	// Node Exporter 메트릭을 정의
	nodeExporterMetrics := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "An example metric for Node Exporter",
		},
		[]string{"server", "os"},
	)

	// 메트릭을 레지스트리에 등록
	prometheusRegistry.MustRegister(nodeExporterMetrics)

	// Gin 핸들러로 Prometheus 메트릭 엔드포인트 노출
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})))

	// 예제 메트릭 업데이트
	go func() {
		for {
			nodeExporterMetrics.WithLabelValues("server01", "windows").Inc()
			nodeExporterMetrics.WithLabelValues("server02", "linux").Dec()
			nodeExporterMetrics.WithLabelValues("server03", "osx").Set(rand.Float64())
			//nodeExporterMetrics.WithLabelValues("infraware002", "linux").Set(40)
			// 실제로 여기에서 시스템 및 하드웨어 메트릭을 수집하여 메트릭을 업데이트해야 합니다.
			// 여기에는 단순한 예제로 숫자를 설정했습니다.
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("--> metric set \n")
		}
	}()

	// 서버 시작
	fmt.Printf("--> run \n")
	if err := r.Run(":9100"); err != nil {
		panic(err)
	}
}
