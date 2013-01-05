package kuler

import (
	"encoding/xml"
	"errors"
	"fmt"
)

var se serviceError

func unmarshalServiceError(data []byte) error {
	if err := xml.Unmarshal(data, &se); err != nil {
		return nil
	}
	return se.Error()
}

type serviceError struct {
	XMLName xml.Name `xml:"root"`
	Code    int      `xml:"error>errorCode"`
	Text    string   `xml:"error>errorText"`
}

func (e *serviceError) String() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Text)
}

func (e *serviceError) Error() error {
	return errors.New(e.String())
}
