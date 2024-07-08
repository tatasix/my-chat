CREATE TABLE `feedback`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `user`        varchar(128) NOT NULL DEFAULT '' COMMENT 'weCom用户标识/customer用户标识',
    `message_id`  varchar(128) NOT NULL DEFAULT '' COMMENT 'message_id customer消息唯一ID',
    `open_kf_id`  varchar(128) NOT NULL DEFAULT '' COMMENT '客服标识',
    `title` varchar(500)  NOT NULL DEFAULT '' COMMENT '用户反馈标题',
    `content` text COMMENT '用户反馈内容',
    `reply` text COMMENT '回复内容',
    `contact` varchar(32) default '' not null comment '联系方式',
    `images` varchar(255) default '' not null comment '图片',
    `status`  tinyint unsigned NOT NULL DEFAULT 1 COMMENT '状态:1待处理，2已处理',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除，0否，1是',
    PRIMARY KEY (`id`),
    KEY           `user_kf_idx` (`user`,`open_kf_id`) USING BTREE,
    KEY           `user_message_idx` (`user`,`message_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;