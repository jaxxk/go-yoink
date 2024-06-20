-- +goose Up
-- Step 1: Drop the existing primary key constraint
ALTER TABLE feeds DROP CONSTRAINT feeds_pkey;

-- Step 2: Add the new primary key column
ALTER TABLE feeds ADD COLUMN id VARCHAR(64) UNIQUE NOT NULL;

-- Step 3: Set the new primary key
ALTER TABLE feeds ADD PRIMARY KEY (id);

-- Step 4: Add created_at and updated_at columns
ALTER TABLE feeds ADD COLUMN created_at TIMESTAMP NOT NULL;
ALTER TABLE feeds ADD COLUMN updated_at TIMESTAMP NOT NULL;

-- +goose Down
-- Step 1: Drop the new primary key constraint
ALTER TABLE feeds DROP CONSTRAINT feeds_pkey;

-- Step 2: Drop the added columns
ALTER TABLE feeds DROP COLUMN id;
ALTER TABLE feeds DROP COLUMN created_at;
ALTER TABLE feeds DROP COLUMN updated_at;

-- Step 3: Restore the previous primary key constraint
ALTER TABLE feeds ADD PRIMARY KEY (name);
