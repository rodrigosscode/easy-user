package paging

import "gorm.io/gorm"

type (
	DbPagerConfig struct {
		Page  int
		Limit int
	}
)

func (p *DbPagerConfig) GetOffset() int {
	offset := (p.Page - 1) * p.Limit
	return offset
}

func (p *DbPagerConfig) PaginateResultScope(gormDb *gorm.DB) *gorm.DB {
	return gormDb.Limit(p.Limit).Offset(p.GetOffset())
}
