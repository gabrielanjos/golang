# Prova


docker run -p 5432:5432 --name dockerPostgres -e POSTGRES_PASSWORD=123456 -d postgres


docker create -v /var/lib/postgresql/data --name PostgresData alpine
docker run -p 5432:5432 --name yourContainerName -e POSTGRES_PASSWORD=yourPassword -d --volumes-from PostgresData postgres
