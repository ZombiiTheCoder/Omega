g="/usr/local/go/bin/go"

echo "Compiling to \"bin/omega.exe\""
$g build -o bin/omega .
echo "Running \"bin/omega example/cur.om cur.Main\""
./bin/omega example/cur.om cur.Main