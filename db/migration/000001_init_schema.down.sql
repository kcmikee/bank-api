-- Drop dependent tables first, then the parent table to satisfy foreign key constraints
DROP TABLE IF EXISTS entries;

DROP TABLE IF EXISTS transfers;

DROP TABLE IF EXISTS accounts;