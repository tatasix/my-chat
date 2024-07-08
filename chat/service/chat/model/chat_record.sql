SET NAMES 'utf8';
CREATE TABLE `chat_record`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `relation_id`          bigint unsigned NOT NULL  DEFAULT 0 COMMENT '关联id',
    `user`        varchar(191) COLLATE utf8mb4_unicode_ci      NOT NULL DEFAULT '' COMMENT 'weCom用户标识/customer用户标识',
    `message_id`  varchar(191) COLLATE utf8mb4_unicode_ci     NOT NULL DEFAULT '' COMMENT 'message_id customer消息唯一ID',
    `open_kf_id`  varchar(191) COLLATE utf8mb4_unicode_ci       NOT NULL DEFAULT '' COMMENT '客服标识',
    `agent_id`    bigint unsigned                   NOT NULL DEFAULT 0 COMMENT '应用ID',
    `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL  COMMENT '消息内容',
    `emoji`    smallint unsigned  NOT NULL DEFAULT 0 COMMENT 'emoji类型',
    `chat_type`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '聊天类型',
    `answer_or_question`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '问题还是答案:1问题，2答案',
    `message_type`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '消息类型',
    `state`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '属于哪一步的聊天记录',
    `state_id` int unsigned  NOT NULL DEFAULT 0 COMMENT '属于哪一步的聊天记录',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `user_agent_idx` (`user`,`agent_id`) USING BTREE,
    KEY           `user_kf_idx` (`user`,`open_kf_id`) USING BTREE,
    KEY           `user_message_idx` (`user`,`message_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;