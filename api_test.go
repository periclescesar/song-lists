package main

import (
	"GoSOLID/pkg"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
	t    *testing.T
}

func (a *apiFeature) resetResponse(*messages.Pickle) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendRequestTo(method, endpoint string, body io.Reader) (err error) {
	r := pkg.SetupRouter()

	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return
	}
	r.ServeHTTP(a.resp, req)

	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	return
}

func (a *apiFeature) theResponseCodeShouldBe(c int) error {
	return assertExpectedAndActual(assert.Equal, c, a.resp.Code)
}

func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) error {
	var expected, actual []byte
	var data interface{}

	if err := json.Unmarshal([]byte(body.Content), &data); err != nil {
		return err
	}
	expected, err := json.Marshal(data)
	if err != nil {
		return err
	}
	actual = a.resp.Body.Bytes()

	if !bytes.Equal(actual, expected) {
		return fmt.Errorf("expected json, does not match actual: %s", string(actual))
	}

	return nil
}

func (a *apiFeature) theBodyShouldIsJson(method, endpoint string, body *godog.DocString) error {
	return a.iSendRequestTo(method, endpoint, strings.NewReader(body.Content))
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}
	ctx.BeforeScenario(api.resetResponse)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, api.iSendRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, api.theResponseShouldMatchJSON)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)" with json:$`, api.theBodyShouldIsJson)
}
