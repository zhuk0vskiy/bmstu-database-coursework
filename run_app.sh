set -x -e
docker container start studios_db
cd backend && go build src/cmd/main.go
mv main ../
