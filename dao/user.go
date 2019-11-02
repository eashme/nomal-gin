package dao

import "database/sql"

type User struct {
	Name string
	Age int
}

type UserDao struct {
	*BaseDao
}

func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{BaseDao:NewBaseDao(db)}
}

// if not implement their method , while panic

func (ud *UserDao) TableName() string {
	return "user"
}

func (ud *UserDao) GetAll(){

}