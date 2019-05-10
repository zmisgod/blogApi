package util

import (
	"time"
)

//PostOrderName 文章排序列表
func PostOrderName() []map[string]string {
	list := make([]map[string]string, 0)
	zero := map[string]string{"name": "文章ID", "value": "id"}
	one := map[string]string{"name": "用户ID", "value": "user_id"}
	two := map[string]string{"name": "文章状态", "value": "post_status"}
	three := map[string]string{"name": "评论状态", "value": "comment_status"}
	four := map[string]string{"name": "文章创建时间", "value": "created_at"}
	five := map[string]string{"name": "文章最后修改时间", "value": "updated_at"}
	six := map[string]string{"name": "文章发布时间", "value": "publish_time"}
	list = append(list, zero, one, two, three, four, five, six)
	return list
}

//TagOrderName 标签排序列表
func TagOrderName() []map[string]string {
	list := make([]map[string]string, 0)
	zero := map[string]string{"name": "标签ID", "value": "tag_id"}
	one := map[string]string{"name": "标签名称", "value": "name"}
	two := map[string]string{"name": "标签总数", "value": "counts"}
	list = append(list, zero, one, two)
	return list
}

//TopicOrderName 主题排序列表
func TopicOrderName() []map[string]string {
	list := make([]map[string]string, 0)
	zero := map[string]string{"name": "主题ID", "value": "id"}
	one := map[string]string{"name": "排序", "value": "sort"}
	two := map[string]string{"name": "创建时间", "value": "created_at"}
	three := map[string]string{"name": "修改时间", "value": "updated_at"}
	list = append(list, zero, one, two, three)
	return list
}

//PostCategories 文章分类
func PostCategories() []map[string]interface{} {
	list := make([]map[string]interface{}, 0)
	one := map[string]interface{}{"id": 1, "name": "不可描述"}
	two := map[string]interface{}{"id": 2, "name": "心情感悟"}
	three := map[string]interface{}{"id": 3, "name": "技术分享"}
	list = append(list, one, two, three)
	return list
}

//PostPostStatus 文章发布状态
func PostPostStatus() []map[string]interface{} {
	list := make([]map[string]interface{}, 0)
	one := map[string]interface{}{"id": 1, "name": "发布"}
	two := map[string]interface{}{"id": 2, "name": "垃圾桶"}
	three := map[string]interface{}{"id": 3, "name": "草稿"}
	four := map[string]interface{}{"id": 4, "name": "私人"}
	list = append(list, one, two, three, four)
	return list
}

//PostPostTyoe 文章类别
func PostPostTyoe() []map[string]interface{} {
	list := make([]map[string]interface{}, 0)
	one := map[string]interface{}{"id": 1, "name": "分类文章"}
	two := map[string]interface{}{"id": 2, "name": "专栏文章"}
	list = append(list, one, two)
	return list
}

//PostCommentStatus 文章评论状态
func PostCommentStatus() []map[string]interface{} {
	list := make([]map[string]interface{}, 0)
	one := map[string]interface{}{"id": 1, "name": "允许"}
	two := map[string]interface{}{"id": 2, "name": "需要审核"}
	three := map[string]interface{}{"id": 3, "name": "不允许"}
	list = append(list, one, two, three)
	return list
}

//PostHistoryTyoe 文章浏览历史类型
var PostHistoryTyoe = map[int]string{1: "article", 2: "home", 3: "category", 4: "link", 5: "comment", 6: "tag"}

//UserStatus 用户状态
var UserStatus = map[int]string{0: "正常", 1: "关闭"}

//UserLinkType 用户链接类型
var UserLinkType = map[int]string{1: "https://github.com/", 2: "https://twitter.com/", 3: "https://facebook.com/", 4: "https://dribbble.com/", 5: "https://gitlab.com/", 6: "https://youtube.com/", 7: "https://www.linkedin.com/in/", 8: "https://weibo.com/", 9: "https://gitee.com/"}

//Disable 全局disable状态
var Disable = map[int]string{0: "有效", 1: "无效"}

//WpsPosts 文章数据结构
type WpsPosts struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	PostTitle     string    `json:"post_title"`
	PostIntro     string    `json:"post_intro"`
	PostType      int       `json:"post_type"`
	PostStatus    int       `json:"post_status"`
	CommentStatus int       `json:"comment_status"`
	CoverURL      string    `json:"cover_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	PublishTime   time.Time `json:"publish_time"`
}

//WpsUsers 用户
type WpsUsers struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//WpsUsersLink 用户链接
type WpsUsersLink struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	LinkType  int       `json:"link_type"`
	Suffix    string    `json:"suffix"`
	Visible   int       `json:"visible"`
	CreatedAt time.Time `json:"created_at"`
	Disabled  int       `json:"disabled"`
}

//WpsUserInfo 用户信息
type WpsUserInfo struct {
	ID        int    `json:"id"`
	NickName  string `json:"nickname"`
	Sex       int    `json:"sex"`
	HeadURL   string `json:"head_url"`
	Introduce string `json:"introduce"`
	birthday  []uint8
	Birthday  string         `json:"birthday"`
	UserLink  []WpsUsersLink `json:"user_link"`
}

//WpsHistory 浏览记录
type WpsHistory struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	UserAgent  string `json:"user_agent"`
	IP         string `json:"ip"`
	visiteTime int
	VisiteTime string `json:"visite_time"`
	params     string
	Params     map[string]interface{}
}

//WpsLink 友链
type WpsLink struct {
	LinkURL         string `json:"link_url"`
	LinkName        string `json:"link_name"`
	LinkImage       string `json:"link_image"`
	LinkDescription string `json:"link_description"`
}
