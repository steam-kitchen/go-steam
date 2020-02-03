package steam

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	token string
	h     *http.Client
}

func NewClient(token string, h *http.Client) *Client {
	if h == nil {
		h = http.DefaultClient
	}
	return &Client{
		token: token,
		h:     h,
	}
}

func (c *Client) SteamUser() *ISteamUser {
	i := new(ISteamUser)
	i.c = c
	i.ifname = "ISteamUser"
	return i
}

func (c *Client) Dota2Match() *IDota2Match {
	i := new(IDota2Match)
	i.c = c
	i.ifname = "IDOTA2Match_570"
	return i
}

func (c *Client) EconDota2() *IEconDota2 {
	i := new(IEconDota2)
	i.c = c
	i.ifname = "IEconDOTA2_570"
	return i
}

func (c *Client) get(u string, params url.Values, v interface{}) error {
	if params == nil {
		params = url.Values{}
	}
	params.Set("format", "json")
	params.Set("key", c.token)
	u = "https://api.steampowered.com/" + u + "?" + params.Encode()
	resp, err := c.h.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}
	if len(body) < 0xff {
		var r struct {
			Result struct {
				Error *string `json:"error,omitempty"`
			}
		}
		if err := json.Unmarshal(body, &r); err != nil {
			return err
		}
		if r.Result.Error != nil {
			return &Error{*r.Result.Error}
		}
	}
	return json.Unmarshal(body, v)
}
