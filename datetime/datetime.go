package datetime

import "time"

func now() time.Time  {
	return time.Now()
}

func nowString()string  {
	return time.Now().Format("2006-01-02 15:04:05")
}