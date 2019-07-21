package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Chinesefoodnutrifromhk struct {
	GroupID         string  `orm:"column(GroupID);size(10)"`
	SubGroupID      string  `orm:"column(SubGroupID);size(10)"`
	FoodID          string  `orm:"column(FoodID);size(30)"`
	FoodChineseName string  `orm:"column(FoodChineseName);size(255);null"`
	FoodName        string  `orm:"column(FoodName);size(255);null"`
	FoodAlias       string  `orm:"column(FoodAlias);size(255);null"`
	DataSource      string  `orm:"column(DataSource);size(4);null"`
	FoodableWeight  string  `orm:"column(FoodableWeight);size(30);null"`
	Energy          float64 `orm:"column(Energy);null;digits(10);decimals(0)"`
	Water           float64 `orm:"column(Water);null;digits(10);decimals(0)"`
	Protein         float64 `orm:"column(Protein);null;digits(10);decimals(0)"`
	Carbohydrate    float64 `orm:"column(Carbohydrate);null;digits(10);decimals(0)"`
	Fat             float64 `orm:"column(Fat);null;digits(10);decimals(0)"`
	Fiber           float64 `orm:"column(Fiber);null;digits(10);decimals(0)"`
	Sugar           float64 `orm:"column(Sugar);null;digits(10);decimals(0)"`
	SaturatedFat    float64 `orm:"column(SaturatedFat);null;digits(10);decimals(0)"`
	UnSaturatedFat  float64 `orm:"column(UnSaturatedFat);null;digits(10);decimals(0)"`
	Cholesterol     float64 `orm:"column(Cholesterol);null;digits(10);decimals(0)"`
	CalciumCa       float64 `orm:"column(CalciumCa);null;digits(10);decimals(0)"`
	CopperCu        float64 `orm:"column(CopperCu);null;digits(10);decimals(0)"`
	IronFe          float64 `orm:"column(IronFe);null;digits(10);decimals(0)"`
	MagnesiumMg     float64 `orm:"column(MagnesiumMg);null;digits(10);decimals(0)"`
	ManganeseMn     float64 `orm:"column(ManganeseMn);null;digits(10);decimals(0)"`
	PhosphorP       float64 `orm:"column(PhosphorP);null;digits(10);decimals(0)"`
	KaliumK         float64 `orm:"column(KaliumK);null;digits(10);decimals(0)"`
	SodiumNa        float64 `orm:"column(SodiumNa);null;digits(10);decimals(0)"`
	ZincZn          float64 `orm:"column(ZincZn);null;digits(10);decimals(0)"`
	VitaminA        float64 `orm:"column(VitaminA);null;digits(10);decimals(0)"`
	VitaminB1       float64 `orm:"column(VitaminB1);null;digits(10);decimals(0)"`
	VitaminB2       float64 `orm:"column(VitaminB2);null;digits(10);decimals(0)"`
	VitaminE        float64 `orm:"column(VitaminE);null;digits(10);decimals(0)"`
	VitaminC        float64 `orm:"column(VitaminC);null;digits(10);decimals(0)"`
	UpdatedFlag     int     `orm:"column(updatedFlag);null"`
	Id              int     `orm:"column(id);auto"`
	CreateTime      int     `orm:"column(create_time);null"`
	UpdateTime      int     `orm:"column(update_time);null"`
}

func (t *Chinesefoodnutrifromhk) TableName() string {
	return "chinesefoodnutrifromhk"
}

func init() {
	orm.RegisterModel(new(Chinesefoodnutrifromhk))
}

// AddChinesefoodnutrifromhk insert a new Chinesefoodnutrifromhk into database and returns
// last inserted Id on success.
func AddChinesefoodnutrifromhk(m *Chinesefoodnutrifromhk) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChinesefoodnutrifromhkById retrieves Chinesefoodnutrifromhk by Id. Returns error if
// Id doesn't exist
func GetChinesefoodnutrifromhkById(id int) (v *Chinesefoodnutrifromhk, err error) {
	o := orm.NewOrm()
	v = &Chinesefoodnutrifromhk{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChinesefoodnutrifromhk retrieves all Chinesefoodnutrifromhk matches certain condition. Returns empty list if
// no records exist
func GetAllChinesefoodnutrifromhk(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Chinesefoodnutrifromhk))
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

	var l []Chinesefoodnutrifromhk
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

// UpdateChinesefoodnutrifromhk updates Chinesefoodnutrifromhk by Id and returns error if
// the record to be updated doesn't exist
func UpdateChinesefoodnutrifromhkById(m *Chinesefoodnutrifromhk) (err error) {
	o := orm.NewOrm()
	v := Chinesefoodnutrifromhk{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChinesefoodnutrifromhk deletes Chinesefoodnutrifromhk by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChinesefoodnutrifromhk(id int) (err error) {
	o := orm.NewOrm()
	v := Chinesefoodnutrifromhk{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Chinesefoodnutrifromhk{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
