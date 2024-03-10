npm i
rm -r venv 
rm requirements.txt
GOARCH=amd64 go build -o ./build/pythonize.exe

printf "build: success\n"