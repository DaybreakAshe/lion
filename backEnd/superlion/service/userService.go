//@program: superlion
//@author: yanjl
//@create: 2023-09-25 14:30
package service

import (
	"fmt"
	"log"
	"superlion/bean"
	"superlion/model"
	"superlion/repository"
	"sync"
)

var (
	tagDao = repository.NewTagDaoInstance()
)

type UserService struct {
}

var userService *UserService

/**在 Do 方法被调用后，该函数将被执行，而且只会执行一次，即使在多个协程同时调用的情况下也是如此*/
var tagDaoOnce sync.Once

// NewUserDaoInstance 单例构建Dao
func NewUserServiceInstance() *UserService {
	tagDaoOnce.Do(
		func() {
			userService = &UserService{}
		})
	return userService
}

// 新增一个标签
func (*UserService) CreateNewTag(tagBean *bean.TagReqBean, login *LionUserInfo) (int, string) {

	if len(tagBean.Description) == 0 {
		return -1, "标签名称不可以为控"
	}

	tagEntity := &model.LionTag{
		Name:        tagBean.Name,
		GroupName:   tagBean.GroupName,
		RefCount:    0,
		Description: tagBean.Description,
		CreatorId:   login.GoId,
		IsDelete:    0,
	}

	row, err := tagDao.AddNewTag(tagEntity)
	if row != 0 {
		return -1, err.Error()
	}

	fmt.Printf("user (%s) creat new tag:[%s]\n", login.GoId, tagBean.Name)
	return 0, ""
}

// 查询标签列表
func (*UserService) GetUserTagList(goId string) ([]bean.TagRspBean, string) {
	if len(goId) == 0 {
		return nil, "goid不可以为控"
	}
	fmt.Printf("find tags by go_id:(%s)\n", goId)
	tags, err := tagDao.FindUserTagsByGoId(goId)
	if err != nil {
		return nil, err.Error()
	}

	rspBeans := make([]bean.TagRspBean, len(tags))

	for index, tag := range tags {
		bean := bean.TagRspBean{
			Id:          tag.Id,
			Name:        tag.Name,
			Description: tag.Description,
			RefCnt:      tag.RefCount,
		}
		rspBeans[index] = bean
	}
	return rspBeans, ""
}

// 删除标签
func (*UserService) DeleteTag(params map[string]string, login *LionUserInfo) (int, string) {

	tagId := params["id"]
	goId := login.GoId
	_, err := tagDao.DeleteTag(tagId, goId)

	if len(err) != 0 {
		log.Printf("delete tag [%s] data failed :%s\n", tagId, err)
		return -1, err
	} else {
		return 0, ""
	}
}
