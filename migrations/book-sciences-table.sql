DROP TABLE IF EXISTS book_sciences;
CREATE TABLE book_sciences (
                      id            INT AUTO_INCREMENT NOT NULL,
                      book_id       INT,
                      science_id    INT,
                      PRIMARY KEY (`id`),
                      foreign key (`book_id`) REFERENCES book(`id`),
                      foreign key (`science_id`) REFERENCES science(`id`)
);
