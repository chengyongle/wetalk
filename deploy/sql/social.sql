CREATE TABLE `friends` (
                           `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                           `user_id` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                           `friend_uid` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                           `remark` varchar(255) DEFAULT NULL,
                           `friend_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 1:好友 2:删除',
                           `add_source`  tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                           `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `friend_requests` (
                                   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                                   `user_id` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                                   `req_uid` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                                   `req_msg` varchar(255) DEFAULT NULL,
                                   `req_time` timestamp  NOT NULL,
                                   `handle_result`  tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                   `handle_msg` varchar(255) DEFAULT NULL,
                                   `handled_time`timestamp NULL DEFAULT NULL,
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `groups` (
                          `id` varchar(24) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                          `name` varchar(255) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                          `icon` varchar(255) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                          `status`  tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                          `creator_uid` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                          `group_type` int(11) NOT NULL DEFAULT '0' ,
                          `member_num` int(11) NOT NULL DEFAULT '0' ,
                          `is_verify` boolean NOT NULL ,
                          `notification` varchar(255) DEFAULT NULL,
                          `notification_uid` varchar(64) DEFAULT NULL,
                          `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `group_members` (
                                 `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                                 `group_id` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                                 `user_id` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                                 `role_level`  tinyint COLLATE utf8mb4_unicode_ci NOT NULL ,
                                 `member_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 1:在群 2:退群',
                                 `join_time` timestamp NULL DEFAULT NULL,
                                 `join_source`  tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                 `inviter_uid` varchar(64) DEFAULT NULL,
                                 `operator_uid` varchar(64) DEFAULT NULL,
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `group_requests` (
                                  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                                  `req_id` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                                  `group_id` varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL ,
                                  `req_msg` varchar(255) DEFAULT NULL,
                                  `req_time` timestamp NULL DEFAULT NULL,
                                  `join_source`  tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                  `inviter_user_id` varchar(64) DEFAULT NULL,
                                  `handle_user_id` varchar(64) DEFAULT NULL,
                                  `handle_time` timestamp NULL DEFAULT NULL,
                                  `handle_result`  tinyint COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

