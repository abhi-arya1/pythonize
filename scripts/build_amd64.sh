npm i
rm -r venv 
rm requirements.txt
GOARCH=amd64 go build -o ./build/pythonize

printf "build: success\n"