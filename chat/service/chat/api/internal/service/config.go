package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars/myerror"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigService struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	ConfigRepository *repository.ConfigRepository
}

func NewConfigService(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigService {
	return &ConfigService{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		ConfigRepository: repository.NewConfigRepository(ctx, svcCtx),
	}
}

func (l *ConfigService) Update(id int64, model *model.Config) error {
	if model.Value != "" {
		exist, err := l.ConfigRepository.GetByName(model.Value, 0)
		if nil != err {
			return err
		}
		// 判断是否已经存在
		if exist != nil && exist.Id != id {
			err = myerror.ConfigExist
			return err
		}
	}
	return l.ConfigRepository.Update(id, model)
}

func (l *ConfigService) GetAll(configType int64) (pos []*model.Config, err error) {
	return l.ConfigRepository.GetAll(configType)
}

func (l *ConfigService) GetById(id int64) (pos *model.Config, err error) {
	return l.ConfigRepository.GetById(id)
}

func (l *ConfigService) Insert(dto *model.Config) (err error) {
	exist, err := l.ConfigRepository.GetByName(dto.Value, dto.ConfigType)
	if nil != err {
		return
	}
	// 判断是否已经存在
	if exist != nil && exist.Id > 0 {
		err = myerror.ConfigExist
		return
	}

	_, err = l.ConfigRepository.Insert(dto)
	return
}

func (l *ConfigService) GetConfigKey() (apiKeys []string, err error) {
	info, err := l.GetAll(repository.ConfigTypeOpenAi)
	if nil != err {
		return
	}
	if info != nil && len(info) > 0 {
		for _, v := range info {
			apiKeys = append(apiKeys, v.Value)
		}
	}
	if len(apiKeys) < 1 {
		err = myerror.ConfigOpenAiKeyEmpty
		return
	}
	return
}
