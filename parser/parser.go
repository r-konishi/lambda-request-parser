package parser

import "encoding/json"

type StringNumber int64

func (s *StringNumber) UnmarshalJSON(b []byte) error {
	var number json.Number
	if err := json.Unmarshal(b, &number); err != nil {
		return err
	}
	i, err := number.Int64()
	if err != nil {
		return err
	}
	*s = StringNumber(i)
	return nil
}

func RequestBodyToStruct(body string, v interface{}) error {
	err := json.Unmarshal([]byte(body), v)
	return err
}

func QueryStringParametersToStruct(params *map[string]string, v interface{}) error {
	jsonStr, err := json.Marshal(&params)
	if err != nil {
		return nil
	}
	return RequestBodyToStruct(string(jsonStr), &v)
}
