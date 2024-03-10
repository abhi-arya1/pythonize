Write-Host "Installing pythonize to your machine..."
Write-Host "Run this script as administrator if it doesn't work."


$source = ".\build\pythonize"
$destination = "$env:ProgramFiles\pythonize"


If(-Not (Test-Path $destination)){
    New-Item -ItemType Directory -Force -Path $destination
}

Copy-Item $source $destination -Force

$systemPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
If(-Not ($systemPath.Split(';') -contains $destination)){
    $newPath = $systemPath + ";" + $destination
    [System.Environment]::SetEnvironmentVariable("Path", $newPath, [System.EnvironmentVariableTarget]::Machine)
}

Write-Host "Installation complete! Have fun! - Abhi"
