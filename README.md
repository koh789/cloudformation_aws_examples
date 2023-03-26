# cloudformation_aws_examples

Example of aws configuration using cloudformation


## AWS Batch
### QuickStart

[Definition file: aws_batch](./aws_batch)

`cd ./aws_batch`

confirmation of make command.

`make help`

1. cloudformation deploy to create Elastic Container Registry. ENV can be selected from dev, stg, prd

```
make ecr-cfn-deploy \ 
    ENV=stg \  
    CFN_DEPLOY_ROLE_ARN=arn:aws:iam::${AccountId}:role/${RoleName}
```

2. docker deploy

```
 make docker-deploy \ 
    DOCKER_IMG=${AccountId}.dkr.ecr.${Region}.amazonaws.com \  
    ENV=stg  \
    IMG_VER=latest
```

3. aws batch deploy. before executing the following command, please create an S3 Bucket for cloudformation package.
   Pass the created S3Bucket as an argument to TEMPLATE_S3_BUCKET

```
 make batch-cfn-deploy \ 
      TEMPLATE_S3_BUCKET=${TemplateS3Bucket}
      CFN_DEPLOY_ROLE_ARN=arn:aws:iam::${AccountId}:role/${RoleName} \
      ENV=stg \ 
      VPC_ID=${VpcId}  \
      SUBNET_IDS="${SubnetId-A},${SubnetId-B}" \
      IMG_VER=${ImgVer}
```

The above command will create the following resources

* ECR Repository
* IAM Role
  * batch service role
  * execution role
  * events role
* Security group
* Compute environment
* Job queue
* Job definitions
* Events rule (Event Bridge)

## ECS task
### QuickStart

[Definition file: ecs_task](./ecs_task)

`cd ./ecs_task`

confirmation of make command.

`make help`

1. cloudformation deploy to create Elastic Container Registry.

```
make ecr-cfn-deploy \ 
    ENV=stg \       
    CFN_DEPLOY_ROLE_ARN=arn:aws:iam::${AccountId}:role/${RoleName}
```

2. docker deploy

```
 make docker-deploy \  
    DOCKER_IMG=${AccountId}.dkr.ecr.${Region}.amazonaws.com \
    ENV=stg  \
    IMG_VER=latest
```

3. ecs deploy. 
   * before executing the following command, please create an S3 Bucket for cloudformation package. 
   * pass the created S3Bucket as an argument to TEMPLATE_S3_BUCKET (bucket_name)
   * also, create an ECS cluster for ECS tasks
   * ECS cluster should be passed ARN as argument of CLUSTER_ARN

```
 make ecs-cfn-deploy \ 
      TEMPLATE_S3_BUCKET=${TemplateS3Bucket}
      CFN_DEPLOY_ROLE_ARN=arn:aws:iam::${AccountId}:role/${RoleName} \
      VPC_ID=${VpcId}  \
      SUBNET_IDS="${SubnetId-A},${SubnetId-B}" \
      IMG_VER=${ImgVer} \ 
      CLUSTER_ARN=${ClusterArn} \ 
      ENV=stg \

```

The above command will create the following resources

* ECR Repository
* LogGroup
* IAM Role
    * execution role
    * events role
* Task definitions
* Events rule (Event Bridge)