CREATE TYPE "video_status" AS ENUM (
    'PENDING',
    'UPLOADED',
    'PROCESSING',
    'PROCESSED',
    'FAILURE',
    'DELETED'
);

CREATE TYPE "resource_storage_provider" AS ENUM (
    'GCP'
);

CREATE TYPE "job_status" AS ENUM (
    'IDLE',
    'PENDING',
    'RUNNING',
    'DONE',
    'FAILURE'
);

CREATE TYPE "job_kind" AS ENUM (
    'TRANSCODE'
);

CREATE TYPE "resource_kind" AS ENUM (
    'RAW_VIDEO',
    'TRANSCODED_VIDEO'
);

CREATE TYPE "resource_status" AS ENUM (
    'ACTIVE',
    'DELETED'
);

CREATE TABLE "video" (
    "id"            UUID PRIMARY KEY NOT NULL,
    "status"        video_status NOT NULL,
    "created_at"    timestamp NOT NULL,
    "updated_at"    timestamp NOT NULL
);

CREATE TABLE "job" (
    "id"            UUID PRIMARY KEY NOT NULL,
    "status"        job_status NOT NULL,
    "kind"          job_kind NOT NULL,
    "video_id"      UUID NOT NULL,
    "resource_id"   UUID NULL,
    "depends_on_id" UUID NULL,
    "error"         VARCHAR(255) NULL,
    "created_at"    timestamp NOT NULL,
    "updated_at"    timestamp NOT NULL
);

CREATE TABLE "resource" (
    "id"                UUID PRIMARY KEY NOT NULL,
    "status"            resource_status NOT NULL,
    "kind"              resource_kind NOT NULL,
    "video_id"          UUID NOT NULL,
    "storage_provider"  resource_storage_provider NOT NULL,
    "size"              INT NULL,
    "path"              VARCHAR(255) NOT NULL,
    "upload_url"        VARCHAR(255) NOT NULL,
    "metadata"          JSONB NOT NULL,
    "created_at"        timestamp NOT NULL,
    "updated_at"        timestamp NOT NULL
);

ALTER TABLE "job" ADD FOREIGN KEY ("depends_on_id") REFERENCES "job" ("id");

ALTER TABLE "job" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");

ALTER TABLE "job" ADD FOREIGN KEY ("resource_id") REFERENCES "resource" ("id");

ALTER TABLE "resource" ADD FOREIGN KEY ("video_id") REFERENCES "video" ("id");
