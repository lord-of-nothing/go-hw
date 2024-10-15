package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Address string
}

func main() {
	address := "http://127.0.0.1:8080"
	cli := Client{address}
	CallVersion(&cli)
	CallDecode("Let's Go", &cli)
	CallHardOp(&cli)
}

func CallVersion(client *Client) {
	resp, err := http.Get(client.Address + "/version")
	if err != nil {
		fmt.Print("/version request failed:" + err.Error())
		return
	}
	respBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("/version request failed:" + err.Error())
		return
	}
	fmt.Println(string(respBytes))
}

func CallDecode(word string, client *Client) {
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
		fmt.Println("/decode request failed:" + err.Error())
		return
	}
	var result Decoded
	_ = json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result.Decoded)
}

func CallHardOp(client *Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", client.Address+"/hard-op", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Println(false)
		} else {
			fmt.Println("/hard-op request failed:" + err.Error())
		}
		return
	}
	fmt.Println(true, resp.StatusCode)
}
