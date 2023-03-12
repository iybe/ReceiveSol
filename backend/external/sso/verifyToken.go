package sso

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VerifyTokenRequest struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

func (c *Client) VerifyToken(user, token string) error {
	url := fmt.Sprintf("%s/token/verify", c.URL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Log.Error(methodVerifyToken, err)
		return err
	}

	req.Header.Set("authorization", c.Authorization)
	req.Header.Set("id", user)
	req.Header.Set("token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.Log.Error(methodVerifyToken, err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var errorResponse errorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			c.Log.Error(methodVerifyToken, err)
			return err
		}

		errR := fmt.Errorf("status code: %d, error: %s", resp.StatusCode, errorResponse.Error)
		c.Log.WarnWithStatusCode(methodVerifyToken, errR, resp.StatusCode)
		return errR
	}

	return nil
}
