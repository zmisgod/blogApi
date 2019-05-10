package models

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/zmisgod/blogApi/util"
)

// type PostDetail struct {
// 	Post     util.WpsPosts    `json:"post"`
// 	UserInfo util.WpsUserInfo `json:"user_info"`
// 	Contents string           `json:"contents"`
// 	Tags     []WpsTag         `json:"tags"`
// 	NumInfo  PostNum          `json:"num_info"`
// 	Media    []WpsPostMedia   `json:"media_info"`
// }

//PostDetail 文章详情
type PostDetail struct {
	util.WpsPosts
	Contents string   `json:"contents"`
	CatID    int      `json:"cat_id"`
	Tags     []string `json:"tags"`
}

//PostList 文章列表
type PostList struct {
	util.WpsPosts
}

//GetArticleLists 获取文章列表
func GetArticleLists(page, pageSize int) ([]PostList, error) {
	var (
		rows *sql.Rows
		err  error
	)
	postList := []PostList{}
	offset := (page - 1) * pageSize
	rows, err = dbConn.Query(fmt.Sprintf("select p.cat_id, p.id,p.user_id, p.post_title,u.name as user_name,c.c_name as category_name,p.post_title,p.post_intro,p.created_at from wps_posts as p left join wps_users as u on p.user_id = u.id left join wps_post_categories as c on c.id = p.cat_id where p.post_status = 1 order by p.created_at desc limit %d,%d", offset, pageSize))
	defer rows.Close()
	if err != nil {
		return postList, err
	}

	// for rows.Next() {
	// 	var aPost PostList
	// 	err = rows.Scan(
	// 		&aPost.CategoryID,
	// 		&aPost.ID,
	// 		&aPost.UserID,
	// 		&aPost.PostTitle,
	// 		&aPost.UserName,
	// 		&aPost.CategoryName,
	// 		&aPost.PostTitle,
	// 		&aPost.PostIntro,
	// 		&aPost.createdAt,
	// 	)
	// 	tm := time.Unix(int64(aPost.createdAt), 0)
	// 	aPost.CreatedAt = tm.Format("2006-01-02 15:04")

	// 	tags, _ := GetPostTagLists(aPost.ID)
	// 	aPost.Tags = tags

	// 	num, _ := GetArticleNumsByPost(aPost.ID)
	// 	aPost.NumInfo = num

	// 	postList = append(postList, aPost)
	// }
	return postList, nil
}

//GetArticleDetail 获取文章详情
func GetArticleDetail(postID int) (PostDetail, error) {
	var post PostDetail
	// err := dbConn.QueryRow(
	// 	fmt.Sprintf("select p.comment_status,p.id,p.user_id, p.post_title,u.name as user_name,c.c_name as category_name,p.post_title,p.post_intro,p.created_at,pc.contents,p.cat_id from wps_posts as p left join wps_users as u on p.user_id = u.id left join wps_post_categories as c on c.id = p.cat_id left join wps_post_contents as pc on pc.id = p.id where p.post_status = 1 and p.id  = %d", postID)).
	// 	Scan(
	// 		&post.CommentStatus,
	// 		&post.ID,
	// 		&post.UserID,
	// 		&post.PostTitle,
	// 		&post.UserName,
	// 		&post.CategoryName,
	// 		&post.PostTitle,
	// 		&post.PostIntro,
	// 		&post.createdAt,
	// 		&post.Contents,
	// 		&post.CategoryID,
	// 	)
	// if err != nil {
	// 	return post, errors.New("empty")
	// }
	// tm := time.Unix(int64(post.createdAt), 0)
	// post.CreatedAt = tm.Format("2006-01-02 15:04")

	// tags, _ := GetPostTagLists(post.ID)
	// post.Tags = tags

	// num, _ := GetArticleNumsByPost(post.ID)
	// post.NumInfo = num

	// userInfo, _ := GetUserInfo(post.UserID)
	// post.UserInfo = userInfo
	return post, nil
}

//AutoSubPostView 文章详情浏览加一操作
func AutoSubPostView(postID int) int64 {
	var (
		err       error
		affectRow int64
	)
	affectRow = 0
	stmt, err := dbConn.Prepare(`update wps_post_nums set view_num = view_num + 1 where post_id = ?`)
	defer stmt.Close()
	if err == nil {
		res, err := stmt.Exec(postID)
		if err == nil {
			num, err := res.RowsAffected()
			if err == nil {
				return num
			}
		}
	}
	return affectRow
}

