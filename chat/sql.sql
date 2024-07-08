
create table if not exists `questionnaire_question`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `open_kf_id`  varchar(191) COLLATE utf8mb4_unicode_ci       NOT NULL DEFAULT '' COMMENT '客服标识',
    `question_type`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '1 判断 2 选择 3 程度',
    `score_type`    tinyint unsigned default 1                 not null comment '计分类型，详情对应customer_config.config字段内容',
    `question` varchar(1000) DEFAULT '' NOT NULL  COMMENT '问题的内容',
    `sort`    tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '排序',
    `option`  varchar(2000) DEFAULT '' NOT NULL  COMMENT '问题的答案',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除',
    PRIMARY KEY (`id`),
    KEY           `kf_idx` (`open_kf_id`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment '问卷问题表';


create table if not exists `questionnaire_response`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `question_id`          bigint unsigned NOT NULL  DEFAULT 0 COMMENT 'question id',
    `relation_id`          bigint unsigned NOT NULL  DEFAULT 0 COMMENT '关联id',
    `message_id`  varchar(191) COLLATE utf8mb4_unicode_ci     NOT NULL DEFAULT '' COMMENT 'message_id customer消息唯一ID',
    `user`        varchar(191) COLLATE utf8mb4_unicode_ci      NOT NULL DEFAULT '' COMMENT 'weCom用户标识/customer用户标识',
    `open_kf_id`  varchar(191) COLLATE utf8mb4_unicode_ci       NOT NULL DEFAULT '' COMMENT '客服标识',
    `question`    varchar(1000) DEFAULT '' NOT NULL  COMMENT '问题',
    `option_id`    int unsigned NOT NULL default 0 COMMENT '选项id',
    `answer`    varchar(1000) DEFAULT '' NOT NULL  COMMENT '选项内容',
    `score`    int DEFAULT 0 NOT NULL  COMMENT '分数',
    `score_type`    tinyint unsigned default 1                 not null comment '计分类型，详情对应customer_config.config字段内容',
    `snapshot_option`  varchar(1000) DEFAULT '' NOT NULL  COMMENT '快照选项',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除',
    PRIMARY KEY (`id`),
    KEY           `user_idx` (`user`) USING BTREE,
    KEY           `message_idx` (`message_id`) USING BTREE,
    KEY           `kf_idx` (`open_kf_id`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment '问卷问题回答表';

create table if not exists `questionnaire_result`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `relation_id`          bigint unsigned NOT NULL  DEFAULT 0 COMMENT '关联id',
    `user`        varchar(191) COLLATE utf8mb4_unicode_ci      NOT NULL DEFAULT '' COMMENT 'weCom用户标识/customer用户标识',
    `open_kf_id`  varchar(191) COLLATE utf8mb4_unicode_ci       NOT NULL DEFAULT '' COMMENT '客服标识',
    `score`    DECIMAL(6, 2)  COMMENT '分数',
    `score_type`  tinyint unsigned default 0                 not null comment '计分类型，详情对应customer_config.config字段内容',
    `status`    tinyint DEFAULT 0 NOT NULL  COMMENT '状态，100 最终态',
    `result`    varchar(1000) DEFAULT '' NOT NULL  COMMENT '结论',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除',
    PRIMARY KEY (`id`),
    KEY           `user_idx` (`user`) USING BTREE,
    KEY           `message_idx` (`relation_id`) USING BTREE,
    KEY           `kf_idx` (`open_kf_id`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment '问卷结果表';

alter table customer_config
    add type tinyint default 1 not null comment '类型；1 常规；2 测试' after prompt_states,
    add config varchar(2000) DEFAULT '' NOT NULL  COMMENT '配置项' after type,
    add summary varchar(1000) DEFAULT '' NOT NULL  COMMENT '概要' after config,
    add description varchar(2000) DEFAULT '' NOT NULL  COMMENT '描述' after summary,
    add note varchar(2000) DEFAULT '' NOT NULL  COMMENT '备注' after description,
    add pc_image varchar(255) DEFAULT '' NOT NULL  COMMENT 'pc 图片' after note,
    add h5_image varchar(255) DEFAULT '' NOT NULL  COMMENT 'h5 图片' after description
;
## 把测试环境的 customer_config 表的数据同步到本地

## 增加mbti结果表
create table if not exists `questionnaire_result_mbti`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `relation_id`          bigint unsigned NOT NULL  DEFAULT 0 COMMENT '关联id',
    `user`        varchar(191)   NOT NULL DEFAULT '' COMMENT 'weCom用户标识/customer用户标识',
    `open_kf_id`  varchar(191)  NOT NULL DEFAULT '' COMMENT '客服标识',
    `score`  varchar(2000)  NOT NULL DEFAULT '' COMMENT '分数',
    `name`  varchar(64)  NOT NULL DEFAULT '' COMMENT '类型名称',
    `name_type`  varchar(32)  NOT NULL DEFAULT '' COMMENT '类型',
    `score_result`  varchar(64)  NOT NULL DEFAULT '' COMMENT '分数结果',
    `nickname`  varchar(64)  NOT NULL DEFAULT '' COMMENT '昵称',
    `simple_result`  varchar(1000)  NOT NULL DEFAULT '' COMMENT '一句话结果',
    `result`    varchar(2000) DEFAULT '' NOT NULL  COMMENT '结论',
    `label`  varchar(64)  NOT NULL DEFAULT '' COMMENT '标签',
    `partner`  varchar(64)  NOT NULL DEFAULT '' COMMENT '伴侣',
    `mate`  varchar(64)  NOT NULL DEFAULT '' COMMENT '拍档',
    `content`    varchar(2000) DEFAULT '' NOT NULL  COMMENT 'ai 生成内容',
    `pc_image` varchar(255) DEFAULT '' NOT NULL  COMMENT 'pc 图片',
    `h5_image` varchar(255) DEFAULT '' NOT NULL  COMMENT 'h5 图片',
    `color` tinyint DEFAULT 0 NOT NULL  COMMENT '颜色',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除',
    PRIMARY KEY (`id`),
    KEY           `user_idx` (`user`) USING BTREE,
    KEY           `message_idx` (`relation_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment 'mbti结果表';


##============================ 2023-11-27  mbti93  db change  begin =================================

ALTER TABLE `chat`.`questionnaire_response`
    ADD COLUMN `mbti` varchar(4) NOT NULL DEFAULT "" COMMENT 'mbti 答案类型 ' AFTER `is_deleted`;



ALTER TABLE `chat`.`questionnaire_result`
    ADD COLUMN `mbti` char(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'mbti 类型' AFTER `is_deleted`;


ALTER TABLE `chat`.`customer_config`
    MODIFY COLUMN `config` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '配置项' AFTER `type`;

## 2023-12-05
ALTER TABLE `chat`.`questionnaire_result`
    MODIFY COLUMN `result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '结论' AFTER `status`;

ALTER TABLE `chat`.`questionnaire_result`
    MODIFY COLUMN `result` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '结论' AFTER `status`;


ALTER TABLE `chat`.`questionnaire_result`
    ADD COLUMN `mbti` char(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'mbti 类型' AFTER `is_deleted`;




##============================ 2023-11-12  =================================

alter table questionnaire_result_mbti
    modify score varchar(2000) default '' not null comment '分数';

## 2023-1-11
ALTER TABLE `chat`.`customer_config`
    ADD COLUMN `quote` varchar(255) DEFAULT '' NOT NULL  COMMENT '引用' AFTER `h5_image`;
