package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(u)
	checkErr(err)
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	rows, err := engine.Rows(new(UserInfo))
	defer rows.Close()
	checkErr(err)
	bean := new(UserInfo)
	var uList []UserInfo
	for rows.Next() {
		err = rows.Scan(bean)
		uList = append(uList, *bean)
	}
	return uList
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	u := new(UserInfo)
	_, err := engine.ID(id).Get(u)
	checkErr(err)
	return u
}
