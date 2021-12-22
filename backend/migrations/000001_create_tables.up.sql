CREATE TABLE IF NOT EXISTS "user"(
                                     id serial NOT NULL PRIMARY KEY,
                                     username varchar(251) NOT NULL DEFAULT '',
                                     email varchar(511) NOT NULL DEFAULT '',
                                     creation_date timestamp NOT NULL
);

COMMENT ON table "user" is 'если будет время)';

CREATE TABLE IF NOT EXISTS  account (
                                        id serial NOT NULL PRIMARY KEY,
                                        name varchar(511) NOT NULL,
                                        user_id int NOT NULL DEFAULT 0,
                                        is_show bool NOT NULL default true,
                                        sum double precision NOT NULL DEFAULT 0
);
COMMENT ON table account is 'банковский аккаунт';

CREATE TABLE IF NOT EXISTS payments_history (
                                               id serial NOT NULL PRIMARY KEY,
                                               account_id int NOT NULL,
                                               payment_id int NOT NULL,
                                               date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                               sum_before double precision NOT NULL
);
COMMENT ON table payments_history is 'история платежей по аккаунту';

CREATE UNIQUE INDEX IF NOT EXISTS payments_history_uniq ON payments_history
(
    account_id,
    payment_id
);

CREATE TABLE IF NOT EXISTS payments (
                                        id serial NOT NULL PRIMARY KEY,
                                        account_id int NOT NULL,
                                        reason varchar(254) NOT NULL,
                                        sum double precision NOT NULL,
                                        date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON table payments is 'платежи и начисления по лицевому счету';
