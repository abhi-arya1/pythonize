npm i
rm -r venv 
rm requirements.txt
go build -o pythonize

./scripts/test.sh

printf "build: success\n"