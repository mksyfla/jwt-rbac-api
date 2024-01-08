CREATE TABLE jobs
(
  id VARCHAR(255) PRIMARY KEY,
  id_user VARCHAR(255) NOT NULL,
  title VARCHAR(255),
  description TEXT,
  end_time TIME,
  reward VARCHAR(255),
  draft BOOLEAN NOT NULL DEFAULT (true),
  
  FOREIGN KEY (id_user) REFERENCES users (id)
)