
-- 创建用户登录信息表
CREATE TABLE lion_user (
        user_id BIGINT auto_increment  NOT NULL COMMENT '用户id',
        nick_name varchar(32) NULL COMMENT '昵称',
        login_name varchar(20) NULL COMMENT '登录名',
        google_id varchar(32) NULL,
        email varchar(32) NULL,
        avatar varchar(100) NULL COMMENT '头像',
        status varchar(2) NULL COMMENT '状态',
        google_token varchar(100) NULL,
        CONSTRAINT lion_user_pk PRIMARY KEY (user_id)
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
