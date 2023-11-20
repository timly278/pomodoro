CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(50) NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "alarm_sound" varchar NOT NULL,
  "repeat_alarm" int NOT NULL
);

CREATE TABLE "types" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "name" varchar(50) NOT NULL,
  "color" varchar NOT NULL,
  "duration" int NOT NULL,
  "shortbreak" int NOT NULL,
  "longbreak" int NOT NULL,
  "longbreakinterval" int NOT NULL,
  "autostart_break" boolean NOT NULL
);

CREATE TABLE "pomodoros" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "type_id" bigserial NOT NULL,
  "task_id" bigserial NOT NULL,
  "focus_degree" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tasks" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "content" varchar(100) NOT NULL,
  "status" int NOT NULL,
  "estimate_pomos" int NOT NULL,
  "progress_pomos" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "completed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "goalperday" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "type_id" bigserial NOT NULL,
  "pomonum" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE INDEX ON "types" ("user_id");

CREATE INDEX ON "pomodoros" ("created_at");

CREATE INDEX ON "pomodoros" ("user_id");

CREATE INDEX ON "tasks" ("user_id");

CREATE INDEX ON "goalperday" ("user_id");

CREATE UNIQUE INDEX ON "goalperday" ("id", "type_id");

ALTER TABLE "types" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "pomodoros" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "pomodoros" ADD FOREIGN KEY ("type_id") REFERENCES "types" ("id");

ALTER TABLE "pomodoros" ADD FOREIGN KEY ("task_id") REFERENCES "tasks" ("id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "goalperday" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "goalperday" ADD FOREIGN KEY ("type_id") REFERENCES "types" ("id");
