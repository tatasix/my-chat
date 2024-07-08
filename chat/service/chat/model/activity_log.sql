create table if not exists activity_log
(
    id          bigint unsigned auto_increment ,
    user    varchar(128)                              not null comment '微信客户的external_userid',
    old_user    varchar(128)                              not null comment '微信客户的external_userid',
    `type` smallint NOT NULL DEFAULT 0,
    `begin_time` varchar(32) NOT NULL default '',
    `end_time`  varchar(32) NOT NULL default '',
    `duration` bigint NOT NULL default 0 comment '持续时间',
    created_at  timestamp       default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at  timestamp       default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`),
    KEY           `idx_begin_time` (`begin_time`) USING BTREE,
    KEY           `idx_user` (`user`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

