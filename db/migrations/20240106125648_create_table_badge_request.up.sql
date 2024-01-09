CREATE TABLE badge_request
(
  id_user VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (id_user) REFERENCES users (id)
)
