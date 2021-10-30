package dao

import (
	"github.com/pkg/errors"
)

type User struct {
	Id            int     `orm:"column(id);pk;auto"`
	Account		  string  `orm:"column(account)"`
	Email		  string  `orm:"column(email)"`
}

func (u *User) Insert() error {
	return errors.Wrap(db.Create(u).Error, "user insert failed")
}

func (u *User) Update() error {
	return errors.Wrap(db.Save(u).Error, "user update failed")
}

func GetUserById(id int) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		// 疑问，这里是否应该过滤 gorm.ErrRecordNotFound 类型的error
		// 若不过滤，那么调用方就与 gorm.ErrRecordNotFound 这个哨兵error强耦合了

		// 方案1：过滤哨兵error
		//if err != gorm.ErrRecordNotFound {
		//	return nil, errors.Wrap(err, "GetUserById failed")
		//}
		//return nil, nil

		// 方案2：所以error都返回
		return nil, errors.Wrap(err, "GetUserById failed")
	}
	return &user, nil
}

func GetUserByAccount(account string) (*User, error) {
	var user User
	if err := db.Where("account=?", account).First(&user).Error; err != nil {
		//if err != gorm.ErrRecordNotFound {
		//	return nil, errors.Wrap(err, "GetUserByAccount failed")
		//}
		//return nil, nil
		return nil, errors.Wrap(err, "GetUserByAccount failed")
	}
	return &user, nil
}