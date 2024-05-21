DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
    CREATE TYPE gender_enum AS ENUM ('Male', 'Female', 'Other');
  END IF;
END$$;

CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY,
    external_id VARCHAR(255),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    age INTEGER,
    phone VARCHAR(20)[],
    mail VARCHAR(50),
    birthday DATE,
    sex gender_enum,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);