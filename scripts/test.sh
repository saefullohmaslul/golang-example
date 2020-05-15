mkdir -p coverage
make test-unit
make test-int

# merge coverage
grep 'mode' coverage/unit_test.txt > coverage/coverage.tmp.txt
grep 'github' coverage/unit_test.txt >> coverage/coverage.tmp.txt
grep 'github' coverage/integration_test.txt >> coverage/coverage.tmp.txt

# remove ignore
grep -A100 'ignore:' ignore_test.yml | grep -v "ignore:" | sed 's/ //g' > coverage/coverage.ignore.txt
grep -v -F -f coverage/coverage.ignore.txt coverage/coverage.tmp.txt > coverage/coverage.txt

go tool cover -func=coverage/coverage.txt
go tool cover -html=coverage/coverage.txt
go tool cover -html=coverage/coverage.txt -o coverage/index.html

rm -f coverage/coverage.tmp* coverage/coverage.ignore*