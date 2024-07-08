# 聊天室
CREATE TABLE IF NOT EXISTS `chat_room`
(
    id          INT auto_increment NOT NULL,
    user  varchar(128) not null default '' comment '用户',
    kf_id         varchar(191)      default ''                not null comment '客服标识',
    title       varchar(100) NOt NULL default '' comment '名称',
    description varchar(255) NOt NULL default '' comment '描述',
    channel_id  bigint     NOT NULL default 0 comment '频道id',
    created_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_deleted tinyint default 0 not null comment '是否删除',
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `idx_channel_id` (`channel_id`),
    KEY         `idx_kf_id` (`kf_id`),
    KEY         `idx_title` (`title`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
