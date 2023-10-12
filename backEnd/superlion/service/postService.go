//@program: superlion
//@author: yanjl
//@create: 2023-09-14 16:31
package service

import (
	"log"
	"superlion/bean"
	"superlion/model"
	"superlion/repository"
	"sync"
)

var (
	postDao = repository.NewPostEntityDaoInstance()

	postService *PostService
	pService    sync.Once
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

// 保存文章（发布）
func (*PostService) PublishPost(postReq *bean.SavePostReq, login *LionUserInfo) (string, string) {

	if login == nil {
		return "", "您还未登录！"
	}
	if len(postReq.Title) == 0 {
		return "", "文章标题不可以为空"
	}
	return "", ""
}
