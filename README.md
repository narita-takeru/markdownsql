
# markdownsql

Create MySQL DDL from markdown text.

markdown exmaple is `test.md`
results DDL is `test.sql`

## Requirements

- go 

## Installation

```bash
$ go get -u github.com/narita-takeru/markdownsql/cmd/markdownsql
```

## Usage
```
markdownsql test.md
```

## example md

------------------------------------------------------------------------

# mydatabase

## users
```
This is users table.
Other description or table comments.
```
### columns
|name|type|null|default|key|comment|
| --- | --- | --- | --- | --- | --- |
|id|bigint|||primary key|user identifier|
|name|varchar(255)|YES|||user name|
| created_at | datetime ||current_timestamp||record created.|
| updated_at | datetime ||current_timestamp on update current_timestamp||record updated.|

### indexes
|columns|unique|
| --- | --- |
|name||

## products
```
This is products table.
```

### columns
|name|type|null|default|key|comment|
| --- | --- | --- | --- | --- | --- |
|id|bigint|||primary key|product identifier|
|category|varchar(255)|YES||||
|name|varchar(255)|YES||||
|price|integer||0|||
| created_at | datetime ||current_timestamp||record created.|
| updated_at | datetime ||current_timestamp on update current_timestamp||record updated.|

### indexes
|columns|unique|
| --- | --- |
|category,name|YES|

------------------------------------------------------------------------
↓
------------------------------------------------------------------------

# example result sql.

```
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL primary key comment 'user identifier',
  `name` varchar(255) NOT NULL comment 'user name',
  `created_at` datetime NOT NULL DEFAULT current_timestamp comment 'record created.',
  `updated_at` datetime NOT NULL DEFAULT current_timestamp on update current_timestamp comment 'record updated.',
  INDEX(`name`)
) ENGINE = InnoDB DEFAULT CHARSET utf8;

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` bigint NOT NULL primary key comment 'product identifier',
  `category` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` integer NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp comment 'record created.',
  `updated_at` datetime NOT NULL DEFAULT current_timestamp on update current_timestamp comment 'record updated.',
  UNIQUE(`category`,`name`)
) ENGINE = InnoDB DEFAULT CHARSET utf8;
```

