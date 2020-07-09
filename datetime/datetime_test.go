package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNowString(t *testing.T)  {
	o := NowString()
	nowstring := time.Now().Format("2006-01-02 15:04:05")
	assert.Equal(t,o,nowstring)
}
