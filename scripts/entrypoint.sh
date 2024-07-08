#!/bin/bash

# Run the setup script
/root/create_user_pool_stack.sh

# Start the main application
exec "$@"