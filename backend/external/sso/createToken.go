package sso

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type createTokenRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type createTokenResponse struct {
	Token string `json:"token"`
}

func (c *Client) CreateToken(user, password string) (*createTokenResponse, error) {
	createTokenReq := createTokenRequest{
		User:     user,
		Password: password,
	}

	body, err := json.Marshal(createTokenReq)
	if err != nil {
		c.Log.Error(methodCreateToken, err)
		return nil, err
	}

	url := fmt.Sprintf("%s/token", c.URL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		c.Log.Error(methodCreateToken, err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", c.Authorization)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.Log.Error(methodCreateToken, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		var errorResponse errorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			c.Log.Error(methodCreateToken, err)
			return nil, err
		}

		errR := fmt.Errorf("status code: %d, error: %s", resp.StatusCode, errorResponse.Error)
		c.Log.WarnWithStatusCode(methodCreateToken, errR, resp.StatusCode)
		return nil, errR
	}

	var createTokenResponse createTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&createTokenResponse)
	if err != nil {
		c.Log.Error(methodCreateToken, err)
		return nil, err
	}

	return &createTokenResponse, nil
}
