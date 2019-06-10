package service

import (
	"errors"
	"webchat/app/model"
	"time"
)

type ContactService struct {

}

func (c *ContactService)AddFriend(userId, addId int ) error  {
	if userId < 1 || addId < 1 {
		return errors.New("传参错误")
	}

	if userId == addId {
		return errors.New("不能添加自己为好友")
	}

	tmp := model.Contact{}
	Db.Where("UserId = ?", userId).And("AddId = ?", addId).And("Type = ?", model.CONCAT_TYPE_USER).Get(&tmp)
	if tmp.ContactId > 0 {
		return  errors.New("该用户已经被添加过啦")
	}
	createTime := time.Now().Format("2006-01-02 15:04:05")
	session := Db.NewSession() // 开启一个事物
	_, e1 := session.InsertOne(model.Contact{
		UserId:userId,
		AddId:addId,
		Type:model.CONCAT_TYPE_USER,
		CreateTime:createTime,
	})
	_, e2 := session.InsertOne(model.Contact{
		UserId:addId,
		AddId:userId,
		Type:model.CONCAT_TYPE_USER,
		CreateTime:createTime,
	})
	if e1 == nil && e2 == nil {
		session.Commit()
		return nil
	} else {
		if e1 != nil {
			session.Rollback()
			return e1
		}
		if e2 != nil {
			session.Rollback()
			return e2
		}
	}
	return nil
}

func (c *ContactService)LoadFriend(userId int) ([]model.User) {
	// 查询符合条件的所有记录
	contacts := make([]model.Contact, 0)
	user := make([]model.User, 0)
	userIds := make([]int,0)
	err := Db.Where("UserId = ? and Type = ?", userId, model.CONCAT_TYPE_USER).Find(&contacts)
	if err != nil {
		return user
	}
	if len(contacts) < 1 {
		return user
	}
	for _,contact := range contacts {
		userIds = append(userIds, contact.AddId)
	}
	Db.In("UserId", userIds).Find(&user)

	return user
}

func (c *ContactService)LoadGroup(userId int) ([]model.Group) {
	// 查询符合条件的所有记录
	contacts := make([]model.Contact, 0)
	group := make([]model.Group, 0)
	groupIds := make([]int,0)
	err := Db.Where("UserId = ? and Type = ?", userId, model.CONCAT_TYPE_GROUP).Find(&contacts)
	if err != nil {
		return group
	}
	if len(contacts) < 1 {
		return group
	}
	for _,contact := range contacts {
		groupIds = append(groupIds, contact.AddId)
	}
	Db.In("GroupId", groupIds).Find(&group)

	return group
}

func (c *ContactService)CreateGroup(group model.Group) (model.Group, error) {
	// 查询符合条件的所有记录
	if group.UserId < 1 {
		return group, errors.New("用户id不能为空")
	}
	if group.Name == "" {
		return group, errors.New("群名称不能为空")
	}
	userOwnGroup := model.Group{
		UserId:group.UserId,
	}
	count,_ := Db.Count(&userOwnGroup)
	if count > 5 {
		return group, errors.New("每个人最多只能创建5个群")
	}
	session := Db.NewSession()
	session.Begin()
	// 先插入创建的数据
	_,err := session.InsertOne(&group)
	if err != nil {
		session.Rollback()
		return group, err
	}
	// 再将自己加入到群里面
	contact := model.Contact{
		UserId:group.UserId,
		AddId : group.GroupId,
		Type : model.CONCAT_TYPE_GROUP,
		CreateTime:time.Now().Format("2006-01-02 15:04:05"),
	}
	_,err = session.InsertOne(&contact)
	if err != nil {
		session.Rollback()
		return group, err
	}
	session.Commit()
	return group, nil
}

func (c *ContactService)JoinGroup(userId, GroupId int) (model.Contact, error) {
	contact := model.Contact{
		UserId:userId,
		AddId:GroupId,
		Type:model.CONCAT_TYPE_GROUP,
		CreateTime:time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err := Db.InsertOne(&contact)
	if err != nil {
		return contact, err
	}

	return contact, nil
}
