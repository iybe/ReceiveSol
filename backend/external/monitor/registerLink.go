package monitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type registerLinkRequest struct {
	Db_id     string  `json:"db_id"`
	Reference string  `json:"reference"`
	Recipient string  `json:"recipient"`
	Amount    float32 `json:"amount"`
	Network   string  `json:"network"`
}

func (c *Client) RegisterLink(db_id, reference, recipient, network string, amount float32) error {
	registerLinkReq := registerLinkRequest{
		Db_id:     db_id,
		Reference: reference,
		Recipient: recipient,
		Amount:    amount,
		Network:   network,
	}

	body, err := json.Marshal(registerLinkReq)
	if err != nil {
		c.Log.Error(methodRegisterLink, err)
		return err
	}

	url := fmt.Sprintf("%s/link", c.URL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		c.Log.Error(methodRegisterLink, err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", c.Authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Log.Error(methodRegisterLink, err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		var errorResponse errorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			c.Log.Error(methodRegisterLink, err)
			return err
		}

		errR := fmt.Errorf("status code: %d, error: %s", resp.StatusCode, errorResponse.Error)
		c.Log.WarnWithStatusCode(methodRegisterLink, errR, resp.StatusCode)
		return errR
	}

	return nil
}
