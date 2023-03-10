CREATE TABLE basic_customers (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL UNIQUE KEY,
  code varchar(255),
  remark varchar(255),
  gender varchar(4) NOT NULL DEFAULT '未知',
    phone varchar(11),
    address varchar(255),
    company varchar(255),
  created_at varchar(16) NOT NULL,
  updated_at varchar(16) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_basic_customers_created_at (created_at),
  KEY idx_basic_customers_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
