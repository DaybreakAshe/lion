
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


