CREATE TABLE IF NOT EXISTS dog_owners (
   dog_owner_id serial PRIMARY KEY,
   name VARCHAR(64) NOT NULL,
   email VARCHAR(255) UNIQUE NOT NULL,
   image TEXT,
   sex CHAR(1),
   reg_at TIMESTAMP,
   upd_at TIMESTAMP
);
