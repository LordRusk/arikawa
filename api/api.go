package api

import (
	"net/http"

	"git.sr.ht/~diamondburned/arikawa/httputil"
)

const (
	BaseEndpoint = "https://discordapp.com/api"
	APIVersion   = "6"

	Endpoint           = BaseEndpoint + "/v" + APIVersion + "/"
	EndpointUsers      = Endpoint + "users/"
	EndpointGateway    = Endpoint + "gateway"
	EndpointGatewayBot = EndpointGateway + "/bot"
	EndpointWebhooks   = Endpoint + "webhooks/"
)

var UserAgent = "DiscordBot (https://github.com/diamondburned/arikawa, v0.0.1)"

type Client struct {
	httputil.Client
	Token string
}

func NewClient(token string) *Client {
	cli := &Client{
		Client: httputil.NewClient(),
		Token:  token,
	}

	tw := httputil.NewTransportWrapper()
	tw.Pre = func(r *http.Request) error {
		if r.Header.Get("Authorization") == "" {
			r.Header.Set("Authorization", cli.Token)
		}

		if r.UserAgent() == "" {
			r.Header.Set("User-Agent", UserAgent)
		}

		return nil
	}

	cli.Client.Transport = tw

	return cli
}
