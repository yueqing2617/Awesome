CREATE TABLE cloth_orders (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '订单名称',
  code varchar(255) NOT NULL DEFAULT '' COMMENT '订单编号' UNIQUE,
  customer_id bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '客户ID',
  delivery_date bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '交货日期',
  order_type varchar(255) DEFAULT '' COMMENT '订单类型',
  salesman_id bigint(20) DEFAULT 0 COMMENT '业务员',
  cloth_style_code varchar(255) NOT NULL DEFAULT '' COMMENT '款式编号',
  cloth_style_name varchar(255) NOT NULL DEFAULT '' COMMENT '款式名称',
  cloth_style_picture varchar(255) NOT NULL DEFAULT '' COMMENT '款式图片',
  cloth_style_colors varchar(255) NOT NULL DEFAULT '' COMMENT '款式颜色',
    cloth_style_sizes varchar(255) NOT NULL DEFAULT '' COMMENT '款式尺寸',
    cloth_style_year varchar(255) NOT NULL DEFAULT '' COMMENT '款式年份',
    cloth_style_season varchar(255) NOT NULL DEFAULT '' COMMENT '款式季节',
    cloth_style_unit_price decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '款式单价',
  total int(11) NOT NULL DEFAULT 0 COMMENT '总数量',
    total_price decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '总金额',
    contains json DEFAULT NULL COMMENT '包含',
    procedures json DEFAULT NULL COMMENT '工序',
    status int(1) NOT NULL DEFAULT 1 COMMENT '状态: 1-未完成, 2-已完成, 3-已关闭',
    remark varchar(255) DEFAULT '' COMMENT '备注',
  created_at bigint(20) NOT NULL DEFAULT 0,
  updated_at bigint(20) NOT NULL DEFAULT 0,
    deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_cloth_orders_created_at (created_at),
  KEY idx_cloth_orders_updated_at (updated_at),
    KEY idx_cloth_orders_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into cloth_orders (name, code, customer_id, delivery_date, order_type, salesman_id, cloth_style_code, cloth_style_name, cloth_style_picture, cloth_style_colors, cloth_style_sizes, cloth_style_year, cloth_style_season, cloth_style_unit_price, total, total_price, contains, procedures, remark, created_at, updated_at) values ('订单1', '202001010001', 1, 1577836800, '普通订单', 1, 'A001', '衬衫', 'https://www.baidu.com/img/bd_logo1.png', '红色,蓝色', 'S,M,L,XL', '2020', '春季', 100.00, 100, 10000.00, '[{"color":"红色","size":"S","num":10},{"color":"红色","size":"M","num":10},{"color":"红色","size":"L","num":10},{"color":"红色","size":"XL","num":10},{"color":"蓝色","size":"S","num":10},{"color":"蓝色","size":"M","num":10},{"color":"蓝色","size":"L","num":10},{"color":"蓝色","size":"XL","num":10}]', '[{"name":"裁剪","sort":1,"price":10.00,"is_completed":false},{"name":"缝制","sort":2,"price":10.00,"is_completed":false},{"name":"包装","sort":3,"price":10.00,"is_completed":true}]', '备注', 1577836800, 1577836800);
insert into cloth_orders (name, code, customer_id, delivery_date, order_type, salesman_id, cloth_style_code, cloth_style_name, cloth_style_picture, cloth_style_colors, cloth_style_sizes, cloth_style_year, cloth_style_season, cloth_style_unit_price, total, total_price, contains, procedures, remark, created_at, updated_at) values ('订单2', '202001010002', 1, 1577836800, '普通订单', 1, 'A001', '衬衫', 'https://www.baidu.com/img/bd_logo1.png', '红色,蓝色', 'S,M,L,XL', '2020', '春季', 100.00, 100, 10000.00, '[{"color":"红色","size":"S","num":10},{"color":"红色","size":"M","num":10},{"color":"红色","size":"L","num":10},{"color":"红色","size":"XL","num":10},{"color":"蓝色","size":"S","num":10},{"color":"蓝色","size":"M","num":10},{"color":"蓝色","size":"L","num":10},{"color":"蓝色","size":"XL","num":10}]', '[{"name":"裁剪","sort":1,"price":10.00,"is_completed":false},{"name":"缝制","sort":2,"price":10.00,"is_completed":false},{"name":"包装","sort":3,"price":10.00,"is_completed":true}]', '备注', 1577836800, 1577836800);