//*********   new     *********//

//PostListsSearch 文章列表
type PostListsSearch struct {
	PostType      int `json:"post_type"`
	UserID        int `json:"user_id"`
	PostStatus    int `json:"post_status"`
	TagID         int `json:"tag_id"`
	CommentStatus int `json:"comment_status"`
	AdminBaseListSearch
}

//AdminGetArticleLists 获取文章列表
func AdminGetArticleLists(search PostListsSearch, userID int) ([]util.WpsPosts, int, map[string]interface{}) {
	var posts []util.WpsPosts
	searchOrigialData := AdminArticleSearchParams()
	var count int

	sql := "select p.id,p.user_id, p.post_title,p.post_intro,p.post_status,p.post_type, p.comment_status,p.cover_url, p.created_at,p.updated_at,p.publish_time from wps_posts as p inner join wps_users as u on p.user_id = u.id where 1=1 "
	countSQL := "select count(1) from wps_posts as p inner join wps_users as u on p.user_id = u.id where 1=1"
	if CheckUserIDAuth(userID, "admin_article_list_search_user_id") {
		if search.UserID != 0 {
			sql += " and p.user_id = " + strconv.Itoa(search.UserID)
			countSQL += " and p.user_id = " + strconv.Itoa(search.UserID)
		}
	} else {
		sql += " and p.user_id = " + strconv.Itoa(userID)
		countSQL += " and p.user_id = " + strconv.Itoa(userID)
	}
	//标签ID文章
	if search.TagID != 0 {
		searchTagSQL := fmt.Sprintf("select post_id from wps_post_tags where tag_id = %d", search.TagID)
		tagRows, err := dbConn.Query(searchTagSQL)
		defer tagRows.Close()
		tagIDs := make([]string, 0)
		if err == nil {
			for tagRows.Next() {
				var tagID int
				err = tagRows.Scan(&tagID)
				if err == nil {
					tagIDs = append(tagIDs, strconv.Itoa(tagID))
				}
			}
		} else {
			fmt.Println(err)
		}
		resTagIDs := "0"
		if len(tagIDs) > 0 {
			resTagIDs = strings.Join(tagIDs, ",")

		}
		sql += " and p.id in (" + resTagIDs + ") "
		countSQL += " and p.id in (" + resTagIDs + ") "
	}
	if search.PostStatus != 0 {
		sql += " and p.post_status = " + strconv.Itoa(search.PostStatus)
		countSQL += " and p.post_status = " + strconv.Itoa(search.PostStatus)
	}
	if search.CommentStatus != 0 {
		sql += " and p.comment_status = " + strconv.Itoa(search.CommentStatus)
		countSQL += " and p.comment_status = " + strconv.Itoa(search.CommentStatus)
	}
	if search.PostType != 0 {
		sql += " and p.post_type = " + strconv.Itoa(search.PostType)
		countSQL += " and p.post_type = " + strconv.Itoa(search.PostType)
	}

	sql += " order by p." + search.OrderbyName + " " + search.OrderType
	sql += " limit " + strconv.Itoa((search.Page-1)*search.PageSize) + " , " + strconv.Itoa(search.PageSize)

	rows, err := dbConn.Query(sql)
	if err != nil {
		return posts, count, searchOrigialData
	}
	defer rows.Close()
	for rows.Next() {
		var post util.WpsPosts
		var createAt, updateAt, publishTime []uint8
		err = rows.Scan(&post.ID, &post.UserID, &post.PostTitle, &post.PostIntro, &post.PostStatus, &post.PostType, &post.CommentStatus, &post.CoverURL, &createAt, &updateAt, &publishTime)
		if err != nil {
			fmt.Println(err)
			continue
		}
		post.CreatedAt = util.ConvertUtf8ToTimeTime(createAt)
		post.UpdatedAt = util.ConvertUtf8ToTimeTime(updateAt)
		post.PublishTime = util.ConvertUtf8ToTimeTime(publishTime)
		posts = append(posts, post)
	}
	err = dbConn.QueryRow(countSQL).Scan(&count)

	return posts, count, searchOrigialData
}

//AdminArticleSearchParams 获取文章搜索的数据
func AdminArticleSearchParams() map[string]interface{} {
	searchOrigialData := make(map[string]interface{}, 0)
	searchOrigialData["postStatusList"] = util.PostPostStatus()
	searchOrigialData["categoryList"] = util.PostCategories()
	searchOrigialData["commentStatusList"] = util.PostCommentStatus()
	searchOrigialData["orderNameList"] = util.PostOrderName()
	searchOrigialData["postTypeList"] = util.PostPostTyoe()
	return searchOrigialData
}

