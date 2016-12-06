package steam

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (c *Client) Dota2Match(id ...int) *IDota2Match {
	i := new(IDota2Match)
	ids := 570
	i.c = c
	if len(id) > 0 {
		ids = id[0]
	}
	i.ifname = fmt.Sprintf("IDOTA2Match_%v", ids)
	return i
}

func (c *Client) EconDota2(id ...int) *IEconDota2 {
	i := new(IEconDota2)
	ids := 570
	i.c = c
	if len(id) > 0 {
		ids = id[0]
	}
	i.ifname = fmt.Sprintf("IEconDOTA2_%v", ids)
	return i
}

func (c *Client) get(u string, params url.Values, v interface{}) error {
	if params == nil {
		params = url.Values{}
	}
	params.Set("format", "json")
	params.Set("key", c.token)
	u = "https://api.steampowered.com/" + u + "?" + params.Encode()
	resp, err := http.Get(u)
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
	return json.Unmarshal(body, v)
}
