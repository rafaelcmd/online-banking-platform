#!/bin/bash

# Create the Cognito User Pool stack
aws cloudformation create-stack \
    --stack-name cognito-user-pool \
    --template-body file://../cloudformation/create_user_pool.yaml \
    --capabilities CAPABILITY_NAMED_IAM > /dev/null 2>&1

# Wait for the stack creation to complete
aws cloudformation wait stack-create-complete --stack-name cognito-user-pool > /dev/null 2>&1

# Get the User Pool Client ID from the stack output
USER_POOL_CLIENT_ID=$(aws cloudformation describe-stacks \
    --stack-name cognito-user-pool \
    --output text \
    --query 'Stacks[0].Outputs[?OutputKey==`UserPoolClientId`].OutputValue')

# Put USER_POOL_CLIENT_ID on Systems Manager Parameter Store
aws ssm put-parameter --name "USER_POOL_CLIENT_ID" --value "$USER_POOL_CLIENT_ID" --type "String" > /dev/null 2>&1