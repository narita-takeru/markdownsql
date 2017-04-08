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


