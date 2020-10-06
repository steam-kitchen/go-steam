package steam

import "strconv"

const idMagicNumber int64 = 76561197960265728

const idAnonymouse = SteamID3(4294967295)

type SteamID interface {
	SteamID64() SteamID64
	SteamID3() SteamID3
}

type SteamID64 int64

func (s SteamID64) SteamID64() SteamID64 {
	return s
}

func (s SteamID64) SteamID3() SteamID3 {
	return SteamID3(int64(s) - idMagicNumber)
}

func (s SteamID64) String() string {
	return strconv.FormatInt(int64(s), 10)
}

type SteamID3 uint32

func (s SteamID3) SteamID64() SteamID64 {
	return SteamID64(int64(s) + idMagicNumber)
}

func (s SteamID3) SteamID3() SteamID3 {
	return s
}

func (s SteamID3) String() string {
	return strconv.Itoa(int(s))
}

func (s SteamID3) IsAnonymous() bool {
	return s == idAnonymouse
}
