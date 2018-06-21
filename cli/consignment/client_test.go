package main

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	c, _ := parseFile(defaultFilename)

	testCases := []struct {
		fieldName string
		got       interface{}
		want      interface{}
	}{
		{fieldName: "Id", got: c.GetId(), want: ""},
		{fieldName: "Description", got: c.GetDescription(), want: "This is a test consignment"},
		{fieldName: "VesselId", got: c.GetVesselId(), want: ""},
		{fieldName: "Weight", got: c.GetWeight(), want: int32(55000)},
		{fieldName: "Containers", got: len(c.GetContainers()), want: 3},
	}

	assertCorrectMessage := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		switch got.(type) {
		case string:
			if got != want {
				t.Errorf("%#v Got %s wanted %s", c, got, want)
			}
		case int32:
		case int:
			if got != want {
				t.Errorf("%#v Got %d wanted %d", c, got, want)
			}
		}

	}

	for _, tt := range testCases {
		t.Run(tt.fieldName, func(t *testing.T) {
			assertCorrectMessage(t, tt.got, tt.want)
		})
	}
}

func TestParseFileWithFakeFileName(t *testing.T) {
	consignment, err := parseFile("fileNotExist.json")
	if consignment != nil {
		t.Error(err)
	}
}
