package kuler

import (
	"io/ioutil"
	"testing"
)

func TestServiceError(t *testing.T) {
	c, _ := ioutil.ReadFile("_testdata/error_0.xml")
	e := unmarshalServiceError(c)
	if e == nil {
		t.Fatalf("failed to unmarshal service error!")
	}
}
