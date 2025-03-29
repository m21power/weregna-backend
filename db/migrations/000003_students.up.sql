CREATE TABLE student (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  profile_pic TEXT,
  name VARCHAR(255) NOT NULL,
  telegram_username VARCHAR(255) UNIQUE,
  head_id INT REFERENCES head(id) ON DELETE SET NULL,
  total_duration INT DEFAULT 0,
  total_active_days INT DEFAULT 0
);