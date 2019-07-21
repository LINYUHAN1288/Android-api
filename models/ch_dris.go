package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ChDris struct {
	Id             int     `orm:"column(id);auto" description:"id "`
	AgeStart       float64 `orm:"column(ageStart)" description:" 两个字段组合成0-6个月、7-12个月、1-3岁、4-6岁、7-10岁、11-13岁、14-18岁、19-49岁、50-64岁、65-80岁和80岁以上"`
	AgeEnd         float64 `orm:"column(ageEnd)" description:"0.6、1、3、6、10、13、18、49、64、79以及80以上"`
	Sex            int8    `orm:"column(sex)" description:"性别 0-男 1－女"`
	NutrNo         string  `orm:"column(NutrNo);size(255);null" description:"营养素编号"`
	NutrName       string  `orm:"column(NutrName);size(255);null" description:"营养素名称"`
	NutrList       string  `orm:"column(NutrList);size(255);null" description:"营养素单位"`
	NutrValueType  int8    `orm:"column(NutrValueType);null" description:"营养素数值计量类型。当NutrValueType为0时，NutrValue字段有效；当NutrValueType为1时，NutrValueStart和NutrValueEnd有效。"`
	NutrValueRNI   float64 `orm:"column(NutrValueRNI);null" description:"营养素值参考摄入量(RNI)"`
	NutrValueUL    float64 `orm:"column(NutrValueUL);null" description:"营养素值最高摄入量(UL)"`
	NutrValueStart float64 `orm:"column(NutrValueStart);null" description:"营养素值范围1"`
	NutrValueEnd   float64 `orm:"column(NutrValueEnd);null" description:"营养素值范围2"`
	PAL            int8    `orm:"column(PAL);null" description:"身体活动水平 0-不适用 1-轻活动水平(1.5) 2-中活动水平(1.75) 3-重活动水平(2)"`
	CreateTime     int     `orm:"column(create_time);null" description:"创建时间"`
	UpdateTime     int     `orm:"column(update_time);null" description:"更新时间"`
}

func (t *ChDris) TableName() string {
	return "ch_dris"
}

func init() {
	orm.RegisterModel(new(ChDris))
}

// AddChDris insert a new ChDris into database and returns
// last inserted Id on success.
func AddChDris(m *ChDris) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChDrisById retrieves ChDris by Id. Returns error if
// Id doesn't exist
func GetChDrisById(id int) (v *ChDris, err error) {
	o := orm.NewOrm()
	v = &ChDris{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChDris retrieves all ChDris matches certain condition. Returns empty list if
// no records exist
func GetAllChDris(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ChDris))
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

	var l []ChDris
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

// UpdateChDris updates ChDris by Id and returns error if
// the record to be updated doesn't exist
func UpdateChDrisById(m *ChDris) (err error) {
	o := orm.NewOrm()
	v := ChDris{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChDris deletes ChDris by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChDris(id int) (err error) {
	o := orm.NewOrm()
	v := ChDris{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ChDris{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
