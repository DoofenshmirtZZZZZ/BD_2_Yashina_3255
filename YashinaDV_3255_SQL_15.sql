/* №1. Выведите на экран любое сообщение */

DO $$
	BEGIN RAISE NOTICE 'Как-то так...';
END; 
$$

/* №2. Выведите на экран текущую дату */

SELECT CURRENT_DATE;

/* №3. Создайте две числовые переменные и присвойте им значение. Выполните математические действия с этими числами и выведите результат на экран. */

DO $$
	DECLARE
	a integer := 3;
	b integer := 2;
	BEGIN
	RAISE NOTICE'a + b = %', a + b;
	RAISE NOTICE'a - b = %', a - b;
	RAISE NOTICE'a * b = %', a * b;
	RAISE NOTICE'a / b = %', a / b;
END; 
$$

/* №4. Написать программу двумя способами 1 - использование IF, 2 - использование CASE. Объявите числовую переменную и присвоейте ей значение. Если число равно 5 - выведите на экран "Отлично". 4 - "Хорошо". 3 - Удовлетворительно". 2 - "Неуд". В остальных случаях выведите на экран сообщение, что введённая оценка не верна. */

DO $$
	DECLARE
	a integer := 1;
	BEGIN
	IF a = 5 THEN
		RAISE NOTICE 'Отлично';
	END IF;
	IF a = 4 THEN
		RAISE NOTICE 'Хорошо';
	END IF;
	IF a = 3 THEN
		RAISE NOTICE 'Удовлетворительно';
	END IF;
	IF a = 2 THEN
		RAISE NOTICE 'Неуд';
	END IF;
	IF ((a != 2) AND (a != 3) AND (a != 4) AND (a != 5)) THEN
		RAISE NOTICE 'введённая оценка не верна';
	END IF;
END; 
$$;
DO $$
	DECLARE
	a integer := 3;
	BEGIN
	CASE a
		WHEN 5 THEN RAISE NOTICE 'Отлично';
		WHEN 4 THEN RAISE NOTICE 'Хорошо';
		WHEN 3 THEN RAISE NOTICE 'Удовлетворительно';
		WHEN 2 THEN RAISE NOTICE 'Неуд';
        ELSE RAISE NOTICE 'все плохо';
	END CASE;
END;
$$;

/* №5. Выведите все квадраты чисел от 20 до 30 3-мя разными способами (LOOP, WHILE, FOR). */

DO $$
	DECLARE
	number integer := 20;
	BEGIN
	LOOP
		RAISE NOTICE '%', number * number;
		number := number + 1;
		IF number > 30 THEN
			EXIT;
		END IF;
	END LOOP;
END; $$;
DO $$
	DECLARE
	number integer := 20;
	BEGIN
	WHILE number <= 30 LOOP
		RAISE NOTICE '%', number * number;
		number := number + 1;
	END LOOP;
END $$;
DO $$
	BEGIN
	FOR i IN 20..30 LOOP
		RAISE NOTICE '%', i * i;
	END LOOP;
END $$;

/* №6. Написать функцию, входной параметр - начальное число, на выходе - количество чисел, пока не получим 1; написать процедуру, которая выводит все числа последовательности. Входной параметр - начальное число. */

CREATE OR REPLACE FUNCTION Col(num int)
	RETURNS integer AS
	$$
	DECLARE
        numbe int := num;
        countt int := 0;
	BEGIN
	WHILE numbe != 1 LOOP
		IF numbe % 2 = 0 THEN
			numbe := numbe / 2;
		ELSE
			numbe := numbe * 3 + 1;
		END IF;
			countt := countt + 1;
	END LOOP;
	RETURN countt;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE PROCEDURE ColProcedure(num int)
	AS $$
	DECLARE
        numbe int := num;
        countt int := 0;
	BEGIN
	WHILE numbe != 1 LOOP
        RAISE NOTICE '%', numbe;
        IF numbe % 2 = 0 THEN
            numbe := numbe / 2;
        ELSE
            numbe := numbe * 3 + 1;
        END IF;
	END LOOP;
	END;
$$ LANGUAGE plpgsql;

/* №7. Написать фунцию, входной параметр - количество чисел, на выходе - последнее число (Например: входной 5, 2 1 3 4 7 - на выходе число 7); написать процедуру, которая выводит все числа последовательности. Входной параметр - количество чисел. */

CREATE OR REPLACE FUNCTION Clown(num int)
    RETURNS integer AS
    $$
    DECLARE
        numbe int := num;
    BEGIN
    IF numbe = 1 THEN
        RETURN 2;
    END IF;
    IF numbe = 2 THEN
        RETURN 1;
    END IF;
    RETURN Clown(numbe - 1) + Clown(numbe - 2);
