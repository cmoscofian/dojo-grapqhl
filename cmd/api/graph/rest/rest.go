package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mercadolibre/fury_post-compra-dojo-graphql/cmd/api/graph/model"
)

const baseUrl string = "http://api.internal.ml.com"

func GetClaim(claimID int) (*model.Claim, error) {
	var claim model.Claim
	err := Get(fmt.Sprintf("/v1/claims/%d", claimID), &claim)
	return &claim, err
}

func Get(uri string, ref interface{}) error {
	url := baseUrl + uri
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Caller-Scopes", "admin")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ref)
	if err != nil {
		return err
	}

	return nil
}
