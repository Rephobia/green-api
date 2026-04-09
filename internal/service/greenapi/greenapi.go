package greenapi

import (
	"errors"
	"fmt"
	"net/http"

	greenapi "github.com/green-api/whatsapp-api-client-golang-v2"
)

var (
	ErrEmptyResponse = errors.New("empty response")
	ErrAPI           = errors.New("greenapi error")
)

type Client struct {
	api greenapi.GreenAPI
}

func New(id, token string) *Client {
	return &Client{
		api: greenapi.GreenAPI{
			APIURL:           "https://api.green-api.com",
			MediaURL:         "https://media.green-api.com",
			IDInstance:       id,
			APITokenInstance: token,
		},
	}
}

func (c *Client) GetSettings() ([]byte, error) {
	r, err := c.api.Account().GetSettings()

	return c.checkResponseErrors(r, err)
}

func (c *Client) GetStateInstance() ([]byte, error) {
	r, err := c.api.Account().GetStateInstance()

	return c.checkResponseErrors(r, err)
}

func (c *Client) SendMessage(chatID, message string) ([]byte, error) {
	r, err := c.api.Sending().SendMessage(chatID, message)

	return c.checkResponseErrors(r, err)
}

func (c *Client) SendFileByUrl(chatID, urlFile, fileName, caption string) ([]byte, error) {
	r, err := c.api.Sending().SendFileByUrl(
		chatID,
		urlFile,
		fileName,
		greenapi.OptionalCaptionSendUrl(caption),
	)

	return c.checkResponseErrors(r, err)
}

func (c *Client) checkResponseErrors(response *greenapi.APIResponse, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, ErrEmptyResponse
	}

	if response.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("%w: %s", ErrAPI, string(response.StatusMessage))
	}

	return response.Body, nil
}
