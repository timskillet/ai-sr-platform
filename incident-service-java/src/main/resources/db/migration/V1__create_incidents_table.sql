CREATE TABLE incidents (
                           id          UUID         PRIMARY KEY,
                           service_name VARCHAR(255) NOT NULL,
                           severity    VARCHAR(50)  NOT NULL,
                           status      VARCHAR(50)  NOT NULL,
                           message     TEXT         NOT NULL,
                           created_at  TIMESTAMP    NOT NULL
);