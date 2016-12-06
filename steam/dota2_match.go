package steam

import (
	"errors"
	"net/url"
	"strconv"
)

type IDota2Match struct {
	c      *Client
	ifname string
}

type Dota2MatchDetails struct {
	Players []struct {
		AccountID         SteamID3 `json:"account_id"`
		PlayerSlot        int      `json:"player_slot"`
		HeroID            int      `json:"hero_id"`
		Item0             int      `json:"item_0"`
		Item1             int      `json:"item_1"`
		Item2             int      `json:"item_2"`
		Item3             int      `json:"item_3"`
		Item4             int      `json:"item_4"`
		Item5             int      `json:"item_5"`
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
		AbilityUpgrades   []struct {
			Ability int `json:"ability"`
			Time    int `json:"time"`
			Level   int `json:"level"`
		} `json:"ability_upgrades"`
		AdditionalUnits []struct {
			UnitName string `json:"unitname"`
			Item0    int    `json:"item_0"`
			Item1    int    `json:"item_1"`
			Item2    int    `json:"item_2"`
			Item3    int    `json:"item_3"`
			Item4    int    `json:"item_4"`
			Item5    int    `json:"item_5"`
		} `json:"additional_units,omitempty"`
	} `json:"players"`
	RadiantWin            bool   `json:"radiant_win"`
	Duration              int    `json:"duration"`
	PreGameDuration       int    `json:"pre_game_duration"`
	StartTime             int    `json:"start_time"`
	MatchID               uint64 `json:"match_id"`
	MatchSeqNum           uint64 `json:"match_seq_num"`
	TowerStatusRadiant    int    `json:"tower_status_radiant"`
	TowerStatusDire       int    `json:"tower_status_dire"`
	BarracksStatusRadiant int    `json:"barracks_status_radiant"`
	BarracksStatusDire    int    `json:"barracks_status_dire"`
	Cluster               int    `json:"cluster"`
	FirstBloodTime        int    `json:"first_blood_time"`
	LobbyType             int    `json:"lobby_type"`
	HumanPlayers          int    `json:"human_players"`
	LeagueID              uint32 `json:"leagueid"`
	PositiveVotes         int    `json:"positive_votes"`
	NegativeVotes         int    `json:"negative_votes"`
	GameMode              int    `json:"game_mode"`
	Flags                 int    `json:"flags"`
	Engine                int    `json:"engine"`
	RadiantScore          int    `json:"radiant_score"`
	DireScore             int    `json:"dire_score"`
}

func (i *IDota2Match) GetMatchDetails(id uint64) (*Dota2MatchDetails, error) {
	v := url.Values{"match_id": {strconv.FormatUint(id, 10)}}
	var r struct {
		Result *Dota2MatchDetails `json:"result,omitempty"`
		Error  *string            `json:"error,omitempty"`
	}
	if err := i.c.get(i.ifname+"/GetMatchDetails/v1/", v, &r); err != nil {
		return nil, err
	}
	if r.Error != nil {
		return nil, errors.New(*r.Error)
	}
	return r.Result, nil
}
