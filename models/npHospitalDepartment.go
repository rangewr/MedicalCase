package models

import "github.com/astaxie/beego/orm"

type NpHospitalDepartment struct {
	Id_hospital_department int64  `orm:"column(id_hospital_department);pk,auto"` //科室编号
	Name                   string `orm:"column(name)"`                           //科室名称
}

// 自定义表名(系统自动调用)
func (u *NpHospitalDepartment) TableName() string {
	return "np_hospital_department"
}

func GetDepartmentList(strWhere string) ([]NpHospitalDepartment, error) {
	var dList []NpHospitalDepartment
	var strSql string = "select b.* from np_hospital_department_doctor a left join np_hospital_department b on a.id_hospital_department=b.id_hospital_department  where 1=1 "
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	strSql += " order by b.id_hospital_department "
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql).QueryRows(&dList)
	if err != nil {
		return dList, err
	}
	return dList, nil
}
