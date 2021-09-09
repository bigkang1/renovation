package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	//登陆的授权错误
	ERROR_LOGIN_SKIP = 401
	ERROR_TOKEN_EXIST      = 404
	ERROR_TOKEN_RUNTIME    = 405
	ERROR_TOKEN_WRONG      = 406
	ERROR_TOKEN_TYPE_WRONG = 407

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003

	ERROR_USER_NO_RIGHT    = 1008

	//图片文件上传
	ERROR_FILE_FORM = 2001
	ERROR_UPLOAD_FAIL = 2002

	// code= 3000... 简历模块的错误
	ERROR_RES_NOT_EXIST = 3001

	// code= 4000... 分类模块的错误
	ERROR_JOBNAME_USED  = 4001
	ERROR_JOB_NOT_EXIST = 4002

	ERROR_VALIDATE_CAPTCHA = 5001
	ERROR_VALIDATE_NOTNULL = 5002
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_LOGIN_SKIP:       "skip login",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_FILE_FORM:       "文件格式错误",
	ERROR_UPLOAD_FAIL:     "文件上传失败",

	ERROR_RES_NOT_EXIST: "简历不存在",

	ERROR_JOBNAME_USED:  "该岗位已存在",
	ERROR_JOB_NOT_EXIST: "该岗位不存在",

	ERROR_VALIDATE_CAPTCHA: "验证码错误",
	ERROR_VALIDATE_NOTNULL: "验证码和vid不能为空",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
