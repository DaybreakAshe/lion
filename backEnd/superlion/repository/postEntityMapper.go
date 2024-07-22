//@program: superlion
//@author: yanjl
//@create: 2023-10-08 17:32
package repository

import (
	"fmt"
	"log"
	"superlion/model"
	"superlion/util"
	"sync"
)

type PostEntityDao struct {
}

var (
	postDao       sync.Once
	postEntityDao *PostEntityDao
)

func NewPostEntityDaoInstance() *PostEntityDao {
	postDao.Do(
		func() {
			postEntityDao = &PostEntityDao{}
		},
	)
	return postEntityDao
}

// 保存文章
func (*PostEntityDao) AddNewPost(post *model.LionPostEntity) (int64, string) {

	// 插入:
	err := db.Create(post).Error

	if err != nil {
		log.Printf("插入失败：%s\n", err.Error())
		return post.Id, err.Error()
	}
	//id := db.
	return post.Id, ""
}

// 根据goid查询文章列表
func (*PostEntityDao) GetMyPostList(condi *model.LionPostEntity) ([]model.LionPostEntity, int64, string) {

	//if len(goId) == 0 {
	//	return nil, 0, ""
	//}

	// 分页参数
	page := &model.PageDto{PageSize: 10, Page: 1}

	var total int64 = 0
	//// 组装查询条件
	//var condi = &model.LionPostEntity{
	//	AuthorId: goId,
	//	AuditState: constants.POST_AUDIT_STATE_YES,
	//}

	// 接收查询结果
	var entitys []model.LionPostEntity

	title := condi.Title
	// status := condi.AuditState
	condi.Title = ""
	// condi.AuditState = ""
	//db := db.Scopes(util.Paginate(page)).
	//	Table("lion_post lp").
	//	Select("lp.*, lt.id TagId,lt.name Tag").
	//	Where("lp.audit_state = ?", status).
	//	Joins("left join lion_tag_posts_mapping ltpm on ltpm.post_id = lp.id").
	//	Joins("left join lion_tag lt on ltpm.tag_id = lt.id").
	//	Where(condi)
	db := db.Scopes(util.Paginate(page)).
		Where(condi)

	log.Printf("like query :%s", title)
	// 模糊搜索标题
	if len(title) != 0 {
		db = db.Where("lion_post.title like ?", fmt.Sprintf("%%%s%%", title))
	}

	//if len(condi.TagIds) != 0 {
	//	log.Printf("select tags :%s", condi.TagIds)
	//	db = db.Where("tags.id in (?)", condi.TagIds)
	//}

	err := db.Preload("Tags").Find(&entitys).Count(&total).Error

	if err != nil {
		return nil, 0, err.Error()
	}

	return entitys, total, ""
}

// 查询文章详细
func (*PostEntityDao) FindPostById(id int64) (*model.LionPostEntity, string) {

	post := &model.LionPostEntity{}

	// var tags []model.Tag
	err := db.Debug().
		//Table("lion_post lp").
		//Select("lp.id, lp.author_id, lp.title, lp.content_type, lp.markdown_content, lp.html_content, lp.views, lp.approvals, lp.head_img, lp.sort ,lp.audit_state , lp.type_id, lp.create_at , lp.update_at ,lp.official , lp.top ,lp.category,lu.avatar, lu.login_name authorName,lt.id as tagId, lt.name as tags").
		//Joins("left join lion_user lu on lp.author_id = lu.go_id").
		//Joins("left join lion_tag_posts_mapping ltpm on lp.id = ltpm.post_id ").
		//Joins("left join lion_tag lt on lt.id = ltpm.tag_id ").
		Where("lion_post.id = ?", id).
		Preload("Tags"). //.Find(&post.Tags).
		Find(post).
		Error
	// Preload("tags").
	// post.Tags = &tags
	if err != nil {
		log.Println("查询文章失败", err.Error())
		return nil, err.Error()
	} else {
		return post, ""
	}
}
