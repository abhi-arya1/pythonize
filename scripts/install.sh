printf "installing pythonize to your machine... \nyou may be prompted to input a password...\n"
./scripts/build.sh > /dev/null
sudo mv pythonize /usr/local/bin 
rm -r venv 
rm requirements.txt 
printf "installation complete! have fun! - abhi\n "
deactivate