create table if not exists rights_record
(
    id          bigint unsigned auto_increment ,
    `rights_id` bigint unsigned NOT NULL default 0,
    `rights_times_id` bigint unsigned NOT NULL default 0,
    `relation_id` bigint unsigned NOT NULL  DEFAULT 0 COMMENT '关联id',
    `user`        varchar(191)  NOT NULL DEFAULT '' COMMENT '用户标识',
    `open_kf_id`  varchar(191)  NOT NULL DEFAULT '' COMMENT '客服标识',
    created_at  timestamp       default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at  timestamp       default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除，0否，1是',
    PRIMARY KEY (`id`),
    KEY           `user_idx` (`user`) USING BTREE,
    KEY           `kf_idx` (`open_kf_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;