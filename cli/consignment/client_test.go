package main

import (
	"strings"
	"testing"
)

const (
	filename = "consignment.json"
)

func TestParseFile(t *testing.T) {
	consignment, _ := parseFile(filename)

	t.Run("Consignment Id", func(t *testing.T) {
		got := consignment.GetId()
		want := ""

		if got != want {
			t.Errorf("Got: %s, wanted: %s", got, want)
		}
	})

	t.Run("Consignment Description", func(t *testing.T) {
		got := consignment.GetDescription()
		want := "consignment"

		if !strings.Contains(got, want) {
			t.Errorf("wanted '%s' to contain '%s' subtring", got, want)
		}

	})

	t.Run("Consignment Weight", func(t *testing.T) {
		got := consignment.GetWeight()
		want := int32(55000)

		if got != want {
			t.Errorf("Got: %d, wanted: %d", got, want)
		}
	})

	t.Run("Consignment Containers", func(t *testing.T) {
		got := len(consignment.GetContainers())
		want := 3

		if got != want {
			t.Errorf("Got: %d, wanted: %d", got, want)
		}
	})

	t.Run("Consignment VesselId", func(t *testing.T) {
		got := consignment.GetVesselId()
		want := ""

		if got != want {
			t.Errorf("Got: %s, wanted: %s", got, want)
		}
	})

}
