CREATE TABLE parking(
 id serial PRIMARY KEY,
 number SMALLINT,
 owner VARCHAR,
 car VARCHAR
);


-- INSERT INTO parking(number, owner, car) VALUES
--     (10, 'Foo Bar', 'Dodge Viper'),
--     (20, 'James Bond', 'Aston Martin');
