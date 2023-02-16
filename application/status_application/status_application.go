package statusapplication

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type StatusApp int64

const (
	Error StatusApp = iota - 1 // -1 - Error - status de erro
	Ok                         // 0 - Status normal

	name_error = "Error"
	name_ok    = "Ok"
)

var (
	status_app_name = map[int64]string{
		-1: name_error,
		0:  name_ok,
	}
	status_app_value = map[string]int64{
		name_error: -1,
		name_ok:    0,
	}

	StatusAppAccepts = func() string {

		descriptionStatusAppAccepts := "Status accepts: "

		for enum, types := range status_app_name {
			descriptionStatusAppAccepts += fmt.Sprintf("%d ou %s | ", enum, types)
		}

		return descriptionStatusAppAccepts
	}
)

func (s StatusApp) String() string {
	switch s {
	case Error:
		return name_error
	case Ok:
		return name_ok
	default:
		panic(fmt.Sprintf("status app is invalid. %s", StatusAppAccepts()))
	}
}

func (s StatusApp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(s.String(), start)
}

func (s *StatusApp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var (
		valueContent string
		stsApp       StatusApp
	)

	d.DecodeElement(&valueContent, &start)

	stsApp, err := s.tryParseValueToStatusApp(valueContent)

	if err != nil {
		return err
	}

	*s = StatusApp(stsApp)

	return err

}

func (s StatusApp) MarshalJSON() ([]byte, error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			panic(fmt.Sprintf("Marshal failed. Status app informed: %d. Failed details: %v", s, errRecover))
		}
	}()

	return []byte(fmt.Sprintf(`"%s"`, s.String())), nil
}

func (s *StatusApp) UnmarshalJSON(bytes []byte) error {
	value, err := s.tryGetValueFromJSON(bytes)
	if err == nil && !strings.EqualFold(value, "") {

		stsApp, err := s.tryParseValueToStatusApp(value)

		if err != nil {
			return err
		}

		*s = StatusApp(stsApp)
	}

	return err
}

func (s StatusApp) Value() (driver.Value, error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			panic(fmt.Sprintf("Value failed. Status app informed: %d. Failed details: %v", s, errRecover))
		}
	}()

	return s.String(), nil
}

func (s *StatusApp) Scan(value interface{}) (err error) {

	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("Scan failed for value %v for status app. Details: %v", value, errRecover)
		}
	}()

	switch data := value.(type) {
	case []uint8:
		str := string([]byte(data))

		stsApp, err := s.tryParseValueToStatusApp(str)

		if err != nil {
			return err
		}

		*s = StatusApp(stsApp)

	case int:
		d := int64(data)
		*s = StatusApp(d)
		_ = s.String()
	case int32:
		d := int64(data)
		*s = StatusApp(d)
		_ = s.String()
	case float32:
		d := int64(data)
		*s = StatusApp(d)
		_ = s.String()
	case float64:
		d := int64(data)
		*s = StatusApp(d)
		_ = s.String()
	case int64:
		*s = StatusApp(data)
		_ = s.String()
	case string:
		stsApp, err := s.tryParseValueToStatusApp(data)
		if err != nil {
			panic(err)
		}
		*s = stsApp
	default:
		_ = s.String()
	}

	return nil
}

func (s *StatusApp) tryParseValueToStatusApp(value string) (stsApp StatusApp, err error) {

	toTitle := cases.Title(language.BrazilianPortuguese, cases.NoLower)

	value = toTitle.String(strings.ToLower(value))

	stsAppINT, ok := status_app_value[value]

	if !ok {
		valueINT, err := strconv.Atoi(value)
		if err != nil {
			return Error, fmt.Errorf("the %s is incorret to status app", value)
		}

		stsAppStr, ok := status_app_name[int64(valueINT)]
		if !ok {
			return Error, fmt.Errorf("the %s not valid status app", value)
		}

		stsAppINT, ok = status_app_value[stsAppStr]
		if !ok {
			return Error, fmt.Errorf("the %s is invalid status app", value)
		}
	}

	return StatusApp(stsAppINT), nil
}

func (s *StatusApp) tryGetValueFromJSON(bytes []byte) (value string, err error) {
	value, err = strconv.Unquote(string(bytes))

	if err != nil {

		valueINT := int(Error)

		valueINT, err = strconv.Atoi(string(bytes))

		if err != nil {
			return
		}

		value = fmt.Sprintf("%d", valueINT)
		err = nil
	}

	return
}
