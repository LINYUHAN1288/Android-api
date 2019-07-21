package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Sporttype struct {
	Id         int     `orm:"column(id);auto"`
	Uid        int     `orm:"column(uid);null" description:"用户id"`
	Name       string  `orm:"column(name)"`
	Type       int8    `orm:"column(type);null"`
	Energy     float64 `orm:"column(energy);null;digits(10);decimals(0)"`
	Unit       string  `orm:"column(unit);size(255);null"`
	Switch     int8    `orm:"column(switch);null"`
	Imgurl     string  `orm:"column(imgurl);size(255);null"`
	CreateTime int     `orm:"column(create_time);null"`
	UpdateTime int     `orm:"column(update_time);null"`
	Common     int8    `orm:"column(common);null" description:"常见运动，0表示不常见，1表示常见"`
}

func (t *Sporttype) TableName() string {
	return "sporttype"
}

func init() {
	orm.RegisterModel(new(Sporttype))
}

// AddSporttype insert a new Sporttype into database and returns
// last inserted Id on success.
func AddSporttype(m *Sporttype) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSporttypeById retrieves Sporttype by Id. Returns error if
// Id doesn't exist
func GetSporttypeById(id int) (v *Sporttype, err error) {
	o := orm.NewOrm()
	v = &Sporttype{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSporttype retrieves all Sporttype matches certain condition. Returns empty list if
// no records exist
func GetAllSporttype(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sporttype))
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

	var l []Sporttype
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

// UpdateSporttype updates Sporttype by Id and returns error if
// the record to be updated doesn't exist
func UpdateSporttypeById(m *Sporttype) (err error) {
	o := orm.NewOrm()
	v := Sporttype{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSporttype deletes Sporttype by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSporttype(id int) (err error) {
	o := orm.NewOrm()
	v := Sporttype{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Sporttype{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
