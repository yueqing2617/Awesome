CREATE TABLE basic_customers (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
    code varchar(255) DEFAULT '' COMMENT '编码',
    phone varchar(255) DEFAULT '' COMMENT '手机号',
    Address varchar(255) DEFAULT '' COMMENT '地址',
    company varchar(255) DEFAULT '' COMMENT '公司',
    gender varchar(5) DEFAULT '未知' COMMENT '性别',
    remark varchar(255) DEFAULT '' COMMENT '备注',
  created_at bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  updated_at bigint(20) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (id),
  KEY idx_basic_customers_created_at (created_at),
  KEY idx_basic_customers_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into basic_customers (id, name, code, phone, Address, company, gender, remark, created_at, updated_at) values (1, '张三', 'zhangsan', '13800138000', '北京市海淀区', '百度', '男', '张三', 0, 0);
