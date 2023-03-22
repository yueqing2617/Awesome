CREATE TABLE basic_colors (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '颜色名称',
    code varchar(255) DEFAULT '' COMMENT '颜色编码',
    remark varchar(255) DEFAULT '' COMMENT '备注',
  created_at bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  updated_at bigint(20) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (id),
  KEY idx_basic_colors_created_at (created_at),
  KEY idx_basic_colors_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into basic_colors (name, code, remark, created_at, updated_at) values ('红色', 'red', '红色', 1546300800, 1546300800);
insert into basic_colors (name, code, remark, created_at, updated_at) values ('绿色', 'green', '绿色', 1546300800, 1546300800);
insert into basic_colors (name, code, remark, created_at, updated_at) values ('蓝色', 'blue', '蓝色', 1546300800, 1546300800);
insert into basic_colors (name, code, remark, created_at, updated_at) values ('黄色', 'yellow', '黄色', 1546300800, 1546300800);