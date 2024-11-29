package db_store

const AddUserQuery = "INSERT INTO shop.users DEFAULT VALUES RETURNING user_id;"
