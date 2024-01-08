ALTER TABLE users DROP COLUMN session_state;
ALTER TABLE users ADD  is_blocked boolean NOT NULL DEFAULT true;