# 聊天记录
CREATE TABLE IF NOT EXISTS chat_room_record
(
    id           INT auto_increment NOT NULL,
    channel_id   bigint  NOT NULL default 0 comment '频道id',
    dialogue_id varchar(128)   NOT NULL default ''  comment '对话id',
    status   tinyint  NOT NULL default 0 comment '状态',
    send_user_id varchar(128) not null default '' comment '用户',
    kf_id         varchar(191)      default ''                not null comment '客服标识',
    message_type    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '消息类型',
    message      text         NOT NULL,
    message_id         varchar(191)      default ''                not null comment 'message_id customer消息唯一ID',
    create_time  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `idx_channel_id` (`channel_id`),
    KEY          `idx_kf_id` (`kf_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
