package models

//AdminBaseListSearch 后台搜索基础搜索
type AdminBaseListSearch struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
	OrderbyName string `json:"order_by_name"`
	OrderType   string `json:"order_type"`
}

//CheckUserIDAuth 检查用户的权限
func CheckUserIDAuth(userID int, authKey string) bool {
	if userID == 1 {
		return true
	}
	return false
}
