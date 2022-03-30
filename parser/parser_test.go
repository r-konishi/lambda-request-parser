package parser

import (
	"fmt"
	"testing"
)

const (
	ID   = 1000
	NAME = "Ryo"
)

type ReqStruct struct {
	ID   StringNumber `json:"id" validate:"required,min=1,max=999"`
	Name string       `json:"name" validate:"required"`
}

func TestRequestBodyToStruct(t *testing.T) {
	reqBody := fmt.Sprintf(`{"id": %d, "name": "%s"}`, ID, NAME)
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
		"id":   fmt.Sprintf("%d", ID),
		"name": fmt.Sprintf("%s", NAME),
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

func TestGetValidationErrors(t *testing.T) {
	t.Run("Regular", func(t *testing.T) {
		reqStruct := ReqStruct{
			ID:   100,
			Name: fmt.Sprintf("%s", NAME),
		}
		err := GetValidationErrors(&reqStruct)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})

	t.Run("Error Check", func(t *testing.T) {
		reqStruct := ReqStruct{
			ID: 1000,
		}
		err := GetValidationErrors(&reqStruct)
		for _, e := range err {
			fmt.Println("========== start ==========")
			fmt.Printf("Field:\t%s\n", e.Field())
			fmt.Printf("Type:\t%s\n", e.Type())
			fmt.Printf("Tag:\t%s\n", e.Tag())
			fmt.Printf("Param:\t%v\n", e.Param())
			fmt.Printf("Value:\t%#v\n", e.Value())
			fmt.Printf("Error:\t%v\n", e)
			fmt.Println("========== end ==========")
		}
		if err != nil && len(err) != 2 {
			t.Errorf("error: %v", err)
		}
	})
}
