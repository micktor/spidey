package api

import (
	"context"
	"encoding/json"
)

type Spidey struct {
	Test string
}

func (sc Client) GetSpidey(ctx context.Context, id int) Spidey {
	response, err := sc.client.Get(sc.baseURL.String() + "/api/v1/spidey/get")
	if err != nil {
		return Spidey{}
	}
	defer response.Body.Close()

	spidey := Spidey{}
	if err := json.NewDecoder(response.Body).Decode(&spidey); err != nil {
		return Spidey{}
	}
	return Spidey{Test: "123"}
}
