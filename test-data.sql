-- schema.sql
CREATE DATABASE db;

\c db;

CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  pass TEXT NOT NULL,
  role VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO users (pass, role, email)
VALUES (crypt('aFAKEpassword', gen_salt('bf')), 'admin', 'fake@gmail.com');

INSERT INTO users (pass, role, email)
VALUES (crypt('aFAKEpassword1', gen_salt('bf')), 'admin', 'fake1@gmail.com');

INSERT INTO users (pass, role, email)
VALUES (crypt('aFAKEp3assword', gen_salt('bf')), 'user', 'fake3@gmail.com');

INSERT INTO users (pass, role, email)
VALUES (crypt('aFAKEp4assword', gen_salt('bf')), 'admin', 'fake4@gmail.com');

INSERT INTO users (pass, role, email)
VALUES (crypt('aFAKEpassgword', gen_salt('bf')), 'admin', 'fake5@gmail.com');
