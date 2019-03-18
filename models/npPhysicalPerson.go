package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type NpPhysicalPerson struct {
	Id_physical_person int64  `orm:"column(id_physical_person);pk,auto"` //患者编号
	Customer_id        string `orm:"column(customer_id)"`                //会员编号
	Name               string `orm:"column(name)"`                       //姓名
	Gender             int64  `orm:"column(gender)"`                     //性别
	Age                int64  `orm:"column(age)"`                        //年龄
	Register_time      int64  `orm:"column(register_time)"`              //注册时间
	Card_id            string `orm:"column(card_id)"`                    //注册时间
	Birthday           string `orm:"column(birthday)"`                   //出生日期
}

// 自定义表名(系统自动调用)
func (u *NpPhysicalPerson) TableName() string {
	return "np_physical_person"
}

func GetMemberInfo(strWhere string) (NpPhysicalPerson, error) {
	var memInfo NpPhysicalPerson
	var strSql string = "select * from np_physical_person where 1=1 "
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	//strSql += " order by id_hospital_department "
	o := orm.NewOrm()
	o.Using("default")
	err := o.Raw(strSql).QueryRow(&memInfo)
	if err != nil {
		fmt.Println(err.Error() == "<QuerySeter> no row found")
		return memInfo, err
	}
	return memInfo, nil
}

func GetMemberNo(memId string) ([]orm.Params, error) {
	var mList []orm.Params
	var strSql string = "select customer_id from np_physical_person where id_physical_person = ?"
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql, memId).Values(&mList)
	if err != nil {
		return mList, err
	}
	return mList, nil
}
