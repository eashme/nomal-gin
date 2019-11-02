package dao

import "database/sql"

type UserDao struct {
	*BaseDao
}

func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{BaseDao:NewBaseDao(db)}
}

// if not implement their method , while panic

func (u *UserDao) TableName() string {
	return "user"
}
