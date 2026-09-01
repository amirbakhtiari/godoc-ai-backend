CREATE TABLE documents
(
    id         TEXT PRIMARY KEY,
    title      TEXT        NOT NULL,
    content    TEXT        NOT NULL,
    source     TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);