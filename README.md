# IMDS CLI

A command-line interface and interactive TUI for accessing EC2 Instance Metadata Service (IMDS) information.

## Installation

```bash
go install github.com/bwagner5/imds/cmd@latest
```

Or build from source:
```bash
go build -o imds ./cmd/main.go
```

## Usage

### Interactive TUI

Launch the interactive explorer by running `imds` with no arguments:

```bash
imds
```

**TUI Features:**
- Browse IMDS data in a filesystem-like tree view
- View key descriptions from AWS documentation
- Full-screen value viewer for detailed inspection
- Global search (`/`) to find keys across all paths

**TUI Navigation:**
- `↑`/`↓` or `j`/`k` - Navigate
- `Enter` or `Space` - Select (enter directory or view value)
- `←`/`Backspace` or `h` - Go back
- `/` - Search all keys
- `Esc` - Cancel search or return to root
- `q` - Quit

### Query Values

Get a specific value using smart lookup (automatically finds the key):

```bash
imds instance-id
# Output: i-1234567890abcdef0

imds region
# Output: us-east-1

imds availability-zone
# Output: us-east-1a
```

Query using explicit paths:

```bash
imds placement/region
imds dynamic/instance-identity/document
```

### Tree View

List all keys recursively (without values):

```bash
imds -r
# Output:
# meta-data/
#   instance-id
#   instance-type
#   placement/
#     availability-zone
#     region
# dynamic/
#   instance-identity/
#     document
# user-data
```

### Dump All Data

Dump all keys with their values:

```bash
imds --dump
# Output:
# meta-data/
#   instance-id: i-1234567890abcdef0
#   instance-type: m5.large
#   placement/
#     availability-zone: us-east-1a
#     region: us-east-1
```

Dump a specific path:

```bash
imds spot --dump
imds events --dump
```

### JSON Output

Export data as JSON:

```bash
# All data
imds --json > imds-data.json

# Specific path
imds events --json

# Specific key
imds instance-id --json
```

### Watch for Changes

Monitor IMDS data for changes:

```bash
# Watch all data
imds --watch

# Watch specific path (useful for spot termination)
imds spot --watch
```

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--recursive` | `-r` | List all paths recursively (tree, keys only) |
| `--dump` | `-d` | Dump all paths with values |
| `--json` | `-j` | Output as JSON |
| `--watch` | `-w` | Watch for changes |
| `--endpoint` | `-e` | IMDS endpoint (default: http://169.254.169.254) |
| `--version` | | Show version information |

## Smart Key Lookup

The CLI automatically searches for keys, so you don't need to remember full paths:

```bash
imds instance-id      # finds meta-data/instance-id
imds region           # finds meta-data/placement/region
imds availability-zone # finds meta-data/placement/availability-zone
imds document         # finds dynamic/instance-identity/document
```

If a key isn't found, similar keys are suggested:

```bash
imds instanc-id
# Key "instanc-id" not found. Did you mean:
#   - instance-id
#   - elastic-inference-accelerator-id
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `IMDS_ENDPOINT` | Override the default IMDS endpoint |

## Examples

```bash
# Launch interactive TUI
imds

# Get instance information
imds instance-id
imds instance-type
imds ami-id

# Get network information
imds local-ipv4
imds public-ipv4
imds mac

# Get placement information
imds region
imds availability-zone

# Get IAM credentials (if available)
imds iam/security-credentials/

# Get instance identity document
imds dynamic/instance-identity/document

# Monitor spot termination
imds spot --watch

# Export all metadata as JSON
imds --json > metadata.json

# View scheduled maintenance events
imds events --dump
```

## Library Usage

The `pkg/imds` package can be used programmatically:

```go
package main

import (
    "context"
    "fmt"
    "github.com/bwagner5/imds/pkg/imds"
)

func main() {
    ctx := context.Background()
    client, _ := imds.NewClient(ctx, "")

    // Get a specific value
    resp, _ := client.Get(ctx, "meta-data/instance-id")
    fmt.Println(string(resp))

    // Get all data recursively
    data := client.GetAll(ctx, "")
    
    // Find a key by name
    path := client.FindKey(ctx, "instance-id")
    
    // Watch for changes
    for update := range client.Watch(ctx, "meta-data/spot") {
        fmt.Printf("Update: %v\n", update)
    }
}
```
