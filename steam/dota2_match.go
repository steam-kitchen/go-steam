package steam

import (
	"net/url"
	"strconv"
)

type IDota2Match struct {
	c      *Client
	ifname string
}

type Dota2MatchDetails struct {
	Players               []Dota2MatchDetailsPlayer `json:"players"`
	RadiantWin            bool                      `json:"radiant_win"`
	Duration              int                       `json:"duration"`
	PreGameDuration       int                       `json:"pre_game_duration"`
	StartTime             int                       `json:"start_time"`
	MatchID               int64                     `json:"match_id"`
	MatchSeqNum           int64                     `json:"match_seq_num"`
	TowerStatusRadiant    uint16                    `json:"tower_status_radiant"`
	TowerStatusDire       uint16                    `json:"tower_status_dire"`
	BarracksStatusRadiant uint8                     `json:"barracks_status_radiant"`
	BarracksStatusDire    uint8                     `json:"barracks_status_dire"`
	Cluster               int                       `json:"cluster"`
	FirstBloodTime        int                       `json:"first_blood_time"`
	LobbyType             int                       `json:"lobby_type"`
	HumanPlayers          int                       `json:"human_players"`
	LeagueID              int32                     `json:"leagueid"`
	PositiveVotes         int                       `json:"positive_votes"`
	NegativeVotes         int                       `json:"negative_votes"`
	GameMode              int                       `json:"game_mode"`
	Flags                 int                       `json:"flags"`
	Engine                int                       `json:"engine"`
	RadiantScore          int                       `json:"radiant_score"`
	DireScore             int                       `json:"dire_score"`
	PicksBans             []struct {
		IsPick bool `json:"is_pick"`
		HeroID int  `json:"hero_id"`
		Team   int  `json:"team"`
		Order  int  `json:"order"`
	} `json:"picks_bans"`
}

type Dota2MatchDetailsPlayer struct {
	AccountID         SteamID3 `json:"account_id"`
	PlayerSlot        byte     `json:"player_slot"`
	HeroID            int      `json:"hero_id"`
	Item0             int      `json:"item_0"`
	Item1             int      `json:"item_1"`
	Item2             int      `json:"item_2"`
	Item3             int      `json:"item_3"`
	Item4             int      `json:"item_4"`
	Item5             int      `json:"item_5"`
	Backpack0         int      `json:"backpack_0"`
	Backpack1         int      `json:"backpack_1"`
	Backpack2         int      `json:"backpack_2"`
	ItemNeutral       int      `json:"item_neutral"`
	Kills             int      `json:"kills"`
	Deaths            int      `json:"deaths"`
	Assists           int      `json:"assists"`
	LeaverStatus      int      `json:"leaver_status"`
	LastHits          int      `json:"last_hits"`
	Denies            int      `json:"denies"`
	GoldPerMin        int      `json:"gold_per_min"`
	XpPerMin          int      `json:"xp_per_min"`
	Level             int      `json:"level"`
	HeroDamage        int      `json:"hero_damage"`
	TowerDamage       int      `json:"tower_damage"`
	HeroHealing       int      `json:"hero_healing"`
	Gold              int      `json:"gold"`
	GoldSpent         int      `json:"gold_spent"`
	ScaledHeroDamage  int      `json:"scaled_hero_damage"`
	ScaledTowerDamage int      `json:"scaled_tower_damage"`
	ScaledHeroHealing int      `json:"scaled_hero_healing"`
	Persona           string   `json:"persona"`
	AbilityUpgrades   []struct {
		Ability int `json:"ability"`
		Time    int `json:"time"`
		Level   int `json:"level"`
	} `json:"ability_upgrades"`
	AdditionalUnits []struct {
		UnitName    string `json:"unitname"`
		Item0       int    `json:"item_0"`
		Item1       int    `json:"item_1"`
		Item2       int    `json:"item_2"`
		Item3       int    `json:"item_3"`
		Item4       int    `json:"item_4"`
		Item5       int    `json:"item_5"`
		Backpack0   int    `json:"backpack_0"`
		Backpack1   int    `json:"backpack_1"`
		Backpack2   int    `json:"backpack_2"`
		ItemNeutral int    `json:"item_neutral"`
	} `json:"additional_units,omitempty"`
}

type Dota2MatchHistory struct {
	Status           int `json:"status"`
	NumResults       int `json:"num_results"`
	TotalResults     int `json:"total_results"`
	ResultsRemaining int `json:"results_remaining"`
	Matches          []struct {
		MatchID       int64 `json:"match_id"`
		MatchSeqNum   int64 `json:"match_seq_num"`
		StartTime     int   `json:"start_time"`
		LobbyType     int   `json:"lobby_type"`
		RadiantTeamID uint  `json:"radiant_team_id"`
		DireTeamID    uint  `json:"dire_team_id"`
		Players       []struct {
			AccountID  SteamID3 `json:"account_id"`
			PlayerSlot int      `json:"player_slot"`
			HeroID     int      `json:"hero_id"`
		} `json:"players"`
	} `json:"matches"`
}

func (i *IDota2Match) GetMatchDetails(id int64, includePersonaNames bool) (*Dota2MatchDetails, error) {
	v := url.Values{
		"match_id":              {strconv.FormatInt(id, 10)},
		"include_persona_names": {"1"},
	}
	var r struct {
		Result *Dota2MatchDetails `json:"result"`
	}
	if err := i.c.get(i.ifname+"/GetMatchDetails/v1/", v, &r); err != nil {
		return nil, err
	}
	return r.Result, nil
}

type Dota2MatchHistoryParams url.Values

func (p Dota2MatchHistoryParams) SetStartAtMatchID(matchID int64) {
	url.Values(p).Set("start_at_match_id", strconv.FormatInt(matchID, 10))
}
func (p Dota2MatchHistoryParams) SetMatchesRequested(length string) {
	url.Values(p).Set("start_at_match_id", length)
}

func (i *IDota2Match) GetMatchHistory(params Dota2MatchHistoryParams) (*Dota2MatchHistory, error) {

	var r struct {
		Result *Dota2MatchHistory `json:"result"`
	}
	if err := i.c.get(i.ifname+"/GetMatchHistory/v1/", url.Values(params), &r); err != nil {
		return nil, err
	}
	return r.Result, nil
}
