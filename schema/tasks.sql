select * from tasks;
select * from employees;
select * from severtity_statuses;
create table employees (
                           id INT AUTO_INCREMENT PRIMARY KEY,
                           last_name VARCHAR(255),
                           first_name VARCHAR(255),
                           three_name VARCHAR(255)
);
drop table tasks;

CREATE DATABASE IF NOT EXISTS task;
USE task;



CREATE TABLE tasks (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       title varchar(140),
                       description TEXT,
                       severity_id int,
                       employee_id int,
                       CONSTRAINT fk_employees_id
                           FOREIGN KEY (employee_id) REFERENCES employees (id),
                       constraint fk_severity_statuses_id
                           foreign key (severity_id) references severtity_statuses (id)

);

create table severtity_statuses (
                                    id INT AUTO_INCREMENT PRIMARY KEY,
                                    status varchar(50)
);
INSERT INTO tasks (title, description, severity_id, employee_id)
VALUES ('Проверить подключение к бд', 'Описание задачи', 1, 2);
SELECT * FROM employees WHERE id = 2;
INSERT INTO employees (last_name, first_name, three_name)
VALUES ('Иванов', 'Иван', 'Иванович');


