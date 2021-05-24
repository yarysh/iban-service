package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

type Case struct {
	Method string
	Path   string
	Query  string
	Status int
	Result interface{}
}
type CaseResponse map[string]interface{}

func doServerTest(srv *Server, cases []Case) error {
	server := httptest.NewServer(srv)

	for _, testCase := range cases {
		req, err := http.NewRequest(testCase.Method, server.URL+testCase.Path+"?"+testCase.Query, nil)
		if err != nil {
			return fmt.Errorf("request error: %v", err)
		}

		client := &http.Client{Timeout: time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("request error: %v", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != testCase.Status {
			return fmt.Errorf("expected http status %v, got %v", testCase.Status, resp.StatusCode)
		}

		var result interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return fmt.Errorf("can't unpack json: %v", err)
		}

		var expected interface{}
		data, err := json.Marshal(testCase.Result)
		json.Unmarshal(data, &expected)

		if !reflect.DeepEqual(result, expected) {
			return fmt.Errorf("results not match\nGot: %#v\nExpected: %#v", result, testCase.Result)
		}
	}
	return nil
}

func TestIBANServiceValidCases(t *testing.T) {
	err := doServerTest(NewIBANService(), []Case{
		{
			Path:   "/",
			Query:  "iban=AZ21NABZ00000000137010001944",
			Status: http.StatusOK,
			Result: CaseResponse{
				"iban":  "AZ21NABZ00000000137010001944",
				"valid": true,
			},
		},
		{
			Path:   "/",
			Query:  "iban=" + url.QueryEscape("XK05 1212 0123 4567 8906"),
			Status: http.StatusOK,
			Result: CaseResponse{
				"iban":  "XK051212012345678906",
				"valid": true,
			},
		},
		{
			Path:   "/",
			Query:  "iban=sc18SScB11010000000000001497usd",
			Status: http.StatusOK,
			Result: CaseResponse{
				"iban":  "SC18SSCB11010000000000001497USD",
				"valid": true,
			},
		},
	})
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestIBANServiceInvalidCases(t *testing.T) {
	err := doServerTest(NewIBANService(), []Case{
		{
			Path:   "/",
			Query:  "iban=AZ21NABZ10000000137010001944",
			Status: http.StatusOK,
			Result: CaseResponse{
				"iban":  "AZ21NABZ10000000137010001944",
				"valid": false,
			},
		},
		{
			Path:   "/",
			Query:  "iban=" + url.QueryEscape("XK05 1212 0133 4567 8906"),
			Status: http.StatusOK,
			Result: CaseResponse{
				"iban":  "XK051212013345678906",
				"valid": false,
			},
		},
		{
			Path:   "/",
			Query:  "iban=sc18SScB11010000000000001497us",
			Status: http.StatusOK,
			Result: CaseResponse{
				"iban":  "SC18SSCB11010000000000001497US",
				"valid": false,
			},
		},
	})
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestIBANServiceBadRequest(t *testing.T) {
	err := doServerTest(NewIBANService(), []Case{
		{
			Path:   "/",
			Query:  "iba=AZ21NABZ10000000137010001944",
			Status: http.StatusBadRequest,
			Result: CaseResponse{
				"error": "iban param required",
			},
		},
		{
			Path:   "/",
			Query:  "",
			Status: http.StatusBadRequest,
			Result: CaseResponse{
				"error": "iban param required",
			},
		},
	})
	if err != nil {
		t.Errorf("%s", err)
	}
}
