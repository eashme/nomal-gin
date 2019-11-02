package dao

import "database/sql"

type BaseDao struct {
	db *sql.DB // 数据库连接
}

func NewBaseDao(db *sql.DB) *BaseDao {
	return &BaseDao{db: db}
}

func (*BaseDao) TableName() string {
	panic("implement me")
}

func (*BaseDao) GetAll(where map[string]interface{}, page, pageSize int) (interface{}, error) {
	panic("implement me")
}

func (*BaseDao) Get(id int64) (interface{}, error) {
	panic("implement me")
}

func (*BaseDao) GetBy(where map[string]interface{}, page, pageSize int) (interface{}, error) {
	panic("implement me")
}

func (*BaseDao) Delete(id int64) (n int64, err error) {
	panic("implement me")
}

func (*BaseDao) Insert(bean interface{}) (ok bool, err error) {
	panic("implement me")
}

func (*BaseDao) InsertMany(bean ...interface{}) (n int64, err error) {
	panic("implement me")
}

func (*BaseDao) Update(bean interface{}, where map[string]interface{}) (n int64, err error) {
	panic("implement me")
}
