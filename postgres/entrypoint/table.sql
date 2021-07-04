CREATE TABLE users (
  id SERIAL NOT NULL,
  created_at timestamp NOT NULL, 
  updated_at timestamp NOT NULL,
  username varchar(100) UNIQUE NOT NULL,
  password varchar(100) NOT NULL,
  PRIMARY KEY (id)
);
