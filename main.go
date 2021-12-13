package main

import (
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

const (
	precision = time.Nanosecond

	controllerMeasurement = "storcli_controller"
	enclosureMeasurement  = "storcli_enclosure"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cli := CLI{}

	list, err := cli.List()
	check(err)

	t := time.Now().UTC()

	for _, c := range list {
		ctrl, err := cli.Show(c.Number)
		check(err)

		adapter_tags := map[string]string{
			"adapter": ctrl.Adapter,
			"model":   ctrl.Model,
			"serial":  ctrl.Serial,
		}

		adapter_fields := map[string]interface{}{
			"status":                      ctrl.Status,
			"roc_temperature":             ctrl.ROCTemperatureDegreeCelsius,
			"memory_correctable_errors":   ctrl.MemoryCorrectableErrors,
			"memory_uncorrectable_errors": ctrl.MemoryUncorrectableErrors,
		}

		p := influxdb2.NewPoint(controllerMeasurement,
			adapter_tags,
			adapter_fields,
			t)

		fmt.Print(write.PointToLineProtocol(p, precision))

		// for _, enc := range ctrl.Enclosures {
		// 	enclosure_tags := adapter_tags
		// 	enclosure_tags["enclosure_id"] = strconv.Itoa(enc.EID)
		// 	enclosure_tags["product_id"] = enc.ProductId
		// 	enclosure_tags["vendor_id"] = enc.VendorSpecific

		// 	enclosure_fields := map[string]interface{}{
		// 		"state": enc.State,
		// 	}

		// 	p := influxdb2.NewPoint(enclosureMeasurement,
		// 		enclosure_tags,
		// 		enclosure_fields,
		// 		t)

		// 	fmt.Print(write.PointToLineProtocol(p, precision))
		// }
	}
}
