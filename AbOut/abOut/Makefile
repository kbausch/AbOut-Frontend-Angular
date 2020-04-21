database_test:
	docker-compose -f docker-compose.database.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.database.test.yml down --volumes

webservice_test:
	docker-compose -f docker-compose.webservice.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.webservice.test.yml down --volumes

frontend_test:
	docker-compose -f docker-compose.frontend.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.frontend.test.yml down --volumes

run:
	docker-compose -f docker-compose.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.yml down --volumes

clean:
	docker-compose down
