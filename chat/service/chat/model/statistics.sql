create table if not exists statistics
(
    id          bigint unsigned auto_increment ,
    `date` date NOT NULL,
    `daily_active` int(11) NOT NULL DEFAULT 0 COMMENT '日活',
    `seven_active` int(11) NOT NULL DEFAULT 0 COMMENT '周活',
    `fifteen_active` int(11) NOT NULL DEFAULT 0 COMMENT '半月活',
    `monthly_active` int(11) NOT NULL DEFAULT 0 COMMENT '月活',
    `total_visitor` int(11) NOT NULL DEFAULT 0 COMMENT '游客',
    `registered_user` int(11) NOT NULL DEFAULT 0 COMMENT '注册用户',
    `add_visitor` int(11) NOT NULL DEFAULT 0 COMMENT '新增游客',
    `add_register` int(11) NOT NULL DEFAULT 0 COMMENT '新增注册用户',
    created_at  timestamp       default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at  timestamp       default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`),
    KEY           `idx_date` (`date`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;