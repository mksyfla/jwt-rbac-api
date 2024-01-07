CREATE TABLE badge_request
(
  id_user VARCHAR(255) NOT NULL,

  FOREIGN KEY (id_user) REFERENCES users (id)
)
