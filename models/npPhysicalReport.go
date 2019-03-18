package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type NpPhysicalReport struct {
	Id                            int    `orm:"column(id_physical_report);pk,auto"`    //体检报告编号
	Id_physical_person            int    `orm:"column(id_physical_person)"`            //患者编号
	Name                          string `orm:"column(name)"`                          //患者姓名
	Gender                        int    `orm:"column(gender)"`                        //患者性别
	Id_hospital_department_doctor int    `orm:"column(id_hospital_department_doctor)"` //医师编号
	Age                           int    `orm:"column(age)"`                           //体检时年龄
	Medical_history               string `orm:"column(medical_history)"`               //既往病史
	Cretate_time                  int64  `orm:"column(cretate_time)"`                  //体检时间
	Report_time                   int64  `orm:"column(report_time)"`                   //体检报告生成时间
	Result                        string `orm:"column(result)"`                        //体检结果
	Status                        int    `orm:"column(status)"`                        //报告状态
	Update_time                   int64  `orm:"column(update_time)"`                   //修改时间
	Create_man                    string `orm:"column(create_man)"`                    //录入人
}

// 自定义表名(系统自动调用)
func (u *NpPhysicalReport) TableName() string {
	return "np_physical_report"
}

func init() {
	orm.RegisterModel(new(NpPhysicalReport)) //把NpPhysicalReport注册到orm中
}

func AddCaseInfo(npr NpPhysicalReport) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(&npr) //向数据库中插入记录
	if err != nil {
		fmt.Println("number=" + strconv.Itoa(int(id)))
		fmt.Println("AddUserInfo Error:" + err.Error())
		return id, err
	}
	return id, nil
}

//根据条件查询病例列表
func GetCaseList(strWhere string) ([]NpPhysicalReport, error) {
	var rList []NpPhysicalReport
	var strSql string = "select * from np_physical_report where 1 = 1 "
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql).QueryRows(&rList)
	if err != nil {
		return rList, err
	}
	return rList, nil
}

//查询数据总数
func GetCaseListCount(strWhere string) ([]NpPhysicalReport, error) {
	var rList []NpPhysicalReport
	var strSql string = "select * from np_physical_report where status = 1"
	//if strWhere != "" {
	//	strSql += " and " + strWhere
	//}
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql).QueryRows(&rList)
	if err != nil {
		return rList, err
	}
	return rList, nil
}

//删除病例信息
func DelCaseInfo(reportId int) (error) {
	//var rList []NpPhysicalReport
	var strSql string = "update np_physical_report set status = 2 where id_physical_report = " + strconv.Itoa(reportId)
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql).Exec()
	if err != nil {
		return err
	}
	return nil
}

//根据id查询订单信息
func GetCaseInfo(reportId int) ([]orm.Params, error) {
	var caseInfo []orm.Params
	var strSql string = "SELECT nhdd.id_hospital hospital," +
		"nhdd.id_hospital_department department," +
		"nhdd.id_hospital_doctor doctor," +
		"npp.customer_id customerid," +
		"npr.name name," +
		"npr.gender gender," +
		"npr.age age," +
		"npr.medical_history history," +
		"npr.result result," +
		"npr.cretate_time cretatetime," +
		"npr.report_time reporttime," +
		"npr.create_man man " +
		"from np_physical_report npr, np_hospital_department_doctor nhdd, np_physical_person npp " +
		"WHERE npr.id_hospital_department_doctor = nhdd.id_hospital_doctor " +
		"AND npp.id_physical_person = npr.id_physical_person " +
		"AND npr.id_physical_report = ?"
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql, reportId).Values(&caseInfo)
	if err != nil {
		return caseInfo, err
	}
	return caseInfo, nil
}
