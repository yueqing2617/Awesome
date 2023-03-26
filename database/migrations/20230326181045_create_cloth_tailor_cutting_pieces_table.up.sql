CREATE TABLE cloth_tailor_cutting_pieces (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  cloth_tailor_id bigint(20) unsigned NOT NULL DEFAULT 0,
  cloth_style_code varchar(255) NOT NULL DEFAULT '',
  bed varchar(255) NOT NULL DEFAULT '',
  number bigint(20) unsigned NOT NULL DEFAULT 0,
  layer int(10) unsigned NOT NULL DEFAULT 0,
  color varchar(255) NOT NULL DEFAULT '',
  size varchar(255) NOT NULL DEFAULT '',
  is_completed tinyint(1) NOT NULL DEFAULT 0,
  created_at bigint(20) NOT NULL DEFAULT 0,
  updated_at bigint(20) NOT NULL DEFAULT 0,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_cloth_tailor_cutting_pieces_created_at (created_at),
  KEY idx_cloth_tailor_cutting_pieces_updated_at (updated_at),
    KEY idx_cloth_tailor_cutting_pieces_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
