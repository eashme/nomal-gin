package services

import "nomal-gin/dao"

type UserService struct {
	ud dao.IUserDao
}

func NewUserService(ud dao.IUserDao) *UserService {
	return &UserService{ud: ud}
}

func (us *UserService) Get(id int64) (*dao.User,){

}