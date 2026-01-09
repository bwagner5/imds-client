#!/bin/bash

# IMDS CLI Usage Examples
# This script demonstrates various ways to use the IMDS CLI

echo "=== IMDS CLI Usage Examples ==="
echo

echo "1. List top-level categories:"
echo "$ imds list"
./imds list
echo

echo "2. Get instance ID (when on EC2):"
echo "$ imds instance-id"
echo "# Would output: i-012345"
echo

echo "3. Get placement information:"
echo "$ imds placement"
echo "# Would list: availability-zone, host-id, partition-number, region"
echo

echo "4. Get specific placement value:"
echo "$ imds placement/availability-zone"
echo "# Would output: us-west-2a"
echo

echo "5. Alternative syntax:"
echo "$ imds placement availability-zone"
echo "# Same as above"
echo

echo "6. List dynamic data:"
echo "$ imds list dynamic"
echo "# Lists all dynamic data categories"
echo

echo "7. Get all data as JSON:"
echo "$ imds --json"
echo "# Dumps all IMDS data as JSON"
echo

echo "8. Watch for spot termination:"
echo "$ imds spot/termination-time --watch"
echo "# Monitors for spot instance termination notices"
echo

echo "9. List all paths recursively:"
echo "$ imds list --recursive"
echo "# Shows complete tree structure"
echo

echo "Note: Most commands will fail when not running on an EC2 instance"
echo "      but the CLI structure and help system work anywhere."
