package model

import (
	"fmt"
	"log"
	"user/server/conf"
	"user/server/data/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var orm *xorm.Engine

func Init() (*xorm.Engine, error) {

	secret := fmt.Sprintf("%s:%s@tcp(%s:%s)%s?charset=utf8", conf.DbUser, conf.DbPassWord, conf.DbHost, conf.DbPort, conf.DbName)
	orm, err := xorm.NewEngine(conf.Db, secret)
	if err != nil {
		log.Fatalf("connect db error:%v", err)
		return nil, err
	}
	defer orm.Close()

	return orm, nil
}

func SelectUser(userName string) (*domain.TblUserInfo, error) {

	orm, err := Init()
	if err != nil {
		log.Fatalf("connect db error:%v", err)
		return nil, err
	}
	var userInfo domain.TblUserInfo
	err = orm.Where("user_name=?", userName).Find(&userInfo)
	if err != nil {
		log.Fatalf("selecr userinfo error:%v", err)
		return nil, err
	}
	return &userInfo, nil
}

func InsertUserInfo(tblUserInfo domain.TblUserInfo) error {

	orm, err := Init()
	if err != nil {
		log.Fatalf("connect db error:%v", err)
		return err
	}

	_, err = orm.Insert(tblUserInfo)
	if err != nil {
		log.Fatalf("insert error:%v", err)
	}
	return nil
}