//AdminGetArticleByID 根据articleID获取相应的文章信息
func AdminGetArticleByID(articleID int) (PostDetail, map[string]interface{}) {
	var post PostDetail

	var createAt, updateAt, publishTime []uint8
	sql := fmt.Sprintf(`select 
	p.id,p.user_id,p.post_title,p.post_intro,p.post_status,p.post_type, p.comment_status,p.cover_url, p.created_at,p.updated_at,p.publish_time, 
	pc.contents, 
	pcate.cat_id
	from wps_posts as p left join wps_users as u on p.user_id = u.id 
	left join wps_post_contents as pc on pc.id = p.id 
	left join wps_post_cate as pcate on p.id = pcate.post_id
	where p.id = %d`, articleID)

	err := dbConn.QueryRow(sql).Scan(&post.ID, &post.UserID, &post.PostTitle, &post.PostIntro, &post.PostStatus, &post.PostType, &post.CommentStatus, &post.CoverURL, &createAt, &updateAt, &publishTime,
		&post.Contents,
		&post.CatID,
	)
	if err != nil {
		fmt.Println(err)
	} else {
		post.CreatedAt = util.ConvertUtf8ToTimeTime(createAt)
		post.UpdatedAt = util.ConvertUtf8ToTimeTime(updateAt)
		post.PublishTime = util.ConvertUtf8ToTimeTime(publishTime)
	}
	rows, err := dbConn.Query(fmt.Sprintf("select t.name from wps_tags as t left join wps_post_tags as pt on pt.tag_id = t.tag_id where pt.post_id = %d ", articleID))
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		tags := []string{}
		for rows.Next() {
			var tag string
			err = rows.Scan(&tag)
			if err != nil {
				fmt.Println(err)
			} else {
				tags = append(tags, tag)
			}
		}
		post.Tags = tags
	}
	searchOrigialData := AdminArticleSearchParams()
	return post, searchOrigialData
}

//PassPost 客户端传递过来的json字符串
type PassPost struct {
	PostID        int64    `json:"post_id"`
	PostTitle     string   `json:"post_title"`
	PostIntro     string   `json:"post_intro"`
	Content       string   `json:"content"`
	PostType      int      `json:"post_type"`
	PostStatus    int      `json:"post_status"`
	CommentStatus int      `json:"comment_status"`
	CoverURL      string   `json:"cover_url"`
	PublishTime   int      `json:"publish_time"`
	CatID         int      `json:"cat_id"`
	Tags          []string `json:"tags"`
}

//SavePost 保存文章
type SavePost struct {
	UserID int `json:"user_id"`
	PassPost
	PublishTime int `json:"publish_time"`
}

//PostTag 文章标签
type PostTag struct {
	ID   int
	Name string
}

