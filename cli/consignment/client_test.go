package main

import (
	"path/filepath"
	"testing"
)

func TestParseFile(t *testing.T) {
	path := filepath.Join("test-fixtures", "consignment_data.json")
	c, err := parseFile(path)

	ok(t, err)

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

	for _, tc := range testCases {
		t.Run(tc.fieldName, func(t *testing.T) {
			equals(t, tc.got, tc.want)
		})
	}
}

func TestParseFileWithFakeFileName(t *testing.T) {
	consignment, err := parseFile("fileNotExist.json")
	if consignment != nil {
		t.Error(err)
	}
}
