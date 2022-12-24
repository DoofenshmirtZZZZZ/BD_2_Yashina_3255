BEGIN;

CREATE TABLE peoplee (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255)
);

 INSERT INTO peoplee (name) VALUES
('Владимир'), ('Владислав'), ('Даниил');

COMMIT;