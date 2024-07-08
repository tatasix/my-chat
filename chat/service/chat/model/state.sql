CREATE TABLE `state`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `user`        varchar(128) NOT NULL DEFAULT '' COMMENT 'weCom用户标识/customer用户标识',
    `kf_id`  varchar(128) NOT NULL DEFAULT '' COMMENT '客服标识',
    `state`  tinyint unsigned NOT NULL DEFAULT 1 COMMENT '状态:1待处理，2已处理',
    `date` date default null  comment '日期' ,
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除，0否，1是',
    PRIMARY KEY (`id`),
    KEY           `user_idx` (`user`) USING BTREE,
    KEY           `kf_id_idx` (`kf_id`) USING BTREE,
    KEY           `date_idx` (`date`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


#alter table state  add `date` date default null  comment '日期' after state;

