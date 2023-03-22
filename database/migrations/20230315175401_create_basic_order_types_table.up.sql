CREATE TABLE basic_order_types (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
  code varchar(255) DEFAULT '' COMMENT '编码',
  remark varchar(255) DEFAULT '' COMMENT '备注',
  created_at bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  updated_at bigint(20) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (id),
  KEY idx_basic_order_types_created_at (created_at),
  KEY idx_basic_order_types_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into basic_order_types (name, code, remark, created_at, updated_at) values ('普通订单', 'normal', '普通订单', 0, 0);
insert into basic_order_types (name, code, remark, created_at, updated_at) values ('团购订单', 'group', '团购订单', 0, 0);