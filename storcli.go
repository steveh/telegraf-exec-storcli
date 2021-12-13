package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type CLI struct {
}

func (c CLI) List() ([]ListController, error) {
	out, err := c.execute("show", "J").Output()
	if err != nil {
		return nil, fmt.Errorf("executing command: %w", err)
	}

	var res ListAllResponse
	if err := json.Unmarshal([]byte(out), &res); err != nil {
		return nil, fmt.Errorf("unmarshaling json: %w", err)
	}

	results := make([]ListController, 0)

	for _, ctrl := range res.Controllers {
		for _, overview := range ctrl.ResponseData.ITSystemOverview {
			overview.Model = strings.TrimSpace(overview.Model)
			overview.Adapter = strings.TrimSpace(overview.Adapter)

			results = append(results, overview)
		}
	}

	return results, nil
}

func (c CLI) Show(number int) (ShowController, error) {
	out, err := c.execute(fmt.Sprintf("/c%d", number), "show", "all", "J").Output()
	if err != nil {
		return ShowController{}, fmt.Errorf("executing command: %w", err)
	}

	var res ShowAllResponse
	if err := json.Unmarshal([]byte(out), &res); err != nil {
		return ShowController{}, fmt.Errorf("unmarshaling json: %w", err)
	}

	if len(res.Controllers) != 1 {
		return ShowController{}, fmt.Errorf("expected 1 controller, got %d", len(res.Controllers))
	}

	ctrl := res.Controllers[0]

	result := ShowController{
		Number:                      number,
		Model:                       strings.TrimSpace(ctrl.ResponseData.Basics.Model),
		Adapter:                     strings.TrimSpace(ctrl.ResponseData.Basics.AdapterType),
		Serial:                      strings.TrimSpace(ctrl.ResponseData.Basics.SerialNumber),
		FirmwareVersion:             ctrl.ResponseData.Version.FirmwareVersion,
		BiosVersion:                 ctrl.ResponseData.Version.BiosVersion,
		DriverVersion:               ctrl.ResponseData.Version.DriverVersion,
		Status:                      ctrl.ResponseData.Status.ControllerStatus,
		MemoryCorrectableErrors:     ctrl.ResponseData.Status.MemoryCorrectableErrors,
		MemoryUncorrectableErrors:   ctrl.ResponseData.Status.MemoryUncorrectableErrors,
		ROCTemperatureDegreeCelsius: ctrl.ResponseData.Hwcfg.ROCTemperatureDegreeCelsius,
		Enclosures:                  make([]ShowEnclosure, 0, len(ctrl.ResponseData.EnclosureInformation)),
	}

	for _, enc := range ctrl.ResponseData.EnclosureInformation {
		enc.ProductId = strings.TrimSpace(enc.ProductId)
		enc.VendorSpecific = strings.TrimSpace(enc.VendorSpecific)

		result.Enclosures = append(result.Enclosures, enc)
	}

	return result, nil
}

func (c CLI) execute(arg ...string) *exec.Cmd {
	return exec.Command("storcli", arg...)
}
