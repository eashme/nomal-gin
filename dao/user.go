package dao

import (
	"database/sql"
)

type User struct {
	Name string
	Age  int
}

type userDao struct {
	*BaseDao
}

func NewUserDao(db *sql.DB) *userDao {
	return &userDao{BaseDao: NewBaseDao(db)}
}

// if not implement their method , will panic

func (ud *userDao) TableName() string {
	return "user"
}

func (ud *userDao) GetAll(where map[string]interface{}, page, pageSize int) (interface{}, error) {
	//cond, vals, err := qb.BuildSelect(ud.TableName(), where, []string{"name", "count(price) as total", "age"})
	//if err != nil {
	//	errlog.Error("sql build failed ", err)
	//}
	//rows, err := ud.db.Query(cond, vals...)
	//if err != nil {
	//	errlog.Error("db execute failed ", err)
	//}
	// todo marshal the rows which from db to users and return
	return nil,nil
}

func (ud *userDao) Get(id int64) (interface{}, error){
	//cond, vals, err := qb.BuildSelect(ud.TableName(), map[string]interface{}{
	//	"id":id,
	//}, []string{"name", "count(price) as total", "age"})
	//if err != nil {
	//	errlog.Error("sql build failed ", err)
	//}
	//rows, err := ud.db.Query(cond, vals...)
	//if err != nil {
	//	errlog.Error("db execute failed ", err)
	//}
	//// todo marshal the rows which from db to users and return

	return nil,nil
}