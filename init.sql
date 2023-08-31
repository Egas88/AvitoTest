CREATE DATABASE Avito;
\c Avito
CREATE TABLE users (Id serial PRIMARY KEY, Name text);
CREATE TABLE segment(Id serial PRIMARY KEY, Name text UNIQUE NOT NULL);
CREATE TABLE user_segment(user_id integer REFERENCES users NOT NULL, segment_id integer REFERENCES segment NOT NULL, PRIMARY KEY(user_id, segment_id));

INSERT INTO segment (name) VALUES
('AVITO_VOICE_MESSAGES'),
('AVITO_PERFORMANCE_VAS'),
('AVITO_DISCOUNT_30'),
('AVITO_DISCOUNT_50');

INSERT INTO users (name) VALUES
('Adam'),
('Johny'),
('Jacky'),
('Stanley'),
('Stanford'),
('Patrick');

INSERT INTO user_segment (user_id, segment_id) VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(2, 2),
(2, 4),
(3, 2),
(4, 1),
(4, 2),
(4, 3);