CREATE DATABASE empresa_internet;

CREATE TABLE
    costumer (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        birth_date DATE NOT NULL,
        state VARCHAR(50) NOT NULL,
        city VARCHAR(50) NOT NULL
    );

CREATE TABLE
    network_plan (
        plan_identity INT PRIMARY KEY AUTO_INCREMENT,
        price INT NOT NULL,
        discount INT NOT NULL,
        velocity INT NOT NULL,
        CONSTRAINT `network_plan_id_costumer_fk` FOREIGN KEY (`id_costumer`) REFERENCES `costumer` (`id`)
    );

INSERT INTO
    costumer (name, last_name, birth_date, state, city)
VALUES
    (
        "Ana",
        "Souza",
        '1999-06-30',
        "São Paulo",
        "São Paulo"
    ),
    (
        "Carlos",
        "Silva",
        '1995-12-15',
        "Rio de Janeiro",
        "Rio de Janeiro"
    ),
    (
        "Maria",
        "Santos",
        '2001-03-20',
        "Minas Gerais",
        "Belo Horizonte"
    ),
    (
        "João",
        "Oliveira",
        '1998-08-10',
        "Bahia",
        "Salvador"
    ),
    (
        "Mariana",
        "Ferreira",
        '1997-05-25',
        "Santa Catarina",
        "Florianópolis"
    ),
    (
        "Pedro",
        "Costa",
        '2000-11-05',
        "Paraná",
        "Curitiba"
    ),
    (
        "Julia",
        "Rodrigues",
        '1996-09-18',
        "Pernambuco",
        "Recife"
    ),
    (
        "Lucas",
        "Almeida",
        '1994-04-12',
        "Ceará",
        "Fortaleza"
    ),
    (
        "Camila",
        "Gomes",
        '1993-07-22',
        "Amazonas",
        "Manaus"
    ),
    (
        "Gustavo",
        "Martins",
        '1992-02-08',
        "Goiás",
        "Goiânia"
    );

INSERT INTO
    network_plan (price, discount, velocity)
VALUES
    (100, 10, 100),
    (200, 20, 200),
    (300, 30, 300),
    (400, 40, 400),
    (500, 50, 500),