docker_compose:
	docker compose up -d 

docker_exec:
	docker exec -it appProduct bash

.PHONY: docker_compose docker_exec