package lib

import (
    "crypto/sha256"
    "fmt"
    "strconv"
	"crypto/md5"
	"encoding/hex"
)

func Pwdhash(str string) string {
    //return Str2Sha256(str)
	return Strtomd5(str)
}

func Str2Sha256(str string) string {
    sha256Hash := sha256.New()
    ret := fmt.Sprintf("%x", sha256Hash.Sum([]byte(str)))
    return ret
}

//create md5 string
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}
