package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	ExamplePusher_Push()
}

func ExamplePusher_Push() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
	completionTime.Set(2)

	// completionTime.Set (200) // set can be set to any value (float64)

	if err := push.New("http://127.0.0.1:9091", ""). // Push.new ("Pushgateway Address", "JOB Name")
								Collector(completionTime). //  Add a label to the indicator, you can add multiple
								Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}

}
