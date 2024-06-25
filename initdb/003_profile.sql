DROP DATABASE IF EXISTS accountdb;
CREATE DATABASE accountdb;

\c accountdb

DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
    account_id VARCHAR(250) PRIMARY KEY NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    username VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX accounts_username_index ON accounts (username);
CREATE INDEX accounts_email_index ON accounts (email);
CREATE INDEX accounts_user_id_index ON accounts (user_id);

DROP TABLE IF EXISTS profiles;
CREATE TABLE profiles (
    profile_id VARCHAR(250) PRIMARY KEY NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    avatar VARCHAR(255) NULL,
    bio TEXT NULL,
    location VARCHAR(100) NULL,
    website VARCHAR(255) NULL,
    birth_date DATE NULL,
    phone_number VARCHAR(20) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX profiles_profile_id_index ON profiles (profile_id);
CREATE INDEX profiles_user_id_index ON profiles (user_id);
