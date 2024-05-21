//@program: superlion
//@author: yanjl
//@create: 2023-09-14 16:31
package service

import (
	"log"
	"superlion/bean"
	"superlion/constants"
	"superlion/model"
	"superlion/repository"
	"sync"
	"time"
)

var (
	postDao    = repository.NewPostEntityDaoInstance()
	tagPostDao = repository.NewTagPostMapInstance()

	postService *PostService
	pService    sync.Once

	postType constants.PostTypeCode
)

type PostService struct {
}

func GetPostServiceInstance() *PostService {
	pService.Do(
		func() {
			postService = &PostService{}
		})
	return postService
}

func (*PostService) AddNewPost() {

}

// FindPostById 查询文章内容
func (*PostService) FindPostById(id int64) (*model.PostEntity, string) {

	log.Print("find post by id:", id)
	if id == 0 {
		return nil, "文章不存在"
	}

	// 查询文章内容明细
	post, err := postDao.FindPostById(id)

	if len(err) != 0 {
		return nil, err
	}
	// 封装文章返回内容

	return post, err
}

// PublishPost 保存文章（发布）
func (*PostService) PublishPost(postReq *bean.SavePostReq, login *LionUserInfo) (int64, string) {

	if login == nil {
		return 0, "您还未登录！"
	}
	if len(postReq.Title) == 0 {
		return 0, "文章标题不可以为空"
	}

	// 文章最多10个标签
	if len(postReq.Tags) > 9 {
		return 0, "文章标签最多10个！"
	}
	post := &model.LionPostEntity{
		Title:           postReq.Title,
		AuthorId:        login.GoId,
		AuditState:      constants.STATUS_OK,
		Category:        constants.POST_PUBLISH,
		ContentType:     postReq.ContentType,
		MarkdownContent: postReq.MarkdownContent,
		HtmlContent:     postReq.HtmlContent,
		Views:           0,
		Approvals:       0, // 点赞量
		Collection:      0, // 收藏量
		Comments:        0,
		TypeId:          0,
		HeadImg:         postReq.HeadImg,
		Official:        0,
		Top:             0,
		Marrow:          0, // 精华
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
	}

	// 保存文章
	postId, err := postDao.AddNewPost(post)
	// 保存文章-标签
	tags := postReq.Tags
	TagPostMaps := make([]model.LionTagPostsMapping, len(tags))

	for i, tag := range tags {
		tagMap := &model.LionTagPostsMapping{
			PostsId:  postId,
			TagId:    tag,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		}
		TagPostMaps[i] = *tagMap
	}
	// 保存文章-标签映射
	err = tagPostDao.SaveTagPostMapList(TagPostMaps)
	if len(err) != 0 {
		return 0, err
	}
	return postId, ""
}

// 查询我的文章列表
func (*PostService) GetMyPostList(login *LionUserInfo) (*bean.PageResult, string) {

	if login == nil || len(login.GoId) == 0 {
		return nil, "login first"
	}

	goId := login.GoId

	entitys, total, err := postDao.GetMyPostList(goId)

	if len(err) != 0 {
		return nil, err
	}

	// 组装结果集
	datas := make([]any, len(entitys))
	for i, post := range entitys {

		// 实体类转bean
		postBean := &bean.PostBeanRsp{
			Id:         post.Id,
			Title:      post.Title,
			HeadImg:    post.HeadImg,
			Official:   post.Official,
			AuditState: post.AuditState,
			Views:      post.Views,
			Approvals:  post.Approvals,  // 点赞量
			Collection: post.Collection, // 收藏量
			Sort:       post.Sort,
			// AuthorId   string `json:"authorId"`
			Preview: "预览内容", // 预览内容
			// Tags    :post. `json:"tags"`    //gorm:"foreignKey:tagId;"
		}

		datas[i] = postBean
	}

	pageRsp := &bean.PageResult{
		Data:  datas,
		Total: total,
		Code:  0,
	}

	return pageRsp, ""
}

//CommentId       0,
//IsDelete        int32,
//CreateAt        : time. jsonTime
//UpdateAt
//Tags
