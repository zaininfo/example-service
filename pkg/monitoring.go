package pkg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ServiceStatus defines a metric for reporting availability status of service
var ServiceStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "service_status",
	Help: "A custom healthcheck that takes into account DB availability.",
}, []string{"app"})

func init() {
	//prometheus.MustRegister(ServiceStatus)
}

// StartMonitoring starts the monitoring server
func StartMonitoring() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", MonitoringPort), nil))
}
