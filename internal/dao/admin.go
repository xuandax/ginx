package dao

var Admin = NewAdminDao()

type AdminDao struct {
	table string
	group string
	Dao
}

func NewAdminDao() *AdminDao {
	return &AdminDao{
		table: "admin",
		group: "default",
	}
}
