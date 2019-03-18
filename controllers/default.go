package controllers

import (
	"MedicalCase/models"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Id   int
	Name string
	Age  int64
	Addr string
	Memo string
}

type TestController struct {
	beego.Controller
}

type GetMemberList struct {
	beego.Controller
}

type GetHospitalList struct {
	beego.Controller
}

type GetDepartmentList struct {
	beego.Controller
}

type GetDoctorList struct {
	beego.Controller
}

type GetSubjectList struct {
	beego.Controller
}

type GetItemList struct {
	beego.Controller
}

type GetItemsInfo struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "src/testVue.html"
}

//测试接口
func (this *TestController) Post() {
	user := User{1, "range", 23, "qiandao", ""}
	wd := this.GetString("wd")
	fmt.Println(user)
	user.Memo = wd
	this.Data["userName"] = user.Name
	this.Data["userAge"] = user.Age
	this.Data["userAddr"] = user.Addr
	this.Data["json"] = user
	this.ServeJSON()

	//this.TplName = "testVue.html"
}

//获取会员信息 Get
func (this *GetMemberList) Get() {
	var result models.ResultInfo
	var strWhere string
	strMemNo := this.GetString("memNo")
	if len(strWhere) > 0 {
		strWhere += " and "
	}
	strWhere += "customer_id = '" + strMemNo + "'"
	memInfo, err := models.GetMemberInfo(strWhere)
	if err != nil {
		result.ErrCode = -2
		result.ErrMsg = "未查询到会员信息"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = memInfo
	this.Data["json"] = result
	this.ServeJSON()
}

//获取医生列表 Get
func (this *GetDoctorList) Get() {
	var result models.ResultInfo
	var strWhere string
	departmentNo := this.GetString("departmentNo")
	hospitalId := this.GetString("hospitalId")
	if departmentNo != "" {
		strWhere += "department.id_hospital_department = " + departmentNo
	}
	if hospitalId != "" {
		strWhere += " and department.id_hospital = " + hospitalId
	}
	doctorList, err := models.GetDoctorList(strWhere)
	if err != nil {
		result.ErrCode = -3
		result.ErrMsg = "获取医师列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = doctorList
	this.Data["json"] = result
	this.ServeJSON()
}

//获取医院列表
func (this *GetHospitalList) Get() {
	var result models.ResultInfo
	var strWhere string
	hospitalList, err := models.GetHospitalList(strWhere)
	if err != nil {
		result.ErrCode = -4
		result.ErrMsg = "获取科目列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = hospitalList
	this.Data["json"] = result
	this.ServeJSON()
}

//获取科室列表
func (this *GetDepartmentList) Get() {
	var result models.ResultInfo
	var strWhere string
	hospitalNo := this.GetString("hospitalId")
	if hospitalNo != "" {
		strWhere += "a.id_hospital = " + hospitalNo
	}
	departmentList, err := models.GetDepartmentList(strWhere)
	if err != nil {
		result.ErrCode = -5
		result.ErrMsg = "获取科室列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = departmentList
	this.Data["json"] = result
	this.ServeJSON()
}

//获取科目列表
func (this *GetSubjectList) Get() {
	var result models.ResultInfo
	var strWhere string
	subjectList, err := models.GetSubjectList(strWhere)
	if err != nil {
		result.ErrCode = -4
		result.ErrMsg = "获取科目列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = subjectList
	this.Data["json"] = result
	this.ServeJSON()
}

//获取项目列表
func (this *GetItemList) Get() {
	var result models.ResultInfo
	var strWhere string
	subNo := this.GetString("subNo")
	if subNo != "" {
		strWhere += "id_physical_subject = " + subNo
	}
	doctorList, err := models.GetItemsList(strWhere)
	if err != nil {
		result.ErrCode = -5
		result.ErrMsg = "获取项目列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = doctorList
	this.Data["json"] = result
	this.ServeJSON()
}

//根据病例编号查询体检项目
func (this *GetItemsInfo) Get() {
	result := models.ResultInfo{}
	//var strWhere string
	reportId, err := this.GetInt("reportId")
	if err != nil {
		result.ErrCode = -19
		result.ErrMsg = ""
		this.Data["json"] = result
		this.ServeJSON()
	}
	//strWhere += " npr.id_physical_report = " + strconv.Itoa(reportId)
	reportInfos, err1 := models.SelectItems(reportId)
	if err1 != nil {
		result.ErrCode = -20
		result.ErrMsg = ""
		this.Data["json"] = result
		this.ServeJSON()
	}
	caseInfos := make([]interface{}, 0)
	for _, val := range reportInfos {
		info := make(map[string]interface{})
		info["Id"] = val["id"]
		info["Subject"] = val["subject"]
		info["Items"] = val["item"]
		info["Result"] = val["result"]
		info["Unit"] = val["unit"]
		info["Normal_range"] = val["normalValue"]
		info["Doctor"] = val["doctor"]
		caseInfos = append(caseInfos, info)
	}

	//if len(reportInfos) > 0 {
	//	info["Id"] = reportInfos[0]["id"]
	//	info["Subject"] = reportInfos[0]["hospital"]
	//	info["Items"] = reportInfos[0]["hospital"]
	//	info["Result"] = reportInfos[0]["hospital"]
	//	info["Unit"] = reportInfos[0]["hospital"]
	//	info["Normal_range"] = reportInfos[0]["hospital"]
	//	info["Doctor"] = reportInfos[0]["hospital"]
	//}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = caseInfos
	this.Data["json"] = result
	this.ServeJSON()
	return
}
