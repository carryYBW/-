package service

import (
	"errors"
	"fmt"
	"helloweb/model"
	"helloweb/util"
	"math/rand"
	"time"
)

type UserService struct {
}

// 注册函数
func (s *UserService) Register(
	mobile,   //手机
	plainpwd, //明文密码
	nickname, //昵称
	avatar, sex string) (user model.User, err error) {

	//检测手机号码是否存在,
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//如果存在则返回提示已经注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}
	//否则拼接插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()
	//token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	//passwd =
	//md5 加密
	//返回新用户信息

	//插入 InserOne
	_, err = DbEngin.InsertOne(&tmp)
	//前端恶意插入特殊字符
	//数据库连接操作失败
	return tmp, err
}

// 登录函数
func (s *UserService) Login(
	mobile, //手机
	plainpwd string) (user model.User, err error) {
	tmp := model.User{}
	//通过 手机号查询用户
	DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	//密码验证
	passwd := tmp.Passwd
	salt := tmp.Salt
	validatePasswd := util.ValidatePasswd(plainpwd, salt, passwd)
	if !validatePasswd {
		return tmp, errors.New("密码不正确")
	}
	//涮新token，安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)

	return tmp, nil
}

//查找某个用户
func (s *UserService) Find(userId int64) (user model.User) {
	tmp := model.User{}
	DbEngin.ID(userId).Get(&tmp)
	return tmp
}
