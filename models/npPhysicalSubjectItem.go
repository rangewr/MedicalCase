package models

import "github.com/astaxie/beego/orm"

type NpPhysicalSubjectItem struct {
	Id_physical_subject_item int64   `orm:"column(id_physical_subject_item);pk,auto"` //项目编号
	Id_physical_subject      int64   `orm:"column(id_physical_subject)"`              //体检科目编号
	Name                     string  `orm:"column(name)"`                             //项目名称
	Price                    float64 `orm:"column(price)"`                            //项目金额
	Normal_range             string  `orm:"column(normal_range)"`                     //正常值范围
	Unit                     string  `orm:"column(unit)"`                             //单位
}

// 自定义表名(系统自动调用)
func (u *NpPhysicalSubjectItem) TableName() string {
	return "np_physical_subject_item"
}

func GetItemsList(strWhere string) ([]NpPhysicalSubjectItem, error) {
	var iList []NpPhysicalSubjectItem
	var strSql string = "select * from np_physical_subject_item where 1=1"
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql).QueryRows(&iList)
	if err != nil {
		return iList, err
	}
	return iList, nil
}
