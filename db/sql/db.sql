-- Active: 1696384805376@@127.0.0.1@3306@sayembara

-- DATABASE
CREATE DATABASE sayembara;

USE sayembara;

DROP DATABASE sayembara;

SHOW TABLES;


-- EXPERIMENT
SELECT * FROM users;
SELECT * FROM jobs;
SELECT * FROM job_files;

SELECT DISTINCT
  jobs.id,
  jobs.title,
  jobs.description,
  job_files.file,
  users.name AS username
FROM
  jobs
INNER JOIN
  users ON jobs.id_user = users.id
LEFT JOIN
  job_files ON jobs.id = job_files.id_job
WHERE
  jobs.draft = 0

SELECT DISTINCT  from job_files;


TRUNCATE TABLE users;
TRUNCATE TABLE umkm;
TRUNCATE TABLE mahasiswa;

DROP TABLE users

CREATE TABLE users
(
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  profile VARCHAR(255) NOT NULL,
  banner VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL
)

CREATE TABLE umkm
(
  id VARCHAR(255) PRIMARY KEY,
  id_user VARCHAR(255) NOT NULL,
  verified TINYINT(1),


  CONSTRAINT umkm_user
    FOREIGN KEY (id_user) REFERENCES users (id)
)

CREATE TABLE mahasiswa
(
  id VARCHAR(255) PRIMARY KEY,
  id_user VARCHAR(255) NOT NULL,
  expert BOOLEAN NOT NULL DEFAULT (false),

  FOREIGN KEY (id_user) REFERENCES users (id)
)