set -x -e
export PATH=$PATH:/usr/local/go/bin
docker container start studios_db
cd backend && go build src/cmd/main.go
mv main ../
