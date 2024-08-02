package admin

// 基类控制器
type BaseController struct {
}

/**
 * api控制器公共变量
 */
var moduleName = "admin" //模块名称

// 构造函数
func init() {

}

// 获取当前模块名称
func (con BaseController) GetModuleName() string {
	return moduleName
}
