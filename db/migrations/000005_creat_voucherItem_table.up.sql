CREATE TABLE voucherItem (
    id SERIAL PRIMARY KEY,
    sl_id INT NOT NULL,
    dl_id INT,
    voucher_id INT NOT NULL,
    debit INT CHECK(debit >= 0),
    credit INT CHECK(credit >= 0),
    FOREIGN KEY (sl_id) REFERENCES SL (id),
    FOREIGN KEY (dl_id) REFERENCES DL (id),
    FOREIGN KEY (voucher_id) REFERENCES Voucher (id)

);
-- migrate -path=migrations -database "postgresql://postgres:1234@localhost:5432/project?sslmode=disable" -verbose up 