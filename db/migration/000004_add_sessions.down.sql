ALTER TABLE users DROP COLUMN is_blocked;
ALTER TABLE users ADD session_state varchar NOT NULL DEFAULT 'logged-in';