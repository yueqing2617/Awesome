CREATE TABLE basic_order_types (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL UNIQUE KEY,
  code varchar(255),
  remark varchar(255),
  created_at varchar(16) NOT NULL,
  updated_at varchar(16) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_basic_order_types_created_at (created_at),
  KEY idx_basic_order_types_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
