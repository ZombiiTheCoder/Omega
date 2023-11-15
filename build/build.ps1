
if ($env:OS -eq "Windows_NT") {
    Write-Host "Compiling to bin/omega.exe"
    go build -o bin/omega.exe .
} elseif ($env:OSTYPE -eq "linux") {
    Write-Host "Compiling to bin/omega"
    go build -o bin/omega .
} elseif ($env:OSTYPE -eq "darwin") {
    Write-Host "Compiling to bin/omega"
    go build -o bin/omega .
} else {
    Write-Host "Unsupported operating system"
}
