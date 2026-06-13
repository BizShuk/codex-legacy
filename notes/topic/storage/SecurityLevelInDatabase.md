# Security Level In Database

### Security Level

1. DB Encryption
   It is on infra level. Generally speaking, it should be done by DB team or infra team to prevent data leak by accessing the hard drive directly.

2. Access Permission
   Common applications will have this. Usually. it depends on the user used to connect DB. It can be archived by minimizing the privileges, somethings like select, insert, update, or create. Application should have select, insert, and update on certain table. But for upgradind, it can have alter and create table, even creat user.

3. Data Encryption
   This is preventing data leak by accessing DB directly. For something like password or confidential data, to store encrypted one, the admin/user of database can not steal the information.
