#!/usr/bin/env bash
#wait for the MySQL Server to come up
sleep 1s

#run the setup script to create the DB and the schema in the DB
mysql -u dialog -pdlgtest -D dlgdblocal < "/docker-entrypoint-initdb.d/001-create-tables.sql"
mysql -u dialog -pdlgtest -D dlgdblocal < "/docker-entrypoint-initdb.d/002-insert-campaign.sql"
