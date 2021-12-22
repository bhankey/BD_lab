INSERT INTO account
    (id, name)
VALUES (1, 'first'), (2, 'second'), (3, 'third');

INSERT INTO payments
    (id, account_id, reason, sum, date)
VALUES (1, 1, 'fun', 123, '2021-10-10'), (2, 1, 'fun', 123, '2021-09-09'),
       (3, 1, 'fun', 131223, '2021-08-08'),  (4, 1, 'fun', -131223, '2021-07-07')

INSERT INTO payments_history (account_id, payment_id, date, sum_before) VALUES
                                (1, 1, '2021-10-10', 0),
                                (1, 2, '2021-09-09', 123),
                                (1, 3, '2021-08-08', 246),
                                (1, 4, '2021-07-07', 131223 + 246);


