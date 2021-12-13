package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
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

	t := time.Now()

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

		p := influxdb2.NewPoint("storcli.controller",
			adapter_tags,
			adapter_fields,
			t)

		fmt.Print(write.PointToLineProtocol(p, time.Second))

		for _, enc := range ctrl.Enclosures {
			enclosure_tags := adapter_tags
			enclosure_tags["enclosure_id"] = strconv.Itoa(enc.EID)
			enclosure_tags["product_id"] = enc.ProductId
			enclosure_tags["vendor_id"] = enc.VendorSpecific

			enclosure_fields := map[string]interface{}{
				"state": enc.State,
			}

			p := influxdb2.NewPoint("storcli.enclosure",
				enclosure_tags,
				enclosure_fields,
				t)

			fmt.Print(write.PointToLineProtocol(p, time.Second))
		}
	}
}
