package models

import "github.com/astaxie/beego/orm"

type NpPhysicalSubject struct {
	Id_physical_subject int64  `orm:"column(id_physical_subject);pk,auto"` //科目编号
	Subject             string `orm:"column(subject)"`                     //科目名称
}

// 自定义表名(系统自动调用)
func (u *NpPhysicalSubject) TableName() string {
	return "np_physical_subject"
}

func GetSubjectList(strWhere string) ([]NpPhysicalSubject, error) {
	var sList []NpPhysicalSubject
	var strSql string = "select * from np_physical_subject where 1=1 "
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	strSql += " order by id_physical_subject "
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql).QueryRows(&sList)
	if err != nil {
		return sList, err
	}
	return sList, nil
}
