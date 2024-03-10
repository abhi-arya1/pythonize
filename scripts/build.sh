npm i
rm -r venv 
rm requirements.txt
go build -o pythonize

printf "running build tests at \"./scripts/test.sh\":\n\n"

./scripts/test.sh

printf "\n\ntests: passed\n"