CREATE USER root WITH PASSWORD 'root';

DROP DATABASE IF EXISTS authdb;
CREATE DATABASE authdb;

GRANT ALL PRIVILEGES ON DATABASE authdb TO root;