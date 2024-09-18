package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

type metrics struct {
	usersOnline              prometheus.Gauge
	usersSingUp              prometheus.Gauge
	producersReserved        prometheus.Gauge
	instrumentalistsReserved prometheus.Gauge
	microphonesReserved      prometheus.Gauge
	guitarsReserved          prometheus.Gauge
	currentReserves          prometheus.Gauge
}

func newMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		usersOnline: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "users_online",
			Help: "Current count of online users",
		}),
		usersSingUp: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "users_singup",
			Help: "Current count of sing up users",
		}),
		producersReserved: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "producers_reserved",
			Help: "Current count of using producers",
		}),
		instrumentalistsReserved: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "instrumentalist_reserved",
			Help: "Current count of using instrumentalists",
		}),
		microphonesReserved: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "microphones_reserved",
			Help: "Current count of using microphones",
		}),
		guitarsReserved: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "guitars_reserved",
			Help: "Current count of using guitars",
		}),
		currentReserves: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "current_reserves",
			Help: "Current count of reserves",
		}),
	}
	reg.MustRegister(m.usersOnline)
	reg.MustRegister(m.usersSingUp)
	reg.MustRegister(m.producersReserved)
	reg.MustRegister(m.instrumentalistsReserved)
	reg.MustRegister(m.microphonesReserved)
	reg.MustRegister(m.guitarsReserved)
	reg.MustRegister(m.currentReserves)

	return m
}

func (m *metrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Header.Get("Method")

	fmt.Println("get ", method)
	switch method {

	case "IncUsersSingUp":
		m.usersSingUp.Inc()

	case "IncUsersOnline":
		m.usersOnline.Inc()
	case "DecUsersOnline":
		m.usersOnline.Dec()

	case "IncProducersReserved":
		m.producersReserved.Inc()
	case "DecProducersReserved":
		m.producersReserved.Dec()

	case "IncInstrumentalistsReserved":
		m.instrumentalistsReserved.Inc()
	case "DecInstrumentalistsReserved":
		m.instrumentalistsReserved.Dec()

	case "IncMicrophonesReserved":
		m.microphonesReserved.Inc()
	case "DecMicrophonesReserved":
		m.microphonesReserved.Dec()

	case "IncGuitarsReserved":
		m.guitarsReserved.Inc()
	case "DecGuitarsReserved":
		m.guitarsReserved.Dec()

	case "IncCurrentReserves":
		m.currentReserves.Inc()
	case "DecCurrentReserves":
		m.currentReserves.Dec()
	}
}

func main() {
	url := os.Getenv("MONITORING_URL")

	url = "0.0.0.0:8080"

	if url == "" {
		fmt.Println("no url")
		return
	}

	reg := prometheus.NewRegistry()
	m := newMetrics(reg)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.Handle("/monitoring", m)
	//http.Post()
	err := http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Errorf("fail to start server on %s", url)
	}

	//fmt.Println("serve on http://" + url)
}
