package workwx

import (
	"time"
)

// ListExternalContact 获取客户列表
func (c *WorkwxApp) ListExternalContact(userID string) ([]string, error) {
	resp, err := c.execExternalContactList(reqExternalContactList{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExternalUserID, nil
}

// GetExternalContact 获取客户详情
func (c *WorkwxApp) GetExternalContact(externalUserID string) (*ExternalContactInfo, error) {
	resp, err := c.execExternalContactGet(reqExternalContactGet{
		ExternalUserID: externalUserID,
	})
	if err != nil {
		return nil, err
	}
	return &resp.ExternalContactInfo, nil
}

// BatchListExternalContact 批量获取客户详情
func (c *WorkwxApp) BatchListExternalContact(userID string, cursor string, limit int) (*BatchListExternalContactsResp, error) {
	resp, err := c.execExternalContactBatchList(reqExternalContactBatchList{
		UserID: userID,
		Cursor: cursor,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	return &BatchListExternalContactsResp{Result: resp.ExternalContactList, NextCursor: resp.NextCursor}, nil
}

// RemarkExternalContact 修改客户备注信息
func (c *WorkwxApp) RemarkExternalContact(req *ExternalContactRemark) error {
	_, err := c.execExternalContactRemark(reqExternalContactRemark{
		Remark: req,
	})
	return err
}

// ListExternalContactCorpTags 获取企业标签库
func (c *WorkwxApp) ListExternalContactCorpTags(tagIDs ...string) ([]ExternalContactCorpTagGroup, error) {
	resp, err := c.execExternalContactListCorpTags(reqExternalContactListCorpTags{
		TagIDs: tagIDs,
	})
	if err != nil {
		return nil, err
	}
	return resp.TagGroup, nil
}

// AddExternalContactCorpTag 添加企业客户标签
func (c *WorkwxApp) AddExternalContactCorpTag(req ExternalContactAddCorpTagGroup) (ExternalContactCorpTagGroup, error) {
	resp, err := c.execExternalContactAddCorpTag(reqExternalContactAddCorpTagGroup{
		ExternalContactAddCorpTagGroup: req,
	})
	if err != nil {
		return ExternalContactCorpTagGroup{}, err
	}
	return resp.TagGroup, nil
}

// EditExternalContactCorpTag 编辑企业客户标签
func (c *WorkwxApp) EditExternalContactCorpTag(id, name string, order uint32) error {
	_, err := c.execExternalContactEditCorpTag(reqExternalContactEditCorpTag{
		ID:    id,
		Name:  name,
		Order: order,
	})
	return err
}

// DelExternalContactCorpTag 删除企业客户标签
func (c *WorkwxApp) DelExternalContactCorpTag(tagID, groupID []string) error {
	_, err := c.execExternalContactDelCorpTag(reqExternalContactDelCorpTag{
		TagID:   tagID,
		GroupID: groupID,
	})
	return err
}

// MarkExternalContactTag 标记客户企业标签
func (c *WorkwxApp) MarkExternalContactTag(userID, externalUserID string, addTag, removeTag []string) error {
	_, err := c.execExternalContactMarkTag(reqExternalContactMarkTag{
		UserID:         userID,
		ExternalUserID: externalUserID,
		AddTag:         addTag,
		RemoveTag:      removeTag,
	})
	return err
}

// ExternalContactUnassigned 离职成员的客户
type ExternalContactUnassigned struct {
	// HandoverUserID 离职成员的userid
	HandoverUserID string
	// ExternalUserID 外部联系人userid
	ExternalUserID string
	// DemissionTime 成员离职时间
	DemissionTime time.Time
}

// ListUnassignedExternalContact 获取离职成员的客户列表
func (c *WorkwxApp) ListUnassignedExternalContact(pageID, pageSize uint32, cursor string) (*ExternalContactUnassignedList, error) {
	resp, err := c.execListUnassignedExternalContact(reqListUnassignedExternalContact{
		PageID:   pageID,
		PageSize: pageSize,
		Cursor:   cursor,
	})
	if err != nil {
		return nil, err
	}
	externalContactUnassignedList := resp.intoExternalContactUnassignedList()
	return &externalContactUnassignedList, nil
}

// TransferExternalContact 分配成员的客户
func (c *WorkwxApp) TransferExternalContact(externalUserID, handoverUserID, takeoverUserID, transferSuccessMsg string) error {
	_, err := c.execTransferExternalContact(reqTransferExternalContact{
		ExternalUserID:     externalUserID,
		HandoverUserID:     handoverUserID,
		TakeoverUserID:     takeoverUserID,
		TransferSuccessMsg: transferSuccessMsg,
	})
	return err
}

// ExternalContactTransferResult 客户接替结果
type ExternalContactTransferResult struct {
	// Status 接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限 5-无接替记录
	Status ExternalContactTransferStatus
	// TakeoverTime 接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
	TakeoverTime time.Time
}

// GetTransferExternalContactResult 查询客户接替结果
func (c *WorkwxApp) GetTransferExternalContactResult(externalUserID, handoverUserID, takeoverUserID string) (*ExternalContactTransferResult, error) {
	resp, err := c.execGetTransferExternalContactResult(reqGetTransferExternalContactResult{
		ExternalUserID: externalUserID,
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
	})
	if err != nil {
		return nil, err
	}
	externalContactTransferResult := resp.intoExternalContactTransferResult()
	return &externalContactTransferResult, nil
}

// ExternalContactTransferGroupChat 离职成员的群再分配
func (c *WorkwxApp) ExternalContactTransferGroupChat(chatIDList []string, newOwner string) ([]ExternalContactGroupChatTransferFailed, error) {
	resp, err := c.execTransferGroupChatExternalContact(reqTransferGroupChatExternalContact{
		ChatIDList: chatIDList,
		NewOwner:   newOwner,
	})
	if err != nil {
		return nil, err
	}
	return resp.FailedChatList, nil
}

// TransferCustomer 在职继承 分配在职成员的客户
// 一次最多转移100个客户
// 为保障客户服务体验，90个自然日内，在职成员的每位客户仅可被转接2次
func (c *WorkwxApp) TransferCustomer(handoverUserID, takeoverUserID string, externalUserIDs []string) (TransferCustomerResult, error) {
	resp, err := c.execTransferCustomer(reqTransferCustomer{
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
		ExternalUserID: externalUserIDs,
	})
	result := resp.intoTransferCustomerResult()
	return result, err
}

type TransferCustomerResult []struct {
	// ExternalUserID 转接客户的外部联系人userid
	ExternalUserID string `json:"external_userid"`
	// Errcode 对此客户进行分配的结果, 具体可参考全局错误码(https://work.weixin.qq.com/api/doc/90000/90135/92125#10649), 0表示成功发起接替,待24小时后自动接替,并不代表最终接替成功
	Errcode int `json:"errcode"`
}

// GetTransferCustomerResult 在职继承 查询客户接替状态
func (c *WorkwxApp) GetTransferCustomerResult(handoverUserID, takeoverUserID, cursor string) (*CustomerTransferResult, error) {
	resp, err := c.execGetTransferCustomerResult(reqGetTransferCustomerResult{
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
		Cursor:         cursor,
	})
	if err != nil {
		return nil, err
	}

	result := resp.intoCustomerTransferResult()
	return &result, nil
}

type CustomerTransferResult struct {
	Customer []struct {
		// ExternalUserID 转接客户的外部联系人userid
		ExternalUserID string `json:"external_userid"`
		// Status 接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限 5-无接替记录
		Status int `json:"status"`
		// TakeoverTime 接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
		TakeoverTime int `json:"takeover_time"`
	} `json:"customer"`
	// NextCursor 下个分页的起始cursor
	NextCursor string `json:"next_cursor"`
}

// ResignedTransferCustomer 离职继承 分配离职成员的客户
// 一次最多转移100个客户
func (c *WorkwxApp) ResignedTransferCustomer(handoverUserID, takeoverUserID string, externalUserIDs []string) (TransferCustomerResult, error) {
	resp, err := c.execTransferResignedCustomer(reqTransferCustomer{
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
		ExternalUserID: externalUserIDs,
	})
	result := resp.intoTransferCustomerResult()
	return result, err
}

// GetTransferResignedCustomerResult 离职继承 查询客户接替状态
func (c *WorkwxApp) GetTransferResignedCustomerResult(handoverUserID, takeoverUserID, cursor string) (*CustomerTransferResult, error) {
	resp, err := c.execGetTransferResignedCustomerResult(reqGetTransferCustomerResult{
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
		Cursor:         cursor,
	})
	if err != nil {
		return nil, err
	}

	result := resp.intoCustomerTransferResult()
	return &result, nil
}

// ExternalContactListFollowUser 获取配置了客户联系功能的成员列表
func (c *WorkwxApp) ExternalContactListFollowUser() (*ExternalContactFollowUserList, error) {
	resp, err := c.execListFollowUserExternalContact(reqListFollowUserExternalContact{})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactFollowUserList, nil
}

// ExternalContactAddContact 配置客户联系「联系我」方式
func (c *WorkwxApp) ExternalContactAddContact(t int, scene int, style int, remark string, skipVerify bool, state string, user []string, party []int, isTemp bool, expiresIn int, chatExpiresIn int, unionID string, conclusions Conclusions) (*ExternalContactAddContact, error) {
	resp, err := c.execAddContactExternalContact(
		reqAddContactExternalContact{
			ExternalContactWay{
				Type:          t,
				Scene:         scene,
				Style:         style,
				Remark:        remark,
				SkipVerify:    skipVerify,
				State:         state,
				User:          user,
				Party:         party,
				IsTemp:        isTemp,
				ExpiresIn:     expiresIn,
				ChatExpiresIn: chatExpiresIn,
				UnionID:       unionID,
				Conclusions:   conclusions,
			},
		})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactAddContact, nil
}

// ExternalContactGetContactWay 获取企业已配置的「联系我」方式
func (c *WorkwxApp) ExternalContactGetContactWay(configID string) (*ExternalContactContactWay, error) {
	resp, err := c.execGetContactWayExternalContact(reqGetContactWayExternalContact{ConfigID: configID})
	if err != nil {
		return nil, err
	}

	return &resp.ContactWay, nil
}

// ExternalContactListContactWayChat 获取企业已配置的「联系我」列表
func (c *WorkwxApp) ExternalContactListContactWayChat(startTime int, endTime int, cursor string, limit int) (*ExternalContactListContactWayChat, error) {
	resp, err := c.execListContactWayChatExternalContact(reqListContactWayExternalContact{
		StartTime: startTime,
		EndTime:   endTime,
		Cursor:    cursor,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}

	return &resp.ExternalContactListContactWayChat, nil
}

// ExternalContactUpdateContactWay 更新企业已配置的「联系我」成员配置
func (c *WorkwxApp) ExternalContactUpdateContactWay(configID string, remark string, skipVerify bool, style int, state string, user []string, party []int, expiresIn int, chatExpiresIn int, unionid string, conclusions Conclusions) error {
	_, err := c.execUpdateContactWayExternalContact(reqUpdateContactWayExternalContact{
		ConfigID:      configID,
		Remark:        remark,
		SkipVerify:    skipVerify,
		Style:         style,
		State:         state,
		User:          user,
		Party:         party,
		ExpiresIn:     expiresIn,
		ChatExpiresIn: chatExpiresIn,
		UnionID:       unionid,
		Conclusions:   conclusions,
	})

	return err
}

// ExternalContactDelContactWay 删除企业已配置的「联系我」方式
func (c *WorkwxApp) ExternalContactDelContactWay(configID string) error {
	_, err := c.execDelContactWayExternalContact(reqDelContactWayExternalContact{ConfigID: configID})

	return err
}

// ExternalContactCloseTempChat 结束临时会话
func (c *WorkwxApp) ExternalContactCloseTempChat(userID, externalUserID string) error {
	_, err := c.execCloseTempChatExternalContact(reqCloseTempChatExternalContact{
		UserID:         userID,
		ExternalUserID: externalUserID,
	})

	return err
}

// AddMsgTemplate 创建企业群发
// https://developer.work.weixin.qq.com/document/path/92135
func (c *WorkwxApp) AddMsgTemplate(chatType ChatType, sender string, externalUserID []string, text Text, attachments []Attachments) (*AddMsgTemplateDetail, error) {
	resp, err := c.execAddMsgTemplate(reqAddMsgTemplateExternalContact{
		AddMsgTemplateExternalContact{
			ChatType:       chatType,
			ExternalUserID: externalUserID,
			Sender:         sender,
			Text:           text,
			Attachments:    attachments,
		},
	})
	if err != nil {
		return nil, err
	}

	return &resp.AddMsgTemplateDetail, nil
}

// GetKFSyncMsg 读取客服消息
// https://developer.work.weixin.qq.com/document/path/94670
func (c *WorkwxApp) GetKFSyncMsg(cursor, token, openKfId string, limit, voiceFormat int) (*KFSyncMsgDetail, error) {
	msg, err := c.execKFSyncMsg(reqKFSyncMsgContact{
		KFSyncMsgContact{
			Cursor:      cursor,
			Token:       token,
			Limit:       limit,
			VoiceFormat: voiceFormat,
			OpenKfid:    openKfId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &msg.KFSyncMsgDetail, nil
}
