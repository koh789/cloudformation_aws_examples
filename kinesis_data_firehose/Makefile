.PHONY: help
help: ## help you can see the list of tasks with `make help`.
	@echo "------- task list ------"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36mmake %-20s\033[0m %s\n", $$1, $$2}'

#############################################################
##                  for build, deploy                      ##
#############################################################
SERVICE_NAME=sample
ENV=dev
PROJECT_NAME=kinesis-data-firehose
SERVICE_PROJECT_NAME=$(ENV)-$(PROJECT_NAME)
TEMPLATE_S3_PREFIX=$(PROJECT_NAME)
