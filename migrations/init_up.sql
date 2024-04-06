-- Create database tables
CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY,
  "email" varchar NOT NULL UNIQUE,
  "username" varchar NOT NULL UNIQUE,
  "name" varchar NOT NULL,
  "points" integer NOT NULL DEFAULT 0,
  "picture" varchar,
  "created_at" timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS "stashes" (
  "id" serial PRIMARY KEY,
  "title" varchar NOT NULL,
  "body" text,
  "points" integer NOT NULL DEFAULT 0,
  "owner_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT current_timestamp,
  "is_public" bool NOT NULL
);

CREATE TABLE IF NOT EXISTS "links" (
  "id" serial PRIMARY KEY,
  "url" text NOT NULL,
  "comment" text,
  "stash_id" integer NOT NULL
);

CREATE TABLE IF NOT EXISTS "comments" (
  "id" serial PRIMARY KEY,
  "author" integer NOT NULL,
  "body" text NOT NULL,
  "stash_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT current_timestamp
);

-- Add relations
ALTER TABLE "stashes" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");
ALTER TABLE "links" ADD FOREIGN KEY ("stash_id") REFERENCES "stashes" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("stash_id") REFERENCES "stashes" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("author") REFERENCES "users" ("id");

