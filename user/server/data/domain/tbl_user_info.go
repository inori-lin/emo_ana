package domain

type TblUserInfo struct {
	ID       int    `xorm:"pk autoincr 'id'"`
	UserName string `xorm:"user_name"`
	Password string `xorm:"password"`
}
