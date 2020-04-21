#!/usr/bin/env bash

cmd="$@"

until mysql -h${db_host} -P${db_port} -u${db_user} -p${db_pass} ${db_name} -e 'select 1'; do
  >&2 echo "mysql is unavailable - sleeping"
  sleep 3
done

>&2 echo "mysql is up - executing command"
exec $cmd