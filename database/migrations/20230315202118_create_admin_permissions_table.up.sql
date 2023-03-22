CREATE TABLE admin_permissions (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '权限名称',
  path text NOT NULL COMMENT '权限路径',
  parent_id bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '父级权限ID',
  remark text NOT NULL COMMENT '备注',
  meta json NOT NULL COMMENT '元数据',
  redirect varchar(255) DEFAULT '' COMMENT '重定向',
  active varchar(255) DEFAULT '' COMMENT '高亮',
  component varchar(255) DEFAULT '' COMMENT '组件',
  sort int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  created_at bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  updated_at bigint(20) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (id),
  KEY idx_admin_permissions_created_at (created_at),
  KEY idx_admin_permissions_updated_at (updated_at),
    KEY idx_admin_permissions_parent_id (parent_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into admin_permissions (name, path, parent_id, remark, meta, redirect, active, component, sort, created_at, updated_at) values ('home', '/home', 0, '首页','{"title":"首页","icon":"el-icon-eleme-filled","type":"menu","affix":false}', '', '', '', 0, 0, 0);
insert into admin_permissions (name, path, parent_id, remark, meta, redirect, active, component, sort, created_at, updated_at) values ('dashboard', '/dashboard', 1, '控制台','{"title":"控制台","icon":"el-icon-menu","type":"iframe","affix":true}', '', '', '', 0, 0, 0);
insert into admin_permissions (name, path, parent_id, remark, meta, redirect, active, component, sort, created_at, updated_at) values ('system', '/system', 0, '系统管理','{"title":"系统管理","icon":"el-icon-setting","type":"menu","affix":false}', '', '', '', 0, 0, 0);