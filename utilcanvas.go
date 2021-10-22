// Package utilcanvas provides various utilities for interfacing with the
// Canvas LMS API.
package utilcanvas

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Some API endpoints.
const (
	ListAnnouncementsEndpoint  = "/api/v1/announcements"
	ListCalendarEventsEndpoint = "/api/v1/calendar_events"
)

// NewClient returns a CanvasClient with the specified
// API URL and API key.
func NewClient(apiUrl, apiKey string) *CanvasClient {
	return &CanvasClient{ApiUrl: apiUrl, ApiKey: apiKey}
}

// CourseIdToContextCodeList creates a slice of ContextCode with
// the only element being a course, identified by the course ID parameter.
func CourseIdToContextCodeList(courseId string) *[]ContextCode {
	return &[]ContextCode{{Type: ContextCodeTypeCourse, Id: courseId}}
}

// coalesceToJson takes a ReadCloser and parses its JSON data
// into the specified interface.
func coalesceToJson(body io.ReadCloser, a interface{}) error {
	objectJson, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(objectJson), a)
	return nil
}

// Get performs a standard GET request for a given endpoint, with
// the specified parameters, and outputs the results into the specified struct.
func (c *CanvasClient) Get(endpoint string, apiParams ApiParameter, targetStruct interface{}) error {
	resp, err := c.Request("GET", endpoint, apiParams)
	if err != nil {
		return err
	}

	err = coalesceToJson(resp.Body, targetStruct)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Request performs an Canvas API request for a given endpoint
// with the specified method and parameters.
func (c *CanvasClient) Request(method, endpoint string, parameters ApiParameter) (*http.Response, error) {
	var respDefaulted *http.Response

	params, err := parameters.Process()
	if err != nil {
		return respDefaulted, err
	}
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest(method, c.ApiUrl+endpoint, body)
	if err != nil {
		return respDefaulted, err
	}
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return respDefaulted, err
	}

	return resp, nil
}
