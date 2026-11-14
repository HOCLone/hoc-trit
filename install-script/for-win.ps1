New-Item -ItemType Directory -Path ".\output\bin" -Force
go build -o ".\output\bin\hoc-trit.exe" ".\src\hoc-trit.go"
