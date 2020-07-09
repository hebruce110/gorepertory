package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMd5(t *testing.T)  {
	md5 := EncodeMD5(time.Now().Format("2006-01-02 15:04:05"))
	assert.NotNil(t,md5)
	assert.Len(t,md5,32)
}

func TestEncodeToJson(t *testing.T)  {
	var a map[string]interface{}
	a=map[string]interface{}{
		"k1":"hellow orld!",
		"k2":2332,
		"k3":false,
		"k4":nil,
		"k5":[]string{"f","i","j"},
		"k6":map[string]interface{}{"foo":"bar"},
	}
	s:=JsonEncode(a)
	assert.NotNil(t,s)
}