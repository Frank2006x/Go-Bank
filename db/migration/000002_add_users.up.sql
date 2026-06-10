CREATE TABLE USERS (
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    password_changed_at timestamptz,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP
);





CREATE INDEX idx_users_username ON USERS(username);
ALTER TABLE ACCOUNTS ADD FOREIGN KEY (owner) REFERENCES USERS(username);
ALTER TABLE ACCOUNTS ADD CONSTRAINT "OWNER_CURRENCY_KEY" UNIQUE (owner, currency); 