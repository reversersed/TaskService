CREATE OR REPLACE FUNCTION update_task_updated_time() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
   NEW.Updated=(now() at time zone 'utc');
   RETURN NEW;
END;
$$;

CREATE OR REPLACE TRIGGER tasks_time_update BEFORE UPDATE
    ON tasks
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE FUNCTION update_task_updated_time();