CREATE TABLE videos (
    id          VARCHAR(255) PRIMARY KEY,
    resource_id VARCHAR(255) NOT NULL,
    file_path   VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE jobs (
    id                  VARCHAR(255) PRIMARY KEY,
    output_bucket_path  VARCHAR(255) NOT NULL,
    status              VARCHAR(255) NOT NULL,
    video_id            VARCHAR(255) NOT NULL,
    error               VARCHAR(255) NULL,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(video_id) REFERENCES videos(id)
);
