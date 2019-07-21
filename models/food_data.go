package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type FoodData struct {
	Id          int     `orm:"column(id);auto" description:"id "`
	Uid         int     `orm:"column(uid)" description:" 用户ID"`
	Name        string  `orm:"column(name);size(200)" description:"食物名称"`
	Foodid      int     `orm:"column(foodid)" description:"ch_dietnutri的id"`
	Type        int8    `orm:"column(type);null" description:"早中晚"`
	ImgUrl      string  `orm:"column(img_url);null" description:"图片地址"`
	Switch      int8    `orm:"column(switch);null" description:"食物是否删除 0，标示删除"`
	Value       float64 `orm:"column(value);null" description:"数值"`
	Energy      float64 `orm:"column(energy);null" description:"消耗能量"`
	Unit        string  `orm:"column(unit);size(200);null" description:"单位"`
	Date        int     `orm:"column(date);null" description:"日期"`
	Source      int8    `orm:"column(source);null" description:"数据读取来源 1 －手动 2-读图片；3-语音输入"`
	Location    string  `orm:"column(location);null" description:"地点位置"`
	CreateTime  int     `orm:"column(create_time);null" description:"创建时间"`
	UpdateTime  int     `orm:"column(update_time);null" description:"更新时间"`
	Data1       string  `orm:"column(data1);size(255);null" description:"bak1"`
	Collection  int8    `orm:"column(collection);null" description:"collect food"`
	Diettime    int8    `orm:"column(diettime);null" description:"早中晚加餐 0－不清楚"`
	Tablesource int8    `orm:"column(tablesource)" description:"食物表来源，0为ch_dietnutri（默认），1为chinesefoodnutrifromhk"`
	Unit0       string  `orm:"column(unit0);size(50);null" description:"单位"`
	Weight0     string  `orm:"column(weight0);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit1       string  `orm:"column(unit1);size(50);null" description:"单位"`
	Weight1     string  `orm:"column(weight1);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit2       string  `orm:"column(unit2);size(50);null" description:"单位"`
	Weight2     string  `orm:"column(weight2);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit3       string  `orm:"column(unit3);size(50);null" description:"单位"`
	Weight3     string  `orm:"column(weight3);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit4       string  `orm:"column(unit4);size(50);null" description:"单位"`
	Weight4     string  `orm:"column(weight4);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit5       string  `orm:"column(unit5);size(50);null" description:"单位"`
	Weight5     string  `orm:"column(weight5);size(50);null" description:"单位对应的重量（食物可食用重量）"`
}

func (t *FoodData) TableName() string {
	return "food_data"
}

func init() {
	orm.RegisterModel(new(FoodData))
}

// AddFoodData insert a new FoodData into database and returns
// last inserted Id on success.
func AddFoodData(m *FoodData) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFoodDataById retrieves FoodData by Id. Returns error if
// Id doesn't exist
func GetFoodDataById(id int) (v *FoodData, err error) {
	o := orm.NewOrm()
	v = &FoodData{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFoodData retrieves all FoodData matches certain condition. Returns empty list if
// no records exist
func GetAllFoodData(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FoodData))
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

	var l []FoodData
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

// UpdateFoodData updates FoodData by Id and returns error if
// the record to be updated doesn't exist
func UpdateFoodDataById(m *FoodData) (err error) {
	o := orm.NewOrm()
	v := FoodData{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFoodData deletes FoodData by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFoodData(id int) (err error) {
	o := orm.NewOrm()
	v := FoodData{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FoodData{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
