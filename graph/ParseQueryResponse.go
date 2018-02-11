package graph

import (
  "encoding/json"
)

func ParseQueryResponse(text string) (QueryResponse, error) {
	qr := QueryResponse{}
	err := json.Unmarshal([]byte(text), &qr)
	if err != nil {
		return qr, err
	}
	return qr, nil
}
