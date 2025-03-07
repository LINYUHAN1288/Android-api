package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type News struct {
	Id         int    `orm:"column(id);auto" description:"id "`
	Uid        int    `orm:"column(uid)" description:"用户id"`
	Foodname   string `orm:"column(foodname);size(200);null" description:"食物名"`
	Url        string `orm:"column(url)" description:"新闻图片"`
	Title      string `orm:"column(title);size(255)" description:"新闻标题"`
	ImgUrl     string `orm:"column(imgUrl);size(255);null" description:"新闻图片"`
	Cause      string `orm:"column(cause);size(255);null" description:"起因"`
	Source     string `orm:"column(source);size(100);null" description:"新闻来源"`
	Date       string `orm:"column(date);size(50);null" description:"新闻发布时间"`
	Address    string `orm:"column(address);size(50);null" description:"新闻发生地"`
	Collect    int8   `orm:"column(collect);null" description:"0收藏，1表示不收藏"`
	CreateTime int    `orm:"column(create_time);null" description:"创建时间"`
	UpdateTime int    `orm:"column(update_time);null" description:"更新时间"`
}

func (t *News) TableName() string {
	return "news"
}

func init() {
	orm.RegisterModel(new(News))
}

// AddNews insert a new News into database and returns
// last inserted Id on success.
func AddNews(m *News) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNewsById retrieves News by Id. Returns error if
// Id doesn't exist
func GetNewsById(id int) (v *News, err error) {
	o := orm.NewOrm()
	v = &News{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNews retrieves all News matches certain condition. Returns empty list if
// no records exist
func GetAllNews(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(News))
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

	var l []News
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

// UpdateNews updates News by Id and returns error if
// the record to be updated doesn't exist
func UpdateNewsById(m *News) (err error) {
	o := orm.NewOrm()
	v := News{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNews deletes News by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNews(id int) (err error) {
	o := orm.NewOrm()
	v := News{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&News{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
