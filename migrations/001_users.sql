-- +goose Up
CREATE TABLE users(
  id INT NOT NULL AUTO_INCREMENT,
  username TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  PRIMARY KEY(id)
);
-- +goose Down
DROP TABLE users;
