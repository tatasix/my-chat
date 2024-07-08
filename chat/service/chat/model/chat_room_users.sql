# 聊天室跟用户关联
CREATE TABLE IF NOT EXISTS chat_room_users
(
    id          INT auto_increment NOT NULL,
    group_id    INT       NOt NULL,
    user  varchar(128) not null default '' comment '用户',
    channel_id  bigint  NOT NULL default 0 comment '频道id',
    is_manager  tinyint(1) NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_deleted tinyint default 0 not null comment '是否删除',
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `idx_group_id` (`group_id`),
    KEY         `idx_user_id` (`user`)
) ENGINE=InnoDB   DEFAULT CHARSET=utf8mb4;
