package service

import "webchat/app/model"

type UserService struct {

}

/**
用户注册逻辑
 */
func (s *UserService)Register(mobile, password, nickname, avatar string, sex int) (model.User, error) {
	user := model.User{}
	Db.Where("Mobile",mobile).Get(&user)


	return model.User{},nil
}
