package main

import (
    "strconv"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
    tests := map[string]struct {
        input  string
        output int
        err    error
    }{
        "successful conversion": {
            input:  "1",
            output: 1,
           err:    nil,
    },
        "invalid integer": {
            input:  "not an integer",
            output: 0,
            err:    &strconv.NumError{},
        },
    }
    for testName, test := range tests {
        t.Logf("Running test case %s", testName)
        output, err := strconv.Atoi(test.input)
        assert.IsType(t, test.err, err)
        assert.Equal(t, test.output, output)
    }
}

func TestExample(t *testing.T) {
    recorder := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "/example", nil)
    assert.Nil(t, err)
    handler := http.HandlerFunc(ExampleHandler)
    handler.ServeHTTP()
    goldie.Assert(t, "example", recorder.Body.Bytes())
}