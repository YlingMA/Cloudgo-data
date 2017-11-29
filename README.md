
#1.用使用 xorm 或 gorm 实现本文的程序！
在老师给的模板上进行修改。
具体修改如下：
##（a）启动mysql
```javascript
db, err := sql.Open("mysql", "root:myl.rl.0676@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
```

使用xorm启动sql如下：
```javascript
Engine, err := xorm.NewEngine("mysql", "root:myl.rl.0676@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
```

##(b)创建数据库中的表格struct
```javascript
type UserInfo struct {
    UID        int `xorm:"id pk autoincr"` //语义标签
	UserName   string
	DepartName string
	CreateAt   *time.Time
}
```

##(c)插入数据
>_, err := engine.Insert(u)

注：这里的u是表格UserInfo的一个实例
那么，像老师写的函数userinfo-service.go中的Save函数可以直接写为
```javascript
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(u)
	checkErr(err)
	return nil
}
```
而不必进入到userinfo-dao.go里面去，再使用对应的接口函数Prepare(),Exec().

##(d)查询所有用户信息
```javascript
func (*UserInfoAtomicService) FindAll() []UserInfo {
	rows, err := engine.Rows(new(UserInfo))
	defer rows.Close()
	checkErr(err)
	bean := new(UserInfo)
	var uList []UserInfo
	for rows.Next() {				//扫描所有的用户，一个用户为一行
		err = rows.Scan(bean)    
		uList = append(uList, *bean)  //将扫描的每一个用户信息append进链表里
	}
	return uList
}
```


##(e)查询某个ID的用户信息
```javascript
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	u := new(UserInfo)
	_, err := engine.ID(id).Get(u)   //对于ID为id的用户，得到它的用户信息
	checkErr(err)
	return u
}
```
Ab测试
测试post：
ab -n 100 -c 10 -p test.txt http://localhost:8080/service/userinfo

![image](https://github.com/YlingMA/Cloudgo-data/blob/master/image/ab-post1.PNG)

如图ab-post-1
如图ab-post-2
如图ab-post-3
其中，test.txt中内容如下图：

测试GET:
ab -n 100 -c 10 http://localhost:8080/service/userinfo?userid=

50%的请求在42ms内完成了，所有的请求在71ms内完成。每一组请求平均耗费47.813ms，每个请求平均花费4.7813秒。每秒平均完成209.15个选择。

#2.设计 GoSqlTemplate 的原型

这个模板是建立在database/sql的基础上。
创建一个名为SQLTemplate的struct，仅包含一个接口SQLExecer.
创建的函数都是调用SQLExcer接口上的函数。
##(a)Template（sqlt）中可调用的函数的功能
Template是在老师给出的demo上进行的修改。
对于sqlt,实现了以下功能：
Select():实际上是展示所有满足条件的对象的信息。
Selectone():展示第一个满足条件的对象的信息
Update():更新对象信息
Delete():删除所有对象信息
##(b)具体实现时使用的功能
* SELECT * FROM username

如图username=xiaoma


* SELECT * FROM userinfo
图findall


* SELECT ID WHERE id = 1
图selectone




* 得到总共的用户数：
图count

##(c)下面介绍template即sqlt的用法：
详情见userinfo-dao.go和userinfo-service.go
```javascript
//FindAll()
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
```
