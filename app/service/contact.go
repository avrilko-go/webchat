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
