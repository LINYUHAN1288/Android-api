package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Foodfrombhcp struct {
	Id            int    `orm:"column(id);auto" description:"自增ID"`
	Fid           string `orm:"column(fid);size(50);null" description:"薄荷网自增ID"`
	Code          string `orm:"column(code);size(50);null" description:"薄荷网的食物ID"`
	Name          string `orm:"column(name);size(100);null" description:"薄荷网的食物的中文名"`
	ThumbImageUrl string `orm:"column(thumb_image_url);size(100);null" description:"食物的小图片"`
	LargeImageUrl string `orm:"column(large_image_url);size(100);null" description:"食物的大图片"`
	HealthLight   string `orm:"column(health_light);size(50);null" description:"食物的警告级别是否是高热量或者高蛋白等"`
	IsLiquid      string `orm:"column(is_liquid);size(50);null" description:"食物是否是液体"`
	Weight        string `orm:"column(weight);size(50);null" description:"标准重量一般是100克"`
	Calory        string `orm:"column(calory);size(50);null" description:"卡路里一般是千卡"`
	Fat           string `orm:"column(fat);size(50);null" description:"脂肪含量"`
	Protein       string `orm:"column(protein);size(50);null" description:"蛋白质含量"`
	FiberDietary  string `orm:"column(fiber_dietary);size(50);null" description:"纤维束含量"`
	Carbohydrate  string `orm:"column(carbohydrate);size(50);null" description:"碳水化合物含量"`
	VitaminA      string `orm:"column(vitamin_a);size(50);null" description:"维生素A"`
	Thiamine      string `orm:"column(thiamine);size(50);null" description:"硫胺素维生素B1"`
	Lactoflavin   string `orm:"column(lactoflavin);size(50);null" description:"核黄素维生素B2"`
	VitaminC      string `orm:"column(vitamin_c);size(50);null" description:"维生素C"`
	VitaminE      string `orm:"column(vitamin_e);size(50);null" description:"维生素E"`
	Carotene      string `orm:"column(carotene);size(50);null" description:"胡萝卜素"`
	Niacin        string `orm:"column(niacin);size(50);null" description:"烟酸"`
	Natrium       string `orm:"column(natrium);size(50);null" description:"钠"`
	Calcium       string `orm:"column(calcium);size(50);null" description:"钙"`
	Iron          string `orm:"column(iron);size(50);null" description:"铁"`
	Kalium        string `orm:"column(kalium);size(50);null" description:"钾"`
	Iodine        string `orm:"column(iodine);size(50);null" description:"碘"`
	Zinc          string `orm:"column(zinc);size(50);null" description:"锌"`
	Selenium      string `orm:"column(selenium);size(50);null" description:"硒"`
	Magnesium     string `orm:"column(magnesium);size(50);null" description:"镁"`
	Copper        string `orm:"column(copper);size(50);null" description:"铜"`
	Manganese     string `orm:"column(manganese);size(50);null" description:"锰"`
	Cholesterol   string `orm:"column(cholesterol);size(50);null" description:"胆固醇"`
	Groupname     string `orm:"column(groupname);size(50);null" description:"类别"`
	Appraise      string `orm:"column(appraise);null" description:"评价"`
	Gi            string `orm:"column(gi);size(50);null" description:"血糖生成指数"`
	Gl            string `orm:"column(gl);size(50);null" description:"血糖贡献指数"`
	Recipe        string `orm:"column(recipe);size(50);null" description:"是否是食谱"`
	RecipeType    string `orm:"column(recipe_type);size(200);null" description:"食谱的类型"`
	Unit0         string `orm:"column(unit0);size(50);null" description:"单位1"`
	Amount0       string `orm:"column(amount0);size(50);null" description:"单位1的数量"`
	Weight0       string `orm:"column(weight0);size(50);null" description:"单位1的重量"`
	Eatweight0    string `orm:"column(eatweight0);size(50);null" description:"单位1可以吃的重量"`
	Calory0       string `orm:"column(calory0);size(50);null" description:"单位1所含的卡路里"`
	Unit1         string `orm:"column(unit1);size(50);null" description:"单位2-5以此类推"`
	Amount1       string `orm:"column(amount1);size(50);null"`
	Weight1       string `orm:"column(weight1);size(50);null"`
	Eatweight1    string `orm:"column(eatweight1);size(50);null"`
	Calory1       string `orm:"column(calory1);size(50);null"`
	Unit2         string `orm:"column(unit2);size(50);null"`
	Amount2       string `orm:"column(amount2);size(50);null"`
	Weight2       string `orm:"column(weight2);size(50);null"`
	Eatweight2    string `orm:"column(eatweight2);size(50);null"`
	Calory2       string `orm:"column(calory2);size(50);null"`
	Unit3         string `orm:"column(unit3);size(50);null"`
	Amount3       string `orm:"column(amount3);size(50);null"`
	Weight3       string `orm:"column(weight3);size(50);null"`
	Eatweight3    string `orm:"column(eatweight3);size(50);null"`
	Calory3       string `orm:"column(calory3);size(50);null"`
	Unit4         string `orm:"column(unit4);size(50);null"`
	Amount4       string `orm:"column(amount4);size(50);null"`
	Weight4       string `orm:"column(weight4);size(50);null"`
	Eatweight4    string `orm:"column(eatweight4);size(50);null"`
	Calory4       string `orm:"column(calory4);size(50);null"`
	Unit5         string `orm:"column(unit5);size(50);null"`
	Amount5       string `orm:"column(amount5);size(50);null"`
	Weight5       string `orm:"column(weight5);size(50);null"`
	Eatweight5    string `orm:"column(eatweight5);size(50);null"`
	Calory5       string `orm:"column(calory5);size(50);null"`
	Lcalory       string `orm:"column(lcalory);size(50);null" description:"是否是高能量食物"`
	Lprotein      string `orm:"column(lprotein);size(50);null" description:"是否是高蛋白食物"`
	Lcarbohydrate string `orm:"column(lcarbohydrate);size(50);null" description:"是否是高碳水化合物"`
	Lfat          string `orm:"column(lfat);size(50);null" description:"是否是高脂肪食物"`
	LfiberDietary string `orm:"column(lfiber_dietary);size(50);null" description:"是否是高纤维食物"`
	Lgi           string `orm:"column(lgi);size(50);null" description:"血糖生成指数的级别"`
	Lgl           string `orm:"column(lgl);size(50);null" description:"血糖贡献的级别"`
	Abnormal      int    `orm:"column(abnormal);null" description:"数据异常的标志"`
}

