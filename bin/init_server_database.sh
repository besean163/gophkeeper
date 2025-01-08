#!/bin/bash

# Запрашиваем логин и пароль суперпользователя
read -p "Введите логин суперпользователя (по умолчанию: postgres): " DB_SUPERUSER
DB_SUPERUSER=${DB_SUPERUSER:-postgres}

read -s -p "Введите пароль суперпользователя: " DB_SUPERUSER_PASSWORD
echo # Переход на новую строку после ввода пароля

# Зашитые параметры подключения
DB_HOST="localhost"
DB_PORT="5432"

# Зашитые параметры создаваемого пользователя и базы
DB_USER="gophkeeper"
DB_USER_PASSWORD="gophkeeper"
DB_NAME="gophkeeper"

# Выполняем команды
export PGPASSWORD=$DB_SUPERUSER_PASSWORD

# создаем пользователя если его нет
psql -U $DB_SUPERUSER -h $DB_HOST -p $DB_PORT <<EOF
DO
\$do\$
BEGIN
   IF NOT EXISTS (
      SELECT
      FROM   pg_catalog.pg_user
      WHERE  usename = '${DB_USER}') THEN
      CREATE USER ${DB_USER} WITH PASSWORD '${DB_USER_PASSWORD}';
   END IF;
END
\$do\$;
EOF

# Создаем бд
psql -U $DB_SUPERUSER -h $DB_HOST -p $DB_PORT -c "CREATE DATABASE ${DB_NAME} OWNER ${DB_USER};"

unset PGPASSWORD
