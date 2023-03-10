CREATE TABLE admins (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  phone varchar(255) NOT NULL UNIQUE KEY,
  password varchar(32) NOT NULL,
  email varchar(255),
  nickname varchar(255),
  gender varchar(4) DEFAULT '未知',
  avatar varchar(255) DEFAULT 'https://avatars.githubusercontent.com/u/16495509?v=4',
  role_name varchar(255) DEFAULT 'guest',
  remark varchar(255),
  created_at varchar(16) NOT NULL,
  updated_at varchar(16) NOT NULL,
    deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_admins_created_at (created_at),
  KEY idx_admins_updated_at (updated_at),
    KEY idx_admins_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

INSERT INTO `admins` (`created_at`,`updated_at`,`phone`,`password`,`email`,`gender`,`nickname`,`avatar`,`role_name`,`remark`,`deleted_at`) VALUES (1678358412,1678358412,'17858361617','1d87793671d44cedd79a738eaad30ab8','804966813@qq.com','男','asdf','122','guest','',NULL);