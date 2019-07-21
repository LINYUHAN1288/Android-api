package models

type Foodsynonym struct {
	Foodname1 string `orm:"column(foodname1);size(255)" description:"食物名称"`
	Foodname2 string `orm:"column(foodname2);size(255)" description:"同义词"`
	Memo      string `orm:"column(memo);size(255);null" description:"备注"`
}
