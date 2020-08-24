package monitoring

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var summaryVectors = make(map[string]*prometheus.SummaryVec)
var summaryVectorsLock sync.Mutex

func getSummaryVector(metricName string, metricHelp string) *prometheus.SummaryVec {
	summaryVectorsLock.Lock()
	defer summaryVectorsLock.Unlock()

	if summaryVector, ok := summaryVectors[metricName]; ok {
		return summaryVector
	} else {
		newSummaryVector := prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Namespace: "backend",
				Name:      metricName,
				Help:      metricHelp,
			},
			[]string{"service"},
		)

		summaryVectors[metricName] = newSummaryVector

		_ = prometheus.Register(newSummaryVector)
		return newSummaryVector
	}
}

func Prometheus(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		summaryVec := getSummaryVector(route.GetName(), path)

		start := time.Now()
		next.ServeHTTP(rw, r) // the call to next in the chain
		duration := time.Since(start)

		// Store duration of request
		summaryVec.WithLabelValues("duration").Observe(duration.Seconds())

		// Store size of response, if possible.
		size, err := strconv.Atoi(rw.Header().Get("Content-Length"))
		if err == nil {
			summaryVec.WithLabelValues("size").Observe(float64(size))
		}
	})
}
