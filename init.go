package lddatabasekit

import (
	"errors"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
)

//var configfile string
//var configtype string
var (
	addrs        string
	database     string
	username     string
	password     string
	timeout      = time.Second * 5
	ConnectError error
)

var mErr error

//GetDataBaseName GetDataBaseName
func GetDataBaseName() string {
	return database
}

// Err get err
func Err() error {
	return mErr
}

func init() {
	if err := initConfig(); err != nil {
		mErr = err
		return
	}
}

func initConfig() error {

	// readfrom ini
	//configfile = getEnv("DB_CONFIGFILE", "conf/database.ini")
	//configtype = getEnv("DB_CONFIGTYPE", "ini")
	addrs = getEnv("DB_MONGO_ADDRS", "127.0.0.1")
	database = getEnv("DB_MONGO_DATABASE", "")
	username = getEnv("DB_MONGO_USERNAME", "")
	password = getEnv("DB_MONGO_PASSWORD", "")

	if t, err := time.ParseDuration(getEnv("DB_MONGO_TIMEOUT", "5s")); err == nil {
		timeout = t
	}

	//if addrs not exist or database name not exist,then read from ini file

	/*
		if addrs == "" || database == "" {
			if err := readFromConfigFile(); err != nil {
				return err
			}
		}

	*/

	//log.Println(addrs, database, username, password)
	if addrs == "" {
		return errors.New("addrs not define in env and configfile")
	} else if database == "" {
		return errors.New("database not define in env and configfile")
	} else if username == "" {
		return errors.New("username not define in env and configfile")
	} else if password == "" {
		return errors.New("password not define in env and configfile")
	}
	dialInfo := &mgo.DialInfo{
		Addrs:     strings.Split(addrs, ","),
		Direct:    false,
		Timeout:   timeout,
		Database:  database,
		Username:  username,
		Password:  password,
		PoolLimit: 4096,
	}
	if ConnectError = ReInitMongoDBSession(dialInfo, mgo.Strong); ConnectError != nil {
		//log.Println("链接主数据库失败", err)
		return ConnectError
	}
	return nil
}

func readFromConfigFile() error {
	// makesure path exist
	/*
		os.MkdirAll(path.Dir(configfile), 0777)
		iniconf, err := config.NewConfig(configtype, configfile)
		if err != nil {
			return errors.New("config file " + configfile + " read error")
		}
		addrs = iniconf.String("mongo::addrs")
		database = iniconf.String("mongo::database")
		username = iniconf.String("mongo::username")
		password = iniconf.String("mongo::password")
		t := iniconf.String("mongo::timeout")
		if t == "" {
			t = "5s"
		}
		if t, err := time.ParseDuration(t); err == nil {
			timeout = t
		}
		/*/
	errors.New("remove ini support")
	return nil
}

//Init Init
func Init() {

}
