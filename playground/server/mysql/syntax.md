# MYSQL syntax

# database
- `show databases`
- `drop database <db_name>`
- `create DATABSE [IF NOT EXISTS] db_name [DEFAULT] CHARACTER SET= utf8`
#: utf8_general_ci , utf8_unicode_ci => unknown character set
- `SET storage_engine=[MYISAM|InnoDB];`

# table
- `show tables`
- `CREATE TABLE t (i INT) ENGINE = INNODB;`


### select

##### Sample:
```
SELECT table_name AS "Table",
round(((data_length + index_length) / 1024 / 1024), 2) "Size in MB"
FROM information_schema.TABLES
WHERE table_schema = "$DB_NAME"
 AND table_name = "$TABLE_NAME";
```

# authority
- `create USER '<username>'@'<host>' IDENTIFIED BY '<passwd>'` , host could be `%` for all source.
- `grant <privileges> ON <db>.<table> TO '<username>'@'<host>'` , dbname and table name could be `*`
    + privileges list:
        + ALL PRIVILEGES
        + CREATE
        + DROP
        + DELETE
        + INSERT
        + SELECT
        + UPDATE
        + GRANT OPTION
- `FLUSH PRIVILEGES` , update privileges
- `drop USER '<username>'@'<host>' `

- `select * from mysql.user` , get mysql account list
- `show grants [for <username>]` , show current user grant
- `rename user '<username>'@'<host>' to '<username>'@'<host>'`
- `set password for '<username>'@'<host>' = PASSWORD('<password>')`


# system info
- `show variables like "%version%"`;
- 