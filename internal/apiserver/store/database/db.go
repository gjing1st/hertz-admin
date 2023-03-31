// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 15:49$

package database

import (
	"database/sql"
	"fmt"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	log1 "log"
	"os"
	"time"
)

var (
	db *gorm.DB
)

const (
	DriverPostgresql = "postgresql"
	DriverMysql      = "mysql"
	DriverMongo      = "mongodb"
)

// InitDB
//	@description:	初始化数据库
//	@param:
//	@author:	GJing
//	@email:		gjing1st@gmail.com
//	@date:		2022/4/6 22:37
//	@success:
func InitDB() {
	var err error
	//数据库驱动
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.UserName,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.DBName,
	)
	newLogger := logger.New(
		log1.New(os.Stdout, "\r\n", log1.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,        // 禁用彩色打印
		},
	)
	fmt.Println("========dsn", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger, //日志
		DisableForeignKeyConstraintWhenMigrating: true,      //外键
		PrepareStmt:                              true,      //预编译缓存
	})
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库连接失败")
		return
	}
	//	读写分离
	slaveDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Slave.UserName,
		config.Config.Slave.Password,
		config.Config.Slave.Host,
		config.Config.Slave.Port,
		config.Config.Slave.DBName,
	)
	err = db.Use(
		dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.Open(slaveDsn)},
		}).
			//设置从库连接池
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(time.Hour).
			SetMaxIdleConns(config.Config.Mysql.MinConns).
			SetMaxOpenConns(config.Config.Mysql.MaxConns),
	)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Error(DriverMysql + "配置数据库读写分离失败")
		return
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库连接失败")
		return
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(config.Config.Mysql.MinConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Info("init db success")
}

// GetDB
//	@description:	获取数据库连接
//	@param:
//	@author:	GJing
//	@email:		gjing1st@gmail.com
//	@date:		2022/4/6 22:38
//	@success:
func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	// 初始化表和表数据
	//initdata.InitData(db)

	return db
}

// MysqlEmptyDsn
//	@description:	mysql配置
//	@param:
//	@author:	GJing
//	@email:		guojing@tna.cn
//	@date:		2022/10/26 18:15
//	@success:
func MysqlEmptyDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.Config.Mysql.UserName,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port)
}

// createDatabase
//	@description:	创建数据库
//	@param:
//	@author:	GJing
//	@email:		guojing@tna.cn
//	@date:		2022/10/26 18:15
//	@success:
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	//fmt.Println("dsn", dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
