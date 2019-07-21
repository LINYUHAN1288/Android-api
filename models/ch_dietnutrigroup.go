package models

type ChDietnutrigroup struct {
	Id_RENAME int    `orm:"column(id)" description:"食品ID"`
	Groupid   string `orm:"column(groupid);size(30)" description:"食品分组ID"`
	Dietid    string `orm:"column(dietid);size(120)" description:"食品编码"`
}
