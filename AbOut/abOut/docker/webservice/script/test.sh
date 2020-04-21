#!/usr/bin/env bash
mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} < ./database/assessment_create.sql
mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} < ./database/assessment_populate.sql
mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} < ./database/assessment_views_procedures.sql
mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} < ./database/permissions_procedures.sql

echo 'Deleting mysql-client...'
apk del mysql-client

echo 'Test application...'
cd ./webService; CGO_ENABLED=0 go test ./...