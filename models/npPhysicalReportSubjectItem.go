package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type NpPhysicalReportSubjectItem struct {
	Id                       int64  `orm:"column(id_physical_report_subject_item);pk,auto"` //关联编号
	Id_physical_report       int64  `orm:"column(id_physical_report)"`                      //关联报告编号
	Value                    string `orm:"column(value)"`                                   //值
	Result                   string `orm:"column(result)"`                                  //单项的结果
	Doctor                   int64  `orm:"column(doctor)"`                                  //体测医师
	Id_physical_subject_item int64  `orm:"column(id_physical_subject_item)"`                //关联体检科目项目表
	Is_increase              int64  `orm:"column(is_increase)"`                             //检测结果与标准值对比
	Normal_value             string `orm:"column(normal_value)"`                            //正常值范围
	Unit                     string `orm:"column(unit)"`                                    //单位
}

// 自定义表名(系统自动调用)
func (u *NpPhysicalReportSubjectItem) TableName() string {
	return "np_physical_report_subject_item"
}

func init() {
	orm.RegisterModel(new(NpPhysicalReportSubjectItem)) //把NpPhysicalReport注册到orm中
}

func AddCaseResult(item NpPhysicalReportSubjectItem) error {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(&item) //向数据库中插入记录
	if err != nil {
		fmt.Println("number=" + strconv.Itoa(int(id)))
		fmt.Println("AddUserInfo Error:" + err.Error())
		return err
	}
	return nil
}

func SelectItems(reportId int) ([]orm.Params, error) {
	var caseInfo []orm.Params
	var strSql string = "SELECT " +
		"nprsi.id_physical_report_subject_item id, " +
		"nps.subject subject, " +
		"npsi.NAME item, " +
		"nprsi.result result, " +
		"nprsi.unit unit, " +
		"nprsi.normal_value normalValue, " +
		"nhd.NAME doctor " +
		"FROM " +
		"np_hospital_doctor nhd, " +
	/*医生*/
		"np_physical_subject_item npsi, " +
	/*项目*/
		"np_physical_report_subject_item nprsi, " +
	/*关联表*/
		"np_physical_subject nps, " +
	/*科目*/
		"np_physical_report npr /*报告*/ " +
		"WHERE " +
		"nprsi.doctor = nhd.id_hospital_doctor " +
		"AND nprsi.id_physical_report = npr.id_physical_report " +
		"AND nprsi.id_physical_subject_item = npsi.id_physical_subject_item " +
		"AND nps.id_physical_subject = npsi.id_physical_subject " +
		"AND npr.id_physical_report = ? " +
		"ORDER BY nprsi.id_physical_report_subject_item"
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw(strSql, reportId).Values(&caseInfo)
	if err != nil {
		return caseInfo, err
	}
	return caseInfo, nil
}
