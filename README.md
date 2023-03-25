# cloudformation_aws_examples

Example of aws configuration using cloudformation


## AWS Batch
### QuickStart

`cd ./aws_batch`

confirmation of make command.

`make help`

1. cloudformation deploy to create Elastic Container Registry.

```
make ecr-cfn-deploy CFN_DEPLOY_ROLE_ARN=arn:aws:iam::${AccountId}:role/${ROLE_NAME}
```

2. docker deploy

```
 make docker-deploy DOCKER_IMG=${AccountId}.dkr.ecr.${Region}.amazonaws.com IMG_VER=latest
```

3. aws batch deploy. before executing the following command, please create an S3 Bucket for cloudformation package.
   Pass the created S3Bucket as an argument to TEMPLATE_S3_BUCKET

```
 make batch-cfn-deploy \ 
      TEMPLATE_S3_BUCKET=${TEMPLATE_S3_BUCKET}
      CFN_DEPLOY_ROLE_ARN=arn:aws:iam::${AccountId}:role/${ROLE_NAME} \ 
      VPC_ID=${VpcId}  \
      SUBNET_IDS="${SubnetId-A},${SubnetId-B}" \
      IMG_VER=${IMG_VER}
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
