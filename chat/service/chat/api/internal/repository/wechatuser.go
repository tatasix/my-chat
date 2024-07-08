package repository

import (
	"chat/common/util"
	"chat/common/wecom"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const (
	UserTypeWechat  = 1
	UserTypeVisitor = 2
	UserTypePc      = 3

	UserStatusSuccess  = 1  //注册成功
	UserStatusMobile   = 2  //填写手机号
	UserStatusPassword = 3  // 填写密码
	UserStatusFail     = 10 //禁用

	UserLevelVip = 1
)

type WechatUserRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatUserRepository(ctx context.Context, svcCtx *svc.ServiceContext) *WechatUserRepository {
	return &WechatUserRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatUserRepository) GetByName(nickname string) (customerPo *model.WechatUser, err error) {

	customerPo, err = l.svcCtx.WechatUserModel.FindOneByQuery(l.ctx,
		l.svcCtx.WechatUserModel.RowBuilder().Where(squirrel.Eq{"nickname": nickname}),
	)
	return
}

func (l *WechatUserRepository) GetByUsers(users []string) (wechatUserPo []*model.WechatUser, err error) {
	return l.svcCtx.WechatUserModel.FindAll(l.ctx,
		l.svcCtx.WechatUserModel.RowBuilder().Where(squirrel.Eq{"user": users}),
	)
}

func (l *WechatUserRepository) GetByUser(user string) (wechatUserPo *model.WechatUser, err error) {
	return l.svcCtx.WechatUserModel.FindOneByQuery(l.ctx,
		l.svcCtx.WechatUserModel.RowBuilder().Where(squirrel.Eq{"user": user}),
	)
}

func (l *WechatUserRepository) InsertWechatUserIfNotExist(user string, isWechatUser bool) {
	ctx := l.ctx
	if user == "" {
		return
	}
	//先查询下
	wechatUserPo, err := l.GetByUser(user)
	if err != nil {
		fmt.Printf("InsertWechatUser FindOneByQuery error: %v", err)
		return
	}
	//已存在就不需要重复去查询并且插入了
	if wechatUserPo != nil && wechatUserPo.Nickname != "" {
		return
	}

	if !isWechatUser {
		if wechatUserPo != nil {
			return
		}
		_, err = l.svcCtx.WechatUserModel.Insert(ctx, &model.WechatUser{
			User:     user,
			UserType: UserTypeVisitor,
		})
		return
	}

	res, err := wecom.GetCustomer([]string{user})
	if err != nil {
		return
	}
	if len(res) > 0 {
		wechatInfo := res[0]
		wechatModel := &model.WechatUser{
			User:     user,
			Nickname: wechatInfo.Nickname,
			Avatar:   wechatInfo.Avatar,
			Gender:   wechatInfo.Gender,
			UserType: UserTypeWechat,
			Unionid:  wechatInfo.Unionid,
		}
		if wechatUserPo != nil && wechatUserPo.Id > 0 {
			wechatUserPo.User = user
			wechatUserPo.Nickname = wechatInfo.Nickname
			wechatUserPo.Avatar = wechatInfo.Avatar
			wechatUserPo.Gender = wechatInfo.Gender
			wechatUserPo.Unionid = wechatInfo.Unionid
			wechatUserPo.UpdatedAt = time.Now()
			wechatUserPo.Mobile = util.Encrypt(wechatUserPo.Mobile)
			err = l.svcCtx.WechatUserModel.Update(ctx, wechatUserPo)
			return
		}
		_, err = l.svcCtx.WechatUserModel.Insert(ctx, wechatModel)
	}
	return
}

func (l *WechatUserRepository) UpdateById(id int64, wechatUser *model.WechatUser) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.WechatUserModel.BuildFiled(old, wechatUser)
	return l.svcCtx.WechatUserModel.Update(context.Background(), wechatUser)
}

func (l *WechatUserRepository) UpdateByUser(user string, wechatUser *model.WechatUser) error {
	old, err := l.GetByUser(user)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}

	l.svcCtx.WechatUserModel.BuildFiled(old, wechatUser)
	return l.svcCtx.WechatUserModel.Update(context.Background(), wechatUser)
}

func (l *WechatUserRepository) GetById(id int64) (wechatUser *model.WechatUser, err error) {
	wechatUser, err = l.svcCtx.WechatUserModel.FindOne(context.Background(), id)
	if err != nil {
		return
	}
	wechatUser.Mobile = util.Decrypt(wechatUser.Mobile)
	return
}

func (l *WechatUserRepository) GetByMobile(mobile string) (wechatUserPo *model.WechatUser, err error) {
	mobile = util.Encrypt(mobile)
	return l.svcCtx.WechatUserModel.FindOneByQuery(l.ctx,
		l.svcCtx.WechatUserModel.RowBuilder().Where(squirrel.Eq{"mobile": mobile}),
	)
}

func (l *WechatUserRepository) Insert(wechatUser *model.WechatUser) (err error) {
	l.svcCtx.WechatUserModel.BuildFiled(nil, wechatUser)
	_, err = l.svcCtx.WechatUserModel.Insert(l.ctx, wechatUser)
	return
}

func (l *WechatUserRepository) Update(old, wechatUser *model.WechatUser) error {
	l.svcCtx.WechatUserModel.BuildFiled(old, wechatUser)
	return l.svcCtx.WechatUserModel.Update(l.ctx, wechatUser)
}

func (l *WechatUserRepository) GetCountUser(startTime, endTime string, checkMobile bool) (count int64, err error) {
	countBuilder := l.svcCtx.WechatUserModel.CountBuilder("id")

	if startTime != "" {
		countBuilder = countBuilder.Where("created_at >= ?", startTime)
	}

	if endTime != "" {
		countBuilder = countBuilder.Where("created_at <= ?", endTime)
	}

	if checkMobile {
		countBuilder = countBuilder.Where("mobile != ''")
	} else {
		countBuilder = countBuilder.Where("mobile = ''")
	}

	return l.svcCtx.WechatUserModel.FindCount(l.ctx, countBuilder)
}

func (l *WechatUserRepository) UpdateLevelById(id, level int64, levelExpire string) error {
	return l.svcCtx.WechatUserModel.UpdateLevelById(l.ctx, id, level, levelExpire)
}

func (l *WechatUserRepository) UpdateLevel() error {
	return l.svcCtx.WechatUserModel.UpdateLevel(l.ctx, time.Now().Format(vars.TimeFormat))
}

func (l *WechatUserRepository) DeleteByUser(user string) error {
	return l.svcCtx.WechatUserModel.DeleteByUser(l.ctx, user)
}

func (l *WechatUserRepository) GetAllHaveMobile() (wechatUserPo []*model.WechatUser, err error) {
	return l.svcCtx.WechatUserModel.FindAll(l.ctx,
		l.svcCtx.WechatUserModel.RowBuilder().Where("mobile != ''"),
	)
}

func (l *WechatUserRepository) UpdateV2(wechatUser *model.WechatUser) error {
	return l.svcCtx.WechatUserModel.Update(l.ctx, wechatUser)
}
