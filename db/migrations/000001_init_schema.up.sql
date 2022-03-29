BEGIN;

CREATE TABLE IF NOT EXISTS quasi_countries
(
    id serial,
    country text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT PK_country PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS quasi_companies
(
    id serial,
    company text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT PK_company PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS quasi_persons
(
    id serial,
    name text NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT PK_person PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS quasi_citizenship
(
    id serial,
    person_id integer NOT NULL REFERENCES quasi_persons (id),
    country_id integer NOT NULL REFERENCES quasi_countries (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT PK_citizenship PRIMARY KEY (id),
    CONSTRAINT FK_person
        FOREIGN KEY (person_id)
            REFERENCES quasi_persons (id),
    CONSTRAINT FK_country
        FOREIGN KEY (country_id)
            REFERENCES quasi_countries (id)
);

CREATE TABLE IF NOT EXISTS quasi_job
(
    id serial,
    person_id integer NOT NULL REFERENCES quasi_persons (id),
    company_id integer NOT NULL REFERENCES quasi_companies (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT PK_job PRIMARY KEY (id),
    CONSTRAINT FK_person
        FOREIGN KEY (person_id)
            REFERENCES quasi_persons (id),
    CONSTRAINT FK_company
        FOREIGN KEY (company_id)
            REFERENCES quasi_companies (id)
);

CREATE TABLE IF NOT EXISTS funny_sanctions
(
    id serial,
    name text NOT NULL,
    body text,
    description text,
    comment text,
    started_at TIMESTAMP,
    ended_at TIMESTAMP,
    person_id integer,
    company_id integer,
    country_id integer,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT PK_funny_sanction PRIMARY KEY (id),
    CONSTRAINT FK_person
        FOREIGN KEY (person_id)
            REFERENCES quasi_persons (id),
    CONSTRAINT FK_company
        FOREIGN KEY (company_id)
            REFERENCES quasi_companies (id),
    CONSTRAINT FK_country
        FOREIGN KEY (country_id)
            REFERENCES quasi_countries (id)
);

COMMIT;