package storage

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

type PostgreConfig struct {
	Host   string
	User   string
	Passwd string
	Dbname string
	Port   int
}

var globalDB *gorm.DB

func loadDBConifg() (dsn string) {
	dbConfig := PostgreConfig{}
	rootDir, _ := filepath.Abs(filepath.Dir("."))
	dbConfigName := path.Join(rootDir, "config", "config.json")
	err := gonfig.GetConf(dbConfigName, &dbConfig)
	if err != nil {
		log.Fatalf("read postgre config error: %s\n", err.Error())
	}
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.User, dbConfig.Passwd, dbConfig.Dbname, dbConfig.Port)
	return
}

func loadLogger() logger.Interface {
	config := logger.Config{
		SlowThreshold: time.Second * 5, // 慢 SQL 阈值
		LogLevel:      logger.Info,     // Log level
		Colorful:      false,           // 禁用彩色打印
	}
	rootDir, _ := filepath.Abs(filepath.Dir("."))
	loggerFile := path.Join(rootDir, "logs", "gorm.log")
	fp, err := os.Open(loggerFile)
	if err != nil {
		log.Println(err)
		fp = os.Stdout
	}
	return logger.New(log.New(fp, "\r\n", log.LstdFlags), config)
}
func migrate(db *gorm.DB) (err error) {
	if err := db.AutoMigrate(&Topic{}, &Question{}, &Answer{}, &Comment{}, &User{}, &VoteAnswer{}, &VoteComment{}, &Follow{}, &Attention{}); err != nil {
		log.Println(err)
	}
	if !db.Migrator().HasTable(&Topic{}) {
		if err = db.Migrator().CreateTable(&Topic{}); err != nil {
			return
		}
	}
	if !db.Migrator().HasTable(&Question{}) {
		if err = db.Migrator().CreateTable(&Question{}); err != nil {
			return
		}
	}
	if !db.Migrator().HasTable(&Answer{}) {
		if err = db.Migrator().CreateTable(&Answer{}); err != nil {
			return
		}
	}
	if !db.Migrator().HasTable(&Comment{}) {
		if err = db.Migrator().CreateTable(&Comment{}); err != nil {
			return
		}
	}
	if !db.Migrator().HasTable(&User{}) {
		if err = db.Migrator().CreateTable(&User{}); err != nil {
			return
		}
	}

	if !db.Migrator().HasTable(&VoteAnswer{}) {
		if err = db.Migrator().CreateTable(&VoteAnswer{}); err != nil {
			return
		}
	}

	if !db.Migrator().HasTable(&VoteComment{}) {
		if err = db.Migrator().CreateTable(&VoteComment{}); err != nil {
			return
		}
	}
	if !db.Migrator().HasTable(&Follow{}) {
		if err = db.Migrator().CreateTable(&Follow{}); err != nil {
			return
		}
	}

	if !db.Migrator().HasTable(&Attention{}) {
		if err = db.Migrator().CreateTable(&Attention{}); err != nil {
			return
		}
	}
	return
}
func init() {
	dsn := loadDBConifg()
	log.Printf("postgre config: %s\n", dsn)
	newLogger := loadLogger()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("open postgre error: %s\n", err.Error())
	}
	//if err := db.AutoMigrate(&Question{}, &Answer{}, &Comment{}, &User{}, &VoteAnswer{}, &VoteComment{}); err != nil {
	//	log.Println(err)
	//}
	if err = migrate(db); err != nil {
		log.Fatalf("migrate error: %v\n", err)
	}
	log.Printf("get db: %#v\n", db)
	globalDB = db
}

func GetDB() *gorm.DB {
	return globalDB
}
