CREATE TABLE cloth_tailors (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  cloth_order_code varchar(255) NOT NULL DEFAULT '' COMMENT '生产订单号',
  cloth_style_code varchar(255) NOT NULL DEFAULT '' COMMENT '布料款号',
  total int(11) NOT NULL DEFAULT 0 COMMENT '总数量',
  completed_num int(11) NOT NULL DEFAULT 0 COMMENT '已完成数量',
  is_completed tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否完成',
  created_at bigint(20) NOT NULL DEFAULT 0,
  updated_at bigint(20) NOT NULL DEFAULT 0,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_cloth_tailors_created_at (created_at),
  KEY idx_cloth_tailors_updated_at (updated_at),
    KEY idx_cloth_tailors_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
