FROM postgres:16.3-alpine3.20

# docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres:16.3-alpine3.20
# docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' some-postgres
# psql "postgres://postgres:@localhost:5432/blogator"