func (t *Foodfrombhcp) TableName() string {
	return "foodfrombhcp"
}

func init() {
	orm.RegisterModel(new(Foodfrombhcp))
}

// AddFoodfrombhcp insert a new Foodfrombhcp into database and returns
// last inserted Id on success.
func AddFoodfrombhcp(m *Foodfrombhcp) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFoodfrombhcpById retrieves Foodfrombhcp by Id. Returns error if
// Id doesn't exist
func GetFoodfrombhcpById(id int) (v *Foodfrombhcp, err error) {
	o := orm.NewOrm()
	v = &Foodfrombhcp{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFoodfrombhcp retrieves all Foodfrombhcp matches certain condition. Returns empty list if
// no records exist
func GetAllFoodfrombhcp(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Foodfrombhcp))
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

	var l []Foodfrombhcp
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

// UpdateFoodfrombhcp updates Foodfrombhcp by Id and returns error if
// the record to be updated doesn't exist
func UpdateFoodfrombhcpById(m *Foodfrombhcp) (err error) {
	o := orm.NewOrm()
	v := Foodfrombhcp{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFoodfrombhcp deletes Foodfrombhcp by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFoodfrombhcp(id int) (err error) {
	o := orm.NewOrm()
	v := Foodfrombhcp{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Foodfrombhcp{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
