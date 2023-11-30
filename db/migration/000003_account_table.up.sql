
ALTER TABLE users ADD email_verified boolean NOT NULL DEFAULT false;
ALTER TABLE users ADD refresh_token varchar NOT NULL DEFAULT 'nil';
ALTER TABLE users ADD session_state varchar NOT NULL DEFAULT 'logged-in';

CREATE INDEX ON "users" ("refresh_token");
CREATE INDEX ON "users" ("id");
