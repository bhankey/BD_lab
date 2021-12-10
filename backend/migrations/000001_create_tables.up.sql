CREATE TABLE IF NOT EXISTS user (
    id serial NOT NULL PRIMARY KEY,
    username varchar(251) NOT NULL DEFAULT '',
    email varchar(511) NOT NULL DEFAULT '',
    creation_date time NOT NULL
) COMMENT 'если будет время)';

CREATE TABLE IF NOT EXISTS  account (
    id serial NOT NULL PRIMARY KEY,
    name varchar(511) NOT NULL,
    user_id int NOT NULL DEFAULT 0,
    is_show bool NOT NULL default true
) COMMENT 'банковский аккаунт';

CREATE TABLE IF NOT EXISTS payment_history (
    id serial NOT NULL PRIMARY KEY,
    account_id int NOT NULL,
    payment_id int NOT NULL,
    date time NOT NULL,
    sum_before double NOT NULL
) COMMENT 'история платежей по аккаунту';

CREATE TABLE IF NOT EXISTS payments (
    id serial NOT NULL PRIMARY KEY,
    account_id int NOT NULL,
    reason varchar(254) NOT NULL,
    sum double NOT NULL,
    date time NOT NULL
) COMMENT 'платежи и начисления по лицевому счету';