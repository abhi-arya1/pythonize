npm i
rm -r venv 
rm requirements.txt
GOARCH=amd64 go build -o ./build/pythonize

./scripts/test.sh

printf "build: success\n"