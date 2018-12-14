package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"github.com/xormplus/xorm"
	"strings"
	"time"
)

type Config struct {
	DB struct {
		Host    string
		User    string
		Pwd     string
		Db      string
		Show    bool
		Port    int
		MaxOpen int
		MaxIdle int
	}
	Redis struct {
		Dns      string
		MinIdle  int
		PoolSize int
	}

	Mgo struct {
		Dns string
	}
}

var (
	EngMgo *mongo.Client
	EngDb  *xorm.Engine
	EngRds *redis.Client
)

func InitConfig(path string) {
	config := Config{}
	Load(path, &config)
	config.loadDb()
	config.loadRedis()
	//config.loadMgo()
}

func (c *Config) loadDb() {
	fmt.Println(c.DB)
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DB.User,
		c.DB.Pwd,
		c.DB.Host,
		c.DB.Db)

	EngDb, err = xorm.NewEngine("mysql", dns)

	ping := EngDb.Ping()
	if ping != nil || err != nil {
		panic(ping)
	}
	EngDb.SetMaxIdleConns(c.DB.MaxIdle)
	EngDb.SetMaxOpenConns(c.DB.MaxOpen)
	EngDb.ShowSQL(c.DB.Show)

}

func (c Config) loadRedis() {

	EngRds = redis.NewClient(&redis.Options{
		Addr:         c.Redis.Dns,
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     c.Redis.PoolSize,
		MinIdleConns: c.Redis.MinIdle,
	})

	_, err := EngRds.Ping().Result()
	type Cluster struct {
		ClusterEnabled int
	}

	a := EngRds.Info()
	demo(fmt.Sprintf("%s", a))
	//s := strings.Replace(
	//	strings.Replace(fmt.Sprintf("%s", a), "info:", "", 1),
	//	"\r", "", -1)
	//sData := strings.Split(s, "#")
	//for _, _v := range sData {
	//	_d := strings.Split(_v, "\n")
	//	//fmt.Println(_k,_d)
	//	demo(_d)
	//}
	//demo(sData)

	if err != nil {
		panic(err)
	}

}

func demo(a string) {
	ret := []string{}
	s := strings.Replace(
		strings.Replace(a, "info:", "", 1),
		"\r", "", -1)
	sData := strings.Split(s, "#")
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
	fmt.Println(m)
}

func (c *Config) loadMgo() {
	var err error
	fmt.Println(c.Mgo.Dns)
	EngMgo, err = mongo.NewClientWithOptions(c.Mgo.Dns)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = EngMgo.Ping(ctx, readpref.Primary())

	if err != nil {
		panic(err)
	}
}
