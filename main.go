package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var spinner = []rune{'|', '/', '-', '\\'}

func main() {
	envName := flag.String("name", "venv", "The name for the virtual environment.")
	packages := flag.String("packages", "", "Comma-delimited list of Python packages to install.")
	flag.Parse()

	cmdName := "python3"

	if runtime.GOOS == "windows" {
		cmdName = "python"
	}

	fmt.Printf("setting up virtual environment: \"%s\"\n", *envName)

	createVirtualEnv(cmdName, *envName)

	fmt.Println("upgrading pip...")
	if err := installPackage(cmdName, *envName, "pip --upgrade"); err != nil {
		fmt.Fprintf(os.Stderr, "error upgrading pip: %s\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat("requirements.txt"); err == nil {
		fmt.Println("requirements.txt found. Installing packages from it...")
		installRequirements(cmdName, *envName)
	} else if *packages != "" {
		generateRequirements(*packages)
		installRequirements(cmdName, *envName)
	} else {
		fmt.Println("No requirements.txt found and no packages specified.")
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current working directory: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("virtual environment \"%s\" is ready at: %s\n", *envName, cwd)
}

func createVirtualEnv(cmdName, envName string) error {
	/* Create a virtual environment */
	args := []string{"-m", "venv", envName}
	cmd := exec.Command(cmdName, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating virtual environment: %s\n", err)
		os.Exit(1)
	}
	return err
}

func installPackage(pythonCmd, envName, pkg string) error {
	/* Install a package in the virtual environment */
	var activateCmd string
	if runtime.GOOS == "windows" {
		activateCmd = fmt.Sprintf(".\\%s\\Scripts\\activate", envName)
	} else {
		activateCmd = fmt.Sprintf("source ./%s/bin/activate", envName)
	}

	cmdString := fmt.Sprintf("%s && %s -m pip install %s", activateCmd, pythonCmd, pkg)
	cmd := exec.Command("bash", "-c", cmdString)

	done := make(chan bool)
	go func() {
		i := 0
		for {
			select {
			case <-done:
				fmt.Printf("\ninstalled %s\n", pkg)
				return
			default:
				fmt.Printf("\rinstalling package: %s %c", pkg, spinner[i%len(spinner)])
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	err := cmd.Run()
	done <- true
	fmt.Print("\r\033[K")

	return err
}

func installRequirements(pythonCmd, envName string) {
	/* Construct the activation command based on the OS */
	var activateCmd string
	if runtime.GOOS == "windows" {
		activateCmd = fmt.Sprintf(".\\%s\\Scripts\\activate", envName)
	} else {
		activateCmd = fmt.Sprintf("source ./%s/bin/activate", envName)
	}

	cmdString := fmt.Sprintf("%s && %s -m pip install -r requirements.txt", activateCmd, pythonCmd)
	cmd := exec.Command("bash", "-c", cmdString)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error installing packages from requirements.txt: %s\n", err)
		os.Exit(1)
	}
}

func generateRequirements(packages string) {
	/* Generate requirements.txt from the specified packages */
	packageList := strings.Split(packages, " ")
	err := os.WriteFile("requirements.txt", []byte(strings.Join(packageList, "\n")+"\n"), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing requirements.txt: %s\n", err)
		os.Exit(1)
	}
}