//SaveArticle 修改、新增用户的文章
func SaveArticle(savePost PassPost, userID int) (int64, error) {
	if userID <= 0 {
		return 0, errors.New("用户id为空")
	}
	if savePost.PostID < 0 {
		return 0, errors.New("文章id不能小于0")
	}
	if savePost.Content == "" {
		return 0, errors.New("文章内容不能为空")
	}
	if savePost.PostTitle == "" || len(savePost.PostTitle) > 240 {
		return 0, errors.New("文章标题不能为空或者不能超过240个字符")
	}
	if cLength := len(savePost.CoverURL); cLength > 240 {
		return 0, errors.New("文章封面url不能超过240个字符")
	}
	if savePost.PostType == 0 {
		return 0, errors.New("文章类型不合法")
	}
	if savePost.PostType == 1 && savePost.CatID == 0 {
		return 0, errors.New("文章分类与类型不合法")
	}
	if savePost.PostType == 2 && savePost.CatID != 0 {
		return 0, errors.New("文章分类与类型不合法")
	}
	if savePost.PostStatus == 0 {
		return 0, errors.New("文章状态不合法")
	}
	if savePost.CommentStatus == 0 {
		return 0, errors.New("文章是否可评论状态有误")
	}
	if savePost.PostIntro == "" {
		savePost.PostIntro = string([]rune(savePost.Content)[:240])
	}
	nowTimestamp := time.Now().Unix()
	trainaction, _ := dbConn.Begin()
	if savePost.PostID == 0 {
		stmt, _ := trainaction.Prepare("insert into wps_posts (user_id, post_title, post_intro, post_type, post_status, comment_status, cover_url, created_at, updated_at, publish_time) values (?,?,?,?,?,?,?,?,?,?)")
		result, err := stmt.Exec(userID, savePost.PostTitle, savePost.PostIntro, savePost.PostType, savePost.PostStatus, savePost.CommentStatus, savePost.CoverURL, nowTimestamp, nowTimestamp, savePost.PublishTime)
		if err != nil {
			fmt.Println(err)
			trainaction.Rollback()
			return 0, errors.New("插入文章失败")
		}
		savePost.PostID, _ = result.LastInsertId()
		stmt, _ = trainaction.Prepare("insert into wps_post_contents (id, contents) values (?,?)")
		result, err = stmt.Exec(savePost.PostID, savePost.Content)
		if err != nil {
			trainaction.Rollback()
			return 0, errors.New("插入文章内容失败")
		}
		if savePost.CatID > 0 {
			stmt, _ = trainaction.Prepare("insert into wps_post_cate (post_id, cat_id, created_at) values (?,?,?)")
			result, err = stmt.Exec(savePost.PostID, savePost.CatID, nowTimestamp)
			if err != nil {
				trainaction.Rollback()
				return 0, errors.New("插入分类失败")
			}
		}
		if len(savePost.Tags) > 0 {
			for _, tag := range savePost.Tags {
				if strings.TrimSpace(tag) != "" {
					var tagID int64
					err = trainaction.QueryRow(fmt.Sprintf("select tag_id from wps_tags where name = '%s'", tag)).Scan(&tagID)
					if err != nil {
						stmt, _ = trainaction.Prepare("insert into wps_tags (name, slug, counts) values (?,?,?)")
						tagRes, _ := stmt.Exec(tag, url.QueryEscape(tag), 1)
						tagID, _ = tagRes.LastInsertId()
					}
					stmt, _ = trainaction.Prepare("insert into wps_post_tags (post_id, tag_id, disabled, create_time) values (?,?,?,?)")
					_, err = stmt.Exec(savePost.PostID, tagID, 0, nowTimestamp)
					if err != nil {
						trainaction.Rollback()
						return 0, errors.New("插入标签失败")
					}
				}
			}
		}
	} else {
		stmt, _ := trainaction.Prepare("update wps_posts set post_title = ?, post_intro =? , post_type = ? , post_status = ?, comment_status = ?, cover_url = ? , updated_at = ?, publish_time = ? where id = ?")
		_, err := stmt.Exec(savePost.PostTitle, savePost.PostIntro, savePost.PostType, savePost.PostStatus, savePost.CommentStatus, savePost.CoverURL, nowTimestamp, savePost.PublishTime, savePost.PostID)
		if err != nil {
			trainaction.Rollback()
			return 0, errors.New("修改文章失败")
		}
		stmt, _ = trainaction.Prepare("update wps_post_contents set contents = ? where id = ?")
		_, err = stmt.Exec(savePost.Content, savePost.PostID)
		if err != nil {
			trainaction.Rollback()
			return 0, errors.New("修改文章内容失败")
		}
		var catID int
		err = trainaction.QueryRow(fmt.Sprintf("select cat_id from wps_post_cate where post_id = %d", savePost.PostID)).Scan(&catID)
		if err != nil {
			stmt, _ = trainaction.Prepare("insert into wps_post_cate (post_id, cat_id, created_at) values (?,?,?)")
			_, err = stmt.Exec(savePost.PostID, savePost.CatID, nowTimestamp)
			if err != nil {
				trainaction.Rollback()
				return 0, errors.New("插入文章分类失败")
			}
		} else {
			if catID != savePost.CatID {
				//检查现在的分类id是否为0，为0 需要删除，不为0需要修改
				if savePost.CatID != 0 {
					stmt, _ = trainaction.Prepare("update wps_post_cate set cat_id = ? where post_id = ?")
					_, err = stmt.Exec(savePost.CatID, savePost.PostID)
					if err != nil {
						fmt.Println(err)
						trainaction.Rollback()
						return 0, errors.New("修改文章分类失败")
					}
				} else {
					stmt, _ = trainaction.Prepare("delete from wps_post_cate where post_id = ? and cat_id = ?")
					_, err = stmt.Exec(savePost.PostID, catID)
					if err != nil {
						trainaction.Rollback()
						return 0, errors.New("删除文章分类失败")
					}
				}
			}
		}
		if len(savePost.Tags) > 0 {
			//先检查老的tag是否存在
			var oldTags []PostTag
			rows, _ := trainaction.Query(fmt.Sprintf("select t.tag_id, t.name from wps_post_tags as pt left join wps_tags as t on pt.tag_id = t.tag_id where pt.post_id = %d", savePost.PostID))
			for rows.Next() {
				var oneTag PostTag
				err := rows.Scan(&oneTag.ID, &oneTag.Name)
				if err == nil {
					oldTags = append(oldTags, oneTag)
				}
			}
			if len(oldTags) == 0 {
				for _, tag := range savePost.Tags {
					if strings.TrimSpace(tag) != "" {
						var tagID int64
						err = trainaction.QueryRow(fmt.Sprintf("select tag_id from wps_tags where name = '%s'", tag)).Scan(&tagID)
						if err != nil {
							stmt, _ = trainaction.Prepare("insert into wps_tags (name, slug, counts) values (?,?,?)")
							tagRes, _ := stmt.Exec(tag, url.QueryEscape(tag), 1)
							tagID, _ = tagRes.LastInsertId()
						}
						stmt, _ = trainaction.Prepare("insert into wps_post_tags (post_id, tag_id, disabled, create_time) values (?,?,?,?)")
						_, err = stmt.Exec(savePost.PostID, tagID, 0, nowTimestamp)
						if err != nil {
							trainaction.Rollback()
							return 0, errors.New("插入标签失败")
						}
					}
				}
			} else {
				insertTag := []string{}
				deleteTagID := []int{}
				//检查老的tag是否在新传过来的数据中，如果没有，需要删除
				for _, tag := range oldTags {
					deleteFlag := 0
					for _, nTag := range savePost.Tags {
						if nTag == tag.Name && tag.Name != "" {
							deleteFlag = 1
						}
					}
					if deleteFlag == 0 {
						deleteTagID = append(deleteTagID, tag.ID)
					}
				}
				//检查新的tag是否已经存在，存在不需要删除
				for _, nTag := range savePost.Tags {
					insertFlag := 1
					for _, oTag := range oldTags {
						if nTag == oTag.Name {
							insertFlag = 0
						}
					}
					if insertFlag == 1 {
						insertTag = append(insertTag, nTag)
					}
				}
				if len(insertTag) > 0 {
					for _, tag := range insertTag {
						if strings.TrimSpace(tag) != "" {
							var tagID int64
							err = trainaction.QueryRow(fmt.Sprintf("select tag_id from wps_tags where name = '%s'", tag)).Scan(&tagID)
							if err != nil {
								stmt, _ = trainaction.Prepare("insert into wps_tags (name, slug, counts) values (?,?,?)")
								tagRes, _ := stmt.Exec(tag, url.QueryEscape(tag), 1)
								tagID, _ = tagRes.LastInsertId()
							}
							stmt, _ = trainaction.Prepare("insert into wps_post_tags (post_id, tag_id, disabled, create_time) values (?,?,?,?)")
							_, err = stmt.Exec(savePost.PostID, tagID, 0, nowTimestamp)
							if err != nil {
								trainaction.Rollback()
								return 0, errors.New("插入标签失败")
							}
						}
					}
				}
				if len(deleteTagID) > 0 && savePost.PostID > 0 {
					stmt, _ = trainaction.Prepare(fmt.Sprintf("delete from wps_post_tags where post_id = ? and tag_id in (%s)", util.ArrayIntToString(deleteTagID, ",")))
					_, err = stmt.Exec(savePost.PostID)
					if err != nil {
						trainaction.Rollback()
						return 0, errors.New("插入删除标签失败")
					}
				}
			}
		}
	}
	if savePost.PostID > 0 {
		//再次检查 wps_post_nums 是否存在
		var postID int
		err = trainaction.QueryRow(fmt.Sprintf("select post_id from wps_post_nums where post_id = %d", savePost.PostID)).Scan(&postID)
		if err != nil {
			stmt, _ := trainaction.Prepare("insert into wps_post_nums (post_id) values (?)")
			_, err = stmt.Exec(savePost.PostID)
			if err != nil {
				trainaction.Rollback()
				return 0, errors.New("插入nums失败")
			}
		}
	}
	trainaction.Commit()
	return savePost.PostID, nil
}
