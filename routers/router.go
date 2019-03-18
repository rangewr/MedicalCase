package routers

import (
	"MedicalCase/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//获取数据使用Get方法,操作数据使用Post方式
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Test", &controllers.MainController{})
	//根目录
	beego.Router("/TestRouter", &controllers.TestController{}) //测试用
	//================登录=============
	//beego.Router("/proc/Login", &controllers.UserLogin{})//登录方法 Post
	//beego.Router("/proc/Logout", &controllers.UserLogout{})//登出方法 Get
	//==============操作病例============
	beego.Router("/illness/GetCaseList", &controllers.GetCaseList{}) //获取病例列表 Get
	beego.Router("/illness/AddCaseInfo", &controllers.AddCaseInfo{}) //添加一个病例 Post
	beego.Router("/illness/GetCaseInfo", &controllers.GetCaseInfo{}) //查询一个病例的详细信息 Get
	beego.Router("/illness/DelCaseInfo", &controllers.DelCaseInfo{}) //删除一个病例的详细信息 Post
	beego.Router("/illness/UpdateCaseInfo", &controllers.UpdateCaseInfo{})//更新一个病例的详细信息 Post
	//============获取辅助字段============
	beego.Router("/illness/GetHospitalList", &controllers.GetHospitalList{})     //获取医院列表 Get
	beego.Router("/illness/GetDepartmentList", &controllers.GetDepartmentList{}) //获取科室列表 Get
	beego.Router("/illness/GetSubjectList", &controllers.GetSubjectList{})       //获取科目列表 Get
	beego.Router("/illness/GetItemList", &controllers.GetItemList{})             //获取项目列表 Get
	beego.Router("/illness/GetDoctorList", &controllers.GetDoctorList{})         //获取医生列表 Get
	beego.Router("/illness/GetMemberInfo", &controllers.GetMemberList{})         //获取会员信息 Get
	beego.Router("/illness/GetItemsInfo", &controllers.GetItemsInfo{})           //获取体检项目信息 Get

}
