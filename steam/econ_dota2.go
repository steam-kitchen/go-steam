package steam

import "net/url"

type IEconDota2 struct {
	c      *Client
	ifname string
}

type Dota2GameItems struct {
	Items []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Cost          int    `json:"cost"`
		SecretShop    int    `json:"secret_shop"`
		SideShop      int    `json:"side_shop"`
		Recipe        int    `json:"recipe"`
		LocalizedName string `json:"localized_name"`
	} `json:"items"`
}

func (i *IEconDota2) GetGameItems(language string) (*Dota2GameItems, error) {
	v := url.Values{}
	if language != "" {
		v.Set("language", language)
	}
	var r struct {
		Result *Dota2GameItems `json:"result,omitempty"`
	}
	if err := i.c.get(i.ifname+"/GetGameItems/v1/", v, &r); err != nil {
		return nil, err
	}
	return r.Result, nil
}

type Dota2Heroes struct {
	Heroes []struct {
		Name          string `json:"name"`
		ID            int    `json:"id"`
		LocalizedName string `json:"localized_name"`
	} `json:"heroes"`
	Count int `json:"count"`
}

func (i *IEconDota2) GetHeroes(language string) (*Dota2Heroes, error) {
	v := url.Values{}
	if language != "" {
		v.Set("language", language)
	}
	var r struct {
		Result *Dota2Heroes `json:"result,omitempty"`
	}
	if err := i.c.get(i.ifname+"/GetHeroes/v1/", v, &r); err != nil {
		return nil, err
	}
	return r.Result, nil
}
