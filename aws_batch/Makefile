.PHONY: help
help: ## help you can see the list of tasks with `make help`.
	@echo "------- task list ------"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36mmake %-20s\033[0m %s\n", $$1, $$2}'

#############################################################
##                  for build, deploy                      ##
#############################################################
DOCKER_PATH=./build/package/Dockerfile
SERVICE_NAME=sample
ENV=dev
PROJECT_NAME=aws-batch
SERVICE_PROJECT_NAME=$(SERVICE_NAME)-$(ENV)-$(PROJECT_NAME)
IMG_VER=latest
TEMPLATE_S3_PREFIX=$(PROJECT_NAME)


.PHONY: ecr-cfn-deploy
ecr-cfn-deploy: ## ecr cfn deploy. used to create an ECR for the first time or to modify an ECR definition [ args: ENV, CFN_DEPLOY_ROLE_ARN ]
	aws cloudformation deploy --capabilities CAPABILITY_NAMED_IAM \
			--template cfn/ecr.yml \
			--stack-name $(SERVICE_PROJECT_NAME)-ecr \
			--no-fail-on-empty-changeset \
			--role-arn $(CFN_DEPLOY_ROLE_ARN) \
			--parameter-overrides Env=$(ENV)  \
			--tags "Name=$(SERVICE_PROJECT_NAME)-ecr-stack" "Group=$(SERVICE_PROJECT_NAME)"


.PHONY: docker-build
docker-build: ## docker build.  [ args: DOCKER_IMG, ENV, IMG_VER ]
	docker build -f $(DOCKER_PATH) -t $(DOCKER_IMG)/$(SERVICE_PROJECT_NAME):$(IMG_VER) . --no-cache

.PHONY: docker-deploy
docker-deploy: ## docker build, deploy. [ args: DOCKER_IMG, ENV, IMG_VER ]
	make docker-build DOCKER_IMG=$(DOCKER_IMG) ENV=$(ENV) IMG_VER=$(IMG_VER)
	aws ecr get-login-password | docker login --username AWS --password-stdin $(DOCKER_IMG)
	docker tag $(DOCKER_IMG)/$(SERVICE_PROJECT_NAME):$(IMG_VER) $(DOCKER_IMG)/$(SERVICE_PROJECT_NAME):$(IMG_VER)
	docker push $(DOCKER_IMG)/$(SERVICE_PROJECT_NAME):$(IMG_VER)

.PHONY: batch-cfn-deploy
batch-cfn-deploy: ## batch cfn deploy. s3Bucket for package processing before deploy. [ args: TEMPLATE_S3_BUCKET, CFN_DEPLOY_ROLE_ARN, ENV, VPC_ID, SUBNET_IDS, IMG_VER ]
	aws cloudformation package \
		--template-file cfn/batch.yml \
		--s3-bucket $(TEMPLATE_S3_BUCKET) \
		--s3-prefix $(TEMPLATE_S3_PREFIX)  \
		--output-template-file tmp.yml
	aws cloudformation deploy --capabilities CAPABILITY_NAMED_IAM \
		--template tmp.yml \
		--stack-name $(SERVICE_PROJECT_NAME) \
		--no-fail-on-empty-changeset \
		--role-arn $(CFN_DEPLOY_ROLE_ARN) \
		--parameter-overrides VpcId=$(VPC_ID) SubnetIds=$(SUBNET_IDS) ImgVer=$(IMG_VER)  \
		--tags "Name=$(PROJECT_NAME)-stack" "Group=$(PROJECT_NAME)"
	@echo "remove tmp.yml.. "
	rm -rf tmp.yml
