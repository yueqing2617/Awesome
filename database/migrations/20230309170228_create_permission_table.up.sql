CREATE TABLE permissions (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  path varchar(255) NOT NULL,
  pid bigint(20) unsigned NOT NULL DEFAULT 0,
  sort int(11) NOT NULL DEFAULT 0,
  created_at varchar(16) NOT NULL DEFAULT 0,
    updated_at varchar(16) NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  KEY idx_permission_created_at (created_at),
  KEY idx_permission_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into permissions (name, path, pid, sort, created_at, updated_at) values ('系统管理', '0', 0, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('管理员管理', '0', 1, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('管理员列表', 'get./admin/system/admin', 2, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('添加管理员', 'post./admin/system/admin', 2, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('编辑管理员', 'put./admin/system/admin/{id}', 2, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('删除管理员', 'delete./admin/system/admin', 2, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('管理员详情', 'get./admin/system/admin/{id}', 2, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('修改管理员密码', 'put./admin/system/admin/{id}/password', 2, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('角色管理', '0', 1, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('角色列表', 'get./admin/system/role', 9, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('添加角色', 'post./admin/system/role', 9, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('编辑角色', 'put./admin/system/role/{id}', 9, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('删除角色', 'delete./admin/system/role', 9, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('角色详情', 'get./admin/system/role/{id}', 9, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('权限管理', '0', 1, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('权限列表', 'get./admin/system/permission', 15, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('添加权限', 'post./admin/system/permission', 15, 0, 0, 0);
insert into permissions (name, path, pid, sort, created_at, updated_at) values ('编辑权限', 'put./admin/system/permission/{id}', 15, 0, 0, 0);