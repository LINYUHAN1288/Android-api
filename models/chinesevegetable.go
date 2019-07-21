package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Chinesevegetable struct {
	Id              int    `orm:"column(id);auto"`
	DietChineseName string `orm:"column(DietChineseName);null"`
	DietID          string `orm:"column(DietID);size(255);null"`
	FoodableWeight  string `orm:"column(FoodableWeight);null"`
	Energy          string `orm:"column(Energy);null"`
	Water           string `orm:"column(Water);null"`
	Protein         string `orm:"column(Protein);null"`
	Fat             string `orm:"column(Fat);null"`
	Fiber           string `orm:"column(Fiber);null"`
	Carbohydrate    string `orm:"column(Carbohydrate);null"`
	VitaminA        string `orm:"column(VitaminA);null"`
	VitaminB1       string `orm:"column(VitaminB1);null"`
	VitaminB2       string `orm:"column(VitaminB2);null"`
	Niacin          string `orm:"column(Niacin);null"`
	VitaminE        string `orm:"column(VitaminE);null"`
	SodiumNa        string `orm:"column(SodiumNa);null"`
	CalciumCa       string `orm:"column(CalciumCa);null"`
	IronFe          string `orm:"column(IronFe);null"`
	GroupID         string `orm:"column(GroupID);size(10);null"`
	VitaminC        string `orm:"column(VitaminC);null"`
	Cholesterol     string `orm:"column(Cholesterol);null"`
	Huluobushu      string `orm:"column(huluobushu);null"`
	Liuanshu        string `orm:"column(liuanshu);null"`
	Hehaungshu      string `orm:"column(hehaungshu);null"`
	Yanshuan        string `orm:"column(yanshuan);null"`
	MagnesiumMg     string `orm:"column(MagnesiumMg);null"`
	CopperCu        string `orm:"column(CopperCu);null"`
	ZincZn          string `orm:"column(ZincZn);null"`
	ManganeseMn     string `orm:"column(ManganeseMn);null"`
	KaliumK         string `orm:"column(KaliumK);null"`
	PhosphorP       string `orm:"column(PhosphorP);null"`
	SeleniumSe      string `orm:"column(SeleniumSe);null"`
	Unit0           string `orm:"column(unit0);null"`
	Unit1           string `orm:"column(unit1);null"`
	Unit2           string `orm:"column(unit2);null"`
	Unit3           string `orm:"column(unit3);null"`
	Category        string `orm:"column(category);null"`
	Value           string `orm:"column(value);null"`
	Evaluation      string `orm:"column(evaluation);null"`
	DietSteps       string `orm:"column(DietSteps);null"`
	ImageSrc        string `orm:"column(ImageSrc);null"`
	Meterial        string `orm:"column(Meterial);null"`
	VitaminB3       int    `orm:"column(VitaminB3);null"`
	DietType        string `orm:"column(DietType);size(255);null"`
	CreateTime      int    `orm:"column(Create_time);null"`
	UpdateTime      int    `orm:"column(Update_time);null"`
	DietWeight      string `orm:"column(DietWeight);size(30);null"`
}

func (t *Chinesevegetable) TableName() string {
	return "chinesevegetable"
}

func init() {
	orm.RegisterModel(new(Chinesevegetable))
}

// AddChinesevegetable insert a new Chinesevegetable into database and returns
// last inserted Id on success.
func AddChinesevegetable(m *Chinesevegetable) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChinesevegetableById retrieves Chinesevegetable by Id. Returns error if
// Id doesn't exist
func GetChinesevegetableById(id int) (v *Chinesevegetable, err error) {
	o := orm.NewOrm()
	v = &Chinesevegetable{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChinesevegetable retrieves all Chinesevegetable matches certain condition. Returns empty list if
// no records exist
func GetAllChinesevegetable(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Chinesevegetable))
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

	var l []Chinesevegetable
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

// UpdateChinesevegetable updates Chinesevegetable by Id and returns error if
// the record to be updated doesn't exist
func UpdateChinesevegetableById(m *Chinesevegetable) (err error) {
	o := orm.NewOrm()
	v := Chinesevegetable{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChinesevegetable deletes Chinesevegetable by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChinesevegetable(id int) (err error) {
	o := orm.NewOrm()
	v := Chinesevegetable{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Chinesevegetable{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
