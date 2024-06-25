\c authdb

DROP TABLE IF EXISTS account_type;
CREATE TABLE account_type (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    account_name VARCHAR(255) NOT NULL
);

CREATE INDEX account_type_id_index ON account_type (id);
CREATE INDEX account_type_account_name_index ON account_type (account_name);

INSERT INTO account_type (id, account_name) 
VALUES ('01J0JRSDYMMXYFCMH834QE9DX3', 'Google'),
       ('01J0JRSRT5MNCA1GKKGWGW5VAR', 'Github'),
       ('01J0JRT1KQ45N5ABQ5S1B4J16H', 'Email');

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    user_id VARCHAR(100) PRIMARY KEY NOT NULL,
    account_type_id VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (account_type_id) REFERENCES account_type(id)
);

CREATE INDEX users_account_type_id_index ON users (account_type_id);
CREATE INDEX users_user_id_index ON users (user_id);
