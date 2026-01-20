CREATE TABLE IF NOT EXISTS personalities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    history TEXT
);

INSERT INTO
    personalities (name, history)
VALUES (
        'Albert Einstein',
        'Físico teórico alemão, conhecido por desenvolver a teoria da relatividade.'
    ),
    (
        'Marie Curie',
        'Cientista polonesa-francesa, pioneira na pesquisa sobre radioatividade.'
    ),
    (
        'Isaac Newton',
        'Físico e matemático inglês, formulador das leis do movimento e da gravitação universal.'
    ),
    (
        'Ada Lovelace',
        'Matemática inglesa, considerada a primeira programadora de computadores.'
    ),
    (
        'Nikola Tesla',
        'Inventor e engenheiro elétrico sérvio-americano, conhecido por suas contribuições ao desenvolvimento da corrente alternada.'
    ),
    (
        'Galileo Galilei',
        'Astrônomo, físico e engenheiro italiano, conhecido como o "pai da ciência moderna".'
    ),
    (
        'Charles Darwin',
        'Naturalista inglês, conhecido por sua teoria da evolução por seleção natural.'
    ),
    (
        'Rosalind Franklin',
        'Química inglesa, cujas pesquisas foram fundamentais para a descoberta da estrutura do DNA.'
    ),
    (
        'Stephen Hawking',
        'Físico teórico e cosmólogo inglês, conhecido por seus trabalhos sobre buracos negros e a origem do universo.'
    );