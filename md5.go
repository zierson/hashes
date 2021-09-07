package strix

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func ToMD5(text string) string {
	hash := md5.Sum([]byte(text))

	return hex.EncodeToString(hash[:])
}

func RandomMD5() string {
	r := int(rand.Int63n(1000000))
	t := int(time.Now().UnixNano())

	return ToMD5(strconv.Itoa(r + t))
}
