package sso

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type createUserRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type createUserResponse struct {
	Id string `json:"id"`
}

func (c *Client) CreateUser(user, password string) (*createUserResponse, error) {
	createUserReq := createUserRequest{
		User:     user,
		Password: password,
	}

	body, err := json.Marshal(createUserReq)
	if err != nil {
		c.Log.Error("CreateUser", err)
		return nil, err
	}

	url := fmt.Sprintf("%s/user", c.URL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		c.Log.Error("CreateUser", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", c.Authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Log.Error("CreateUser", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		var errorResponse errorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			c.Log.Error("CreateUser", err)
			return nil, err
		}

		errR := fmt.Errorf("status code: %d, error: %s", resp.StatusCode, errorResponse.Error)
		c.Log.WarnWithStatusCode("CreateUser", errR, resp.StatusCode)
		return nil, errR
	}

	var createUserResp createUserResponse
	err = json.NewDecoder(resp.Body).Decode(&createUserResp)
	if err != nil {
		c.Log.Error("CreateUser", err)
		return nil, err
	}

	return &createUserResp, nil
}
