package monitor

import (
	"adiris/pkg/config"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func RedisInfo() {
	info := config.EngRds.Info()

	m := DealRedisInfo(fmt.Sprintf("%s", info))
	c := config.EngMgo.DB("rules").C("RedisInfo")
	fmt.Println("insert into data")
	e := c.Insert(&m)
	fmt.Println(e)

}
func DealRedisInfo(a string) interface{} {
	ret := []string{}
	sData := strings.Split(strings.Replace(
		strings.Replace(a, "info:", "", 1),
		"\r", "", -1), "#")
	for _, _v := range sData {
		_d := strings.Split(_v, "\n")
		key := strings.Replace(_d[0], " ", "", -1)

		if len(key) < 2 {
			continue
		}
		s := []string{}
		for _, v := range _d[1:] {

			if strings.Contains(v, ":") {
				_s := strings.Split(v, ":")
				s = append(s, "\""+_s[0]+"\":\""+_s[1]+"\"")
			}
		}
		ret = append(ret, "\""+key+"\":{"+strings.Join(s, ",")+"}")
	}

	json_str := "{" + strings.Join(ret, ",") + "}"

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(json_str), &m)
	fmt.Println(err)
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	m["_id"] = timeStr
	return m

}
