package services

import (
	"database/sql"
	"nomal-gin/cache"
	"nomal-gin/dao"
)

type userService struct {
	*BaseService
	ud dao.IUserDao
}

func NewUserService(db *sql.DB, cache cache.ICache) *userService {
	return &userService{NewBaseService(cache), dao.NewUserDao(db)}
}

func (us *userService) Get(id int64) (*dao.User, *WarpErr) {
	return nil,nil
}

func (us *userService)GetAll(page,pageSize int) (IPageNator,*WarpErr){
	return nil,nil
}