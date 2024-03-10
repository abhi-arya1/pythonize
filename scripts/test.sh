printf "running build tests at \"./scripts/test.sh\":\n\n"

arch=$(uname -m)

if [ "$arch" = "aarch64" ]; then
    ./pythonize3 --name venv --packages "numpy, matplotlib"
else
    ./pythonize --name venv --packages "numpy, matplotlib"
fi

printf "\n\ntests: passed\n"
