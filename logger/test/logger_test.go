package test

import (
	"github.com/heyuan110/gorepertory/logger"
	"testing"
)


func TestLogger(t *testing.T)  {
	logger.Info("demo")
	logger.Info([]int{23,23,3})
	logger.Info([]string{"a","b","c"})
	logger.Info(map[string]string{"key1":"hello","key2":"today"})
}