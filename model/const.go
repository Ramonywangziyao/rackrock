package model

var OK = 0
var BAD = -1

var VISITOR = 0
var ADMIN = 1

var Cities = []string{"北京", "上海", "杭州", "成都", "广州", "天津", "深圳", "长春", "西安", "沈阳"}
var EventPageSize = 12
var RankingPageSize = 10

var CONSIGNMENT_EVENT_TYPE = 1
var CONSIGNMENT_EVENT_TYPE_LABEL = "代销"
var PURCHASED_EVENT_TYPE = 2
var PURCHASED_EVENT_TYPE_LABEL = "自采"

var RequestSuccessMsg string = "请求成功"

var RecordExistErrorCode = 1000
var SqlQueryErrorCode = 1001
var SqlInsertionErrorCode = 1002
var RegisterErrorCode = 1003

var RequestParameterErrorCode = 2000
var RequestBodyErrorCode = 2001
var ImportFileErrorCode = 2002
var RequestParameterMissingErrorCode = 2003

var ExcelParseErrorCode = 3000
var DataTypeConversionErrorCode = 3001
var ReportNotReadyErrorCode = 3002

var NotAuthorizedErrorCode = 4000
var PasswordErrorCode = 4001
var InvitationCodeErrorCode = 4002

var RecordExistError string = "记录已存在错误"
var RegisterError string = "注册错误"
var SqlQueryError string = "数据库查询错误"
var SqlInsertionError string = "数据库插入错误"

var RequestParameterError string = "请求参数错误"
var RequestBodyError string = "请求体结构错误"
var ImportFileError string = "上传文件错误"
var RequestParameterMissingError = "请求参数缺失错误"

var ExcelParseError string = "表格转换错误"
var NotAuthorizedError string = "无权限错误"
var ReportNotReadyError string = "报告页未就绪错误"

var PasswordError string = "密码错误"
var InvitationCodeError string = "邀请码错误"
var DataTypeConversionError string = "数据格式转化错误"

var InvitationCode string = "OPRKBGIN"
var Pirvatekey = `-----BEGIN 私钥-----
MIIEpAIBAAKCAQEAk+89V7vpOj1rG6bTAKYM56qmFLwNCBVDJ3MltVVtxVUUByqc
5b6u909MmmrLBqS//PWC6zc3wZzU1+ayh8xbUAEZuA3EjlPHIaFIVIz04RaW10+1
xnby/RQE23tDqsv9a2jv/axjE/27b62nzvCWeItu1kNQ3MGdcuqKjke+LKhQ7nWP
RCOd/ffVqSuRvG0YfUEkOz/6UpsPr6vrI331hWRB4DlYy8qFUmDsyvvExe4NjZWb
lXCqkEXRRAhi2SQRCl3teGuIHtDUxCskRIDiaMD+Qt2Yp+Vvbz6hUiqIWSIH1BoH
Jer/JOq2/O6X3cmuppU4AdVNgy8Bq236iXvrMQIDAQABAoIBAQCCbxZvHMfvCeg+
YUD5+W63dMcq0QPMdLLZPbWpxMEclH8sMm5UQ2SRueGY5UBNg0WkC/R64BzRIS6p
jkcrZQu95rp+heUgeM3C4SmdIwtmyzwEa8uiSY7Fhbkiq/Rly6aN5eB0kmJpZfa1
6S9kTszdTFNVp9TMUAo7IIE6IheT1x0WcX7aOWVqp9MDXBHV5T0Tvt8vFrPTldFg
IuK45t3tr83tDcx53uC8cL5Ui8leWQjPh4BgdhJ3/MGTDWg+LW2vlAb4x+aLcDJM
CH6Rcb1b8hs9iLTDkdVw9KirYQH5mbACXZyDEaqj1I2KamJIU2qDuTnKxNoc96HY
2XMuSndhAoGBAMPwJuPuZqioJfNyS99x++ZTcVVwGRAbEvTvh6jPSGA0k3cYKgWR
NnssMkHBzZa0p3/NmSwWc7LiL8whEFUDAp2ntvfPVJ19Xvm71gNUyCQ/hojqIAXy
tsNT1gBUTCMtFZmAkUsjqdM/hUnJMM9zH+w4lt5QM2y/YkCThoI65BVbAoGBAMFI
GsIbnJDNhVap7HfWcYmGOlWgEEEchG6Uq6Lbai9T8c7xMSFc6DQiNMmQUAlgDaMV
b6izPK4KGQaXMFt5h7hekZgkbxCKBd9xsLM72bWhM/nd/HkZdHQqrNAPFhY6/S8C
IjRnRfdhsjBIA8K73yiUCsQlHAauGfPzdHET8ktjAoGAQdxeZi1DapuirhMUN9Zr
kr8nkE1uz0AafiRpmC+cp2Hk05pWvapTAtIXTo0jWu38g3QLcYtWdqGa6WWPxNOP
NIkkcmXJjmqO2yjtRg9gevazdSAlhXpRPpTWkSPEt+o2oXNa40PomK54UhYDhyeu
akuXQsD4mCw4jXZJN0suUZMCgYAgzpBcKjulCH19fFI69RdIdJQqPIUFyEViT7Hi
bsPTTLham+3u78oqLzQukmRDcx5ddCIDzIicMfKVf8whertivAqSfHytnf/pMW8A
vUPy5G3iF5/nHj76CNRUbHsfQtv+wqnzoyPpHZgVQeQBhcoXJSm+qV3cdGjLU6OM
HgqeaQKBgQCnmL5SX7GSAeB0rSNugPp2GezAQj0H4OCc8kNrHK8RUvXIU9B2zKA2
z/QUKFb1gIGcKxYr+LqQ25/+TGvINjuf6P3fVkHL0U8jOG0IqpPJXO3Vl9B8ewWL
cFQVB/nQfmaMa4ChK0QEUe+Mqi++MwgYbRHx1lIOXEfUJO+PXrMekw==
-----END 私钥-----
`
