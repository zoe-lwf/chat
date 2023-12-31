create database go_chat;
use go_chat;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `phone_number` varchar(20)     NOT NULL COMMENT '手机号',
    `nickname`     varchar(20)     NOT NULL COMMENT '昵称',
    `password`     varchar(255)    NOT NULL COMMENT '密码',
    `create_time`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_phone_number` (`phone_number`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`     bigint unsigned NOT NULL COMMENT '用户id',
    `friend_id`   bigint unsigned NOT NULL COMMENT '好友id',
    `create_time` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_friend_id` (`user_id`, `friend_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `group`;
CREATE TABLE `group`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `name`        varchar(50)     NOT NULL COMMENT '群组名称',
    `owner_id`    bigint unsigned NOT NULL COMMENT '群主id',
    `create_time` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `group_user`;
CREATE TABLE `group_user`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `group_id`    bigint unsigned NOT NULL COMMENT '群组id',
    `user_id`     bigint unsigned NOT NULL COMMENT '用户id',
    `create_time` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_group_id_user_id` (`group_id`, `user_id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id`      bigint unsigned NOT NULL COMMENT '用户id，指接受者用户id',
    `sender_id`    bigint unsigned NOT NULL COMMENT '发送者用户id',
    `session_type` tinyint         NOT NULL COMMENT '聊天类型，群聊/单聊',
    `receiver_id`  bigint unsigned NOT NULL COMMENT '接收者id，群聊id/用户id',
    `message_type` tinyint         NOT NULL COMMENT '消息类型,语言、文字、图片',
    `content`      blob            NOT NULL COMMENT '消息内容',
    `seq`          bigint unsigned NOT NULL COMMENT '消息序列号',
    `send_time`    datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '消息发送时间',
    `create_time`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_seq` (`user_id`, `seq`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `uid`;
CREATE TABLE `uid`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `business_id` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '业务id',
    `max_id`      bigint unsigned                  DEFAULT NULL COMMENT '最大id',
    `step`        int unsigned                     DEFAULT NULL COMMENT '步长',
    `create_time` datetime                         DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime                         DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_business_id` (`business_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='分布式自增主键';