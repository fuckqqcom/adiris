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
	s := strings.Split(strings.Replace(fmt.Sprintf("%s", a), "info: ", "", 1), "#")
	demo(s)

	if err != nil {
		panic(err)
	}

}

func demo(data []string) {
	s := " {"
	for _, value := range data {
		_t := strings.Split(strings.Replace(value, "\n", "", -1), "\r")
		_s := strings.TrimSpace(_t[0])
		if len(_s) > 1 {
			s += " \"" + _s + "\":{"
			fmt.Println(_t[1:])
			for i, _v := range _t[1:] {
				org := strings.Split(_v, ":")
				if len(org) > 1 {
					if len(_t[1:]) == (i) {
						s += "\"" + strings.Trim(org[0], "\n") + "\":\"" + org[1] + "\""
					} else {
						s += "\"" + strings.Trim(org[0], "\n") + "\":\"" + org[1] + "\","
					}
				}
			}
			s += "},"
		}
	}
	fmt.Println(s + "}")
	//
	var f map[string]interface{}
	err := json.Unmarshal([]byte(s), &f)
	fmt.Println(err)
	fmt.Println(f)

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
