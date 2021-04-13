package mock

type DB1 struct {
}

func NewDB1() *DB1 {
	return &DB1{}
}

func (*DB1) Get(key string) (int, error) {
	//todo 查询数据库
	return 1, nil
}

func GetFromeDB1(key string) int {
	db := NewDB1()
	// 常规调用
	if r, err := db.Get(key); err == nil {
		return r
	}
	return -1
}

// 将依赖项抽象为接口
type DB interface {
	Get(key string) (int, error)
	Delete(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	//依赖注入：将依赖传递给调用方
	if r, err := db.Get(key); err == nil {
		return r
	}
	return -1
}
