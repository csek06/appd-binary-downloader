package privlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

/*
ParseJSON will attempt to parse an io.Reader and place it on a passed interface{}
*/
func ParseJSON(jsonText io.Reader, thisStruct interface{}) {
	// parse JSON response to our AgentSearch Struct
	//var searchresults AgentSearch
	dec := json.NewDecoder(jsonText)
	dec.DisallowUnknownFields()
	err2 := dec.Decode(&thisStruct)
	if err2 != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err2, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			fmt.Println(msg + "\n" + err2.Error())

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err2, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			fmt.Println(msg + "\n" + err2.Error())
		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err2, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			fmt.Println(msg + "\n" + err2.Error())
		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err2.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err2.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			fmt.Println(msg + "\n" + err2.Error())
		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err2, io.EOF):
			msg := "Request body must not be empty"
			fmt.Println(msg + "\n" + err2.Error())
		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err2.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			fmt.Println(msg + "\n" + err2.Error())
		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			log.Println(err2.Error())
		}

	}
	err2 = dec.Decode(&struct{}{})
	if err2 != io.EOF {
		msg := "Request body must only contain a single JSON object"
		fmt.Println(msg)
	}
}
