CREATE TABLE "users" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "username" varchar(50) NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "alarm_sound" varchar NOT NULL DEFAULT ('Kitchen'),
  "repeat_alarm" int NOT NULL DEFAULT 1
);

CREATE TABLE "types" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "name" varchar(50) NOT NULL,
  "color" varchar NOT  NULL DEFAULT '#fc5603',
  "duration" int NOT NULL DEFAULT 25,
  "shortbreak" int NOT NULL DEFAULT 5,
  "longbreak" int NOT NULL DEFAULT 15,
  "longbreakinterval" int NOT NULL DEFAULT 4,
  "autostart_break" boolean NOT NULL DEFAULT false
);

CREATE TABLE "pomodoros" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "type_id" bigint NOT NULL,
  "task_id" bigint NULL,
  "focus_degree" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tasks" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "content" varchar(100) NOT NULL,
  "status" int NOT NULL,
  "estimate_pomos" int NOT NULL,
  "progress_pomos" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "completed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "goalperday" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "type_id" bigint UNIQUE NOT NULL,
  "pomonum" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE INDEX ON "types" ("user_id");

CREATE INDEX ON "pomodoros" ("created_at");

CREATE INDEX ON "pomodoros" ("user_id");

CREATE INDEX ON "tasks" ("user_id");

CREATE INDEX ON "goalperday" ("user_id");

ALTER TABLE "types" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "pomodoros" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "pomodoros" ADD FOREIGN KEY ("type_id") REFERENCES "types" ("id") ON DELETE CASCADE;

ALTER TABLE "pomodoros" ADD FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON DELETE CASCADE;

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "goalperday" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "goalperday" ADD FOREIGN KEY ("type_id") REFERENCES "types" ("id") ON DELETE CASCADE;
