drop table persons;

create table persons (
  name varchar(255),
  age integer
)

-- psql -U test_user -f example_p sql.sql -d test_dbで作成
-- psql test_dbで操作
-- \dでテーブル一覧を確認
--  drop database test_db;でテーブル削除
-- drop user test_user;でユーザーを削除