CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "bio" varchar NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("name");
