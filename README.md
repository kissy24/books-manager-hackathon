# books-manager-hackathon

![](https://img.shields.io/github/go-mod/go-version/kissy24/go-reading-books)
![](https://img.shields.io/github/license/kissy24/go-reading-books)

System for managing books API.

I created this for a hackathon at my company.

## API Reference

- `/` : SELECT * from books
- `/add` : INSERT INTO books VALUES(title, isbn, publisher)

## Database

### Books

|cid|name|type|notnull|dflt_value|pk|
|--|--|--|--|--|--|
|0|id|integer|0||1|
|1|created_at|datetime|0||0|
|2|updated_at|datetime|0||0|
|3|deleted_at|datetime|0||0|
|4|title|varchar(255)|0||0|
|5|isbn|integer|0||0|
|6|publisher|varchar(255)|0||0|

## Requirements

- Go 1.20
- github.com/gin-gonic/gin v1.9.1
- github.com/jinzhu/gorm v1.9.16
- github.com/mattn/go-sqlite3 v1.14.17

## License

MIT

## Author

kissy24(yuhei-kishida)
