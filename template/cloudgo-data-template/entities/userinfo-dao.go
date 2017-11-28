package entities

import (
	"github.com/YlingMA/Cloudgo-data/template/sqlt"
)

type userInfoDao struct {
	sqlt.SQLTemplate
}

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
	return dao.Insert(userInfoInsertStmt, &u.UID, u.UserName, u.DepartName, u.CreateAt)
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"
var userInfoCount = "SELECT count(*) FROM userinfo"
var userInfoQueryByName = "SELECT * FROM userinfo where username = ?"

func getUserInfoMapper(ul *[]UserInfo) sqlt.RowMapperCallback {
	return func(row sqlt.RowScanner) error {
		u := UserInfo{}
		err := row.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
		if err != nil {
			return err
		}
		*ul = append(*ul, u)
		return nil
	}
}

func getUserInfoOnceMapper(u *UserInfo) sqlt.RowMapperCallback {
	return func(row sqlt.RowScanner) error {
		err := row.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
		return err
	}
}

// FindAll .
func (dao *userInfoDao) FindAll() ([]UserInfo, error) {
	ulist := make([]UserInfo, 0, 0)
	err := dao.Select(userInfoQueryAll, getUserInfoMapper(&ulist))
	return ulist, err
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) (*UserInfo, error) {
	u := UserInfo{}
	err := dao.SelectOne(userInfoQueryByID, getUserInfoOnceMapper(&u), id)
	return &u, err
}

// Count .
func (dao *userInfoDao) Count() (int, error) {
	count := 0
	f := func(row sqlt.RowScanner) error {
		err := row.Scan(&count)
		return err
	}
	err := dao.SelectOne(userInfoCount, f)
	return count, err
}

//userInfoQueryByName
func (dao *userInfoDao) FindByName(username string) ([]UserInfo, error) {
	u :=  make([]UserInfo, 0, 0)
	err := dao.Select(userInfoQueryByName,  getUserInfoMapper(&u), username)
	return u, err
}
