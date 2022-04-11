package errInfo

/*
   功能说明: 错误提示信息
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 13:45
*/

type ErrInfo = string

const (
	SUCCESS               ErrInfo = "success"
	FormValidateFailed    ErrInfo = "表单验证失败"
	NotFound              ErrInfo = "访问路径不存在"
	TransError            ErrInfo = "翻译器错误"
	RequestNotAllow       ErrInfo = "请求方式不被允许"
	FileIsTooBig          ErrInfo = "文件太大"
	FileGetFailed         ErrInfo = "文件获取失败"
	FileNameIsInvalid     ErrInfo = "文件名无效"
	FileReadFailed        ErrInfo = "文件读取失败"
	FileCreateFailed      ErrInfo = "文件创建失败"
	FileWriteFailed       ErrInfo = "文件写入失败"
	FileNotFound          ErrInfo = "文件不存在"
	Base64DecodeFailed    ErrInfo = "base64解码失败"
	CreatedSuccess        ErrInfo = "数据创建成功"
	CreatedFailed         ErrInfo = "数据创建失败"
	UpdatesSuccess        ErrInfo = "数据修改成功"
	UpdatedFailed         ErrInfo = "数据修改失败"
	Exist                 ErrInfo = "记录已存在"
	PrimaryRequire        ErrInfo = "主键ID必传"
	RecordDoesNotExist    ErrInfo = "记录不存在"
	DontResetStatus       ErrInfo = "记录不能被重置"
	ResetStatusSuccess    ErrInfo = "记录重置成功"
	ResetStatusFailed     ErrInfo = "记录重置成功"
	SerializationFailed   ErrInfo = "序列化数据失败"
	DeserializationFailed ErrInfo = "反序列化数据失败"
	QueryGameNodeFailed   ErrInfo = "查询游戏服节点失败"
	NoUserGameNode        ErrInfo = "无可用的游戏服节点"
	RequestGameFailed     ErrInfo = "请求游戏服失败"
	ByteFlowToMapFailed   ErrInfo = "字符串转map失败"
	DomainIDRequired      ErrInfo = "游戏ID必传"
)
