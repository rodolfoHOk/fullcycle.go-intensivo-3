CREATE TABLE users (
	id integer PRIMARY KEY,
	name text,
	email text,
	password text
);

INSERT INTO users values(1, 'Rudolf', 'rudy@email.com', 'super_secret');
INSERT INTO users values(2, 'Marianne', 'marianne@email.com', 'super_secret');

SELECT * FROM users;
