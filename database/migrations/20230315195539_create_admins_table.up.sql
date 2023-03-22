CREATE TABLE admins (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    nickname varchar(255) NOT NULL DEFAULT '' COMMENT '管理员名称',
    phone varchar(255) NOT NULL DEFAULT '' COMMENT '手机号',
    password varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
    avatar text COMMENT '头像',
    email varchar(255) DEFAULT '' COMMENT '邮箱',
    role_id bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '角色ID',
    role text COMMENT '角色',
    gender varchar(5) DEFAULT '未知' COMMENT '性别',
    remark varchar(255) DEFAULT '' COMMENT '备注',
  created_at bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  updated_at bigint(20) NOT NULL DEFAULT 0 COMMENT '更新时间',
  deleted_at datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  KEY idx_admins_created_at (created_at),
  KEY idx_admins_updated_at (updated_at),
  KEY idx_admins_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into admins (id, nickname, phone, password, avatar, email, role_id, role, gender, remark, created_at, updated_at) VALUE (1, 'admin', '17758361517', '$argon2id$v=19$m=65536,t=4,p=1$NwSHKuoiAtqqpsGqjT0FqA$JeWQ09B9f3i0IO5+9Q6w+Cr9UkcJysyE0SmYZuolR5A', '', '', 1, '超级管理员', '未知', '', 0, 0);
insert into admins (id, nickname, phone, password, avatar, email, role_id, role, gender, remark, created_at, updated_at) VALUE (2, 'admin', '17758361527', '$argon2id$v=19$m=65536,t=4,p=1$NwSHKuoiAtqqpsGqjT0FqA$JeWQ09B9f3i0IO5+9Q6w+Cr9UkcJysyE0SmYZuolR5A', '', '', 2, '管理员', '未知', '', 0, 0);
