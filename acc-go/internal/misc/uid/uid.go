package uid

import (
	"github.com/speps/go-hashids/v2"
	"strconv"
)

const (
	alphabet  = "0123456789abcdef"
	salt      = "acc@2022#com"
	base      = 16
	minLength = 3
	bitSize   = 64
)

var hashID *hashids.HashID

func init() {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	hd.Alphabet = alphabet

	hashID, _ = hashids.NewWithData(hd)
}

func Id2Uid(id int64) int64 {
	hash, _ := hashID.EncodeInt64([]int64{id})
	uid, _ := strconv.ParseInt(hash, base, bitSize)
	return uid
}

func Uid2Id(uid int64) int64 {
	hex := strconv.FormatInt(uid, base)
	id, _ := hashID.DecodeInt64WithError(hex)
	return id[0]
}
