## initial db step

```bash
1. create db
    create database <db_name>
2. create user for connection
    CREATE USER '<username>'@'localhost' IDENTIFIED BY '<pwd>';
    GRANT ALL PRIVILEGES ON db_name.* TO '<username>'@'localhost';
    FLUSH PRIVILEGES;
3. write sql file for table schemas
4. show schema
```
