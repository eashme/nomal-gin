package services

import (
	"database/sql"
	"nomal-gin/cache"
	"nomal-gin/dao"
)

type UserService struct {
	*BaseService
	ud dao.IUserDao
}

func NewUserService(db *sql.DB, cache cache.ICache) *UserService {
	return &UserService{NewBaseService(cache), dao.NewUserDao(db)}
}

func (us *UserService) Get(id int64) (*dao.User, ) {

}

func (us *UserService)GetAll(page,pageSize int) (IPageNator,*WarpErr){

}