create table if not exists resource_usage
(
    `id`    bigint unsigned auto_increment,
    `user`  varchar(128) not null comment '微信客户的external_userid',
    `date`  date NOT NULL,
    `hour`  tinyint NOT NULL DEFAULT 0 COMMENT '时间',
    `times` int(11) NOT NULL DEFAULT 0 COMMENT '频次',
    `token` bigint NOT NULL DEFAULT 0 COMMENT 'token',
    `created_at`  timestamp  default CURRENT_TIMESTAMP null comment '创建时间',
    `updated_at`  timestamp  default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user` (`user`) USING BTREE,
    KEY `idx_date_hour` (`date`,`hour`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;