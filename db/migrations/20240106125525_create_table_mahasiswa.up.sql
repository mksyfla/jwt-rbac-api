CREATE TABLE mahasiswa
(
  id VARCHAR(255) PRIMARY KEY,
  id_user VARCHAR(255) NOT NULL,
  expert BOOLEAN NOT NULL DEFAULT (false),

  FOREIGN KEY (id_user) REFERENCES users (id)
)