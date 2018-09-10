cover:
	@go test -v -coverprofile .testCoverage.txt ./...


cover-html:
	mkdir -p cover
	@go test -v -tags test  -coverprofile=.testCoverage.txt ./... && go tool cover -html=.testCoverage.txt -o cover/coverage.html
