#!/usr/bin/env bash

echo 'Restoring database state...'
mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} < ./assessment_create.sql
mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} < ./assessment_populate.sql
mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} < ./assessment_views_procedures.sql
mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} < ./permissions_procedures.sql

echo 'Test application...'
dotnet test