CREATE TABLE cloth_tailor_cutting_piece_process (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '',
  sort int(11) NOT NULL DEFAULT 0,
  price decimal(10,2) NOT NULL DEFAULT 0.00,
  cutting_piece_id bigint(20) unsigned NOT NULL DEFAULT 0,
  is_completed tinyint(1) NOT NULL DEFAULT 0,
  employee_id bigint(20) unsigned NOT NULL DEFAULT 0,
  completed_at bigint(20) NOT NULL DEFAULT 0,
  created_at bigint(20) NOT NULL DEFAULT 0,
  updated_at bigint(20) NOT NULL DEFAULT 0,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_cloth_tailor_cutting_piece_process_created_at (created_at),
  KEY idx_cloth_tailor_cutting_piece_process_updated_at (updated_at),
    KEY idx_cloth_tailor_cutting_piece_process_deleted_at (deleted_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
