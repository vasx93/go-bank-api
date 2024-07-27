createdb:
	docker exec -it bank-db createdb --username=root --owner=root bank

dropdb:
	docker exec -it bank-db dropdb bank

.PHONY: createdb