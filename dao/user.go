package dao

import (
	"database/sql"
	qb "github.com/didi/gendry/builder"
	"nomal-gin/utils/errlog"
)

type User struct {
	Name string
	Age  int
}

type UserDao struct {
	*BaseDao
}

func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{BaseDao: NewBaseDao(db)}
}

// if not implement their method , will panic

func (ud *UserDao) TableName() string {
	return "user"
}

func (ud *UserDao) GetAll(where map[string]interface{}, page, pageSize int) (interface{}, error) {
	cond, vals, err := qb.BuildSelect(ud.TableName(), where, []string{"name", "count(price) as total", "age"})
	if err != nil {
		errlog.Error("sql build failed ", err)
	}
	rows, err := ud.db.Query(cond, vals...)
	if err != nil {
		errlog.Error("db execute failed ", err)
	}
	// todo marshal the rows which from db to users and return
}
