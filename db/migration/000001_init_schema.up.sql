CREATE TABLE IF NOT EXISTS test
(
    id SERIAL,
    name text NOT NULL,
    CONSTRAINT test_pkey PRIMARY KEY (id)
)