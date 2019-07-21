package models

type ChineseFoodSubGroup struct {
	GroupID      string `orm:"column(GroupID);size(10)"`
	SubGroupID   string `orm:"column(SubGroupID);size(10)"`
	SubGroupName string `orm:"column(SubGroupName);size(255);null"`
}
