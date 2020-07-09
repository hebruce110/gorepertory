package strings

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func IntToString(val int) string  {
	return strconv.Itoa(val)
}

func Int64ToString(val int64) string  {
	return strconv.FormatInt(val,10)
}

func StringToInt(val string)int  {
	r,_ := strconv.Atoi(val)
	return r
}

func StringToInt64(val string)int64  {
	r,_ := strconv.ParseInt(val,10,64)
	return r
}

func EncodeMD5(value string) string  {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

