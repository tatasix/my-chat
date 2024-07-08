create table if not exists user_portrait
(
    id          bigint unsigned auto_increment ,
    `user`      varchar(128)  default  ''   not null comment '微信客户的external_userid',
    `call`      varchar(64)   default  ''   not null comment '称呼',
    age         tinyint unsigned   default  0    not null comment '年龄',
    address    varchar(512)   default  ''           not null comment '居住地址',
    career    varchar(64)   default  ''                      not null comment '职业',
    gender tinyint  default 0 not null comment '性别',
    sleep_is_normal tinyint  default 0 not null comment '睡眠是否正常；1正常，2不正常',
    appetite_is_normal tinyint  default 0 not null comment '胃口是否正常；1正常，2不正常',
    psychological_stress tinyint  default 0 not null comment '心理压力程度，1-10的数字',
    created_at  timestamp       default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at  timestamp       default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `idx_user` (`user`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '用户画像信息收集表';

