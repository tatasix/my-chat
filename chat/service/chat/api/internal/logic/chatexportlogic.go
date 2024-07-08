package logic

import (
	"chat/common/redis"
	"chat/common/util"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/vars"
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const MaxExportNumber = 5000

type ChatExportLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	chatService *service.ChatService
}

func NewChatExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatExportLogic {
	return &ChatExportLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		chatService: service.NewChatService(ctx, svcCtx),
	}
}

func (l *ChatExportLogic) ChatExport(req *types.GetChatListRequest) (resp *types.ChatHistoryExportReply, err error) {
	paramMD5 := util.MD5(req.Agent + req.User + req.Customer)
	key := l.ChatHistoryExportKey(paramMD5)
	fileUrl, err := redis.Rdb.Get(l.ctx, key).Result()
	if err == nil {
		return &types.ChatHistoryExportReply{File: fileUrl}, nil
	}
	req.Page = 1
	req.PageSize = MaxExportNumber
	list, count, err := l.chatService.GetExportList(req.User, req.ChatRecordUser, req.Customer, req.Agent, req.StartCreatedAt, req.EndCreatedAt, req.ChatType, "created_at asc", 0, req.Page, req.PageSize)
	if err != nil {
		fmt.Printf("GetSystemConfig error: %v", err)
		return
	}
	if list == nil || count <= 0 || len(list) <= 0 {
		return nil, fmt.Errorf("没有要导出的数据")
	}

	if count > MaxExportNumber {
		return nil, fmt.Errorf("导出数据太多，请加入筛选条件")
	}

	dirName := l.GetDirName()
	nowTime := time.Now().Format("150405")
	fileName := paramMD5 + nowTime + ".csv"
	fullFilePath := dirName + fileName
	fileHandle, err := os.Create(fullFilePath)
	if err != nil {
		fmt.Printf("create file error %v", err)
		return
	}
	defer fileHandle.Close()
	writer := csv.NewWriter(fileHandle)
	defer writer.Flush()

	// 写入CSV头部
	headers := []string{"ID", "user", "content", "time"}
	err = writer.Write(headers)
	if err != nil {
		return
	}

	// 写入CSV行
	i := 1
	for _, chatSon := range list {
		var user string
		//客户
		if chatSon.AnswerOrQuestion == repository.AnswerOrQuestionQuestion {
			user = chatSon.UserName
		} else {
			//机器人聊天
			if chatSon.ChatType == repository.ChatTypeApplication {
				user = chatSon.AgentName
			} else {
				user = chatSon.OpenKfName
			}
		}

		row1 := []string{strconv.Itoa(i), user, chatSon.Content, chatSon.CreatedAt}
		err = writer.Write(row1)
		if err != nil {
			return nil, err
		}
		i++
	}
	lastFile := l.svcCtx.Config.Domain + "/api/download/chat/history?file=" + fileName
	redis.Rdb.Set(l.ctx, key, lastFile, 5*time.Minute)
	return &types.ChatHistoryExportReply{File: lastFile}, nil
}

func (l *ChatExportLogic) ChatHistoryExportKey(key string) string {
	return fmt.Sprintf(redis.ChatHistoryExportKey, key)
}

func (l *ChatExportLogic) GetDirName() string {
	dir := vars.ChatHistoryDirectory + time.Now().Format("20060102")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Printf("Cannot create a file when that file already exists %v \n ", err)
	}
	return dir + "/"
}
