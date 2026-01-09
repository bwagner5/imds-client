# IMDS CLI

A comprehensive command-line interface for accessing EC2 Instance Metadata Service (IMDS) information.

## Installation

```bash
go build -o imds ./cmd/main.go
```

## Usage

### Basic Queries

Get a specific value:
```bash
# Get instance ID (smart lookup - finds meta-data/instance-id automatically)
imds instance-id
# Output: i-012345

# Get availability zone (smart lookup - finds meta-data/placement/availability-zone)
imds availability-zone
# Output: us-east-1a

# Get region (smart lookup - finds meta-data/placement/region)
imds region
# Output: us-east-1

# Traditional explicit path still works
imds placement/availability-zone
# or
imds placement availability-zone
```

### Listing Available Data

List top-level categories:
```bash
imds
# Output:
# meta-data/
# dynamic/
# user-data

# Or explicitly use list command
imds list
# Output:
# meta-data/
# dynamic/
# user-data
```

List specific category contents:
```bash
# List all dynamic data categories
imds list dynamic

# List all meta-data categories  
imds list meta-data
```

List recursively:
```bash
# List all paths recursively from root
imds list --recursive
# Output:
# meta-data/
# meta-data/instance-id
# meta-data/instance-type
# meta-data/placement/
# meta-data/placement/availability-zone
# meta-data/placement/region
# dynamic/
# dynamic/instance-identity/
# dynamic/instance-identity/document
# user-data

# List dynamic paths recursively
imds list dynamic --recursive
```

### Directory Navigation

Query directories to see their contents:
```bash
# List placement information
imds placement
# Output:
# availability-zone
# host-id
# partition-number
# region

# Then query specific values
imds placement/host-id
```

### JSON Output

Get all IMDS data as JSON:
```bash
# Dump everything to JSON
imds --json > imds-data.json

# Get specific path as JSON
imds placement --json
```

### Watching for Changes

Monitor IMDS data for changes:
```bash
# Watch all data
imds --watch

# Watch specific path
imds spot/termination-time --watch
```

## Command Reference

### Global Flags

- `--endpoint, -e`: IMDS endpoint (default: http://169.254.169.254)
- `--json, -j`: Output in JSON format
- `--recursive, -r`: List all paths recursively (with list command)
- `--watch, -w`: Watch for changes
- `--version`: Show version information

### Commands

#### `imds [path]`
Query a specific IMDS path or dump all data if no path provided.

Examples:
```bash
imds instance-id
imds placement/host-id
imds dynamic/instance-identity/document
```

#### `imds list [category]`
List available paths and categories.

Examples:
```bash
imds list                    # List top-level categories
imds list dynamic            # List dynamic data paths
imds list meta-data          # List meta-data paths
imds list --recursive        # List all paths recursively
```

## Path Shortcuts

The CLI automatically handles path prefixes and includes smart key lookup:

- `imds instance-id` → automatically finds `meta-data/instance-id`
- `imds region` → automatically finds `meta-data/placement/region`
- `imds availability-zone` → automatically finds `meta-data/placement/availability-zone`
- `imds placement/host-id` → `imds meta-data/placement/host-id`
- `imds dynamic/instance-identity` → `imds dynamic/instance-identity`

The smart lookup feature searches recursively through the IMDS tree to find the first occurrence of a key, so you don't need to remember the full path structure.

## Environment Variables

- `IMDS_ENDPOINT`: Override the default IMDS endpoint

## Examples

### Common Use Cases

```bash
# Get basic instance information
imds instance-id
imds instance-type
imds ami-id

# Get network information
imds local-ipv4
imds public-ipv4
imds mac

# Get placement information
imds placement/availability-zone
imds placement/region

# Get security credentials (if available)
imds iam/security-credentials/

# Get dynamic data
imds dynamic/instance-identity/document

# Monitor spot instance termination
imds spot/termination-time --watch

# Export all metadata
imds --json > instance-metadata.json
```

### Advanced Usage

```bash
# List all available paths
imds list --recursive | grep -E "(network|placement)"

# Get specific network interface information
imds network/interfaces/macs/$(imds mac)/

# Watch for maintenance events
imds events/maintenance/scheduled --watch
```

## Output Formats

The CLI intelligently formats output based on the data type:

- **Simple values**: Plain text (e.g., "i-012345")
- **JSON data**: Pretty-printed JSON
- **Directory listings**: One item per line
- **Arrays**: JSON array format

Use `--json` flag to force JSON output for all responses.
