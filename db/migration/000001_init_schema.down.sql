
ALTER TABLE IF EXISTS "goalPerday" DROP CONSTRAINT IF EXISTS "goalPerDay_type_id_fkey";
ALTER TABLE IF EXISTS "goalPerday" DROP CONSTRAINT IF EXISTS "goalPerDay_user_id_fkey";
ALTER TABLE IF EXISTS "pomodoros" DROP CONSTRAINT IF EXISTS "pomodoros_task_id_fkey";
ALTER TABLE IF EXISTS "pomodoros" DROP CONSTRAINT IF EXISTS "pomodoros_type_id_fkey";
ALTER TABLE IF EXISTS "pomodoros" DROP CONSTRAINT IF EXISTS "pomodoros_user_id_fkey";
ALTER TABLE IF EXISTS "tasks" DROP CONSTRAINT IF EXISTS "tasks_user_id_fkey";
ALTER TABLE IF EXISTS "types" DROP CONSTRAINT IF EXISTS "types_user_id_fkey";

DROP TABLE IF EXISTS settings cascade;
DROP TABLE IF EXISTS goalperday cascade;
DROP TABLE IF EXISTS pomodoros cascade;
DROP TABLE IF EXISTS types cascade;
DROP TABLE IF EXISTS tasks cascade;
DROP TABLE IF EXISTS users;