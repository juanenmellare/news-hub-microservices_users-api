go test $(go list ./... | grep -v mocks | grep -v main) -cover
