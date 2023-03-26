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
PROJECT_NAME=ecs-executor
SERVICE_PROJECT_NAME=$(SERVICE_NAME)-$(ENV)-$(PROJECT_NAME)
CLUSTER_ARN=
IMG_VER=latest

.PHONY: ecr-cfn-deploy
ecr-cfn-deploy: ## ecr cfn deploy. used to create an ECR for the first time or to modify an ECR definition [ args: ENV, FN_DEPLOY_ROLE_ARN ]
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
	make docker-build DOCKER_IMG=$(DOCKER_IMG) IMG_VER=$(IMG_VER)
	aws ecr get-login-password | docker login --username AWS --password-stdin $(DOCKER_IMG)
	docker tag $(DOCKER_IMG)/$(SERVICE_PROJECT_NAME):$(IMG_VER) $(DOCKER_IMG)/$(SERVICE_PROJECT_NAME):$(IMG_VER)
	docker push $(DOCKER_IMG)/$(SERVICE_PROJECT_NAME):$(IMG_VER)

.PHONY: ecs-cfn-deploy
ecs-cfn-deploy: ## ecs cfn deploy. s3Bucket for package processing before deploy. [ args: TEMPLATE_S3_BUCKET, CFN_DEPLOY_ROLE_ARN, VPC_ID, SUBNET_IDS, IMG_VER, CLUSTER_ARN, ENV ]
	aws cloudformation deploy --capabilities CAPABILITY_NAMED_IAM \
		--template cfn/task.yml \
		--stack-name $(SERVICE_PROJECT_NAME) \
		--no-fail-on-empty-changeset \
		--role-arn $(CFN_DEPLOY_ROLE_ARN) \
		--parameter-overrides VpcId=$(VPC_ID) SubnetIds=$(SUBNET_IDS) ImgVer=$(IMG_VER) ClusterArn=$(CLUSTER_ARN) Env=$(ENV)  \
		--tags "Name=$(SERVICE_PROJECT_NAME)-stack" "Group=$(SERVICE_PROJECT_NAME)"
