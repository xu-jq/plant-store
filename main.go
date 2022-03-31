package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"

	_ "bee_project/routers"
)

//initMySQL 初始化mysql
func initMySQL() error {
	// 注册MySQL驱动
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		return err
	}

	type MySqlConf struct {
		Username string
		Password string
		Uri      string
	}
	var conf MySqlConf
	var err error

	if conf.Username, err = web.AppConfig.String("mysql::user"); err != nil {
		return err
	}

	if conf.Password, err = web.AppConfig.String("mysql::password"); err != nil {
		return err
	}

	if conf.Uri, err = web.AppConfig.String("mysql::uri"); err != nil {
		return err
	}

	if err := orm.RegisterDataBase("default", "mysql", "root:xu123..123@/plant?charset=utf8"); err != nil {
		return err
	}

	return nil
}

//initLogger 初始化日志配置
func initLogger() error {

	type LogConf struct {
		Filename string `json:"filename,omitempty"`
		MaxLines int    `json:"maxlines,omitempty"`
		MaxSize  int    `json:"maxsize,omitempty"`
		Daily    bool   `json:"daily,omitempty"`
		MaxDays  int    `json:"maxdays,omitempty"`
		Rotate   bool   `json:"rotate,omitempty"`
		Level    int    `json:"level,omitempty"`
		Color    bool   `json:"color,omitempty"`
	}
	const (
		KDefaultFilename = "error.log"
		KDefaultMaxLines = 1000000
		KDefaultMaxSize  = 1 << 28
		KDefaultMaxDays  = 7
	)
	conf := LogConf{
		Filename: web.AppConfig.DefaultString("log::filepath", KDefaultFilename),
		MaxLines: web.AppConfig.DefaultInt("log::max_lines", KDefaultMaxLines),
		MaxSize:  web.AppConfig.DefaultInt("log::max_size", KDefaultMaxSize),
		Daily:    web.AppConfig.DefaultBool("log::daily", true),
		MaxDays:  web.AppConfig.DefaultInt("log::max_days", KDefaultMaxDays),
		Rotate:   web.AppConfig.DefaultBool("log::rotate", true),
		Level:    web.AppConfig.DefaultInt("log::level", logs.LevelDebug),
		Color:    web.AppConfig.DefaultBool("log::color", true),
	}
	logAdapter := web.AppConfig.DefaultString("log::adapter", logs.AdapterFile)

	confStr, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	if err := logs.SetLogger(logAdapter, string(confStr)); err != nil {
		return err
	}

	return nil
}

//loadAppConfig 加载app.conf
func loadAppConfig() error {
	var conf string

	// 配置文件路径，默认为conf/app.conf
	flag.StringVar(&conf, "c", "conf/app.conf", "config file")
	// 解析参数
	flag.Parse()

	if err := web.LoadAppConfig("ini", conf); err != nil {
		return err
	}
	return nil
}

func main() {
	//TODO...过滤器这一部分放到routers那里
	web.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//TODO...这种系统运行信息使用info级别，不要用debug，因为实际生产环境会调整日志级别，不会输出debug级别的日志
	logs.Debug("load app.conf...")
	if err := loadAppConfig(); err != nil {
		logs.Error("load app conf failed, err: %s", err.Error())
		os.Exit(1)
	}

	//TODO...使用info
	logs.Debug("init logger...")
	if err := initLogger(); err != nil {
		logs.Error("init logger conf failed, err: %s", err.Error())
		os.Exit(1)
	}

	//TODO...使用info
	logs.Debug("init mysql...")
	if err := initMySQL(); err != nil {
		logs.Error("init mysql failed, err: %s", err.Error())
		os.Exit(1)
	}

	//TODO...使用info
	logs.Debug("beego start...")
	web.Run()
}
