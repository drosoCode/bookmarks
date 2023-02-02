DROP TABLE IF EXISTS "tag_link";
DROP TABLE IF EXISTS "tag";
DROP TABLE IF EXISTS "token";
DROP TABLE IF EXISTS "bookmark";
DROP TABLE IF EXISTS "users";

CREATE TABLE "public"."users" (
    "id" bigserial primary key,
    "username" text NOT NULL,
    "name" text NOT NULL
) WITH (oids = false);

CREATE TABLE "public"."bookmark" (
    "id" bigserial primary key,
    "link" text NOT NULL,
    "name" text NOT NULL,
    "description" text NOT NULL,
    "save" boolean NOT NULL,
    "add_date" bigint NOT NULL,
    "id_user" BIGINT NOT NULL REFERENCES "users"("id") ON DELETE CASCADE
) WITH (oids = false);

CREATE TABLE "public"."tag" (
    "id" bigserial primary key,
    "name" text NOT NULL,
    "color" text NOT NULL,
    "id_user" BIGINT NOT NULL REFERENCES "users"("id") ON DELETE CASCADE
) WITH (oids = false);

CREATE TABLE "public"."token" (
    "id" bigserial primary key,
    "name" text NOT NULL,
    "add_date" bigint NOT NULL,
    "value" text NOT NULL,
    "id_user" BIGINT NOT NULL REFERENCES "users"("id") ON DELETE CASCADE
) WITH (oids = false);

CREATE TABLE "public"."tag_link" (
    "id_bookmark" BIGINT NOT NULL REFERENCES "bookmark"("id") ON DELETE CASCADE,
    "id_tag" BIGINT NOT NULL REFERENCES "tag"("id") ON DELETE CASCADE,
    PRIMARY KEY ("id_bookmark", "id_tag")
) WITH (oids = false);
