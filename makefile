run:
	sh .dev_environment/scripts/run.sh

format:
	sh .dev_environment/scripts/format_code.sh

tests:
	sh .dev_environment/scripts/test_code.sh

tests-coverage:
	sh .dev_environment/scripts/test_code_coverage.sh

update-mocks:
	mockery --all --keeptree --case=snake