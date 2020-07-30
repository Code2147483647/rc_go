package main

import (
	"fmt"
	ConfigService "rc_go/configs"
	"rc_go/models"
	"rc_go/repositories"
	routes "rc_go/routes"
	"runtime"
	_ "time"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// ormDBInit
func dbInit() orm.Ormer {
	configSrv := ConfigService.GetInstance()

	mysqlCon := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		configSrv.GetString("mysql.user"),
		configSrv.GetString("mysql.password"),
		configSrv.GetString("mysql.host"),
		configSrv.GetInt32("mysql.port"),
		configSrv.GetString("mysql.database"),
		configSrv.GetString("mysql.charset"),
	)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(configSrv.GetString("mysql.dbAlias"), "mysql", mysqlCon)

	orm.RegisterModel(new(models.Account))
	orm.RegisterModel(new(models.Action))
	orm.RegisterModel(new(models.ActionTag))
	orm.RegisterModel(new(models.Brand))
	orm.RegisterModel(new(models.RiskLevel))
	orm.RegisterModel(new(models.RuleAction))
	orm.RegisterModel(new(models.RuleCondition))
	orm.RegisterModel(new(models.Rule))
	orm.RegisterModel(new(models.SystemSetting))

	orm.Debug = false
	instance := orm.NewOrm()
	instance.Using("rc")
	return instance
}

// func dbInit() *sqlx.DB {
// 	db, err := sqlx.Connect("mysql", mysqlCon)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db.SetMaxOpenConns(1000)
// 	db.SetMaxIdleConns(10)
// 	db.SetConnMaxLifetime(0)
// 	return db
// }

func main() {
	runtime.GOMAXPROCS(1)
	// go func() {
	// 	for range time.Tick(1 * time.Second) {
	// 		fmt.Println("runtime.NumGoroutine: ", runtime.NumGoroutine())
	// 	}
	// }()
	dbInstance := dbInit()
	repositories.InitRep(dbInstance)
	r := gin.Default()
	routes.InitUserRoute(r)
	r.Run(":5000")
}
