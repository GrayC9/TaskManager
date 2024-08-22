CREATE DATABASE IF NOT EXISTS task;
USE task;
CREATE TABLE tasks (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             last_name VARCHAR(255),
                             first_name VARCHAR(255),
                             three_name VARCHAR(255),
                             tekst_obrasheniya TEXT,
                             time_to_complete TIME
);