END;
$$ LANGUAGE plpgsql;
    CREATE OR REPLACE PROCEDURE ClownProcedure(num int)
    AS $$
    DECLARE
        numbe int := num;
        clown_1 int := 2;
        clown_2 int := 1;
        tmp int;
    BEGIN
    RAISE NOTICE '%', clown_1;
    RAISE NOTICE '%', clown_2;
    FOR i IN 0..num - 3 LOOP
        tmp := clown_1;
        clown_1 := clown_2;
        clown_2 := tmp + clown_2;
        RAISE NOTICE '%', clown_2;
    END LOOP;
    END;
$$ LANGUAGE plpgsql;

/* №8. Напишите функцию, которая возвращает количество человек родившихся в заданном году. */

CREATE OR REPLACE FUNCTION BornYear(yr int)
RETURNS int AS $$
DECLARE
    countt int;
BEGIN
SELECT count(*) INTO countt
FROM people
WHERE EXTRACT(YEAR FROM people.birth_date) = yr;
RETURN countt;
END
$$ LANGUAGE plpgsql;

/* №9. Напишите функцию, которая возвращает количество человек с заданным цветом глаз. */

CREATE OR REPLACE FUNCTION EyeColor(color varchar)
RETURNS int AS $$
DECLARE
    countt int;
BEGIN
SELECT count(*) INTO countt
FROM people
WHERE people.eyes = color;
RETURN countt;
END
$$ LANGUAGE plpgsql;

/* №10. Напишите функцию, которая возвращает ID самого молодого человека в таблице. */

CREATE OR REPLACE FUNCTION YoungMan()
RETURNS int AS $$
DECLARE
    man int;
BEGIN
SELECT people.id INTO man
FROM people
WHERE birth_date = (SELECT max(birth_date) FROM people);
RETURN man;
END
$$ LANGUAGE plpgsql;

/* №11. Напишите процедуру, которая возвращает людей с индексом массы тела больше заданного. ИМТ = масса в кг / (рост в м)^2. */

CREATE OR REPLACE PROCEDURE BigMan(imt int)
AS $$
DECLARE
    mass people%ROWTYPE;
BEGIN
FOR mass IN SELECT * FROM people LOOP
    IF mass.weight / (mass.growth / 100)^2 > imt THEN
        RAISE NOTICE 'id: %, name: %, surname: %', mass.id, mass.name, mass.surname;
    END IF;
END LOOP;
END
$$ LANGUAGE plpgsql;

/* №12. Измените схему БД так, чтобы в БД можно было хранить родственные связи между людьми. Код должен быть представлен в виде транзакции (Например (добавление атрибута): BEGIN; ALTER TABLE people ADD COLUMN leg_size REAL; COMMIT;). Дополните БД данными. */

BEGIN;
CREATE TABLE family_people(
people_id int REFERENCES people(id),
relative_id int REFERENCES people(id));
INSERT INTO family_people(people_id, relative_id)
VALUES (1, 1), (2, 2), (3, 3);
COMMIT;

/* №13. Напишите процедуру, которая позволяет создать в БД нового человека с указанным родством. */

CREATE OR REPLACE PROCEDURE AddPeople
(name varchar, surname varchar, birth_date date, growth real, weight real, eyes varchar, hair varchar, relative_id int, people_1_id int, people_2_id int)
AS $$
DECLARE
    person_id int;
BEGIN
INSERT INTO people (name, surname, birth_date, growth, weight, eyes, hair) VALUES (name, surname, birth_date, growth, weight, eyes, hair) RETURNING id INTO person_id;
INSERT INTO family_people (people_id, relative_id) VALUES (person_id, relative_id);
INSERT INTO family_people (people_id, relative_id) VALUES (people_1_id, person_id);
INSERT INTO family_people (people_id, relative_id) VALUES (people_2_id, person_id);
END
$$ LANGUAGE plpgsql;
CALL AddPeople('Ivan', 'Ivanov', '01.01.2000', 150.0, 50.0, 'black', 'black', 1, 2, 3)

/* №14. Измените схему БД так, чтобы в БД можно было хранить время актуальности данных человека (выполнить также, как п.12). */

BEGIN;
ALTER TABLE people
ADD data_relevance timestamp NOT NULL DEFAULT NOW();
COMMIT;

/* №15. Напишите процедуру, которая позволяет актуализировать рост и вес человека. */

CREATE OR REPLACE PROCEDURE NewGrowthWeight(person_id int, new_growth real, new_weight real)
LANGUAGE plpgsql
AS $$
BEGIN
UPDATE people
SET growth = new_growth, weight = new_weight, data_relevance = NOW()
WHERE people.id = person_id;
END
$$;