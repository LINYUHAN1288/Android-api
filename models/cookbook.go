package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Cookbook struct {
	Id         int       `orm:"column(id);auto" description:"id "`
	Uid        int       `orm:"column(uid)" description:" 用户ID"`
	Foodid     string    `orm:"column(foodid);size(50)" description:"食物id"`
	Diettime   int8      `orm:"column(diettime);null" description:"早中晚"`
	Source     string    `orm:"column(source);size(255)" description:"b表名"`
	Week       int8      `orm:"column(week)" description:"周一至周日"`
	CreateTime int       `orm:"column(create_time);null" description:"创建时间"`
	UpdateTime int       `orm:"column(update_time);null" description:"更新时间"`
	Fooddate   time.Time `orm:"column(fooddate);type(date);null" description:"日期"`
	Foodname   string    `orm:"column(foodname);size(255);null" description:"食物名称"`
	Indexid    int       `orm:"column(indexid);null" description:"食物ID，对应到食物数据表的主键"`
	Rule       string    `orm:"column(rule);size(255)" description:"rule name"`
	Num        float64   `orm:"column(num);digits(10);decimals(2)" description:"数量，单位为克"`
	Goodorbad  int       `orm:"column(goodorbad);null" description:"0表示用户不喜欢，1表示用户不喜欢"`
	IsPurchase int       `orm:"column(isPurchase);null" description:"0表示用户不购买，1表示用户购买"`
	Unit       string    `orm:"column(unit);size(255);null"`
	Unitweight float64   `orm:"column(unitweight);null"`
}

func (t *Cookbook) TableName() string {
	return "cookbook"
}

func init() {
	orm.RegisterModel(new(Cookbook))
}

// AddCookbook insert a new Cookbook into database and returns
// last inserted Id on success.
func AddCookbook(m *Cookbook) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCookbookById retrieves Cookbook by Id. Returns error if
// Id doesn't exist
func GetCookbookById(id int) (v *Cookbook, err error) {
	o := orm.NewOrm()
	v = &Cookbook{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCookbook retrieves all Cookbook matches certain condition. Returns empty list if
// no records exist
func GetAllCookbook(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cookbook))
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

	var l []Cookbook
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

// UpdateCookbook updates Cookbook by Id and returns error if
// the record to be updated doesn't exist
func UpdateCookbookById(m *Cookbook) (err error) {
	o := orm.NewOrm()
	v := Cookbook{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCookbook deletes Cookbook by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCookbook(id int) (err error) {
	o := orm.NewOrm()
	v := Cookbook{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Cookbook{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
