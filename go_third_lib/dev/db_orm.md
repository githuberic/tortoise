# DB ORM

## sqlc

Generate type safe Go from SQL

https://github.com/kyleconroy/sqlc

## sql-migrate

SQL schema migration tool for Go.

https://github.com/rubenv/sql-migrate

## gorm 

The fantastic ORM library for Golang, aims to be developer friendly

gorm是一个使用Go语言编写的ORM框架，文档齐全，对开发者友好，并且支持主流的数据库：MySQL, PostgreSQL, SQlite, SQL Server。

gorm最大的好处在于它是由国人开发，中文文档齐全，上手很快，目前大多数企业也都在使用gorm。

我们来一下gorm的特性：
- 全功能 ORM
- 关联 (Has One，Has Many，Belongs To，Many To Many，多态，单表继承)
- Create，Save，Update，Delete，Find 中钩子方法
- 支持 Preload、Joins 的预加载
- 事务，嵌套事务，Save Point，Rollback To Saved Point
- Context、预编译模式、DryRun 模式
- 批量插入，FindInBatches，Find/Create with Map，使用 SQL 表达式、Context Valuer 进行 CRUD
- SQL 构建器，Upsert，数据库锁，Optimizer/Index/Comment Hint，命名参数，子查询
- 复合主键，索引，约束
- Auto Migration
- 灵活的可扩展插件 API：Database Resolver（多数据库，读写分离）、Prometheus…

github地址： https://github.com/go-gorm/gorm

官方文档：https://gorm.io/zh_CN/docs/index.html



