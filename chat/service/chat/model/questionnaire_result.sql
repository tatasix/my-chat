create table if not exists `questionnaire_result`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `relation_id`          bigint unsigned NOT NULL  DEFAULT 0 COMMENT '关联id',
    `user`        varchar(191) COLLATE utf8mb4_unicode_ci      NOT NULL DEFAULT '' COMMENT 'weCom用户标识/customer用户标识',
    `open_kf_id`  varchar(191) COLLATE utf8mb4_unicode_ci       NOT NULL DEFAULT '' COMMENT '客服标识',
    `score`    DECIMAL(6, 2)  COMMENT '分数',
    `score_type`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '计分类型，详情对应customer_config.config字段内容',
    `status`    tinyint DEFAULT 0 NOT NULL  COMMENT '状态，100 最终态',
    `result`    varchar(1000) DEFAULT '' NOT NULL  COMMENT '结论',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除',
    `mbti` char(4) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'mbti 类型',
    PRIMARY KEY (`id`),
    KEY           `user_idx` (`user`) USING BTREE,
    KEY           `message_idx` (`relation_id`) USING BTREE,
    KEY           `kf_idx` (`open_kf_id`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment '问卷结果表';
