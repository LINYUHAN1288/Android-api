package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Chinesefoodnutri struct {
	FoodName       string  `orm:"column(FoodName);null"`
	FoodableWeight float64 `orm:"column(FoodableWeight);null;digits(10);decimals(0)"`
	Energy         int     `orm:"column(Energy);null"`
	Water          float64 `orm:"column(Water);null;digits(10);decimals(0)"`
	Protein        float64 `orm:"column(Protein);null;digits(10);decimals(0)"`
	Fat            float64 `orm:"column(Fat);null;digits(10);decimals(0)"`
	Fiber          float64 `orm:"column(Fiber);null;digits(10);decimals(0)"`
	Carbohydrate   float64 `orm:"column(Carbohydrate);null;digits(10);decimals(0)"`
	VitaminA       float64 `orm:"column(VitaminA);null;digits(10);decimals(0)"`
	VitaminB1      float64 `orm:"column(VitaminB1);null;digits(10);decimals(0)"`
	VitaminB2      float64 `orm:"column(VitaminB2);null;digits(10);decimals(0)"`
	Niacin         float64 `orm:"column(Niacin);null;digits(10);decimals(0)"`
	VitaminE       float64 `orm:"column(VitaminE);null;digits(10);decimals(0)"`
	SodiumNa       float64 `orm:"column(SodiumNa);null;digits(10);decimals(0)"`
	CalciumCa      float64 `orm:"column(CalciumCa);null;digits(10);decimals(0)"`
	IronFe         float64 `orm:"column(IronFe);null;digits(10);decimals(0)"`
	GroupID        string  `orm:"column(GroupID);size(30);null"`
	GroupName      string  `orm:"column(GroupName);size(120);null"`
	SubGroupID     string  `orm:"column(SubGroupID);size(30);null"`
	SubGroupName   string  `orm:"column(SubGroupName);size(120);null"`
	VitaminC       float64 `orm:"column(VitaminC);null;digits(10);decimals(0)"`
	Cholesterol    float64 `orm:"column(Cholesterol);null;digits(10);decimals(0)"`
	Id             int     `orm:"column(id);auto"`
}

func (t *Chinesefoodnutri) TableName() string {
	return "chinesefoodnutri"
}

func init() {
	orm.RegisterModel(new(Chinesefoodnutri))
}

// AddChinesefoodnutri insert a new Chinesefoodnutri into database and returns
// last inserted Id on success.
func AddChinesefoodnutri(m *Chinesefoodnutri) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChinesefoodnutriById retrieves Chinesefoodnutri by Id. Returns error if
// Id doesn't exist
func GetChinesefoodnutriById(id int) (v *Chinesefoodnutri, err error) {
	o := orm.NewOrm()
	v = &Chinesefoodnutri{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChinesefoodnutri retrieves all Chinesefoodnutri matches certain condition. Returns empty list if
// no records exist
func GetAllChinesefoodnutri(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Chinesefoodnutri))
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

	var l []Chinesefoodnutri
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

// UpdateChinesefoodnutri updates Chinesefoodnutri by Id and returns error if
// the record to be updated doesn't exist
func UpdateChinesefoodnutriById(m *Chinesefoodnutri) (err error) {
	o := orm.NewOrm()
	v := Chinesefoodnutri{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChinesefoodnutri deletes Chinesefoodnutri by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChinesefoodnutri(id int) (err error) {
	o := orm.NewOrm()
	v := Chinesefoodnutri{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Chinesefoodnutri{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
