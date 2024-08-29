CREATE SEQUENCE task_id_sequence
INCREMENT 1
START 1
MINVALUE 1;

CREATE TABLE IF NOT EXISTS tasks (
    Id int PRIMARY KEY NOT NULL DEFAULT nextval('task_id_sequence'),
    Title varchar(255) NOT NULL,
    Description varchar(1023) NOT NULL,
    Due timestamp,
    Created timestamp DEFAULT (now() at time zone 'utc'),
    Updated timestamp DEFAULT (now() at time zone 'utc')
);
