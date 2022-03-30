package parser

import (
	"testing"
)

const (
	ID   = 10
	NAME = "Ryo"
)

type ReqStruct struct {
	ID   StringNumber `json:"id"`
	Name string       `json:"name"`
}

func TestRequestBodyToStruct(t *testing.T) {
	reqBody := `{"id": 10, "name": "Ryo"}`
	reqStruct := ReqStruct{}

	t.Run("Regular", func(t *testing.T) {
		err := RequestBodyToStruct(reqBody, &reqStruct)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		if reqStruct.ID != ID {
			t.Errorf("struct ID is not %d: %#v", ID, reqStruct.ID)
		}

		if reqStruct.Name != NAME {
			t.Errorf("struct Name is not %s: %#v", NAME, reqStruct.Name)
		}
	})
}

func TestQueryStringParametersToStruct(t *testing.T) {
	queryParams := map[string]string{
		"id":   "10",
		"name": "Ryo",
	}
	reqStruct := ReqStruct{}

	t.Run("Regular", func(t *testing.T) {
		err := QueryStringParametersToStruct(&queryParams, &reqStruct)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		if reqStruct.ID != ID {
			t.Errorf("struct ID is not %d: %#v", ID, reqStruct.ID)
		}

		if reqStruct.Name != NAME {
			t.Errorf("struct Name is not %s: %#v", NAME, reqStruct.Name)
		}
	})
}
