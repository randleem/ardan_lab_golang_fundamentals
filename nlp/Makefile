test:
	go tool staticcheck .
	go test -v .

fuzz: 
	go test -run NOE -fuzz . -fuzztime 10s