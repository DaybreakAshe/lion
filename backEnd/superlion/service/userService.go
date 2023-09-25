//@program: superlion
//@author: yanjl
//@create: 2023-09-25 14:30
package service

import (
	"fmt"
	"superlion/bean"
	"superlion/model"
	"superlion/repository"
)

var (
	tagDao = repository.NewTagDaoInstance()
)

// 新增一个标签
func CreateNewTag(tagBean *bean.TagReqBean, login *LionUserInfo) (int, string) {

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
func GetUserTagList(goId string) ([]bean.TagRspBean, string) {
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
