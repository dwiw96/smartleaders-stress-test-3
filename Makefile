pgRun:
	sudo docker compose -f ./docker-compose.yml up -d
pgStart:
	sudo docker container start pg-smartleaders-system-test
pgExec:
	sudo docker exec -it pg-smartleaders-system-test psql -U dwiw rental_store
pgStop:
	sudo docker container stop pg-smartleaders-system-test

pgCreateTable:
	sudo docker exec -i pg-smartleaders-system-test psql -U dwiw rental_store < table_normalization.sql
pgDropTable:
	sudo docker exec -i pg-smartleaders-system-test psql -U dwiw rental_store < drop_table_normalization.sql

pgInsertData:
	sudo docker exec -i pg-smartleaders-system-test psql -U dwiw rental_store < insert_data.sql
pgSelectData:
	sudo docker exec -i pg-smartleaders-system-test psql -U dwiw rental_store < select_data.sql
