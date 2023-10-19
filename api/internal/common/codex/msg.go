package codex

var codeMsgMap = map[ResCode]string{
	CodeSuccess:                 "操作成功",
	CodeInvalidParams:           "请求参数错误",
	CodeNeedLogin:               "请先登陆",
	CodeCasbinError:             "没有该API接口访问权限",
	CodeBusinessTypeErr:         "业务类型错误",
	CodeInvalidToken:            "无效的token",
	CodeWrongPassword:           "密码错误",
	CodeWrongUserNameOrPassword: "用户名或密码错误",
	CodeWrongSmsCode:            "验证码不正确",
	CodeSendSmsCodeErr:          "验证码发送失败",
	CodeCacheSmsCodeErr:         "验证码缓存失败",
	CodeDelCacheSmsCodeErr:      "删除验证码缓存失败",
	CodePasswordNotEqual:        "两次密码不一致",
	CodeCacheEmailCodeErr:       "邮箱验证码缓存出错",
	CodeSendEmailErr:            "发送邮箱验证码出错",
	CodeWrongEmailCode:          "邮箱验证码错误",
	CodeConfirmPasswordError:    "两次密码输入不一致",
	CodeOldPasswordError:        "旧密码错误",
	CodeEmailCodeExpired:        "邮箱验证码已过期",
	CodeInviteLinkExpired:       "邀请链接过期",
	CodeSMSCodeExpired:          "短信验证码已过期",
	CodeRoleTypeErr:             "注册用户角色错误",
	CodeNullInfo:                "信息不能为空",

	CodeUserNotExist:     "该用户不存在",
	CodeUserExist:        "该用户已存在",
	CodeEmailExist:       "该邮箱已注册",
	CodePhoneExist:       "该手机号已注册",
	CodeCompanyExist:     "公司已经存在，请登录",
	CodeInternalErr:      "服务器开小差啦，稍后再来试一试",
	CodeGenTokenErr:      "生成Token异常",
	CodeParseFormFileErr: "解析上传文件异常",
	CodeFileFormErr:      "文件格式错误",
	CodeCreateFileErr:    "创建文件失败",
	CodeQueryUserErr:     "查询用户信息出错",
	CodeDatabaseError:    "数据库错误",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeInternalErr]
	}
	return msg
}
