package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type UserMonitor struct {
	Id          int     `orm:"column(id);auto" description:"id "`
	Uid         int     `orm:"column(uid)" description:" 用户ID"`
	MonitorId   int     `orm:"column(monitor_id)" description:"监控ID"`
	MonitorName string  `orm:"column(monitor_name);size(200);null" description:"监控名称"`
	ValueTime   int     `orm:"column(value_time)" description:"当前时间戳"`
	Value       float64 `orm:"column(value)" description:"当前时间戳的输入的值"`
	CreateTime  int     `orm:"column(create_time)" description:"创建时间"`
	Value1      float64 `orm:"column(value1);null" description:"第二个值"`
	Comment     string  `orm:"column(comment);size(255);null"`
	UpdateTime  int     `orm:"column(update_time);null" description:"更新时间"`
}

func (t *UserMonitor) TableName() string {
	return "user_monitor"
}

func init() {
	orm.RegisterModel(new(UserMonitor))
}

// AddUserMonitor insert a new UserMonitor into database and returns
// last inserted Id on success.
func AddUserMonitor(m *UserMonitor) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserMonitorById retrieves UserMonitor by Id. Returns error if
// Id doesn't exist
func GetUserMonitorById(id int) (v *UserMonitor, err error) {
	o := orm.NewOrm()
	v = &UserMonitor{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserMonitor retrieves all UserMonitor matches certain condition. Returns empty list if
// no records exist
func GetAllUserMonitor(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserMonitor))
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

	var l []UserMonitor
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

// UpdateUserMonitor updates UserMonitor by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserMonitorById(m *UserMonitor) (err error) {
	o := orm.NewOrm()
	v := UserMonitor{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserMonitor deletes UserMonitor by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserMonitor(id int) (err error) {
	o := orm.NewOrm()
	v := UserMonitor{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserMonitor{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
