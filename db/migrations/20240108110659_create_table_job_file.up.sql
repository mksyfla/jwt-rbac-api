CREATE TABLE job_files
(
  id VARCHAR(255) PRIMARY KEY,
  id_job VARCHAR(255) NOT NULL,
  file VARCHAR(255),
  
  FOREIGN KEY (id_job) REFERENCES jobs (id)
)