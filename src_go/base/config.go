package base

import (
	"../utils"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

// 数据库配置
type DataBaseConfig struct {
	Dialect  string `json:"dialect"`
	User     string `json:"user"`
	Dbname   string `json:"dbname"`
	Password string `json:"password"`
	Sslmode  string `json:"sslmode"`
	Init     bool   `json:"init"`
}

// 网络配置
type WebConfig struct {
	Port string `json:"port"`
}

// 系统配置信息
type Config struct {
	File string         `json:"-"`
	Db   DataBaseConfig `json:"db"`
	Web  WebConfig      `json:"web"`
}

// 加密数据库信息
func (db *DataBaseConfig) encode() bool {
	old := db.Password
	var err error
	db.Password, err = utils.Decrypt(db.Password, "xuender@gmail.com")
	if err != nil {
		db.Password, _ = utils.Encrypt(old, "xuender@gmail.com")
		return true
	}
	return false
}

// 数据库连接
func (db *DataBaseConfig) GetSource() string {
	return fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", db.User, db.Dbname, db.Password, db.Sslmode)
}

// 读取配置文件
func (config *Config) Read(file string) error {
	config.File = file
	log.WithFields(log.Fields{
		"File": file,
	}).Debug("读取配置文件")
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("读取文件错误")
		return err
	}
	if err := json.Unmarshal(bytes, config); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("解析JSON错误")
		return err
	}
	if config.Db.encode() {
		config.Save()
	}
	log.WithFields(log.Fields{
		"db":  config.Db,
		"web": config.Web,
	}).Debug("config")
	return nil
}

// 保存配置文件
func (config *Config) Save() error {
	log.WithFields(log.Fields{
		"db": config.Db,
	}).Debug("save")
	bs, err := json.MarshalIndent(config, " ", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(config.File, bs, 0600)
}

var BaseConfig Config
