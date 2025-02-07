package model

import (
	"gin_work/wrap/database"
	"gorm.io/gorm"
)

type User struct {
	Id            uint       `gorm:"primaryKey"`
	Uuid          string     `gorm:"uuid;<-:create"`
	Username      string     `gorm:"username;<-:create"`
	Password      string     `gorm:"password"`
	Salt          string     `gorm:"salt"`
	Nickname      string     `gorm:"nickname"`
	RealName      string     `gorm:"real_name"`
	Phone         string     `gorm:"phone"`
	Email         string     `gorm:"email"`
	SignUpIp      string     `gorm:"sign_up_ip;<-:create"`
	CreateTime    int64      `gorm:"create_time;autoCreateTime"`
	LastLoginIp   string     `gorm:"last_login_ip"`
	LastLoginTime int64      `gorm:"last_login_time"`
	Avatar        string     `gorm:"avatar"`
	SignUpType    int        `gorm:"sign_up_type;<-:create"`
	State         int        `gorm:"state"`
	DeleteTime    int64      `gorm:"delete_time"`
	softDelete    SoftDelete `default:"delete_time"`
}

const (
	UsernamePrefix = "y_"

	SaltLen = 10

	SignTypePhone   = 1
	SignTypeEmail   = 2
	SignTypeAccount = 3
	UserStateOff    = 0
	UserStateVerify = 1
	UserStateNormal = 2
)

func (*User) TableName() string {
	return "y_user"
}

func (u *User) FindById(id uint) (int64, error) {
	data := database.DB.Where(SoftWhere(u, "id = ?"), id).Find(&u)
	return data.RowsAffected, data.Error
}

func (u *User) FindByUuid(uuid string) (int64, error) {
	data := database.DB.Where(SoftWhere(u, "uuid = ?"), uuid).First(&u)
	return data.RowsAffected, data.Error
}

func (u *User) FindByUsername(username string) (int64, error) {
	data := database.DB.Where(SoftWhere(u, "username = ?"), username).First(&u)
	return data.RowsAffected, data.Error
}

func (u *User) FindByPhone(phone string) (int64, error) {
	data := database.DB.Where(SoftWhere(u, "phone = ?"), phone).First(&u)
	return data.RowsAffected, data.Error
}

func (u *User) FindByEmail(email string) (int64, error) {
	data := database.DB.Where(SoftWhere(u, "email = ?"), email).First(&u)
	return data.RowsAffected, data.Error
}

func (u *User) GetBySearch(where [][]interface{}, field interface{}, page, limit uint) (*[]User, int64) {
	var users []User
	var count int64
	sqlWhere, params := Or(where)
	database.DB.Model(&u).Select("id").Where(SoftWhere(u, sqlWhere), params...).Count(&count)
	database.DB.Select(field).Where(SoftWhere(u, sqlWhere), params...).Limit(int(limit)).Offset(int((page - 1) * limit)).Find(&users)
	return &users, count
}

func (u *User) Create() (int64, error) {
	result := database.DB.Create(u)
	return result.RowsAffected, result.Error
}

func (u *User) Update(data ...interface{}) (int64, error) {
	var result *gorm.DB
	if len(data) > 0 {
		err := database.DB.Model(&u).Where(SoftWhere(u, "id = ?"), u.Id).Updates(data[0]).Error
		if err != nil {
			return 0, err
		}
		return 1, err
	} else {
		result = database.DB.Model(&u).Where(SoftWhere(u, "id = ?"), u.Id).Updates(u)
	}
	return result.RowsAffected, result.Error
}

func (u *User) SoftDeleteById(ids ...uint) (int64, error) {
	tx := database.DB
	if len(ids) > 0 {
		return SoftDel(tx, u, "id IN ?", ids)
	} else {
		return SoftDel(tx, u, "id = ?", u.Id)
	}
}
