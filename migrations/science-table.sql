DROP TABLE IF EXISTS science;
CREATE TABLE science (
                      id         INT AUTO_INCREMENT NOT NULL,
                      name      VARCHAR(128) NOT NULL UNIQUE,
                      PRIMARY KEY (`id`)
);
