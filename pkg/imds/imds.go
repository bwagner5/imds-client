package imds

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/smithy-go/transport/http"
)

const (
	IPv4Mode        = "ipv4"
	IPv6Mode        = "ipv6"
	spotITNPath     = "spot/termination-time"
	scheduledEvents = "events/maintenance/scheduled"
)

type IMDS struct {
	client *imds.Client
}

type ScheduledEventDetail struct {
	NotBefore   string `json:"NotBefore"`
	Code        string `json:"Code"`
	Description string `json:"Description"`
	EventID     string `json:"EventId"`
	NotAfter    string `json:"NotAfter"`
	State       string `json:"State"`
}

type InstanceAction struct {
	Action string `json:"action"`
	Time   string `json:"time"`
}

type RebalanceRecommendation struct {
	NoticeTime string `json:"noticeTime"`
}

func NewClient(ctx context.Context, endpoint string, ipMode string) (*IMDS, error) {
	cfg, err := config.LoadDefaultConfig(ctx, withIMDSEndpoint(endpoint), withIPMode(ipMode))
	if err != nil {
		return nil, err
	}
	return &IMDS{
		client: imds.NewFromConfig(cfg),
	}, nil
}

func withIMDSEndpoint(imdsEndpoint string) func(*config.LoadOptions) error {
	return func(lo *config.LoadOptions) error {
		lo.EC2IMDSEndpoint = imdsEndpoint
		return nil
	}
}

func withIPMode(ipMode string) func(*config.LoadOptions) error {
	return func(lo *config.LoadOptions) error {
		if ipMode == IPv6Mode {
			lo.EC2IMDSEndpointMode = imds.EndpointModeStateIPv6
		} else if ipMode == IPv4Mode {
			lo.EC2IMDSEndpointMode = imds.EndpointModeStateIPv4
		} else {
			return fmt.Errorf("invalid IMDS IP Mode \"%s\"", ipMode)
		}
		return nil
	}
}

func (i IMDS) GetAMIID(ctx context.Context) (string, error) {
	amiID, err := i.GetMetadata(ctx, "ami-id")
	if err != nil {
		return "", err
	}
	return amiID, nil
}

func (i IMDS) GetAMILaunchIndex(ctx context.Context) (int, error) {
	amiLaunchIndex, err := i.GetMetadata(ctx, "ami-launch-index")
	if err != nil {
		return 0, err
	}
	launchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		return 0, fmt.Errorf("unable to convert ami-launch-index of %s to integer: %w", amiLaunchIndex, err)
	}
	return launchIndexNum, nil
}

func (i IMDS) GetAMIManifestPath(ctx context.Context) (string, error) {
	amiManifestPath, err := i.GetMetadata(ctx, "ami-manifest-path")
	if err != nil {
		return "", err
	}
	return amiManifestPath, nil
}

func (i IMDS) GetHostname(ctx context.Context) (string, error) {
	hostname, err := i.GetMetadata(ctx, "hostname")
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func (i IMDS) GetInstanceAction(ctx context.Context) (string, error) {
	instanceAction, err := i.GetMetadata(ctx, "instance-action")
	if err != nil {
		return "", err
	}
	return instanceAction, nil
}

func (i IMDS) GetMetadata(ctx context.Context, path string) (string, error) {
	out, err := i.client.GetMetadata(ctx, &imds.GetMetadataInput{
		Path: path,
	})
	if err != nil {
		return "", fmt.Errorf("unable to retrieve \"%s\" metadata: %w", path, err)
	}
	content, err := io.ReadAll(out.Content)
	if err != nil {
		return "", fmt.Errorf("unable to read response of \"%s\" metadata: %w", path, err)
	}
	return string(content), nil
}

// TODO(bwagner5): use spot/instance-action instead
func (i IMDS) GetSpotInterruptionNotification(ctx context.Context) (*time.Time, bool, error) {
	output, err := i.client.GetMetadata(ctx, &imds.GetMetadataInput{Path: spotITNPath})
	var re *http.ResponseError
	if errors.As(err, &re) && re.HTTPStatusCode() == 404 {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("IMDS Failed to get \"%s\": %w", spotITNPath, err)
	}
	termTimeBytes := new(bytes.Buffer)
	termTimeBytes.ReadFrom(output.Content)
	termTime, err := time.Parse("2006-01-02T15:04:05Z", termTimeBytes.String())
	if err != nil {
		return nil, true, fmt.Errorf("invalid time received from \"%s\": %w", spotITNPath, err)
	}
	return &termTime, true, nil
}

//TODO(bwagner5): Make this work
// func (i IMDS) GetMaintenanceEvent(ctx context.Context) (bool, error) {
// 	output, err := i.client.GetMetadata(ctx, &imds.GetMetadataInput{Path: scheduledEvents})
// 	if err != nil {
// 		return false, fmt.Errorf("IMDS Failed to get \"%s\": %w", scheduledEvents, err)
// 	}
// 	return true, nil
// }
