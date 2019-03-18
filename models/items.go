package models

type Items struct {
	Subject      int64  `json:"subject"`      //体检科目
	Items        int64  `json:"items"`        //体检项目
	Result       string `json:"result"`       //检查结果
	Unit         string `json:"unit"`         //单位
	Normal_range string `json:"normal_range"` //正常值范围
	Doctor       int64  `json:"doctor"`       //检测医师
}
