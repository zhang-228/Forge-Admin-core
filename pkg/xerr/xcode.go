package xerr

// 系统预设错误
var (
	JWTTokenExpired = "Token is expired"            // token过期
	JwtNotExists    = "no token present in request" // token不存在

	Canceled         = Add(498, "CANCELED")         // 请求取消
	ServerErr        = Add(500, "INTERNAL_ERROR")   // 服务器内部错误
	Deadline         = Add(504, "")                 // 请求超时
	NotAuth          = Add(1, "not auth")           // 未授权
	AuthExpired      = Add(2, "auth expired")       // 授权信息过期
	ResourceNotFound = Add(3, "resource not found") // 资源未找到
	RequestParamsErr = Add(4, "request params err") // 请求参数错误
	AccessDenied     = Add(5, "access denied")      // 拒绝访问(权限不

	PasswordErr           = Add(10001, "密码错误请重新输入")        // 密码错误
	AccountDisabled       = Add(10002, "账号已被禁用,请联系管理员")    // 账号已被禁用
	VerifyCodeError       = Add(10007, "验证码错误")            // 验证码错误
	VerifyCodeFrequently  = Add(10008, "验证码已发送,请勿频繁获取验证码") // 获取验证码过于频繁
	UnsupportedLoginType  = Add(10003, "未支持的登录方式")         // 未支持的登录方式
	AuthorityError        = Add(10004, "权限不足")             // 用户权限不足
	UpdateCasbinRuleErr   = Add(10005, "更新角色接口权限失败")       // 更新casbin rule失败
	NotAllowCasbinRuleErr = Add(10006, "存在非法接口")           // 使用了不存在的接口
	UploadFileEmpty       = Add(20000, "上传文件为空")           // 上传文件为空
	UploadFileFail        = Add(20001, "上传文件失败")           // 上传文件失败
	UploadFileMaxSize     = Add(20001, "上传文件大小超过限制")       // 上传文件大小超过限制

	TemplateFileNotExist     = Add(20002, "指定的模板文件文件不存在")
	TemplateServiceExist     = Add(20003, "该接口已存在,请修改后重新提交")
	TemplateRouterGroupExist = Add(20004, "路由组已存在,请修改后重新提交")
	TemplateServiceNotExist  = Add(20004, "接口不存在")
	TableNotExist            = Add(30001, "指定的表不存在")                // 指定的表不存在
	OpenTemplateFileFail     = Add(30002, "打开模板文件失败")               // 打开模板文件失败
	GenerateCodeFail         = Add(30003, "生成代码失败")                 // 生成代码失败
	TableMustHavePrimaryKey  = Add(30002, "必须指定一个主键,否则会导致无法生成代码")   // 表必须有一个主键
	TableOnlyOnePrimaryKey   = Add(30003, "只能指定一个主键,请删除多余的主键后重新提交") // 只能指定一个主键
)
