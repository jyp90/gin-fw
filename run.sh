pid="$(lsof -t -i :8090 -s TCP:LISTEN)";
kill -9 $pid

go run main.go