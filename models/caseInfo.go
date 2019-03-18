package models

type CaseInfo struct {
	Department   int     `json:"department"`   //科室
	Doctor       int     `json:"doctor"`       //主治医师
	CheckTime    string  `json:"checkTime"`    //体检时间
	HistoryInfo  string  `json:"historyInfo"`  //既往病史
	MemNo        string  `json:"memno"`        //患者编号
	UserName     string  `json:"userName"`     //用户名
	UserGender   string  `json:"userGender"`   //用户性别
	CheckItems   []Items `json:"ckeckItems"`   //体检项目
	DoctorResult string  `json:"doctorResult"` //检测结果
	CreateName   string  `json:"createName"`   //录入人
	CreateTime   string  `json:"createTime"`   //录入时间
}
