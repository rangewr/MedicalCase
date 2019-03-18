package controllers

import (
	"MedicalCase/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
)

type GetCaseList struct {
	beego.Controller
}

type AddCaseInfo struct {
	beego.Controller
}

type DelCaseInfo struct {
	beego.Controller
}

type GetCaseInfo struct {
	beego.Controller
}

type UpdateCaseInfo struct {
	beego.Controller
}

//获取病例列表
func (this *GetCaseList) Get() {
	var result models.ResultInfo
	var strWhere string
	pageNo, err1 := this.GetInt("pageNo")
	pageSize, err2 := this.GetInt("pageSize")
	if err1 != nil || err2 != nil {
		result.ErrCode = -9
		result.ErrMsg = "获取参数失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	strWhere += " status = 1 "
	if pageNo > 0 && pageSize > 0 {
		strWhere += " order by update_time limit " + (strconv.Itoa((pageNo - 1) * pageSize)) + "," + strconv.Itoa(pageSize)
	}
	//查询病例列表
	rList, err3 := models.GetCaseList(strWhere)
	if err3 != nil {
		result.ErrCode = -10
		result.ErrMsg = "获取病例列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	caseInfos := make([]interface{}, 0)
	for _, val := range rList {
		info := make(map[string]interface{})
		info["Id"] = val.Id
		info["Update_time"] = time.Unix(val.Update_time, 0).Format("2006-01-02 03:04:05")
		info["Customer_name"] = val.Name
		info["Customer_gender"] = val.Gender
		info["Customer_age"] = val.Age
		info["Cretate_time"] = time.Unix(val.Report_time, 0).Format("2006-01-02 03:04:05")
		var doctorId = val.Id_hospital_department_doctor
		doctorInfo, err4 := models.GetHospitalDepartmentDoctor(strconv.Itoa(doctorId))
		if err4 != nil {
			result.ErrCode = -11
			result.ErrMsg = "获取医师部门医院失败"
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
		info["Hospital_name"] = doctorInfo[0]["hospitalName"]
		info["Hospital_department_name"] = doctorInfo[0]["departmentName"]
		//todo 查询会员编号
		Customer_id, err5 := models.GetMemberNo(strconv.Itoa(val.Id_physical_person))
		if err5 != nil {
			result.ErrCode = -12
			result.ErrMsg = "获取患者会员编号失败"
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
		if len(Customer_id) > 0 {
			info["Customer_id"] = Customer_id[0]["customer_id"]
		}
		caseInfos = append(caseInfos, &info)
	}
	listCount, err01 := models.GetCaseListCount("")
	if err01 != nil {
		result.ErrCode = -13
		result.ErrMsg = "获取医师部门医院失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	results := make(map[string]interface{})
	results["ListCount"] = len(listCount)
	results["caseInfos"] = caseInfos
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = results
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//删除一个病例
func (this *DelCaseInfo) Post() {
	var result models.ResultInfo
	reportId, err := this.GetInt("reportId")
	if err != nil {
		result.ErrCode = -14
		result.ErrMsg = "获取病例编号错误"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if reportId > 0 {
		err1 := models.DelCaseInfo(reportId)
		if err1 != nil {
			result.ErrCode = -14
			result.ErrMsg = "删除病例异常"
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
	} else {
		result.ErrCode = -14
		result.ErrMsg = "删除病例异常"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//查询一个病例详情
func (this *GetCaseInfo) Get() {
	var result models.ResultInfo
	//var strWhere string
	reportId, err := this.GetInt("reportId")
	if err != nil {
		result.ErrCode = -15
		result.ErrMsg = "获取参数异常"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	reportInfos, err1 := models.GetCaseInfo(reportId)
	if err1 != nil {
		result.ErrCode = -16
		result.ErrMsg = "获取病例详情异常"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	info := make(map[string]interface{})
	if len(reportInfos) > 0 {
		info["hospital"] = reportInfos[0]["hospital"]
		info["department"] = reportInfos[0]["department"]
		info["doctor"] = reportInfos[0]["doctor"]
		checkTime, errCheckTime := strconv.ParseInt(reportInfos[0]["cretatetime"].(string), 10, 64)
		if errCheckTime != nil {
			result.ErrCode = -17
			result.ErrMsg = "获取病例创建时间异常"
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
		info["checkTime"] = time.Unix(checkTime, 0).Format("2006-01-02 15:04:05")
		info["memNo"] = reportInfos[0]["customerid"]
		info["userName"] = reportInfos[0]["name"]
		info["userGender"] = reportInfos[0]["gender"]
		info["userAge"] = reportInfos[0]["age"]
		info["historyInfo"] = reportInfos[0]["history"]
		info["doctorResult"] = reportInfos[0]["result"]
		createTime, errCreateTime := strconv.ParseInt(reportInfos[0]["reporttime"].(string), 10, 64)
		if errCreateTime != nil {
			result.ErrCode = -18
			result.ErrMsg = "获取体检时间异常"
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
		info["createTime"] = time.Unix(createTime, 0).Format("2006-01-02 15:04:05")
		info["createName"] = reportInfos[0]["man"]
		//	time.Unix(reportInfos[0]["cretatetime"], 0).Format("2006-01-02 15:04:05")
		//
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = info
	this.Data["json"] = result
	this.ServeJSON()
	return

}

//修改病例详情
func (this *UpdateCaseInfo) Post() {
	var result models.ResultInfo

	result.ErrCode = 0
	result.ErrMsg = ""
	//result.Result = info
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//添加体检信息到数据库
func (this *AddCaseInfo) Post() {
	fmt.Println("enter the AddCaseInfo function ")
	var result models.ResultInfo
	//var strWhere string
	//department, err := this.GetInt("department")   //科室名称
	doctor, err1 := this.GetInt("doctor")          //主治医师
	checkTime := this.GetString("checkTime")       //体检时间
	historyInfo := this.GetString("historyInfo")   //既往病史
	memNo, err0 := this.GetInt("memNo")            //患者编号
	userName := this.GetString("userName")         //患者姓名
	userGender, err2 := this.GetInt("userGender")  //患者性别
	userAge, err3 := this.GetInt("userAge")        //患者年龄
	checkItems := this.GetString("checkItems")     //体检项目
	doctorResult := this.GetString("doctorResult") //诊断结论
	createMan := this.GetString("createName")      //录入人
	createTime := this.GetString("createTime")     //录入时间
	var is []models.Items
	json.Unmarshal([]byte(checkItems), &is)
	//checkTime = strings.Replace(checkTime, "z", "UTC", -1)
	//fmt.Println("checkTime ==============")
	//fmt.Println(checkTime)
	//createTime = strings.Replace(createTime, "z", "UTC", -1)
	//fmt.Println("createTime ==============")
	//fmt.Println(createTime)
	checkTime1, _ := time.Parse("2006-01-02 03:04:05", checkTime)
	createTime1, _ := time.Parse("2006-01-02 03:04:05", createTime)
	if err1 != nil && err2 != nil && err3 != nil && err0 != nil {
		result.ErrCode = -6
		result.ErrMsg = "参数传递错误"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	var npr models.NpPhysicalReport
	npr.Id_physical_person = memNo
	npr.Id_hospital_department_doctor = doctor
	npr.Name = userName
	npr.Gender = userGender
	npr.Age = userAge
	npr.Medical_history = historyInfo
	npr.Cretate_time = checkTime1.Unix()
	npr.Report_time = createTime1.Unix()
	npr.Result = doctorResult
	npr.Status = 1
	npr.Update_time = time.Now().Unix()
	npr.Create_man = createMan
	id, err4 := models.AddCaseInfo(npr)
	if err4 != nil {
		result.ErrCode = -7
		result.ErrMsg = "保存数据失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	for _, its := range is {
		var nprsi models.NpPhysicalReportSubjectItem
		nprsi.Result = its.Result                  //检测结果
		nprsi.Doctor = its.Doctor                  //检测医师
		nprsi.Id_physical_report = id              //关联报告id
		nprsi.Id_physical_subject_item = its.Items //关联体检项目id
		nprsi.Normal_value = its.Normal_range      //正常值范围
		nprsi.Unit = its.Unit                      //单位
		//判断检查结果是否正常
		Range := strings.Split(its.Normal_range, "-")
		if len(Range) == 1 { //阴阳性
			nprsi.Is_increase = 0 //无法比较
		} else { //范围
			min, error1 := strconv.ParseFloat(Range[0], 64)
			max, error2 := strconv.ParseFloat(Range[1], 64)
			if error1 != nil && error2 != nil {
				fmt.Println("正常值范围转换异常,最大值最小值异常")
			}
			Result, error3 := strconv.ParseFloat(its.Result, 64)
			if error3 != nil {
				fmt.Println("检查结果转换为int出错!")
			}
			if Result >= min && Result <= max {
				nprsi.Is_increase = 3 //正常值
			} else if Result > max {
				nprsi.Is_increase = 1 //值偏高
			} else {
				nprsi.Is_increase = 2 //值偏低
			}
		}
		errAdd := models.AddCaseResult(nprsi)
		if errAdd != nil {
			fmt.Println("插入数据库异常")
			result.ErrCode = -8
			result.ErrMsg = "保存数据失败"
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	this.Data["json"] = result
	this.ServeJSON()
}
