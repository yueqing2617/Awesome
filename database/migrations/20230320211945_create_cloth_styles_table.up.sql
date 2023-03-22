CREATE TABLE cloth_styles (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
  code varchar(255) NOT NULL UNIQUE COMMENT '编码',
  picture text COMMENT '图片',
  colors text NOT NULL COMMENT '颜色',
  sizes text NOT NULL COMMENT '尺码',
  year varchar(255) NOT NULL DEFAULT '' COMMENT '年份',
  season varchar(255) NOT NULL DEFAULT '' COMMENT '季节',
  unit_price decimal(10,2) NOT NULL DEFAULT 0 COMMENT '单价',
  procedures json NOT NULL COMMENT '工序',
  remark varchar(255) DEFAULT '' COMMENT '备注',
  created_at bigint(20) NOT NULL DEFAULT 0,
  updated_at bigint(20) NOT NULL DEFAULT 0,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_cloth_styles_created_at (created_at),
  KEY idx_cloth_styles_updated_at (updated_at),
    KEY idx_cloth_styles_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

insert into cloth_styles (name, code, picture, colors, sizes, year, season, unit_price, procedures, remark, created_at, updated_at, deleted_at) values ('衬衫', 'shirt', 'http://dummyimage.com/400x400', '白色,黑色,蓝色', 'S,M,L,XL', '2020', '春季', 100, '[{"name":"裁剪","price":10,"sort":1,"is_completed":false},{"name":"缝制","price":20,"sort":2,"is_completed":false},{"name":"烫熨","price":30,"sort":3,"is_completed":false},{"name":"包装","price":40,"sort":4,"is_completed":true}]', '备注', 1600000000, 1600000000, null);
INSERT INTO cloth_styles (name, code, picture, colors, sizes, year, season, unit_price, procedures, remark, created_at, updated_at, deleted_at) VALUES ('裤子', 'pants', 'http://dummyimage.com/400x400', '白色,黑色,蓝色', 'S,M,L,XL', '2020', '春季', 100, '[{"name":"裁剪","price":10,"sort":1,"is_completed":false},{"name":"缝制","price":20,"sort":2,"is_completed":false},{"name":"烫熨","price":30,"sort":3,"is_completed":false},{"name":"包装","price":40,"sort":4,"is_completed":true}]', '备注', 1600000000, 1600000000, null);
INSERT INTO cloth_styles (name, code, picture, colors, sizes, year, season, unit_price, procedures, remark, created_at, updated_at, deleted_at) VALUES ('外套', 'coat', 'http://dummyimage.com/400x400', '白色,黑色,蓝色', 'S,M,L,XL', '2020', '春季', 100, '[{"name":"裁剪","price":10,"sort":1,"is_completed":false},{"name":"缝制","price":20,"sort":2,"is_completed":false},{"name":"烫熨","price":30,"sort":3,"is_completed":false},{"name":"包装","price":40,"sort":4,"is_completed":true}]', '备注', 1600000000, 1600000000, null);
INSERT INTO cloth_styles (name, code, picture, colors, sizes, year, season, unit_price, procedures, remark, created_at, updated_at, deleted_at) VALUES ('裙子', 'skirt', 'http://dummyimage.com/400x400', '白色,黑色,蓝色', 'S,M,XL', '2020', '春季', 100, '[{"name":"裁剪","price":10,"sort":1,"is_completed":false},{"name":"缝制","price":20,"sort":2,"is_completed":false},{"name":"烫熨","price":30,"sort":3,"is_completed":false},{"name":"包装","price":40,"sort":4,"is_completed":true}]', '备注', 1600000000, 1600000000, null);
INSERT INTO cloth_styles (name, code, picture, colors, sizes, year, season, unit_price, procedures, remark, created_at, updated_at, deleted_at) VALUES ('连衣裙', 'dress', 'http://dummyimage.com/400x400', '白色,黑色,蓝色', 'S,M,L,XL', '2020', '春季', 100, '[{"name":"裁剪","price":10,"sort":1,"is_completed":false},{"name":"缝制","price":20,"sort":2,"is_completed":false},{"name":"烫熨","price":30,"sort":3,"is_completed":false},{"name":"包装","price":40,"sort":4,"is_completed":true}]', '备注', 1600000000, 1600000000, null);
INSERT INTO cloth_styles (name, code, picture, colors, sizes, year, season, unit_price, procedures, remark, created_at, updated_at, deleted_at) VALUES ('卫衣', 'sweater', 'http://dummyimage.com/400x400', '白色,黑色,蓝色', 'S,M,L', '2020', '春季', 100, '[{"name":"裁剪","price":10,"sort":1,"is_completed":false},{"name":"缝制","price":20,"sort":2,"is_completed":false},{"name":"烫熨","price":30,"sort":3,"is_completed":false},{"name":"包装","price":40,"sort":4,"is_completed":true}]', '备注', 1600000000, 1600000000, null);