package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ChDietnutriMeterial struct {
	GroupID         string  `orm:"column(GroupID);size(30)"`
	DietID          string  `orm:"column(DietID);size(120)"`
	DietName        string  `orm:"column(DietName);size(255);null"`
	DietChineseName string  `orm:"column(DietChineseName);size(255);null"`
	DietType        string  `orm:"column(DietType);size(255);null"`
	DietWeight      string  `orm:"column(DietWeight);size(30);null"`
	Energy          float64 `orm:"column(Energy);null;digits(10);decimals(2)"`
	Protein         float64 `orm:"column(Protein);null;digits(10);decimals(2)"`
	Carbohydrate    float64 `orm:"column(Carbohydrate);null;digits(10);decimals(2)"`
	Fat             float64 `orm:"column(Fat);null;digits(10);decimals(2)"`
	Fiber           float64 `orm:"column(Fiber);null;digits(10);decimals(2)"`
	Sugar           float64 `orm:"column(Sugar);null;digits(10);decimals(2)"`
	VitaminC        float64 `orm:"column(VitaminC);null;digits(10);decimals(2)"`
	DietSteps       string  `orm:"column(DietSteps);size(8000);null"`
	Meterial        string  `orm:"column(Meterial);null" description:"食材"`
	ImageSrc        string  `orm:"column(ImageSrc);null"`
	Source          uint    `orm:"column(Source);null"`
	VitaminA        float64 `orm:"column(VitaminA);null;digits(10);decimals(2)"`
	VitaminE        float64 `orm:"column(VitaminE);null;digits(10);decimals(2)"`
	VitaminB1       float64 `orm:"column(VitaminB1);null;digits(10);decimals(2)"`
	VitaminB2       float64 `orm:"column(VitaminB2);null;digits(10);decimals(2)"`
	VitaminB3       float64 `orm:"column(VitaminB3);null;digits(10);decimals(2)"`
	Cholesterol     float64 `orm:"column(Cholesterol);null;digits(10);decimals(2)"`
	MagnesiumMg     float64 `orm:"column(MagnesiumMg);null;digits(10);decimals(2)"`
	CalciumCa       float64 `orm:"column(CalciumCa);null;digits(10);decimals(2)"`
	IronFe          float64 `orm:"column(IronFe);null;digits(10);decimals(2)"`
	ZincZn          float64 `orm:"column(ZincZn);null;digits(10);decimals(2)"`
	CopperCu        float64 `orm:"column(CopperCu);null;digits(10);decimals(2)"`
	ManganeseMn     float64 `orm:"column(ManganeseMn);null;digits(10);decimals(2)"`
	KaliumK         float64 `orm:"column(KaliumK);null;digits(10);decimals(2)"`
	PhosphorP       float64 `orm:"column(PhosphorP);null;digits(10);decimals(2)"`
	SodiumNa        float64 `orm:"column(SodiumNa);null;digits(10);decimals(2)"`
	SeleniumSe      float64 `orm:"column(SeleniumSe);null;digits(10);decimals(2)"`
	Id              int     `orm:"column(id);auto"`
	CreateTime      int     `orm:"column(Create_time);null"`
	UpdateTime      int     `orm:"column(Update_time);null"`
	BarCode         string  `orm:"column(BarCode);size(50);null" description:"条码"`
	Evaluation      string  `orm:"column(Evaluation);size(255);null"`
	Brand           string  `orm:"column(Brand);size(255)"`
	Unit0           string  `orm:"column(unit0);size(50);null" description:"单位"`
	Weight0         string  `orm:"column(weight0);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit1           string  `orm:"column(unit1);size(50);null" description:"单位"`
	Weight1         string  `orm:"column(weight1);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit2           string  `orm:"column(unit2);size(50);null" description:"单位"`
	Weight2         string  `orm:"column(weight2);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit3           string  `orm:"column(unit3);size(50);null" description:"单位"`
	Weight3         string  `orm:"column(weight3);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit4           string  `orm:"column(unit4);size(50);null" description:"单位"`
	Weight4         string  `orm:"column(weight4);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Unit5           string  `orm:"column(unit5);size(50);null" description:"单位"`
	Weight5         string  `orm:"column(weight5);size(50);null" description:"单位对应的重量（食物可食用重量）"`
	Carotenes       float64 `orm:"column(Carotenes);null" description:"胡萝卜素"`
	Niacin          float64 `orm:"column(Niacin);null" description:"烟酸"`
	Iodine          float64 `orm:"column(Iodine);null" description:"碘"`
	HealthLight     string  `orm:"column(HealthLight);size(50);null" description:"健康等级"`
	HealthAppraise  string  `orm:"column(HealthAppraise);size(255);null" description:"食物的健康评价"`
	GI              float64 `orm:"column(GI);null;digits(10);decimals(2)" description:"血糖生成指数GI"`
	GL              float64 `orm:"column(GL);null;digits(10);decimals(2)" description:"血糖生成指数GL"`
	HEnergy         string  `orm:"column(HEnergy);size(50);null" description:"高能量食物标志"`
	HProtein        string  `orm:"column(HProtein);size(50);null" description:"高蛋白质食物标志"`
	HFat            string  `orm:"column(HFat);size(50);null" description:"高脂肪食物标志"`
	HCarbohydrat    string  `orm:"column(HCarbohydrat);size(50);null" description:"高碳水化合物食物标志"`
	HFiber          string  `orm:"column(HFiber);size(50);null" description:"高纤维食物标志"`
	HGL             string  `orm:"column(HGL);size(50);null" description:"高血糖食物标志GL"`
	HGI             string  `orm:"column(HGI);size(50);null" description:"高血糖食物标志GI"`
	Common          int8    `orm:"column(Common);null" description:"常见食物，0表示不常见，1表示常见"`
	Purchase        int8    `orm:"column(Purchase);null" description:"是否可购买，0表示可购买，1表示不可购买"`
}

func (t *ChDietnutriMeterial) TableName() string {
	return "ch_dietnutri_meterial"
}

func init() {
	orm.RegisterModel(new(ChDietnutriMeterial))
}

// AddChDietnutriMeterial insert a new ChDietnutriMeterial into database and returns
// last inserted Id on success.
func AddChDietnutriMeterial(m *ChDietnutriMeterial) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChDietnutriMeterialById retrieves ChDietnutriMeterial by Id. Returns error if
// Id doesn't exist
func GetChDietnutriMeterialById(id int) (v *ChDietnutriMeterial, err error) {
	o := orm.NewOrm()
	v = &ChDietnutriMeterial{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChDietnutriMeterial retrieves all ChDietnutriMeterial matches certain condition. Returns empty list if
// no records exist
func GetAllChDietnutriMeterial(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ChDietnutriMeterial))
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

	var l []ChDietnutriMeterial
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

// UpdateChDietnutriMeterial updates ChDietnutriMeterial by Id and returns error if
// the record to be updated doesn't exist
func UpdateChDietnutriMeterialById(m *ChDietnutriMeterial) (err error) {
	o := orm.NewOrm()
	v := ChDietnutriMeterial{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChDietnutriMeterial deletes ChDietnutriMeterial by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChDietnutriMeterial(id int) (err error) {
	o := orm.NewOrm()
	v := ChDietnutriMeterial{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ChDietnutriMeterial{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
