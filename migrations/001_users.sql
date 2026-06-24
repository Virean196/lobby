-- +goose Up

CREATE TABLE users {
  id int NOT NULL,
  username text NOT NULL,
  PRIAMRY KEY(id)
};



-- +goose Down
DROP TABLE users;
