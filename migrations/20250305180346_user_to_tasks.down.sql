ALTER TABLE tasks DROP CONSTRAINT IF EXISTS tasks_user_id_fkey;
ALTER TABLE tasks DROP COLUMN IF EXISTS user_id;
