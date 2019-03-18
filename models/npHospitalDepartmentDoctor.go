package models

import "github.com/astaxie/beego/orm"

type NpHospitalDepartmentDoctor struct {
	Id_hospital_department_doctor int64 `orm:"column(id_hospital_department_doctor)";pk,auto` //关联表编号
	Id_hospital                   int64 `orm:"column(id_hospital)"`                           //医院机构编号
	Id_hospital_department        int64 `orm:"column(id_hospital_department)"`                //医院科室编号
	Id_hospital_doctor            int64 `orm:"column(id_hospital_doctor)"`                    //医师编号
	Create_time                   int64 `orm:"column(create_time)"`                           //创建时间
}

type DoctorInfo struct {
	HospitalName   string //医院名称
	DepartmentName string //部门名称
}

// 自定义表名(系统自动调用)
func (u *NpHospitalDepartmentDoctor) TableName() string {
	return "np_hospital_department_doctor"
}

func GetHospitalDepartmentDoctor(doctorId string) ([]orm.Params, error) {
	var dList []orm.Params
	var strSql string = "select nh.name hospitalName,nhd.name departmentName from np_hospital nh,np_hospital_department nhd,np_hospital_department_doctor nhdd where " +
		"nh.id_hospital = nhdd.id_hospital and nhd.id_hospital_department = nhdd.id_hospital_department and nhdd.id_hospital_doctor = ?"
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql, doctorId).Values(&dList)
	if err != nil {
		return dList, err
	}
	return dList, nil
}
