CREATE TABLE IF NOT EXISTS
  go_sample.samples (
      id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
      name VARCHAR(255),
      content VARCHAR(255),
      created_at DATETIME,
      updated_at DATETIME,
      PRIMARY KEY(id)
  ) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;