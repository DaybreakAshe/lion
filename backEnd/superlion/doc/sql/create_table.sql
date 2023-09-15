
-- 创建用户登录信息表
CREATE TABLE `lion_user` (
     `go_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户id',
     `go_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '昵称',
     `login_name` varchar(20) DEFAULT NULL COMMENT '登录名',
     `user_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
     `avatar` varchar(100) DEFAULT NULL COMMENT '头像',
     `status` varchar(2) DEFAULT NULL COMMENT '状态',
     `go_token` varchar(320) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
     `go_email` varchar(20) DEFAULT NULL COMMENT '谷歌email',
     `go_verified_email` tinyint(1) DEFAULT NULL COMMENT '谷歌验证',
     `go_picture` varchar(100) DEFAULT NULL,
     `go_locale` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
     `create_time` datetime DEFAULT NULL,
     `signature` varchar(1024) DEFAULT NULL COMMENT '个人简介',
     PRIMARY KEY (`go_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `lion_post` (
     `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
     `audit_state` varchar(5) COMMENT '审核状态',
     `category` varchar(5) COMMENT '类别',
     `author_id` bigint(11) COMMENT '作者ID',
     `title` varchar(100) COMMENT '标题',
     `content_type` varchar(5) COMMENT '内容类型',
     `markdown_content` longtext  COMMENT 'markdown内容',
     `html_content` longtext  COMMENT 'html内容',
     `views` bigint(11) NOT NULL DEFAULT '0' COMMENT '浏览量',
     `approvals` bigint(11) NOT NULL DEFAULT '0' COMMENT '点赞量/收藏量',
     `comments` bigint(11) NOT NULL DEFAULT '0' COMMENT '评论量',
     `type_id` bigint(11) NOT NULL DEFAULT '0' COMMENT '文章类型ID',
     `head_img` varchar(8192)  DEFAULT '' COMMENT '文章头图',
     `official` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '官方',
     `top` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '置顶',
     `sort` int(4) NOT NULL DEFAULT '1000' COMMENT '排序',
     `marrow` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '精华',
     `comment_id` bigint(11) COMMENT '问答最佳答案ID',
     `is_delete` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '删除标识（0:未删除、1:已删除）',
     `create_at` datetime COMMENT '记录创建时间',
     `update_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='帖子文章表';