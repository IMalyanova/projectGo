package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const AvgSleep = 50

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/", loadPostHandle)

	siteHandler := timingMiddleware(siteMux)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(" :8080", siteHandler)
}




func trackContextTinings (ctx context.Context, metricName string, start time.Time) {

	// получаем тайминги из контекста
	// поскольку там пустой интерфейс, то нам надо преобразовать к нужному типу

	timings, ok := ctx.Value(timingsKey).(*ctxTimings)

	if !ok {
		return
	}

	elapsed := time.Since(start)
	//лочимся на случай конкурентной записи в мэп

	timings.Lock()
	defer timings.Unlock

	if metric, metricExist := timings.Data[metricName]; !metricExist {
		timings.Data[metricName] = &Timing {
			Count:
			Duration: elapsed,
		}
	} else {
		metric.Count++
		metric.Duration += elapsed
	}
}



type Timing struct {
	Count    int
	Duration time.Duration
}



type ctxTimings struct {
	sync.Mutex
	Data map[string]*Timing
}