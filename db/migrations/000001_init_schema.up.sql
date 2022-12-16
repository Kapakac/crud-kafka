BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    email text UNIQUE,
    name text NOT NULL,
    hash text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX ON users ((lower(email)));

CREATE TABLE IF NOT EXISTS quasi_countries
(
    id serial PRIMARY KEY,
    country text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS quasi_companies
(
    id serial PRIMARY KEY,
    company text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS quasi_persons
(
    id serial PRIMARY KEY,
    name text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS quasi_citizenship
(
    id serial PRIMARY KEY,
    person_id integer NOT NULL REFERENCES quasi_persons (id),
    country_id integer NOT NULL REFERENCES quasi_countries (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT FK_person
        FOREIGN KEY (person_id)
            REFERENCES quasi_persons (id),
    CONSTRAINT FK_country
        FOREIGN KEY (country_id)
            REFERENCES quasi_countries (id)
);

CREATE TABLE IF NOT EXISTS quasi_job
(
    id serial PRIMARY KEY,
    person_id integer NOT NULL REFERENCES quasi_persons (id),
    company_id integer NOT NULL REFERENCES quasi_companies (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT FK_person
        FOREIGN KEY (person_id)
            REFERENCES quasi_persons (id),
    CONSTRAINT FK_company
        FOREIGN KEY (company_id)
            REFERENCES quasi_companies (id)
);

COMMIT;