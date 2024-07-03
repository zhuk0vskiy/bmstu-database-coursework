package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

type metrics struct {
}

//func NewMetrics(reg prometheus.Registerer) *metrics {
//	m := &metrics{
//		devices: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: "myapp",
//			Name:      "connected_devices",
//			Help:      "Number of currently connected devices.",
//		}),
//	}
//	reg.MustRegister(m.devices)
//	return m
//}

func main() {

	usersOnline := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "users_online",
		})

	go func() {
		for {
			for i := 0; i < 10000; i++ {
				usersOnline.Set(float64(i)) // or: Inc(), Dec(), Add(5), Dec(5)
				fmt.Println(i)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8086", nil)
}
