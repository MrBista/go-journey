create table sample (
    id integer primary key,
    name text not null,
    created_at timestamp default current_timestamp
);


SELECT * FROM sample;


CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    password VARCHAR(100),
    name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)

DROP TABLE users;


select * from users;



ALTER TABLE users 
RENAME COLUMN name TO first_name;

ALTER TABLE users
ADD COLUMN last_name VARCHAR(100) AFTER middle_name;


ALTER TABLE users
ADD COLUMN middle_name VARCHAR(100) AFTER first_name;


SELECT * FROM users;