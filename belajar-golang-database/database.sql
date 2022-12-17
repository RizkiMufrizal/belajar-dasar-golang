create table customer(
    id varchar(100) not null,
    name varchar(100)not null,
    primary key(id)
)engine=InnoDb;

DELETE from customer;

ALTER TABLE customer
    ADD COLUMN email varchar(100),
    ADD COLUMN balance integer DEFAULT 0,
    ADD COLUMN rating double DEFAULT 0.0,
    ADD COLUMN created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT false;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES('rizki', 'Rizki', 'rizki@gmail.com', 10000, 5.0, '1999-8-9', true);

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES('budi', 'Budi', 'budi@gmail.com', 5000, 3.0, '2000-8-9', false);

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES('joko', 'Joko', NULL, 5000, 3.0, NULL, false);

-- Table User
CREATE TABLE user(
    username varchar(100) not null,
    password varchar(100) not null,
    primary key (username)
) engine=InnoDb;

INSERT INTO user(username, password)
VALUES('rizki', 'rizki');

-- Table Comment

CREATE TABLE comments(
    id int not null auto_increment,
    email varchar(100) not null,
    comment text,
    primary key(id)
)engine=InnoDb;