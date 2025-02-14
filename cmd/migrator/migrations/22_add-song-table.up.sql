CREATE TABLE IF not exists song (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       "group" VARCHAR(255) NOT NULL,
                       release_date DATE DEFAULT NULL,
                       text TEXT DEFAULT NULL,
                       link VARCHAR(255) DEFAULT NULL
);
