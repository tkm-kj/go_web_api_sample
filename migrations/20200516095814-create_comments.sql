
-- +migrate Up
CREATE TABLE comments (
    id INT NOT NULL AUTO_INCREMENT,
    post_id INT NOT NULL,
    name VARCHAR(256) NOT NULL,
    body TEXT NOT NULL,
    deleted_at datetime,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);

-- +migrate Down

DROP TABLE comments;
