package model

var OK = 0
var BAD = -1

var VISITOR = 0
var ADMIN = 1

var Cities = []string{"北京", "上海", "杭州", "成都", "广州", "天津", "深圳", "长春", "西安", "沈阳"}
var EventPageSize = 12
var CONSIGNMENT_EVENT_TYPE = 1
var CONSIGNMENT_EVENT_TYPE_LABEL = "代销"
var PURCHASED_EVENT_TYPE = 2
var PURCHASED_EVENT_TYPE_LABEL = "自采"

var RequestSuccessMsg string = "请求成功"

var RecordExistError string = "记录已存在错误"
var SqlQueryError string = "数据库查询错误"
var SqlInsertionError string = "数据库插入错误"
var RequestParameterError string = "请求参数错误"
var RequestBodyError string = "请求体结构错误"
var ImportFileError string = "上传文件错误"
var ExcelParseError string = "表格转换错误"
var NotAuthorizedError string = "无权限错误"

var DataTypeConversionError string = "数据格式转化错误"
