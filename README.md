# asterisk_db_migrate

1. Загружаем бекап игнорируя существующие триггеры
```shell
sed 's/TRIGGER/TRIGGER IF NOT EXISTS/g' aviacc-pbx1-mysql-20250530.sql | mysql asterisk_backup
```

2. Даем права пользователю до базы с бекапом
```sql
GRANT ALL PRIVILEGES ON asterisk_backup.* TO 'asteriskuser'@'localhost';
FLUSH PRIVILEGES;
```

4. Запускаем миграцию
```shell
DB_PASSWORD=asteriskpassword ./asterisk_db_migrate
```
