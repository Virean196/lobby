-- +goose Up

CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  username text NOT NULL,
  password text NOT NULL,
  PRIMARY KEY(id)
  );

-- +goose Down
DROP TABLE users;
