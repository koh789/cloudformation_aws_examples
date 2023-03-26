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
SERVICE_PROJECT_NAME=$(SERVICE_NAME)-$(ENV)-$(PROJECT_NAME)


# ------------------------ internal ------------------------
.PHONY: internal-cfn-deploy
internal-cfn-deploy: # [ args: CFN_PATH, STACK_NAME, CFN_DEPLOY_ROLE_ARN, ENV, PARAM ]
	aws cloudformation deploy --capabilities CAPABILITY_NAMED_IAM \
		--template $(CFN_PATH) \
		--stack-name $(STACK_NAME) \
		--no-fail-on-empty-changeset \
		--role-arn $(CFN_DEPLOY_ROLE_ARN) \
		--parameter-overrides $(PARAM) \
		--tags "Name=$(STACK_NAME)" "Group=$(SERVICE_PROJECT_NAME)"

.PHONY: 1-cfn-s3-deploy
1-cfn-s3-deploy: ## cfn s3 deploy. [ args: CFN_DEPLOY_ROLE_ARN, ENV ]
	make internal-cfn-deploy CFN_PATH=cfn/1_s3.yml \
		STACK_NAME=$(SERVICE_PROJECT_NAME)-s3-stack \
 		CFN_DEPLOY_ROLE_ARN=$(CFN_DEPLOY_ROLE_ARN) \
 		ENV=$(ENV) \
 		PARAM="Env=$(ENV)"



.PHONY: 2-cfn-glue-deploy
2-cfn-glue-deploy: ## cfn glue deploy. [ args: CFN_DEPLOY_ROLE_ARN, ENV, GLUE_DATABASE_NAME ]
	make internal-cfn-deploy CFN_PATH=cfn/2_glue.yml \
		STACK_NAME=$(SERVICE_PROJECT_NAME)-glue-stack \
		CFN_DEPLOY_ROLE_ARN=$(CFN_DEPLOY_ROLE_ARN) \
		ENV=$(ENV) \
		PARAM="Env=$(ENV) GlueDatabaseName=$(GLUE_DATABASE_NAME)"

.PHONY: 3-cfn-api-kinesis-deploy
3-cfn-api-kinesis-deploy: ## cfn api, kinesis deploy. [ args: CFN_DEPLOY_ROLE_ARN, ENV, S3_KEY_ARN, ALLOW_IP_ADDRESS_LIST, BASE_API_KEY ]
	make internal-cfn-deploy CFN_PATH=cfn/3_api_gateway_kinesis.yml \
		STACK_NAME=$(SERVICE_PROJECT_NAME)-api-kinesis-stack \
		CFN_DEPLOY_ROLE_ARN=$(CFN_DEPLOY_ROLE_ARN) \
		ENV=$(ENV) \
		PARAM="Env=$(ENV) S3KeyArn=$(S3_KEY_ARN) AllowIpAddressList=$(ALLOW_IP_ADDRESS_LIST) BaseApiKey=$(BASE_API_KEY)"

.PHONY: api-gateway-deploy
api-gateway-deploy: ## redeploy the api when you change the definition regarding the api gateway [ args: ENV, REGION ]
	$(eval REST_API_ID :=  `aws cloudformation describe-stacks --stack-name $(SERVICE_PROJECT_NAME)-api-kinesis-stack | jq -r '.Stacks[] | .Outputs[] | select(.OutputKey == "RestApiId") | .OutputValue'`)
	$(eval STAGE := `aws cloudformation describe-stacks --stack-name $(SERVICE_PROJECT_NAME)-api-kinesis-stack | jq -r '.Stacks[] | .Parameters[] | select(.ParameterKey == "ApiGatewayStage") | .ParameterValue'`)
	@echo "REST_API_ID is $(REST_API_ID)"
	@echo "STAGE is $(STAGE)"
	 aws apigateway create-deployment --region $(REGION) \
	   --rest-api-id $(REST_API_ID) \
	   --stage-name $(STAGE)
