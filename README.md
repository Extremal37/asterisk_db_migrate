# asterisk_db_migrate

1. Загружаем бекап игнорируя существующие триггеры
```shell
sed 's/TRIGGER/TRIGGER IF NOT EXISTS/g' aviacc-pbx1-mysql-20250530.sql | mysql asterisk_backup
```