// DON'T CHANGE THIS FILE MANUALLY

package i18n

const (
    FbTokenError = "FbTokenError" // Facebook信息获取出现错误
    GoogleTokenError = "GoogleTokenError" // Google信息获取出现错误
    NotLoginError = "NotLoginError" // 尚未登录，请先登录
    ParamNullError = "ParamNullError" // 参数不能为空
    PhoneAlreadyBindFb = "PhoneAlreadyBindFb" // 该手机号码已经绑定Facebook账号
    PhoneAlreadyBindGoogle = "PhoneAlreadyBindGoogle" // 该手机号码已经绑定Google账号
    PhoneFormatError = "PhoneFormatError" // 请输入正确的手机号码
    SystemError = "SystemError" // 系统出小差了，请稍后再试
    VerifyCodeError = "VerifyCodeError" // 验证码不存在或者已经过期
    Success = "success" // 成功

)

const LanJson = `{
    "success": {
        "zh": "成功",
        "en": "success."
    },
    "SystemError": {
        "zh": "系统出小差了，请稍后再试",
        "en": "System error, please try later."
    },
    "PhoneFormatError": {
        "zh": "请输入正确的手机号码",
        "en": "Please input correct phone number."
    },
    "ParamNullError": {
        "zh": "参数不能为空",
        "en": "Param can not be empty."
    },
    "VerifyCodeError": {
        "zh": "验证码不存在或者已经过期",
        "en": "Verification code error or expired."
    },
    "NotLoginError": {
        "zh": "尚未登录，请先登录",
        "en": "Need login first."
    },
    "GoogleTokenError": {
        "zh": "Google信息获取出现错误",
        "en": "Google info get error."
    },
    "FbTokenError": {
        "zh": "Facebook信息获取出现错误",
        "en": "Facebook info get error."
    },
    "PhoneAlreadyBindGoogle": {
        "zh": "该手机号码已经绑定Google账号",
        "en": "Phone number already bind other google account."
    },
    "PhoneAlreadyBindFb": {
        "zh": "该手机号码已经绑定Facebook账号",
        "en": "Phone number already bind other facebook account."
    }
}`