CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    email    VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)        NOT NULL
);

CREATE TABLE events
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    date        DATE         NOT NULL,
    address     VARCHAR(255) NOT NULL,
    image_url   VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL
);

INSERT INTO events(title, date, address, image_url, description)
VALUES ('Доброволческа акция: Почисти природата',
        '2025-04-15',
        'Парк „Борисова градина“, София',
        'https://i.postimg.cc/4xptQzXm/image1.avif',
        'Присъединете се към нас в еднодневна инициатива за почистване на парка! Ще осигурим ръкавици и чували, а в края на деня ще има пикник за участниците.'),
       ('Как да създаваме екологични навици?',
        '2025-04-25',
        'Зала „Грийн Лаб“, Пловдив',
        'https://i.postimg.cc/TYVgrVdh/image2.avif',
        'Практически семинар за изграждане на устойчиви навици в ежедневието – от намаляване на отпадъците до екологично пазаруване. Включени са интерактивни задачи и подаръци за участниците.'),
       ('Фестивал на уличното изкуство',
        '2025-05-05',
        'Център на Варна',
        'https://i.postimg.cc/dtPGkHKz/image3.avif',
        ' Потопи се в света на уличното изкуство! Очакват те графити демонстрации, музика на живо и арт базар с местни творци.');

CREATE TABLE registrations
(
    id      SERIAL PRIMARY KEY,
    user_id  INT NOT NULL REFERENCES users (id),
    event_id INT NOT NULL REFERENCES events (id)
);



