CREATE TABLE admin_roles (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  display_name varchar(255) NOT NULL,
  remark varchar(255),
  permissions json default NULL,
  created_at varchar(16) NOT NULL,
  updated_at varchar(16) NOT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_admin_roles_created_at (created_at),
  KEY idx_admin_roles_updated_at (updated_at),
  KEY idx_admin_roles_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

# 插入默认角色
INSERT INTO admin_roles (name, display_name, remark, permissions, created_at, updated_at) VALUES ('super_admin', '超级管理员', '超级管理员', '["*"]', '00', '0');
INSERT INTO admin_roles (name, display_name, remark, permissions, created_at, updated_at) VALUES ('admin', '管理员', '管理员', '["get./admin/system/role","post./admin/system/role","put./admin/system/role/{id}","delete./admin/system/role","get./admin/system/role/{id}"]', '0', '0');