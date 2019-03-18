package models

import "github.com/astaxie/beego/orm"

type NpHospitalDoctor struct {
	Id_hospital_doctor  int64  `orm:"column(id_hospital_doctor)";pk,auto` //医生编号
	Name                string `orm:"column(name)"`                       //姓名
	Detail_link         string `orm:"column(detail_link)"`                //对应文章链接
	Id_physical_subject string `orm:"column(id_physical_subject)"`        //对应科目编号
}

// 自定义表名(系统自动调用)
func (u *NpHospitalDoctor) TableName() string {
	return "np_hospital_doctor"
}

func GetDoctorList(strWhere string) ([]NpHospitalDoctor, error) {
	var dList []NpHospitalDoctor
	var strSql string = "select doctor.* from np_hospital_doctor doctor,np_hospital_department_doctor department where department.id_hospital_doctor = doctor.id_hospital_doctor"
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql).QueryRows(&dList)
	if err != nil {
		return dList, err
	}
	return dList, nil
}
