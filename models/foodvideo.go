package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Foodvideo struct {
	Id             int    `orm:"column(id);auto"`
	Chinesename    string `orm:"column(chinesename);size(255);null" description:"食物的中文名"`
	Imgurl         string `orm:"column(imgurl);size(255);null" description:"食物的图片地址"`
	Catid          string `orm:"column(catid);size(50);null" description:"食物的原网页id"`
	Videourl       string `orm:"column(videourl);size(255);null" description:"食物视屏自动播放的Url"`
	Sourceurl      string `orm:"column(sourceurl);size(255);null" description:"食物视屏的Url"`
	Catname        string `orm:"column(catname);size(255);null" description:"视屏的标记"`
	Source         string `orm:"column(source);size(50);null" description:"原网站网页名称"`
	Videosourceurl string `orm:"column(videosourceurl);size(255);null" description:"下载链接"`
	Playerurl      string `orm:"column(playerurl);size(255);null" description:"501服务器地址"`
}

func (t *Foodvideo) TableName() string {
	return "foodvideo"
}

func init() {
	orm.RegisterModel(new(Foodvideo))
}

// AddFoodvideo insert a new Foodvideo into database and returns
// last inserted Id on success.
func AddFoodvideo(m *Foodvideo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFoodvideoById retrieves Foodvideo by Id. Returns error if
// Id doesn't exist
func GetFoodvideoById(id int) (v *Foodvideo, err error) {
	o := orm.NewOrm()
	v = &Foodvideo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFoodvideo retrieves all Foodvideo matches certain condition. Returns empty list if
// no records exist
func GetAllFoodvideo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Foodvideo))
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

	var l []Foodvideo
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

// UpdateFoodvideo updates Foodvideo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFoodvideoById(m *Foodvideo) (err error) {
	o := orm.NewOrm()
	v := Foodvideo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFoodvideo deletes Foodvideo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFoodvideo(id int) (err error) {
	o := orm.NewOrm()
	v := Foodvideo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Foodvideo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
