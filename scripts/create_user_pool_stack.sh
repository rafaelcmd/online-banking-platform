#!/bin/bash

# Create the Cognito User Pool stack
echo "Creating Cognito User Pool stack..."
/usr/local/bin/aws cloudformation create-stack \
    --stack-name cognito-user-pool \
    --template-body file:///root/create_user_pool.yaml \
    --capabilities CAPABILITY_NAMED_IAM

# Check if the stack creation command succeeded
if [ $? -ne 0 ]; then
    echo "Failed to initiate stack creation."
    exit 1
fi

# Wait for the stack creation to complete
echo "Waiting for stack creation to complete..."
/usr/local/bin/aws cloudformation wait stack-create-complete --stack-name cognito-user-pool

# Get the User Pool Client ID from the stack output
echo "Retrieving User Pool Client ID..."
USER_POOL_CLIENT_ID=$(/usr/local/bin/aws cloudformation describe-stacks \
    --stack-name cognito-user-pool \
    --output text \
    --query 'Stacks[0].Outputs[?OutputKey==`UserPoolClientId`].OutputValue')

# Put USER_POOL_CLIENT_ID on Systems Manager Parameter Store
aws ssm put-parameter --name "USER_POOL_CLIENT_ID" --value "$USER_POOL_CLIENT_ID" --type "String"

# Check if the parameter was stored successfully
if [ $? -ne 0 ]; then
    echo "Failed to store User Pool Client ID in SSM Parameter Store."
    exit 1
fi

echo "Script completed successfully."