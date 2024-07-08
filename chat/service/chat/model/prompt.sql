CREATE TABLE `prompt`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `state_id`  bigint unsigned NOT NULL default 0 comment '状态',
    `kf_id`  varchar(128) NOT NULL DEFAULT '' COMMENT '客服',
    `title` varchar(500)  NOT NULL DEFAULT '' COMMENT '用户反馈标题',
    `prompt`  varchar(2000) NOT NULL DEFAULT '' COMMENT 'prompt',
    `updated_by` varchar(32) default '' not null comment '更新',
    `created_by` varchar(32) default '' not null comment '创建',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除，0否，1是',
    PRIMARY KEY (`id`),
    KEY           `kf_id_idx` (`kf_id`) USING BTREE,
    KEY           `state_id_idx` (`state_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;