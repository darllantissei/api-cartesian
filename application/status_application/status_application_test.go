package statusapplication

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatusApp_String(t *testing.T) {

	var stsApp StatusApp = -1

	require.Equal(t, name_error, stsApp.String(), "Must be Error")

	stsApp = Ok

	require.Equal(t, name_ok, stsApp.String(), "Must be Ok")

	stsApp = 3

	require.Panics(t, func() {
		_ = stsApp.String()
	}, "Status app don't mapped")
}

func TestStatusApp_MarshalJSON(t *testing.T) {

	var stsApp StatusApp = -1

	contentJsonExpected := `"Error"`

	parsedJson, err := stsApp.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Error")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Error has been success")

	stsApp = Ok

	contentJsonExpected = `"Ok"`

	parsedJson, err = stsApp.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Ok")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Ok has been success")

	stsApp = 3

	require.Panics(t, func() {
		parsedJson, err = stsApp.MarshalJSON()
	}, "Status app don't mapped, then will happen panic")

}

func TestStatusApp_UnmarshalJSON(t *testing.T) {

	var stsApp StatusApp

	possibilities := []string{
		`"Error"`,
		`"error"`,
		`"ERROR"`,
		`"-1"`,
		`-1`,
	}

	for _, possibility := range possibilities {

		err := stsApp.UnmarshalJSON([]byte(possibility))

		typeExpected := Error

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, stsApp, "Content parsed with value Error has been success")
	}

	possibilities = []string{
		`"Ok"`,
		`"ok"`,
		`"OK"`,
		`"0"`,
		`0`,
	}

	for _, possibility := range possibilities {

		err := stsApp.UnmarshalJSON([]byte(possibility))

		typeExpected := Ok

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, stsApp, "Content parsed with value Ok has been success")
	}

	possibilities = []string{
		`"sdklfjsfd"`,
		`"Asdfg"`,
		`"SDFSDF"`,
		`"445"`,
		`4345345`,
	}

	for _, possibility := range possibilities {

		err := stsApp.UnmarshalJSON([]byte(possibility))

		require.NotNil(t, err, "Content parsed with error, the error must be not nil")

	}

}

func TestStatusApp_Value(t *testing.T) {
	var stsApp StatusApp = -1

	contentValueExpected := `Error`

	parsedValue, err := stsApp.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Error")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Error has been success")

	stsApp = Ok

	contentValueExpected = `Ok`

	parsedValue, err = stsApp.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Ok")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Ok has been success")

	stsApp = 3

	require.Panics(t, func() {
		parsedValue, err = stsApp.Value()
	}, "Status app don't mapped, then will happen panic")
}

func TestStatusApp_Scan(t *testing.T) {
	var stsApp StatusApp

	possibilities := []interface{}{
		`Error`,
		`error`,
		`ERROR`,
		`-1`,
		[]uint8(`Error`),
		[]uint8(`error`),
		[]uint8(`ERROR`),
		[]uint8(`-1`),
		int(-1),
		int32(-1),
		int64(-1),
		float32(-1),
		float64(-1),
		StatusApp(-1),
	}

	for _, possibility := range possibilities {

		err := stsApp.Scan(possibility)

		typeExpected := Error

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, stsApp, "Content parsed with value Error has been success")
	}

	possibilities = []interface{}{
		`Ok`,
		`ok`,
		`OK`,
		`0`,
		[]uint8(`Ok`),
		[]uint8(`ok`),
		[]uint8(`OK`),
		[]uint8(`0`),
		int(0),
		int32(0),
		int64(0),
		float32(0),
		float64(0),
		StatusApp(0),
	}

	for _, possibility := range possibilities {

		err := stsApp.Scan(possibility)

		typeExpected := Ok

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, stsApp, "Content parsed with value Ok has been success")
	}

	possibilities = []interface{}{
		`Adfsf`,
		`sdfsdf`,
		`SFDDSFDS`,
		`645`,
		[]uint8(`Adfsf`),
		[]uint8(`sdfsdf`),
		[]uint8(`SFDDSFDS`),
		[]uint8(`756`),
		int(8657),
		int32(8657),
		int64(8657),
		float32(8657),
		float64(8657),
		StatusApp(8657),
	}

	for _, possibility := range possibilities {

		err := stsApp.Scan(possibility)

		require.NotNil(t, err, "Content parsed with error, the error must be not nil")

	}
}

func TestStatusApp_MarshalXML(t *testing.T) {
	possibilities := []interface{}{
		-1,
		0,
		"Error",
		"error",
		"ERROR",
		"Ok",
		"ok",
		"OK",
		"-1",
		"0",
		Error,
		Ok,
	}

	var stsApp StatusApp

	for _, possibility := range possibilities {

		valueStsApp, err := stsApp.tryParseValueToStatusApp(fmt.Sprint(possibility))

		require.Nil(t, err, "Error must be nil when parse possibility to status application")

		expected := []byte(fmt.Sprintf(`<status>%s</status>`, valueStsApp))

		contentOUT := new(bytes.Buffer)

		err = valueStsApp.MarshalXML(xml.NewEncoder(contentOUT), xml.StartElement{Name: xml.Name{Space: "", Local: "status"}})

		require.Nil(t, err, "Error must be returned nil when parsed status application")

		require.Equal(t, expected, contentOUT.Bytes(), "The parse XML must be equal expected")

	}
}

func TestStatusApp_UnmarshalXML(t *testing.T) {
	possibilities := []interface{}{
		-1,
		0,
		"Error",
		"error",
		"ERROR",
		"Ok",
		"ok",
		"OK",
		"-1",
		"0",
		Error,
		Ok,
	}

	var stsAppUnmarshalXML StatusApp

	for _, possibility := range possibilities {

		expectStatus, err := stsAppUnmarshalXML.tryParseValueToStatusApp(fmt.Sprint(possibility))

		require.Nil(t, err, "Error must be nil")

		contentIN := []byte(`<StatusApp>` + fmt.Sprint(possibility) + `</StatusApp>`)

		err = xml.Unmarshal(contentIN, &stsAppUnmarshalXML)

		require.Nil(t, err, "Error must be returned nil when parsed status application")

		require.Equal(t, expectStatus, stsAppUnmarshalXML, "The status parsed must be equal expected")
	}

	contentIN := []byte(`<StatusApp></StatusApp>`)

	err := xml.Unmarshal(contentIN, &stsAppUnmarshalXML)

	require.NotNil(t, err, "Error must be returned when parsed status application empty")

}
