CREATE TYPE "video_status" AS ENUM (
  'PENDING',
  'UPLOADED',
  'PROCESSING',
  'PROCESSED',
  'FAILURE'
);

CREATE TYPE "storage_provider" AS ENUM (
  'GCP'
);

CREATE TYPE "job_status" AS ENUM (
  'IDLE',
  'PENDING',
  'RUNNING',
  'DONE',
  'FAILURE'
);

CREATE TYPE "job_type" AS ENUM (
  'TRANSCODE'
);

CREATE TYPE "resource_type" AS ENUM (
  'RAW_VIDEO',
  'TRANSCODED_VIDEO'
);

CREATE TABLE "video" (
  "id" string PRIMARY KEY NOT NULL,
  "status" video_status NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "storage" (
  "id" string PRIMARY KEY NOT NULL,
  "provider" storage_provider NOT NULL,
  "path" string NOT NULL,
  "upload_url" string NOT NULL
);

CREATE TABLE "job" (
  "id" string PRIMARY KEY NOT NULL,
  "status" job_status NOT NULL,
  "type" job_type NOT NULL,
  "video_id" string NOT NULL,
  "resource_id" string,
  "depends_on_id" string,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "resource" (
  "id" string PRIMARY KEY NOT NULL,
  "type" resource_type NOT NULL,
  "video_id" string NOT NULL,
  "storage_id" string NOT NULL,
  "metadata" text NOT NULL,
  "created_at" timestamp NOT NULL
);

ALTER TABLE "job" ADD FOREIGN KEY ("depends_on_id") REFERENCES "job" ("id");

ALTER TABLE "job" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");

ALTER TABLE "job" ADD FOREIGN KEY ("resource_id") REFERENCES "resource" ("id");

ALTER TABLE "resource" ADD FOREIGN KEY ("storage_id") REFERENCES "storage" ("id");

ALTER TABLE "resource" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");
