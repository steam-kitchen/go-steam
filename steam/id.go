package steam

import "strconv"

const idMagicNumber uint64 = 76561197960265728

const idAnonymouse = SteamID3(4294967295)

type SteamID interface {
	SteamID() uint64
	FriendID() uint32
	SteamID64() SteamID64
	SteamID3() SteamID3
}

type SteamID64 uint64

func (s SteamID64) SteamID() uint64 {
	return uint64(s)
}

func (s SteamID64) FriendID() uint32 {
	return uint32(uint64(s) - idMagicNumber)
}

func (s SteamID64) SteamID64() SteamID64 {
	return s
}

func (s SteamID64) SteamID3() SteamID3 {
	return SteamID3(s.FriendID())
}

func (s SteamID64) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

type SteamID3 uint32

func (s SteamID3) SteamID() uint64 {
	return uint64(s) + idMagicNumber
}

func (s SteamID3) FriendID() uint32 {
	return uint32(s)
}

func (s SteamID3) SteamID64() SteamID64 {
	return SteamID64(s.SteamID())
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
