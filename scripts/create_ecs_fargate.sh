#!/bin/bash

# Create ECR repository and capture the URI
REPO_URI=$(/usr/local/bin/aws ecr create-repository --repository-name online-bank-auth-service --region us-east-2 --query 'repository.repositoryUri' --output text)

# Tag the Docker image with the repository URI
docker tag my-online-bank-auth-service:latest $REPO_URI:latest

# Push the Docker image to ECR
$(/usr/local/bin/aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin $REPO_URI)
docker push $REPO_URI:latest

# Full image name
FULL_IMAGE_NAME=$REPO_URI:latest

# Create CloudWatch log group
/usr/local/bin/aws logs create-log-group --log-group-name /ecs/online-bank-platform --region us-east-2

# Create or update the CloudFormation stack with the full image name as a parameter
/usr/local/bin/aws cloudformation create-stack \
  --stack-name online-bank-auth-service-stack \
  --template-body file://./cloudformation/create_ecs_fargate.yaml \
  --capabilities CAPABILITY_IAM \
  --parameters ParameterKey=ContainerImage,ParameterValue=$FULL_IMAGE_NAME