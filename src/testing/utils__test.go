package main

import (
      "testing"

	  "test-case-engine/utils"
)

func TestParseConfigs(t *testing.T) {
	expected := "Sample_Application"
	filename := "../input.sample.yaml"

	config := utils.ParseConfigs(filename)

	if expected != config.Title {
		t.Error("ParseConfigs failed.")
	}
}