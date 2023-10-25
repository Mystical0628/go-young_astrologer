CREATE TABLE IF NOT EXISTS apod (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255),
    explanation     TEXT,
    url             VARCHAR(255),
    hd_url          VARCHAR(255),
    date            DATE,
    media_type      VARCHAR(10),
    service_version VARCHAR(10)
)