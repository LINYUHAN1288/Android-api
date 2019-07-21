package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UserBasic struct {
	Id          int    `orm:"column(id);auto" description:"id "`
	Name        string `orm:"column(name);size(40);null" description:"用户名"`
	Password    string `orm:"column(password);size(40);null" description:"用户密码"`
	TrueName    string `orm:"column(trueName);size(40);null" description:"用户真实姓名"`
	PersonImage string `orm:"column(personImage);null" description:"用户图像"`
	IDCard      string `orm:"column(IDCard);size(20);null" description:"身份证号码"`
	Email       string `orm:"column(email);size(40);null" description:"用户邮箱"`
	Mobile      string `orm:"column(mobile);size(20);null" description:"用户手机号"`
	SourceType  int    `orm:"column(SourceType);null" description:"0-手工注册；1-微博；2-QQ；3-微信 账户注册类型 默认0"`
	Sex               int       `orm:"column(sex);null" description:"1-男；2-女"`
	BirthDate         time.Time `orm:"column(birthDate);type(date);null" description:"出生日期"`
	Age               float64   `orm:"column(age);null" description:"年龄"`
	Height            int       `orm:"column(height);null" description:"年龄"`
	Weight            int       `orm:"column(weight);null" description:"体重kg"`
	Memo              string    `orm:"column(memo);null" description:"备注，用户可再此输入和自己相关的一些信息"`
	UserType          int       `orm:"column(userType);null" description:"0-用户 1-营养师 用户类型，不同类型的用户注册时必须输入的信息不同"`
	CertNo            string    `orm:"column(certNo);size(30);null" description:"资格证编号，仅对用户类型为1的用户有效"`
	CertDate          time.Time `orm:"column(certDate);type(date);null" description:"取得营养师资格的时间"`
	CreateTime        int       `orm:"column(create_time);null" description:"创建时间"`
	UpdateTime        int       `orm:"column(update_time);null" description:"更新时间"`
	PAL               int       `orm:"column(PAL);null"`
	GluteninAllery    int       `orm:"column(gluteninAllery);null"`
	FlowerAllery      int       `orm:"column(flowerAllery);null"`
	PeanutAllery      int       `orm:"column(peanutAllery);null"`
	EggAllery         int       `orm:"column(eggAllery);null"`
	TreeNutAllery     int       `orm:"column(treeNutAllery);null"`
	FishAllery        int       `orm:"column(fishAllery);null"`
	EatHeartHealth    int       `orm:"column(eatHeartHealth);null"`
	EatLowCholesterol int       `orm:"column(eatLowCholesterol);null"`
	EatLowBloodPres   int       `orm:"column(eatLowBloodPres);null"`
	EatForPregnant    int       `orm:"column(eatForPregnant);null"`
	EatForMomBaby     int       `orm:"column(eatForMomBaby);null"`
	Pregnantlevel     int8      `orm:"column(Pregnantlevel);null"`
	AvoidGMO          int8      `orm:"column(AvoidGMO);null"`
	AvoidPork         int       `orm:"column(avoidPork);null"`
	AvoidMeatAndFish  int       `orm:"column(avoidMeatAndFish);null"`
	EatVegan          int       `orm:"column(eatVegan);null"`
	PaleoDiet         int       `orm:"column(paleoDiet);null"`
	SpecialFood       string    `orm:"column(specialFood);size(255);null"`
	AutoGenerateStps  int       `orm:"column(autoGenerateStps);null"`
	AutoDate          time.Time `orm:"column(autoDate);type(time);null"`
	MonitorId         int       `orm:"column(monitor_id);null" description:"首页监控id"`
	Openid            string    `orm:"column(openid);size(40);null" description:"第三方登录返回openid"`
}

func (t *UserBasic) TableName() string {
	return "user_basic"
}

func init() {
	orm.RegisterModel(new(UserBasic))
}

// AddUserBasic insert a new UserBasic into database and returns
// last inserted Id on success.
func AddUserBasic(m *UserBasic) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserBasicById retrieves UserBasic by Id. Returns error if
// Id doesn't exist
func GetUserBasicById(id int) (v *UserBasic, err error) {
	o := orm.NewOrm()
	v = &UserBasic{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserBasicByMoblie retrieves UserBasic by Mobile. Returns error if
// Mobile doesn't exist
func GetUserBasicByMobile(mobile string) (v *UserBasic, err error) {
	o := orm.NewOrm()
	v = &UserBasic{Mobile: mobile}
	if err = o.Read(v, "mobile"); err == nil {
		return v, nil
	}
	err = errors.New("Error")
	return nil, err
}



// GetAllUserBasic retrieves all UserBasic matches certain condition. Returns empty list if
// no records exist
func GetAllUserBasic(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserBasic))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []UserBasic
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUserBasic updates UserBasic by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserBasicById(m *UserBasic) (err error) {
	o := orm.NewOrm()
	v := UserBasic{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateUserBasic updates Password and returns error if
// the record to be updated doen't exist
func UpdateUserPwd(v *UserBasic, oldpwd string, newpwd string) (err error) {
	o := orm.NewOrm()
	//v := &UserBasic{Password: oldpwd}
	if v.Password == oldpwd {
		v.Password = newpwd
		if num, err := o.Update(v, "password"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
		return err
	} else {
		err = errors.New("Error")
		return err
	}
	/*
	if err = o.Read(v, "password"); err == nil {
		v.Password = newpwd
		if num, err := o.Update(v, "password"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	*/
}


// DeleteUserBasic deletes UserBasic by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserBasic(id int) (err error) {
	o := orm.NewOrm()
	v := UserBasic{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserBasic{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
