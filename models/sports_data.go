package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SportsData struct {
	Id           int     `orm:"column(id);auto" description:"id "`
	Uid          int     `orm:"column(uid)" description:" 用户ID"`
	Name         string  `orm:"column(name);size(200)" description:"运动名称"`
	Type         int8    `orm:"column(type);null" description:"sport操作来源类型 1-系统， 2-自定义"`
	ImgUrl       string  `orm:"column(img_url);size(255);null" description:"图片地址"`
	Switch       int8    `orm:"column(switch);null" description:"运动记录是否删除 0，标示删除"`
	Value        float64 `orm:"column(value);null" description:"数值"`
	Energy       float64 `orm:"column(energy);null" description:"消耗能量"`
	Unit         string  `orm:"column(unit);size(200);null" description:"单位"`
	Date         int     `orm:"column(date);null" description:"日期"`
	Source       int8    `orm:"column(source);null" description:"数据读取来源 1 －手动 2-读仪器"`
	CreateTime   int     `orm:"column(create_time);null" description:"创建时间"`
	UpdateTime   int     `orm:"column(update_time);null" description:"更新时间"`
	Isstepupdate int8    `orm:"column(isstepupdate);null"`
	Step         int     `orm:"column(step);null"`
}

func (t *SportsData) TableName() string {
	return "sports_data"
}

func init() {
	orm.RegisterModel(new(SportsData))
}

// AddSportsData insert a new SportsData into database and returns
// last inserted Id on success.
func AddSportsData(m *SportsData) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSportsDataById retrieves SportsData by Id. Returns error if
// Id doesn't exist
func GetSportsDataById(id int) (v *SportsData, err error) {
	o := orm.NewOrm()
	v = &SportsData{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSportsData retrieves all SportsData matches certain condition. Returns empty list if
// no records exist
func GetAllSportsData(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SportsData))
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

	var l []SportsData
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

// UpdateSportsData updates SportsData by Id and returns error if
// the record to be updated doesn't exist
func UpdateSportsDataById(m *SportsData) (err error) {
	o := orm.NewOrm()
	v := SportsData{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSportsData deletes SportsData by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSportsData(id int) (err error) {
	o := orm.NewOrm()
	v := SportsData{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SportsData{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
