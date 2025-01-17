package db

import (
	"context"
	"errors"

	db "github.com/rodrigosscode/easy-user/infrastructure/db/entity"
	"github.com/rodrigosscode/easy-user/infrastructure/db/paging"
	zapLogger "github.com/rodrigosscode/easy-user/infrastructure/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	maxOpenConns    = 10
	maxIdleConns    = 5
	connMaxLifetime = 5
)

type DbConnection struct {
	gormDb *gorm.DB
}

func NewDbConnection(dsn string) (*DbConnection, error) {
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		zapLogger.Warn("Failed to connect to db", zap.Error(err))
		return nil, errors.New("fail to connect db")
	}

	sqlDB, err := gormDb.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	gormDb.AutoMigrate(&db.User{})
	return &DbConnection{gormDb: gormDb}, nil
}

func (d *DbConnection) FindById(ctx *context.Context, id int) (db.User, error) {
	u := db.User{ID: id}
	result := d.gormDb.First(&u)

	if result.Error != nil {
		return db.User{}, result.Error
	}

	return u, nil
}

func (d *DbConnection) Save(ctx *context.Context, u db.User) (db.User, error) {
	uExisting := db.User{}
	resultExisting := d.gormDb.Where(&db.User{Email: u.Email}).First(&uExisting)

	if resultExisting.RowsAffected > 0 {
		return db.User{}, gorm.ErrDuplicatedKey
	}

	if resultExisting.Error != nil && !errors.Is(resultExisting.Error, gorm.ErrRecordNotFound) {
		return db.User{}, resultExisting.Error
	}

	result := d.gormDb.Create(&u)

	if result.Error != nil {
		return db.User{}, result.Error
	}

	return u, nil
}

func (d *DbConnection) DeleteById(ctx *context.Context, id int) error {
	u := db.User{ID: id}
	result := d.gormDb.Delete(&u)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *DbConnection) Update(ctx *context.Context, u db.User) (db.User, error) {
	result := d.gormDb.Updates(&u)

	if result.Error != nil {
		return db.User{}, result.Error
	}

	return u, nil
}

func (d *DbConnection) FindByPage(ctx *context.Context, page, limit int) ([]db.User, int64, error) {
	var users []db.User
	var totalRecords int64

	p := paging.DbPagerConfig{Page: page, Limit: limit}

	if err := d.gormDb.Model(&db.User{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := d.gormDb.Scopes(p.PaginateResultScope).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, totalRecords, nil
}
