CREATE TABLE job_skills
(
  id VARCHAR(255) PRIMARY KEY,
  id_job VARCHAR(255) NOT NULL,
  id_skill VARCHAR(255) NOT NULL,
  
  FOREIGN KEY (id_job) REFERENCES jobs (id),
  FOREIGN KEY (id_skill) REFERENCES skills (id)
)