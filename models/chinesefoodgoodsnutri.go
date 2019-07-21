package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Chinesefoodgoodsnutri struct {
	Id             int     `orm:"column(id);auto"`
	FoodName       string  `orm:"column(FoodName);null"`
	FoodableWeight string  `orm:"column(FoodableWeight);null"`
	Energy         float64 `orm:"column(Energy);null;digits(10);decimals(2)"`
	Water          float64 `orm:"column(Water);null;digits(10);decimals(2)"`
	Protein        float64 `orm:"column(Protein);null;digits(10);decimals(2)"`
	Fat            float64 `orm:"column(Fat);null;digits(10);decimals(2)"`
	Fiber          float64 `orm:"column(Fiber);null;digits(10);decimals(2)"`
	Carbohydrate   float64 `orm:"column(Carbohydrate);null;digits(10);decimals(2)"`
	VitaminA       float64 `orm:"column(VitaminA);null;digits(10);decimals(2)"`
	VitaminB1      float64 `orm:"column(VitaminB1);null;digits(10);decimals(2)"`
	VitaminB2      float64 `orm:"column(VitaminB2);null;digits(10);decimals(2)"`
	Niacin         float64 `orm:"column(Niacin);null;digits(10);decimals(2)"`
	VitaminE       float64 `orm:"column(VitaminE);null;digits(10);decimals(2)"`
	SodiumNa       float64 `orm:"column(SodiumNa);null;digits(10);decimals(2)"`
	CalciumCa      float64 `orm:"column(CalciumCa);null;digits(10);decimals(2)"`
	IronFe         float64 `orm:"column(IronFe);null;digits(10);decimals(2)"`
	GroupID        string  `orm:"column(GroupID);size(10);null"`
	VitaminC       float64 `orm:"column(VitaminC);null;digits(10);decimals(2)"`
	Cholesterol    float64 `orm:"column(Cholesterol);null;digits(10);decimals(2)"`
	Huluobushu     float64 `orm:"column(huluobushu);null;digits(10);decimals(2)"`
	Liuanshu       float64 `orm:"column(liuanshu);null;digits(10);decimals(2)"`
	Hehaungshu     float64 `orm:"column(hehaungshu);null;digits(10);decimals(2)"`
	Yanshuan       float64 `orm:"column(yanshuan);null;digits(10);decimals(2)"`
	Mei            float64 `orm:"column(mei);null;digits(10);decimals(2)"`
	Tong           float64 `orm:"column(tong);null;digits(10);decimals(2)"`
	Xin            float64 `orm:"column(xin);null;digits(10);decimals(2)"`
	Meng           float64 `orm:"column(meng);null;digits(10);decimals(2)"`
	Jia            float64 `orm:"column(jia);null;digits(10);decimals(2)"`
	Lin            float64 `orm:"column(lin);null;digits(10);decimals(2)"`
	Xi             float64 `orm:"column(xi);null;digits(10);decimals(2)"`
	Unit0          string  `orm:"column(unit0);null"`
	Unit1          string  `orm:"column(unit1);null"`
	Unit2          string  `orm:"column(unit2);null"`
	Unit3          string  `orm:"column(unit3);null"`
	Category       string  `orm:"column(category);null"`
	Value          string  `orm:"column(value);null"`
	Evaluation     string  `orm:"column(evaluation);null"`
	Imgurl         string  `orm:"column(imgurl);null"`
}

func (t *Chinesefoodgoodsnutri) TableName() string {
	return "chinesefoodgoodsnutri"
}

func init() {
	orm.RegisterModel(new(Chinesefoodgoodsnutri))
}

// AddChinesefoodgoodsnutri insert a new Chinesefoodgoodsnutri into database and returns
// last inserted Id on success.
func AddChinesefoodgoodsnutri(m *Chinesefoodgoodsnutri) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChinesefoodgoodsnutriById retrieves Chinesefoodgoodsnutri by Id. Returns error if
// Id doesn't exist
func GetChinesefoodgoodsnutriById(id int) (v *Chinesefoodgoodsnutri, err error) {
	o := orm.NewOrm()
	v = &Chinesefoodgoodsnutri{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChinesefoodgoodsnutri retrieves all Chinesefoodgoodsnutri matches certain condition. Returns empty list if
// no records exist
func GetAllChinesefoodgoodsnutri(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Chinesefoodgoodsnutri))
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

	var l []Chinesefoodgoodsnutri
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

// UpdateChinesefoodgoodsnutri updates Chinesefoodgoodsnutri by Id and returns error if
// the record to be updated doesn't exist
func UpdateChinesefoodgoodsnutriById(m *Chinesefoodgoodsnutri) (err error) {
	o := orm.NewOrm()
	v := Chinesefoodgoodsnutri{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChinesefoodgoodsnutri deletes Chinesefoodgoodsnutri by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChinesefoodgoodsnutri(id int) (err error) {
	o := orm.NewOrm()
	v := Chinesefoodgoodsnutri{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Chinesefoodgoodsnutri{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
