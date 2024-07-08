create table if not exists config
(
    id  bigint unsigned auto_increment ,
    config_type  smallint unsigned default 0 not null comment '配置类型：1，openai key',
    name VARCHAR(64) default '' not null comment '配置名',
    description VARCHAR(256) default '' not null comment '配置描述',
    value VARCHAR(5000) default '' not null comment '配置内容',
    created_by VARCHAR(64) default '' not null comment '创建人',
    updated_by VARCHAR(64) default '' not null comment '更新人',
    created_at  timestamp       default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at  timestamp       default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted tinyint default 0 not null comment '是否删除',
    PRIMARY KEY (`id`),
    KEY           `idx_config_type` (`config_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
