create table if not exists `questionnaire_question`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `open_kf_id`  varchar(191) COLLATE utf8mb4_unicode_ci       NOT NULL DEFAULT '' COMMENT '客服标识',
    `question_type`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '1 判断 2 选择 3 程度',
    `score_type`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '计分类型，详情对应customer_config.config字段内容',
    `question` varchar(1000) DEFAULT '' NOT NULL  COMMENT '问题的内容',
    `sort`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '排序',
    `option`  varchar(2000) DEFAULT '' NOT NULL  COMMENT '问题的答案',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除',
    PRIMARY KEY (`id`),
    KEY           `kf_idx` (`open_kf_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment '问卷问题表';
