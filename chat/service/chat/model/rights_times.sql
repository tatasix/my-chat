create table if not exists rights_times
(
    id          bigint unsigned auto_increment ,
    `rights_id` bigint unsigned NOT NULL default 0,
    `user`        varchar(191) NOT NULL DEFAULT '' COMMENT '用户标识',
    `open_kf_id`  varchar(191) NOT NULL DEFAULT '' COMMENT '客服标识',
    `pay_type` tinyint NOT NULL DEFAULT 0 COMMENT '付款方式：1 次卡;2 月卡; 3年卡',
    `period` tinyint NOT NULL DEFAULT 0 COMMENT '计算周期：1 按天算次数;2 不按时间计算次数',
    `start`  DATETIME COMMENT '开始时间',
    `end`    DATETIME COMMENT '结束时间',
    `total` smallint NOT NULL DEFAULT 0 COMMENT '总次数',
    `remain` smallint NOT NULL DEFAULT 0 COMMENT '剩余次数',
    `used` smallint NOT NULL DEFAULT 0 COMMENT '使用次数',
    created_at  timestamp       default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at  timestamp       default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    `is_deleted` tinyint default 0 not null comment '是否删除，0否，1是',
    PRIMARY KEY (`id`),
    KEY           `user_idx` (`user`) USING BTREE,
    KEY           `kf_idx` (`open_kf_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;