package controllers

//ErrorController 错误处理控制器
type ErrorController struct {
	BaseController
}

//Error404 api page not found
func (e *ErrorController) Error404() {
	e.SendError("api not found")
}

//Error500 Internal Error
func (e *ErrorController) Error500() {
	e.SendInternalError("Internal Error")
}

//ErrorDb DB ERROR
func (e *ErrorController) ErrorDb() {
	e.SendInternalError("service can not connect to db")
}

//Error501 501 ERROR
func (e *ErrorController) Error501() {
	e.SendInternalError("service Not Implemented")
}
