package redis

import (
	"adiris/pkg/config"
	"adiris/tools/common"
	"time"
)

func SetKey(key string, data interface{}, ex int) bool {
	s, err := config.EngRds.Set(key, data, time.Duration(ex)).Result()

	if !commons.CheckErr(err, s) {
		return false
	}

	return true
}

func GetKey(key string) bool {
	s, err := config.EngRds.Get(key).Result()
	if !commons.CheckErr(err, s) {
		return false
	}

	return true
}

func DeleteKey(key string) bool {
	s, err := config.EngRds.Del(key).Result()
	if !commons.CheckErr(err, s) {
		return false
	}
	return true
}
