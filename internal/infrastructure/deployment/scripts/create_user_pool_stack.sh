#!/bin/bash

# Create the Cognito User Pool stack
aws cloudformation create-stack \
    --stack-name cognito-user-pool \
    --template-body file://../cloudformation/create_user_pool.yaml \
    --capabilities CAPABILITY_NAMED_IAM

# Wait for the stack creation to complete
aws cloudformation wait stack-create-complete --stack-name cognito-user-pool

# Get the User Pool Client ID from the stack output
USER_POOL_CLIENT_ID=$(aws cloudformation describe-stacks \
    --stack-name cognito-user-pool \
    --output text \
    --query 'Stacks[0].Outputs[?OutputKey==`UserPoolClientId`].OutputValue')

# Print the environment variable to verify it's set
echo "USER_POOL_CLIENT_ID: $USER_POOL_CLIENT_ID"