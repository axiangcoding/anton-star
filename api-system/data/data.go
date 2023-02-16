package data

import (
	"github.com/axiangcoding/antonstar-bot/data/dal"
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"time"
)

var _db *gorm.DB

func InitData() {
	_db = initDB()
}

func initDB() *gorm.DB {
	source := settings.C().App.Data.Db.Source
	dial := postgres.Open(source)
	dbLogger := zapgorm2.New(logging.L())
	dbLogger.SetAsDefault()
	dbLogger.IgnoreRecordNotFoundError = true
	dbLogger.SlowThreshold = 1 * time.Second
	dbLogger.LogLevel = logger.Silent
	db, err := gorm.Open(dial,
		&gorm.Config{
			NamingStrategy: &schema.NamingStrategy{SingularTable: true},
			Logger:         dbLogger})
	if err != nil {
		logging.L().Fatal("can't connect to db",
			logging.Error(err),
			logging.Any("source", source))
	} else {
		logging.L().Info("database connected success")
	}
	setProperties(db)
	autoMigrate(db)
	dal.SetDefault(db)
	return db
}

func GetDB() *gorm.DB {
	return _db
}

// 自动更新表结构
func autoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&table.Mission{},
		&table.GameUser{},
		&table.QQGroupConfig{},
		&table.QQUserConfig{},
		&table.GlobalConfig{},
	); err != nil {
		logging.L().Fatal("auto migrate error", logging.Error(err))
	} else {
		logging.L().Info("auto migrate db tables success")
	}
}

func setProperties(db *gorm.DB) {
	s, err := db.DB()
	if err != nil {
		logging.L().Fatal("open db failed while set properties",
			logging.Error(err))
	}
	s.SetMaxOpenConns(settings.C().App.Data.Db.MaxOpenConn)
	s.SetMaxIdleConns(settings.C().App.Data.Db.MaxIdleConn)
}

func GenCode(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./data/dal",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable: true,
	})

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(table.Mission{},
		table.GameUser{},
		table.QQGroupConfig{},
		table.QQUserConfig{},
		table.GlobalConfig{})

	// Execute the generator
	g.Execute()
}
