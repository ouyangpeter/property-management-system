package lib

import (
    "crypto/md5"
    "crypto/sha256"
    "encoding/hex"
    "strconv"
)

func Pwdhash(str string) string {
    salt := "pms_#880"
    return str2Sha256(salt + str)
}

func str2Sha256(str string) string {
    //log.Println(str + "hehe")
    sha256Hash := sha256.New()
    sha256Hash.Write([]byte(str))
    ret := hex.EncodeToString(sha256Hash.Sum(nil))
    //log.Println(ret)
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
