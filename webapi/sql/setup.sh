#!/bin/bash

# データベースの接続情報
DB_USER="paimon"          # MySQLのユーザー名
DB_PASS="paimon1234"      # MySQLのパスワード
DB_HOST="127.0.0.1"     # MySQLサーバーのホスト名

# SQLファイルのパス
CREATE_DB_SQL="create_database.sql"
SCHEMA_SQL="schema.sql"

# データベースの作成
echo "Creating database..."
mysql -u $DB_USER -p$DB_PASS -h $DB_HOST < $CREATE_DB_SQL

# スキーマの作成
echo "Creating schema..."
mysql -u $DB_USER -p$DB_PASS -h $DB_HOST genshindb < $SCHEMA_SQL

echo "Database setup completed successfully."