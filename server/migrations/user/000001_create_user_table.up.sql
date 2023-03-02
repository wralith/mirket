CREATE TABLE IF NOT EXISTS "user" (
	id varchar PRIMARY KEY,
	username varchar UNIQUE NOT NULL,
	email varchar UNIQUE NOT NULL,
	hashed_password bytea UNIQUE NOT NULL,
	about varchar UNIQUE NOT NULL,

	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now()),
	deleted_at timestamptz DEFAULT NULL
);

CREATE INDEX ON "user" (username);
CREATE INDEX ON "user" (deleted_at);
