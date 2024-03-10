npm i
rm -r venv 
rm requirements.txt
GOARCH=arm64 go build -o ./build/pythonize3

./scripts/test.sh

printf "build: success\n"