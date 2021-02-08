CREATE DATABASE IF NOT EXISTS test_db;

USE test_db;

CREATE TABLE IF NOT EXISTS test_db.people (
    name        VARCHAR(100),
    title       VARCHAR(10),
    description VARCHAR(100),
    PRIMARY KEY (name)
);

DELETE FROM test_db.people;

INSERT INTO test_db.people VALUES ('Gru', 'Felonius', 'Where are the minions?');
INSERT INTO test_db.people VALUES ('Nefario', 'Dr.', 'Why ... why are you so old?');
INSERT INTO test_db.people VALUES ('Agnes', '', 'Your unicorn is so fluffy!');
INSERT INTO test_db.people VALUES ('Edith', '', "Don't touch anything!");
INSERT INTO test_db.people VALUES ('Vector', '', 'Committing crimes with both direction and magnitude!');
INSERT INTO test_db.people VALUES ('Dave', 'Minion', 'Ngaaahaaa! Patalaki patalaku Big Boss!!');