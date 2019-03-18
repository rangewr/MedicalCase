package models

//与前台交互的json结构
type ResultInfo struct {
	Result  interface{} //返回数据
	ErrCode int         //状态码
	ErrMsg  string      //错误信息
}
