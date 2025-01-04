#!/bin/bash

# Default configuration
FIX_MODE=false
CONFIG_FILE=".golangci.yml"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --fix)
            FIX_MODE=true
            shift
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Check if config file exists
if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: $CONFIG_FILE not found"
    exit 1
fi

# Base command with config file
CMD="golangci-lint run --config $CONFIG_FILE"

# Add --fix flag if in fix mode
if [ "$FIX_MODE" = true ]; then
    CMD="$CMD --fix"
fi

# Run golangci-lint
echo "Running golangci-lint with $CONFIG_FILE..."
$CMD

# Check the exit status
exit_status=$?
if [ $exit_status -eq 0 ]; then
    echo "Linting completed successfully"
else
    echo "Linting failed with exit status $exit_status"
    exit $exit_status
fi