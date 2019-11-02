package dao

type IBaseDao interface {
	TableName() string
	GetAll(where map[string]interface{}, page, pageSize int) (interface{}, error)
	Get(id int64) (interface{}, error)
	GetBy(where map[string]interface{}, page, pageSize int) (interface{}, error)
	Delete(id int64) (n int64, err error)
	Insert(bean interface{}) (ok bool, err error)
	InsertMany(bean ...interface{}) (n int64, err error)
	Update(bean interface{}, where map[string]interface{}) (n int64, err error)
}

type IUserDao interface {
	IBaseDao
	//... Special Method
}
