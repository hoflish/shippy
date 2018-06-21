package main

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	consignment, _ := parseFile(defaultFilename)

	assertCorrectMessage := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		switch got.(type) {
		case string:
			if got != want {
				t.Errorf("Got: %s, wanted: %s", got, want)
			}
		case int32:
		case int:
			if got != want {
				t.Errorf("Got: %d, wanted: %d", got, want)
			}
		}

	}
	t.Run("Consignment Id", func(t *testing.T) {
		got := consignment.GetId()
		want := ""
		assertCorrectMessage(t, got, want)
	})

	t.Run("Consignment Description", func(t *testing.T) {
		got := consignment.GetDescription()
		want := "This is a test consignment"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Consignment Weight", func(t *testing.T) {
		got := consignment.GetWeight()
		want := int32(55000)
		assertCorrectMessage(t, got, want)
	})

	t.Run("Consignment Containers", func(t *testing.T) {
		got := len(consignment.GetContainers())
		want := 3
		assertCorrectMessage(t, got, want)
	})

	t.Run("Consignment VesselId", func(t *testing.T) {
		got := consignment.GetVesselId()
		want := ""
		assertCorrectMessage(t, got, want)
	})

}

func TestParseFileWithFakeFile(t *testing.T) {
	consignment, err := parseFile("foo.json")

	if consignment != nil {
		t.Error(err)
	}
}
