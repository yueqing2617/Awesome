CREATE TABLE basic_sizes (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '尺码名称',
  code varchar(255) DEFAULT '' COMMENT '尺码编码',
  remark varchar(255) DEFAULT '' COMMENT '备注',
      created_at bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  updated_at bigint(20) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (id),
  KEY idx_basic_sizes_created_at (created_at),
  KEY idx_basic_sizes_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into basic_sizes (name, code, remark, created_at, updated_at) values ('S', 'S', 'S', 0, 0);
insert into basic_sizes (name, code, remark, created_at, updated_at) values ('M', 'M', 'M', 0, 0);
insert into basic_sizes (name, code, remark, created_at, updated_at) values ('L', 'L', 'L', 0, 0);
insert into basic_sizes (name, code, remark, created_at, updated_at) values ('XL', 'XL', 'XL', 0, 0);
insert into basic_sizes (name, code, remark, created_at, updated_at) values ('XXL', 'XXL', 'XXL', 0, 0);
insert into basic_sizes (name, code, remark, created_at, updated_at) values ('XXXL', 'XXXL', 'XXXL', 0, 0);
