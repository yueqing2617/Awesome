CREATE TABLE admin_roles (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL DEFAULT '' COMMENT '角色名称',
    display_name varchar(255) NOT NULL DEFAULT '' COMMENT '角色显示名称',
    permissions text NOT NULL COMMENT '角色权限',
    remark varchar(255) DEFAULT '' COMMENT '备注',
  created_at bigint(20) NOT NULL DEFAULT 0,
  updated_at bigint(20) NOT NULL DEFAULT 0,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_admin_roles_created_at (created_at),
  KEY idx_admin_roles_updated_at (updated_at),
    KEY idx_admin_roles_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

INSERT INTO admin_roles (id, name, display_name, permissions, remark, created_at, updated_at, deleted_at) VALUES (1, 'admin', '超级管理员', '["*"]', '', 0, 0, NULL);
insert into admin_roles (id, name, display_name, permissions, remark, created_at, updated_at, deleted_at) values (2, 'user', '普通用户', '["user"]', '', 0, 0, NULL);