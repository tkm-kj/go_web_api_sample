
-- +migrate Up
CREATE TABLE posts (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(256) NOT NULL,
    body TEXT NOT NULL,
    deleted_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE posts;
