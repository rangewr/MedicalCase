package models

import "github.com/astaxie/beego/orm"

type NpHospital struct {
	Id_hospital int64  `orm:"column(id_hospital);pk,auto"` //医院编号
	Name        string `orm:"column(name)"`                //医院名称
}

// 自定义表名(系统自动调用)
func (u *NpHospital) TableName() string {
	return "np_hospital"
}

func GetHospitalList(strWhere string) ([]NpHospital, error) {
	var hList []NpHospital
	var strSql string = "select * from np_hospital where 1=1 "
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	strSql += " order by id_hospital "
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql).QueryRows(&hList)
	if err != nil {
		return hList, err
	}
	return hList, nil
}
