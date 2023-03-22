CREATE TABLE admin_permission_policies (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '权限名称',
  path varchar(255) NOT NULL DEFAULT '' COMMENT '权限路径',
  created_at bigint(20) NOT NULL DEFAULT 0,
  updated_at bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  KEY idx_admin_permission_policies_created_at (created_at),
  KEY idx_admin_permission_policies_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into admin_permission_policies (name, path, created_at, updated_at) values ('permission.list', 'get./manage/system/permission', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('permission.create', 'post./manage/system/permission', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('permission.update', 'put./manage/system/permission/{id}', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('permission.delete', 'delete./manage/system/permission', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('permission.show', 'get./manage/system/permission/{id}', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('role.list', 'get./manage/system/role', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('role.create', 'post./manage/system/role', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('role.update', 'update./manage/system/role/{id}', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('role.delete', 'delete./manage/system/role', 0, 0);
insert into admin_permission_policies (name, path, created_at, updated_at) values ('role.show', 'get./manage/system/role/{id}', 0, 0);