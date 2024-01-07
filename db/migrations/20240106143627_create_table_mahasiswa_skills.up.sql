-- Active: 1696384805376@@127.0.0.1@3306@sayembara
CREATE TABLE mahasiswa_skills
(
  id VARCHAR(255) PRIMARY KEY,
  id_mahasiswa VARCHAR(255) NOT NULL,
  id_skills VARCHAR(255) NOT NULL,
  
  FOREIGN KEY (id_mahasiswa) REFERENCES mahasiswa (id),
  FOREIGN KEY (id_skills) REFERENCES skills (id)
)