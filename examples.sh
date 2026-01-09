#!/bin/bash

# IMDS CLI Usage Examples
# This script demonstrates various ways to use the IMDS CLI

echo "=== IMDS CLI Usage Examples ==="
echo

echo "1. Launch interactive TUI:"
echo "$ imds"
echo "# Opens interactive explorer with tree view and search"
echo

echo "2. Get instance ID (smart lookup):"
echo "$ imds instance-id"
echo "# Output: i-1234567890abcdef0"
echo

echo "3. Get region (finds it automatically):"
echo "$ imds region"
echo "# Output: us-east-1"
echo

echo "4. Get value using explicit path:"
echo "$ imds placement/availability-zone"
echo "# Output: us-east-1a"
echo

echo "5. Tree view of all keys:"
echo "$ imds -r"
echo "# Shows recursive tree structure (keys only)"
echo

echo "6. Dump all data with values:"
echo "$ imds --dump"
echo "# Shows all keys with their values"
echo

echo "7. Dump specific path:"
echo "$ imds spot --dump"
echo "# Dumps spot termination info"
echo

echo "8. Get all data as JSON:"
echo "$ imds --json"
echo "# Dumps all IMDS data as JSON"
echo

echo "9. Get specific path as JSON:"
echo "$ imds events --json"
echo "# Dumps events data as JSON"
echo

echo "10. Watch for changes:"
echo "$ imds spot --watch"
echo "# Monitors for spot instance termination notices"
echo

echo "11. Fuzzy key lookup (typo correction):"
echo "$ imds instanc-id"
echo "# Suggests: instance-id"
echo

echo "Note: Most commands require running on an EC2 instance."
echo "      Use -e flag to specify a custom IMDS endpoint for testing."
