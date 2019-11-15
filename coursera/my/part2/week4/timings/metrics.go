package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	timings = prometheus.NewSummaryVec(
		prometheus.SummaryOpts {
			Name: "method_timing",
			Help: "Per method timing",
		},
		[]string{"method"},
	)
	counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "method_counter",
			Help: "Per method counter",
		},
		[]string{"method"},
	)
)



func init(){
	prometheus.MustRegister(timings)
	prometheus.MustRegister(counter)
}



func mainPage(){
	rnd := time.Duration(rand.Intn(50))
	time.Sleep(time.Millisecond * rnd)
	w.Write([]byte("hello"))
}



func timeTrackingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()

		next.ServeHTTP(w, r)

		timings.
			WithLabelValues(r.URL.Path).
			Observe(float64(time.Since(start).Seconds()))
		counter.
			WithLabelValues(r.URL.Path).
			Inc()
	})
}




func main() {
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/", mainPage)
	siteMux.Handle("/metrics", promhttp.Handler())

	siteHandler := timeTrackingMiddleware(siteMux)

	fmt.Println("starting server at :8083")
	http.ListenAndServe(":8083", siteHandler)
}
