CREATE TABLE task (
     id INT NOT NULL AUTO_INCREMENT,
     goal VARCHAR(255) NOT NULL,
     completed BOOLEAN DEFAULT false
     PRIMARY KEY (id)
);