package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type UserGoal struct {
	Id             int     `orm:"column(id);auto" description:"id "`
	Uid            int     `orm:"column(uid)" description:" 用户ID"`
	Name           string  `orm:"column(name);size(200);null" description:"目标名称"`
	Type           int8    `orm:"column(type);null" description:"目标类型 1-目标， 2-监控"`
	ImgUrl         string  `orm:"column(img_url);size(255);null" description:"图片地址"`
	Switch         int8    `orm:"column(switch);null" description:"目标是否满足"`
	Unit           string  `orm:"column(unit);size(255);null"`
	Lastid         int     `orm:"column(lastid);null"`
	Start          float64 `orm:"column(start);null" description:"开始值"`
	End            float64 `orm:"column(end);null" description:"目标值"`
	WeeklyChange   float64 `orm:"column(weekly_change);null" description:"每周变化"`
	Increase       int8    `orm:"column(increase);null" description:"0-表示增加 1-表示减少"`
	Date           int     `orm:"column(date);null" description:"截止日期"`
	Display        int     `orm:"column(display);null" description:"是什么"`
	MinValue       float64 `orm:"column(min_value);null" description:"Min Value 这里的微量元素都是mg为单位"`
	MaxValue       float64 `orm:"column(max_value);null" description:"Max Value 以毫克为单位"`
	DailyEntry     float64 `orm:"column(daily_entry);null" description:"推荐的能量摄入数量"`
	DailyFoodpoint float64 `orm:"column(daily_foodpoint);null" description:"推荐的食物份数摄入量"`
	CreateTime     int     `orm:"column(create_time);null" description:"创建时间"`
	UpdateTime     int     `orm:"column(update_time);null" description:"更新时间"`
}

func (t *UserGoal) TableName() string {
	return "user_goal"
}

func init() {
	orm.RegisterModel(new(UserGoal))
}

// AddUserGoal insert a new UserGoal into database and returns
// last inserted Id on success.
func AddUserGoal(m *UserGoal) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserGoalById retrieves UserGoal by Id. Returns error if
// Id doesn't exist
func GetUserGoalById(id int) (v *UserGoal, err error) {
	o := orm.NewOrm()
	v = &UserGoal{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserGoal retrieves all UserGoal matches certain condition. Returns empty list if
// no records exist
func GetAllUserGoal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserGoal))
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

	var l []UserGoal
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

// UpdateUserGoal updates UserGoal by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserGoalById(m *UserGoal) (err error) {
	o := orm.NewOrm()
	v := UserGoal{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserGoal deletes UserGoal by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserGoal(id int) (err error) {
	o := orm.NewOrm()
	v := UserGoal{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserGoal{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
