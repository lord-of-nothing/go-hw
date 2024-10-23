package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Address string
}

func (client Client) GetVersion() (string, error) {
	resp, err := http.Get(client.Address + "/version")
	if err != nil {
		return "", err
	}
	respBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	//fmt.Println(string(respBytes))
	return string(respBytes), nil
}

func (client Client) Decode(word string) (string, error) {
	type Encoded struct {
		Encoded string `json:"encoded"`
	}

	type Decoded struct {
		Decoded string `json:"decoded"`
	}

	wordEncoded := Encoded{base64.StdEncoding.EncodeToString([]byte(word))}
	body, _ := json.Marshal(wordEncoded)
	resp, err := http.Post(client.Address+"/decode", "application/json", bytes.NewBuffer(body))
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	var result Decoded
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.Decoded, nil
}

func (client Client) CallHardOp() (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", client.Address+"/hard-op", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return false, -1, nil
		} else {
			return false, -1, err
		}
	}
	return true, resp.StatusCode, nil
}
