CREATE TABLE IF NOT EXISTS tasks (
    id int PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    description varchar(1023) NOT NULL,
    due timestamp,
    created timestamp
    updated timestamp
);