package model

var OK = 0
var BAD = -1

var VISITOR = 0
var ADMIN = 1

var READY uint = 1
var NOT_READY uint = 0

var Cities = []string{"北京", "上海", "杭州", "成都", "广州", "天津", "深圳", "长春", "西安", "沈阳"}
var CitiesEnglish = []string{"Beijing", "Shanghai", "Hangzhou", "Chengdu", "Guangzhou", "Tianjin", "Shenzhen", "Changchun", "Xian", "Shenyang"}
var EventPageSize = 12
var RankingPageSize = 10

var ItemColumns = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
var MemberColumns = []string{"A", "B", "C", "D", "E", "F"}
var ReturnColumns = []string{"A", "B", "C"}
var SaleDetailColumns = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"}
var SaleDetailColumnsNames = []string{"Order ID", "Order Time", "Brand", "SKU", "Barcode", "Color", "Category", "Season", "Size", "Retail Price", "Sale Price", "Discount", "Paid Price", "Coupon Used", "Is Return"}
var SheetName = "Sheet1"

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
var TokenMissingErrorCode = 4003
var NotLoggedInErrorCode = 4003

var RecordExistError string = "记录已存在错误"
var RegisterError string = "注册错误"
var SqlQueryError string = "数据库查询错误"
var SqlInsertionError string = "数据库插入错误"
var SqlUpdateError string = "数据库更新错误"

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
var TokenMissingError string = "无有效TOKEN错误"
var NotLoggedInError = "未登录，或TOKEN已失效错误"

var InvitationCode string = "OPRKBGIN"
var Publickey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAz5TOzCB0DXwuYgftCFc2
7KxvJDPmNvDqVcqBPnn1UGmNkwcnZQSd+LSg1laDHwNui6dd/69pthE5Cj06SPKq
/tXVazW7t5ycOfrRLrO22bym2ZiskndhxyF1k7/LqoCnLhIFm82bNkihcUbmAbQM
H6c4zqOVKJ5Hp8y4rd3oIk/zW/YyPQ+7ibFPEl2+2YUs4RDMwtghJqOv83nUryKP
yo+zItq8qSzKDxrjNI5G/Ormlxn/nTt6jJtOn3klbJG6CbtmOnX4P7gM/oJHRBq1
r//P6Lcrr1OZESkUJ4+2/Q1JCiL9wVSU+EmfyIBvY+xTlQ7UOegUJ3/mvzHEufDf
zQIDAQAB
-----END PUBLIC KEY-----`

var Pirvatekey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAz5TOzCB0DXwuYgftCFc27KxvJDPmNvDqVcqBPnn1UGmNkwcn
ZQSd+LSg1laDHwNui6dd/69pthE5Cj06SPKq/tXVazW7t5ycOfrRLrO22bym2Zis
kndhxyF1k7/LqoCnLhIFm82bNkihcUbmAbQMH6c4zqOVKJ5Hp8y4rd3oIk/zW/Yy
PQ+7ibFPEl2+2YUs4RDMwtghJqOv83nUryKPyo+zItq8qSzKDxrjNI5G/Ormlxn/
nTt6jJtOn3klbJG6CbtmOnX4P7gM/oJHRBq1r//P6Lcrr1OZESkUJ4+2/Q1JCiL9
wVSU+EmfyIBvY+xTlQ7UOegUJ3/mvzHEufDfzQIDAQABAoIBAD1HjcD+96OffEXe
VyA2NvWpden3FEg12MfYz0y1TjEd5/h2jS+qLERmdnCv+2dlaPX7Q6mejBN+hBs8
tf8g/E/cqnNK2o66wffvzl7+GMWwhoUIKDHY4lmZzA8A+MvtzOyxz0wOZ3qf+GDr
cC0ijM2vXPrLmdXy2+5yZjaVoti1v1+IywDN37pckZ4dmhs+m1ZovNzO8kf0u117
exf3PGMRAC/6aHWLrwvmzUKLtsCUzGwNXPFqECnMmCJZPjQqkx1pQWaVKo1E+DKF
4r7ViBxbYlMZmMZ9f1mVNLA152aQbX8jiPs/A20023zFaoFpQbAXgQ7f70wbe98n
ytiqUQECgYEA6iQVAk2MAp2vSJfQ8TDVO29R4PwDdD3a+9FpRmPnmVtSafKePDU3
IIrjET9nf+e5jA/rIfbZYSUX5ckPdN64o6G8FuqErOUqUKJAE8KVflzMfPMLAwGO
ioUnYxtfnR1GyNQTA6VMkzx/UYFvvtRpPUNwVQ+uxlfKUpAJ/u8JnkECgYEA4vX0
mtDlss47X4jPrNGvE0qoSmenw8aKEnzgyziwGeEOYAAu1y0gbvhiTljodtA7wQMX
MZwuZOGzeBhG4dMiQcSM6IBnajR9z70hEIz2YItS/hVCJUOGtaqUFk6J5hPpbrzG
+Wygq0JlWXas4HneG+XrzzoW0ar+SgKe7U+DNo0CgYB/X0meixkTgzyLvSsJSot1
XcWpIu+uGMg8HVur00V2g9t9j2LNVhW7OlL0Ww2u4xxpOW+sdmEjG864Tnx+E3tW
aPGtdb7fX3t5igpZtY0lxM3pWz4uUHZ+nJkkrQuCqR6MufHuFcpmfo60hDmKEnt9
vGYrn/BwLen+qCUH7nnJAQKBgQDFL2G8HCBk8C6/etLL2EWeoi+CrXohat5M37hC
d9bwNQtTNvV7N5bFMwHeBfq6N4Ki17eP/5yDQ2C0x4rV6qUJtOWjnuO6by6bjTsr
8Pyhtop9fCTC0V85eKE+nC/M+KHH9zV8QPd6s63wQ15BjT/+xwzQNyzaLxDNZmeD
0KA0hQKBgQDk046TF1xRSr3nCMZm/v9RQ3miC9P0/QohweL+gUAKsPlXqowm0zDm
DDCxPie0pOhW7x+pw8vtOuEf7CY6JskVcMAHxMTa9C1i0tqLwJ2bgWquK/KggKPZ
KH+/Dx4wBgnnTrdcWdLs1bb39x6EqbZXikGb5+D4hHFL2AIUIvpsuQ==
-----END RSA PRIVATE KEY-----`
