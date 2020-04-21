#!/bin/bash
mysql -uweb_user -presu_bew assessment < assessment_create.sql
mysql -uweb_user -presu_bew assessment < assessment_populate.sql
mysql -uweb_user -presu_bew assessment < assessment_views_procedures.sql
mysql -uweb_user -presu_bew assessment < permissions_procedures.sql