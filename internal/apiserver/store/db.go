// $
// Created by bestTeam.
// Author: GJing
// Date: 2022/9/9$ 15:49$

package store

import (
	"database/sql"
	"fmt"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/database/initdata"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
	"github.com/gjing1st/hertz-admin/pkg/global"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	//dm "github.com/nfjBill/gorm-driver-dm"
	//"gorm.io/driver/clickhouse"
	log "github.com/sirupsen/logrus"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

// InitDB
// @description: 初始化数据库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 22:37
// @success:
func InitDB() {
	var err error
	var dsn, createSql string
	if config.Config.Base.DBType == global.DriverMysql {
		dsn = MysqlEmptyDsn()
		createSql = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", config.Config.Database.DBName)
	} else if config.Config.Base.DBType == global.DriverPostgresql || config.Config.Base.DBType == global.DriverOpenGauss {
		creatDsn := createPostgresDsn()
		createSql = fmt.Sprintf("CREATE DATABASE %s   WITH OWNER = %s  ENCODING = 'UTF8'   TEMPLATE template0;", config.Config.Database.DBName, config.Config.Database.UserName)
		if err = createPostgresqlDatabase(creatDsn, createSql); err != nil {
			log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.Base.DBType + "数据库创建失败")
			return
		}
	} else if config.Config.Base.DBType == global.DriverKingBase {
		creatDsn := createKingBaseDsn()
		createSql = fmt.Sprintf("CREATE DATABASE %s   WITH OWNER = %s  ENCODING = 'UTF8'   TEMPLATE template0;", config.Config.Database.DBName, config.Config.Database.UserName)
		if err = createPostgresqlDatabase(creatDsn, createSql); err != nil {
			log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.Base.DBType + "数据库创建失败")
			return
		}
	}
	if len(createSql) > 0 {
		// 创建数据库
		if err = createDatabase(dsn, config.Config.Base.DBType, createSql); err != nil {
			log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.Base.DBType + "数据库创建失败")
			return
		}
	}
	switch config.Config.Base.DBType {
	case global.DriverMysql:
		//数据库驱动
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config.Database.UserName,
			config.Config.Database.Password,
			config.Config.Database.Host,
			config.Config.Database.Port,
			config.Config.Database.DBName,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case global.DriverDM:
	//dsn = fmt.Sprintf("dm://%s:%s@%s:%s?ignoreCase=false&statEnable=false&schema=%s&autoCommit=true",
	//	config.Config.Database.UserName,
	//	config.Config.Database.Password,
	//	config.Config.Database.Host,
	//	config.Config.Database.Port,
	//	config.Config.Database.DBName,
	//)
	//db, err = gorm.Open(dm.Open(dsn), &gorm.Config{
	//	DisableForeignKeyConstraintWhenMigrating: true,
	//})
	case global.DriverPostgresql, global.DriverOpenGauss:
		//数据库驱动
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
			config.Config.Database.Host,
			config.Config.Database.UserName,
			config.Config.Database.Password,
			config.Config.Database.DBName,
			config.Config.Database.Port,
			config.Config.Database.SSLMode,
		)
		//db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case global.DriverKingBase:
		//数据库驱动
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
			config.Config.Database.Host,
			config.Config.Database.UserName,
			config.Config.Database.Password,
			config.Config.Database.DBName,
			config.Config.Database.Port,
			config.Config.Database.SSLMode,
		)
		//db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case global.DriverSqlite:
		fmt.Println("path", config.Config.Database.Host)
	//db, err = gorm.Open(sqlite.Open(config.Config.Database.Host), &gorm.Config{})
	case global.DriverClickHouse:
		//dsn = fmt.Sprintf("tcp://%s:%s?database=%s&username=%s&password=%s",
		//	config.Config.Database.Host,
		//	config.Config.Database.Port,
		//	config.Config.Database.DBName,
		//	config.Config.Database.UserName,
		//	config.Config.Database.Password,
		//)
		//dsn = "clickhouse://tna:Dked@213@192.168.200.83:9000/gaf?dial_timeout=10s&read_timeout=20s"

		dsn = fmt.Sprintf("clickhouse://%s:%s@%s:%s/%s?dial_timeout=10s&read_timeout=20s",
			config.Config.Database.UserName,
			config.Config.Database.Password,
			config.Config.Database.Host,
			config.Config.Database.Port,
			config.Config.Database.DBName,
		)
		fmt.Println("dsn=======", dsn)

		//db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.Base.DBType + "数据库连接失败")
		return
	}
	//db.Use(dbresolver.Register(dbresolver.Config{
	//	// `db2` 作为 sources，`db3`、`db4` 作为 replicas
	//	Sources:  []gorm.Dialector{mysql.Open("dsn")},
	//	Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
	//	// sources/replicas 负载均衡策略
	//	Policy: dbresolver.RandomPolicy{},
	//}))

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.Base.DBType + "数据库连接失败")
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.Base.DBType + "数据库连接失败")
		return
	}

	// SetMaxIdleConns 设置MySQL的最大空闲连接数。
	sqlDB.SetMaxIdleConns(config.Config.Database.MinConns)
	// SetMaxOpenConns 设置MySQL的最大连接数。
	sqlDB.SetMaxOpenConns(config.Config.Database.MaxConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	log.Info("init db success")
}

// GetDB
// @description: 获取数据库连接
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 22:38
// @success:
func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	// 初始化表和表数据
	initdata.InitData(db)

	return db
}

// MysqlEmptyDsn
// @description: mysql配置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/26 18:15
// @success:
func MysqlEmptyDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.Config.Database.UserName,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port)
}

// createDatabase
// @description: 创建数据库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/26 18:15
// @success:
func createDatabase(dsn string, driver string, createSql string) error {
	db1, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db1)
	if err = db1.Ping(); err != nil {
		return err
	}
	_, err = db1.Exec(createSql)
	return err
}

// @Description postgres系统dns连接,此处基于postgres数据库做跳板，金仓需要修改dbname=kingbase
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/1/11 10:05
func createPostgresDsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.UserName,
		config.Config.Database.Password,
	)
}

func createKingBaseDsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=kingbase sslmode=disable",
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.UserName,
		config.Config.Database.Password,
	)
}

// createDatabase
// @description: 创建postgres数据库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/1/11 18:15
// @success:
func createPostgresqlDatabase(dsn, createSql string) error {
	fmt.Println("------dsn--", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("err===", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("ping--err===", err)
	}

	// 创建数据库
	_, err = db.Exec(createSql)

	fmt.Println("create---err==", err)
	if err != nil {
		log.Println("create--db-err==", err)
	}
	return nil
}
