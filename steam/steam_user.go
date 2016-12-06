package steam

import (
	"net/url"
	"strings"
)

type ISteamUser struct {
	c      *Client
	ifname string
}

type PlayerSummaries struct {
	Players []PlayerSummary `json:"players"`
}

type PlayerSummary struct {
	SteamID                  SteamID64 `json:"steamid,string"`
	CommunityVisibilityState int       `json:"communityvisibilitystate"`
	ProfileState             int       `json:"profilestate"`
	PersonaName              string    `json:"personaname"`
	LastLogOff               Timestamp `json:"lastlogoff"`
	ProfileURL               string    `json:"profileurl"`
	Avatar                   string    `json:"avatar"`
	AvatarMedium             string    `json:"avatarmedium"`
	AvatarFull               string    `json:"avatarfull"`
	PersonaState             int       `json:"personastate"`
	RealName                 string    `json:"realname"`
	PrimaryClanID            string    `json:"primaryclanid"`
	TimeCreated              Timestamp `json:"timecreated"`
	PersonaStateFlags        int       `json:"personastateflags"`
	LocCountryCode           string    `json:"loccountrycode"`
	LocStateCode             string    `json:"locstatecode"`
	LocCityID                int       `json:"loccityid"`
}

func (i *ISteamUser) GetPlayerSummaries(steamids []SteamID) (*PlayerSummaries, error) {
	ids := make([]string, len(steamids))
	for i, id := range steamids {
		ids[i] = id.SteamID64().String()
	}
	v := url.Values{
		"steamids": {strings.Join(ids, ",")},
	}
	var r struct {
		Response PlayerSummaries `json:"response"`
	}
	if err := i.c.get(i.ifname+"/GetPlayerSummaries/v2/", v, &r); err != nil {
		return nil, err
	}
	return &r.Response, nil
}
