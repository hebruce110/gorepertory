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
