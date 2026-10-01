ALTER TABLE users
ADD username VARCHAR(100) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT unique_username UNIQUE (username);