
ALTER TABLE types DROP COLUMN goalperday;

CREATE TABLE "goalperday" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "type_id" bigint UNIQUE NOT NULL,
  "pomonum" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

ALTER TABLE "goalperday" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "goalperday" ADD FOREIGN KEY ("type_id") REFERENCES "types" ("id") ON DELETE CASCADE;
