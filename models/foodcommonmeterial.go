package models

type Foodcommonmeterial struct {
	Foodname string `orm:"column(foodname);size(255);null" description:"材料你名称"`
	Type     int8   `orm:"column(type);null" description:"0-蔬菜,1-家畜类,2-家禽类3-鱼，4-其他海产品"`
	Memo     string `orm:"column(memo);size(255);null" description:"备注"`
}
