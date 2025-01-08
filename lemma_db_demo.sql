show databases;
create database lemma_db;
use lemma_db;

drop table registerusers_tbl;

select*from registerusers_tbl;




CREATE TABLE registerusers_tbl (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    name VARCHAR(100) NOT NULL,          
    email VARCHAR(100) NOT NULL UNIQUE,  
    username VARCHAR(50) NOT NULL UNIQUE, 
    password VARCHAR(255) NOT NULL,       
    mobile VARCHAR(15) NOT NULL,       
    role ENUM('user', 'admin', 'Super Admin') DEFAULT 'user'
);






