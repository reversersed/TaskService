CREATE TABLE IF NOT EXISTS tasks (
    Id int PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    Title varchar(255) NOT NULL,
    Description varchar(1023) NOT NULL,
    Due timestamp,
    Created timestamp DEFAULT SYSUTCDATETIME(),
    Updated timestamp DEFAULT SYSUTCDATETIME(),
);

CREATE OR REPLACE FUNCTION update_task_updated_time() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
   UPDATE tasks SET Updated=SYSUTCDATETIME() WHERE Id=NEW.Id
END;
$$

CREATE OR REPLACE TRIGGER tasks_time_update AFTER UPDATE
    ON tasks
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE FUNCTION update_task_updated_time();
