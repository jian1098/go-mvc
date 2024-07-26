package common

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/gookit/validate"
)

// 验证器 验证规则参考https://github.com/gookit/validate/blob/master/README.zh-CN.md
type customValidator struct {
	Scene string //验证场景
}

func (c *customValidator) ValidateStruct(ptr any) error {
	v := validate.Struct(ptr)
	v.AtScene(c.Scene) // 设置验证场景
	v.Validate()       // 调用验证

	return v.Errors
}

func (c *customValidator) Engine() any {
	return nil
}

// InitValidator 初始化自定义验证器实例
func InitValidator() {
	// 更改全局选项
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = true
		opt.SkipOnEmpty = true
	})

	//全局错误消息
	validate.AddGlobalMessages(map[string]string{
		"required":  "{field}不能为空",
		"len":       "{field}长度不正确",
		"minLen":    "{field}长度不能小于%d位",
		"maxLen":    "{field}长度不能大于%d位",
		"cn_mobile": "{field}格式不正确",
	})

	binding.Validator = &customValidator{}
}

/*
*
全局配置选项
*/
type GlobalOption struct {
	// FilterTag 结构体中的过滤规则标签名称。默认 'filter`
	FilterTag string
	// ValidateTag 结构体中的验证规则标签名称。默认 'binding`
	ValidateTag string
	// FieldTag 定义结构体字段验证错误时的输出名字。默认使用 json
	FieldTag string
	// LabelTag 定义结构体字段验证错误时的输出翻译名称。默认使用 label
	// - 等同于设置 字段 translate
	LabelTag string
	// MessageTag define error message for the field. default: message
	MessageTag string
	// StopOnError 如果为 true，则出现第一个错误时，将停止继续验证。默认 true
	StopOnError bool
	// SkipOnEmpty 跳过对字段不存在或值为空的检查。默认 true
	SkipOnEmpty bool
	// UpdateSource Whether to update source field value, useful for struct validate
	UpdateSource bool
	// CheckDefault Whether to validate the default value set by the user
	CheckDefault bool
	// CheckZero Whether validate the default zero value. (intX,uintX: 0, string: "")
	CheckZero bool
	// CheckSubOnParentMarked 当字段是一个结构体时，仅在当前字段配置了验证tag时才收集子结构体的规则
	CheckSubOnParentMarked bool
}
