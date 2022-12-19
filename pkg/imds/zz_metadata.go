/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package imds

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"fmt"
	"strconv"
)

type metadata struct {
	ramDiskID                       string   `imds:"path=meta-data/ramdisk-id"`
	reservationID                   string   `imds:"path=meta-data/reservation-id"`
	securityGroups                  []string `imds:"path=meta-data/security-groups"`
	availabilityZone                string   `imds:"path=meta-data/placement/availability-zone"`
	availabilityZoneID              string   `imds:"path=meta-data/placement/availability-zone-id"`
	groupName                       string   `imds:"path=meta-data/placement/group-name"`
	hostID                          string   `imds:"path=meta-data/placement/host-id"`
	partitionNumber                 int      `imds:"path=meta-data/placement/partition-number"`
	region                          string   `imds:"path=meta-data/placement/region"`
	productCodes                    []string `imds:"path=meta-data/product-codes"`
	publicHostname                  string   `imds:"path=meta-data/public-hostname"`
	publicIPv4                      string   `imds:"path=meta-data/public-ipv4"`
	localHostname                   string   `imds:"path=meta-data/local-hostname"`
	localIPv4                       string   `imds:"path=meta-data/local-ipv4"`
	mac                             string   `imds:"path=meta-data/mac"`
	instanceAction                  string   `imds:"path=meta-data/instance-action"`
	instanceID                      string   `imds:"path=meta-data/instance-id"`
	instanceLifecycle               string   `imds:"path=meta-data/instance-life-cycle"`
	instanceType                    string   `imds:"path=meta-data/instance-type"`
	kernelID                        string   `imds:"path=meta-data/kernel-id"`
	amiID                           string   `imds:"path=meta-data/ami-id"`
	amiLaunchIndex                  int      `imds:"path=meta-data/ami-launch-index"`
	amiManifestPath                 string   `imds:"path=meta-data/ami-manifest-path"`
	ancestorAMIIDs                  []string `imds:"path=meta-data/ancestor-ami-ids"`
	autoscalingTargetLifecycleState string   `imds:"path=meta-data/autoscaling/target-lifecycle-state"`
	blockDeviceMappingAMI           string   `imds:"path=meta-data/block-device-mapping/ami"`
	blockDeviceMappingRoot          []string `imds:"path=meta-data/block-device-mapping/root"`
	eventsMaintenanceHistory        string   `imds:"path=meta-data/events/maintenance/history"`
	eventsMaintenanceScheduled      string   `imds:"path=meta-data/events/maintenance/scheduled"`
	eventsRecommendationsRebalance  string   `imds:"path=meta-data/events/recommendations/rebalance"`
	iamInfo                         string   `imds:"path=meta-data/iam/info"`
}

func (i Client) MustGetRamDiskIDWithContext(ctx context.Context) string {
	ramDiskID, err := i.GetMetadata(ctx, "ramdisk-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch ramDiskID: %v", err))
	}
	return ramDiskID
}

func (i Client) MustGetRamDiskID() string {
	ctx := context.Background()
	ramDiskID, err := i.GetMetadata(ctx, "ramdisk-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch ramDiskID: %v", err))
	}
	return ramDiskID
}

func (i Client) GetRamDiskIDWithContext(ctx context.Context) (string, error) {
	ramDiskID, err := i.GetMetadata(ctx, "ramdisk-id")
	if err != nil {
		return "", err
	}
	return ramDiskID, nil
}

func (i Client) GetRamDiskID() (string, error) {
	ctx := context.Background()
	ramDiskID, err := i.GetMetadata(ctx, "ramdisk-id")
	if err != nil {
		return "", err
	}
	return ramDiskID, nil
}

func (i Client) MustGetReservationIDWithContext(ctx context.Context) string {
	reservationID, err := i.GetMetadata(ctx, "reservation-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch reservationID: %v", err))
	}
	return reservationID
}

func (i Client) MustGetReservationID() string {
	ctx := context.Background()
	reservationID, err := i.GetMetadata(ctx, "reservation-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch reservationID: %v", err))
	}
	return reservationID
}

func (i Client) GetReservationIDWithContext(ctx context.Context) (string, error) {
	reservationID, err := i.GetMetadata(ctx, "reservation-id")
	if err != nil {
		return "", err
	}
	return reservationID, nil
}

func (i Client) GetReservationID() (string, error) {
	ctx := context.Background()
	reservationID, err := i.GetMetadata(ctx, "reservation-id")
	if err != nil {
		return "", err
	}
	return reservationID, nil
}

func (i Client) MustGetAvailabilityZoneWithContext(ctx context.Context) string {
	availabilityZone, err := i.GetMetadata(ctx, "placement/availability-zone")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch availabilityZone: %v", err))
	}
	return availabilityZone
}

func (i Client) MustGetAvailabilityZone() string {
	ctx := context.Background()
	availabilityZone, err := i.GetMetadata(ctx, "placement/availability-zone")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch availabilityZone: %v", err))
	}
	return availabilityZone
}

func (i Client) GetAvailabilityZoneWithContext(ctx context.Context) (string, error) {
	availabilityZone, err := i.GetMetadata(ctx, "placement/availability-zone")
	if err != nil {
		return "", err
	}
	return availabilityZone, nil
}

func (i Client) GetAvailabilityZone() (string, error) {
	ctx := context.Background()
	availabilityZone, err := i.GetMetadata(ctx, "placement/availability-zone")
	if err != nil {
		return "", err
	}
	return availabilityZone, nil
}

func (i Client) MustGetAvailabilityZoneIDWithContext(ctx context.Context) string {
	availabilityZoneID, err := i.GetMetadata(ctx, "placement/availability-zone-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch availabilityZoneID: %v", err))
	}
	return availabilityZoneID
}

func (i Client) MustGetAvailabilityZoneID() string {
	ctx := context.Background()
	availabilityZoneID, err := i.GetMetadata(ctx, "placement/availability-zone-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch availabilityZoneID: %v", err))
	}
	return availabilityZoneID
}

func (i Client) GetAvailabilityZoneIDWithContext(ctx context.Context) (string, error) {
	availabilityZoneID, err := i.GetMetadata(ctx, "placement/availability-zone-id")
	if err != nil {
		return "", err
	}
	return availabilityZoneID, nil
}

func (i Client) GetAvailabilityZoneID() (string, error) {
	ctx := context.Background()
	availabilityZoneID, err := i.GetMetadata(ctx, "placement/availability-zone-id")
	if err != nil {
		return "", err
	}
	return availabilityZoneID, nil
}

func (i Client) MustGetGroupNameWithContext(ctx context.Context) string {
	groupName, err := i.GetMetadata(ctx, "placement/group-name")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch groupName: %v", err))
	}
	return groupName
}

func (i Client) MustGetGroupName() string {
	ctx := context.Background()
	groupName, err := i.GetMetadata(ctx, "placement/group-name")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch groupName: %v", err))
	}
	return groupName
}

func (i Client) GetGroupNameWithContext(ctx context.Context) (string, error) {
	groupName, err := i.GetMetadata(ctx, "placement/group-name")
	if err != nil {
		return "", err
	}
	return groupName, nil
}

func (i Client) GetGroupName() (string, error) {
	ctx := context.Background()
	groupName, err := i.GetMetadata(ctx, "placement/group-name")
	if err != nil {
		return "", err
	}
	return groupName, nil
}

func (i Client) MustGetHostIDWithContext(ctx context.Context) string {
	hostID, err := i.GetMetadata(ctx, "placement/host-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch hostID: %v", err))
	}
	return hostID
}

func (i Client) MustGetHostID() string {
	ctx := context.Background()
	hostID, err := i.GetMetadata(ctx, "placement/host-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch hostID: %v", err))
	}
	return hostID
}

func (i Client) GetHostIDWithContext(ctx context.Context) (string, error) {
	hostID, err := i.GetMetadata(ctx, "placement/host-id")
	if err != nil {
		return "", err
	}
	return hostID, nil
}

func (i Client) GetHostID() (string, error) {
	ctx := context.Background()
	hostID, err := i.GetMetadata(ctx, "placement/host-id")
	if err != nil {
		return "", err
	}
	return hostID, nil
}

func (i Client) MustGetPartitionNumberWithContext(ctx context.Context) int {
	partitionNumber, err := i.GetMetadata(ctx, "placement/partition-number")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch partitionNumber: %v", err))
	}
	partitionNumberNum, err := strconv.Atoi(partitionNumber)
	if err != nil {
		panic(fmt.Sprintf("unable to convert placement/partition-number of %s to integer: %v", partitionNumber, err))
	}
	return partitionNumberNum
}

func (i Client) MustGetPartitionNumber() int {
	ctx := context.Background()
	partitionNumber, err := i.GetMetadata(ctx, "placement/partition-number")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch partitionNumber: %v", err))
	}
	partitionNumberNum, err := strconv.Atoi(partitionNumber)
	if err != nil {
		panic(fmt.Sprintf("unable to convert placement/partition-number of %s to integer: %v", partitionNumber, err))
	}
	return partitionNumberNum
}

func (i Client) GetPartitionNumberWithContext(ctx context.Context) (int, error) {
	partitionNumber, err := i.GetMetadata(ctx, "placement/partition-number")
	if err != nil {
		return 0, err
	}
	partitionNumberNum, err := strconv.Atoi(partitionNumber)
	if err != nil {
		return 0, fmt.Errorf("unable to convert placement/partition-number of %s to integer: %w", partitionNumber, err)
	}
	return partitionNumberNum, nil
}

func (i Client) GetPartitionNumber() (int, error) {
	ctx := context.Background()
	partitionNumber, err := i.GetMetadata(ctx, "placement/partition-number")
	if err != nil {
		return 0, err
	}
	partitionNumberNum, err := strconv.Atoi(partitionNumber)
	if err != nil {
		return 0, fmt.Errorf("unable to convert placement/partition-number of %s to integer: %w", partitionNumber, err)
	}
	return partitionNumberNum, nil
}

func (i Client) MustGetRegionWithContext(ctx context.Context) string {
	region, err := i.GetMetadata(ctx, "placement/region")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch region: %v", err))
	}
	return region
}

func (i Client) MustGetRegion() string {
	ctx := context.Background()
	region, err := i.GetMetadata(ctx, "placement/region")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch region: %v", err))
	}
	return region
}

func (i Client) GetRegionWithContext(ctx context.Context) (string, error) {
	region, err := i.GetMetadata(ctx, "placement/region")
	if err != nil {
		return "", err
	}
	return region, nil
}

func (i Client) GetRegion() (string, error) {
	ctx := context.Background()
	region, err := i.GetMetadata(ctx, "placement/region")
	if err != nil {
		return "", err
	}
	return region, nil
}

func (i Client) MustGetPublicHostnameWithContext(ctx context.Context) string {
	publicHostname, err := i.GetMetadata(ctx, "public-hostname")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch publicHostname: %v", err))
	}
	return publicHostname
}

func (i Client) MustGetPublicHostname() string {
	ctx := context.Background()
	publicHostname, err := i.GetMetadata(ctx, "public-hostname")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch publicHostname: %v", err))
	}
	return publicHostname
}

func (i Client) GetPublicHostnameWithContext(ctx context.Context) (string, error) {
	publicHostname, err := i.GetMetadata(ctx, "public-hostname")
	if err != nil {
		return "", err
	}
	return publicHostname, nil
}

func (i Client) GetPublicHostname() (string, error) {
	ctx := context.Background()
	publicHostname, err := i.GetMetadata(ctx, "public-hostname")
	if err != nil {
		return "", err
	}
	return publicHostname, nil
}

func (i Client) MustGetPublicIPv4WithContext(ctx context.Context) string {
	publicIPv4, err := i.GetMetadata(ctx, "public-ipv4")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch publicIPv4: %v", err))
	}
	return publicIPv4
}

func (i Client) MustGetPublicIPv4() string {
	ctx := context.Background()
	publicIPv4, err := i.GetMetadata(ctx, "public-ipv4")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch publicIPv4: %v", err))
	}
	return publicIPv4
}

func (i Client) GetPublicIPv4WithContext(ctx context.Context) (string, error) {
	publicIPv4, err := i.GetMetadata(ctx, "public-ipv4")
	if err != nil {
		return "", err
	}
	return publicIPv4, nil
}

func (i Client) GetPublicIPv4() (string, error) {
	ctx := context.Background()
	publicIPv4, err := i.GetMetadata(ctx, "public-ipv4")
	if err != nil {
		return "", err
	}
	return publicIPv4, nil
}

func (i Client) MustGetLocalHostnameWithContext(ctx context.Context) string {
	localHostname, err := i.GetMetadata(ctx, "local-hostname")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch localHostname: %v", err))
	}
	return localHostname
}

func (i Client) MustGetLocalHostname() string {
	ctx := context.Background()
	localHostname, err := i.GetMetadata(ctx, "local-hostname")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch localHostname: %v", err))
	}
	return localHostname
}

func (i Client) GetLocalHostnameWithContext(ctx context.Context) (string, error) {
	localHostname, err := i.GetMetadata(ctx, "local-hostname")
	if err != nil {
		return "", err
	}
	return localHostname, nil
}

func (i Client) GetLocalHostname() (string, error) {
	ctx := context.Background()
	localHostname, err := i.GetMetadata(ctx, "local-hostname")
	if err != nil {
		return "", err
	}
	return localHostname, nil
}

func (i Client) MustGetLocalIPv4WithContext(ctx context.Context) string {
	localIPv4, err := i.GetMetadata(ctx, "local-ipv4")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch localIPv4: %v", err))
	}
	return localIPv4
}

func (i Client) MustGetLocalIPv4() string {
	ctx := context.Background()
	localIPv4, err := i.GetMetadata(ctx, "local-ipv4")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch localIPv4: %v", err))
	}
	return localIPv4
}

func (i Client) GetLocalIPv4WithContext(ctx context.Context) (string, error) {
	localIPv4, err := i.GetMetadata(ctx, "local-ipv4")
	if err != nil {
		return "", err
	}
	return localIPv4, nil
}

func (i Client) GetLocalIPv4() (string, error) {
	ctx := context.Background()
	localIPv4, err := i.GetMetadata(ctx, "local-ipv4")
	if err != nil {
		return "", err
	}
	return localIPv4, nil
}

func (i Client) MustGetMacWithContext(ctx context.Context) string {
	mac, err := i.GetMetadata(ctx, "mac")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch mac: %v", err))
	}
	return mac
}

func (i Client) MustGetMac() string {
	ctx := context.Background()
	mac, err := i.GetMetadata(ctx, "mac")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch mac: %v", err))
	}
	return mac
}

func (i Client) GetMacWithContext(ctx context.Context) (string, error) {
	mac, err := i.GetMetadata(ctx, "mac")
	if err != nil {
		return "", err
	}
	return mac, nil
}

func (i Client) GetMac() (string, error) {
	ctx := context.Background()
	mac, err := i.GetMetadata(ctx, "mac")
	if err != nil {
		return "", err
	}
	return mac, nil
}

func (i Client) MustGetInstanceActionWithContext(ctx context.Context) string {
	instanceAction, err := i.GetMetadata(ctx, "instance-action")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceAction: %v", err))
	}
	return instanceAction
}

func (i Client) MustGetInstanceAction() string {
	ctx := context.Background()
	instanceAction, err := i.GetMetadata(ctx, "instance-action")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceAction: %v", err))
	}
	return instanceAction
}

func (i Client) GetInstanceActionWithContext(ctx context.Context) (string, error) {
	instanceAction, err := i.GetMetadata(ctx, "instance-action")
	if err != nil {
		return "", err
	}
	return instanceAction, nil
}

func (i Client) GetInstanceAction() (string, error) {
	ctx := context.Background()
	instanceAction, err := i.GetMetadata(ctx, "instance-action")
	if err != nil {
		return "", err
	}
	return instanceAction, nil
}

func (i Client) MustGetInstanceIDWithContext(ctx context.Context) string {
	instanceID, err := i.GetMetadata(ctx, "instance-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceID: %v", err))
	}
	return instanceID
}

func (i Client) MustGetInstanceID() string {
	ctx := context.Background()
	instanceID, err := i.GetMetadata(ctx, "instance-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceID: %v", err))
	}
	return instanceID
}

func (i Client) GetInstanceIDWithContext(ctx context.Context) (string, error) {
	instanceID, err := i.GetMetadata(ctx, "instance-id")
	if err != nil {
		return "", err
	}
	return instanceID, nil
}

func (i Client) GetInstanceID() (string, error) {
	ctx := context.Background()
	instanceID, err := i.GetMetadata(ctx, "instance-id")
	if err != nil {
		return "", err
	}
	return instanceID, nil
}

func (i Client) MustGetInstanceLifecycleWithContext(ctx context.Context) string {
	instanceLifecycle, err := i.GetMetadata(ctx, "instance-life-cycle")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceLifecycle: %v", err))
	}
	return instanceLifecycle
}

func (i Client) MustGetInstanceLifecycle() string {
	ctx := context.Background()
	instanceLifecycle, err := i.GetMetadata(ctx, "instance-life-cycle")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceLifecycle: %v", err))
	}
	return instanceLifecycle
}

func (i Client) GetInstanceLifecycleWithContext(ctx context.Context) (string, error) {
	instanceLifecycle, err := i.GetMetadata(ctx, "instance-life-cycle")
	if err != nil {
		return "", err
	}
	return instanceLifecycle, nil
}

func (i Client) GetInstanceLifecycle() (string, error) {
	ctx := context.Background()
	instanceLifecycle, err := i.GetMetadata(ctx, "instance-life-cycle")
	if err != nil {
		return "", err
	}
	return instanceLifecycle, nil
}

func (i Client) MustGetInstanceTypeWithContext(ctx context.Context) string {
	instanceType, err := i.GetMetadata(ctx, "instance-type")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceType: %v", err))
	}
	return instanceType
}

func (i Client) MustGetInstanceType() string {
	ctx := context.Background()
	instanceType, err := i.GetMetadata(ctx, "instance-type")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceType: %v", err))
	}
	return instanceType
}

func (i Client) GetInstanceTypeWithContext(ctx context.Context) (string, error) {
	instanceType, err := i.GetMetadata(ctx, "instance-type")
	if err != nil {
		return "", err
	}
	return instanceType, nil
}

func (i Client) GetInstanceType() (string, error) {
	ctx := context.Background()
	instanceType, err := i.GetMetadata(ctx, "instance-type")
	if err != nil {
		return "", err
	}
	return instanceType, nil
}

func (i Client) MustGetKernelIDWithContext(ctx context.Context) string {
	kernelID, err := i.GetMetadata(ctx, "kernel-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch kernelID: %v", err))
	}
	return kernelID
}

func (i Client) MustGetKernelID() string {
	ctx := context.Background()
	kernelID, err := i.GetMetadata(ctx, "kernel-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch kernelID: %v", err))
	}
	return kernelID
}

func (i Client) GetKernelIDWithContext(ctx context.Context) (string, error) {
	kernelID, err := i.GetMetadata(ctx, "kernel-id")
	if err != nil {
		return "", err
	}
	return kernelID, nil
}

func (i Client) GetKernelID() (string, error) {
	ctx := context.Background()
	kernelID, err := i.GetMetadata(ctx, "kernel-id")
	if err != nil {
		return "", err
	}
	return kernelID, nil
}

func (i Client) MustGetAmiIDWithContext(ctx context.Context) string {
	amiID, err := i.GetMetadata(ctx, "ami-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiID: %v", err))
	}
	return amiID
}

func (i Client) MustGetAmiID() string {
	ctx := context.Background()
	amiID, err := i.GetMetadata(ctx, "ami-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiID: %v", err))
	}
	return amiID
}

func (i Client) GetAmiIDWithContext(ctx context.Context) (string, error) {
	amiID, err := i.GetMetadata(ctx, "ami-id")
	if err != nil {
		return "", err
	}
	return amiID, nil
}

func (i Client) GetAmiID() (string, error) {
	ctx := context.Background()
	amiID, err := i.GetMetadata(ctx, "ami-id")
	if err != nil {
		return "", err
	}
	return amiID, nil
}

func (i Client) MustGetAmiLaunchIndexWithContext(ctx context.Context) int {
	amiLaunchIndex, err := i.GetMetadata(ctx, "ami-launch-index")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiLaunchIndex: %v", err))
	}
	amiLaunchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		panic(fmt.Sprintf("unable to convert ami-launch-index of %s to integer: %v", amiLaunchIndex, err))
	}
	return amiLaunchIndexNum
}

func (i Client) MustGetAmiLaunchIndex() int {
	ctx := context.Background()
	amiLaunchIndex, err := i.GetMetadata(ctx, "ami-launch-index")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiLaunchIndex: %v", err))
	}
	amiLaunchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		panic(fmt.Sprintf("unable to convert ami-launch-index of %s to integer: %v", amiLaunchIndex, err))
	}
	return amiLaunchIndexNum
}

func (i Client) GetAmiLaunchIndexWithContext(ctx context.Context) (int, error) {
	amiLaunchIndex, err := i.GetMetadata(ctx, "ami-launch-index")
	if err != nil {
		return 0, err
	}
	amiLaunchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		return 0, fmt.Errorf("unable to convert ami-launch-index of %s to integer: %w", amiLaunchIndex, err)
	}
	return amiLaunchIndexNum, nil
}

func (i Client) GetAmiLaunchIndex() (int, error) {
	ctx := context.Background()
	amiLaunchIndex, err := i.GetMetadata(ctx, "ami-launch-index")
	if err != nil {
		return 0, err
	}
	amiLaunchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		return 0, fmt.Errorf("unable to convert ami-launch-index of %s to integer: %w", amiLaunchIndex, err)
	}
	return amiLaunchIndexNum, nil
}

func (i Client) MustGetAmiManifestPathWithContext(ctx context.Context) string {
	amiManifestPath, err := i.GetMetadata(ctx, "ami-manifest-path")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiManifestPath: %v", err))
	}
	return amiManifestPath
}

func (i Client) MustGetAmiManifestPath() string {
	ctx := context.Background()
	amiManifestPath, err := i.GetMetadata(ctx, "ami-manifest-path")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiManifestPath: %v", err))
	}
	return amiManifestPath
}

func (i Client) GetAmiManifestPathWithContext(ctx context.Context) (string, error) {
	amiManifestPath, err := i.GetMetadata(ctx, "ami-manifest-path")
	if err != nil {
		return "", err
	}
	return amiManifestPath, nil
}

func (i Client) GetAmiManifestPath() (string, error) {
	ctx := context.Background()
	amiManifestPath, err := i.GetMetadata(ctx, "ami-manifest-path")
	if err != nil {
		return "", err
	}
	return amiManifestPath, nil
}

func (i Client) MustGetAutoscalingTargetLifecycleStateWithContext(ctx context.Context) string {
	autoscalingTargetLifecycleState, err := i.GetMetadata(ctx, "autoscaling/target-lifecycle-state")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch autoscalingTargetLifecycleState: %v", err))
	}
	return autoscalingTargetLifecycleState
}

func (i Client) MustGetAutoscalingTargetLifecycleState() string {
	ctx := context.Background()
	autoscalingTargetLifecycleState, err := i.GetMetadata(ctx, "autoscaling/target-lifecycle-state")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch autoscalingTargetLifecycleState: %v", err))
	}
	return autoscalingTargetLifecycleState
}

func (i Client) GetAutoscalingTargetLifecycleStateWithContext(ctx context.Context) (string, error) {
	autoscalingTargetLifecycleState, err := i.GetMetadata(ctx, "autoscaling/target-lifecycle-state")
	if err != nil {
		return "", err
	}
	return autoscalingTargetLifecycleState, nil
}

func (i Client) GetAutoscalingTargetLifecycleState() (string, error) {
	ctx := context.Background()
	autoscalingTargetLifecycleState, err := i.GetMetadata(ctx, "autoscaling/target-lifecycle-state")
	if err != nil {
		return "", err
	}
	return autoscalingTargetLifecycleState, nil
}

func (i Client) MustGetBlockDeviceMappingAMIWithContext(ctx context.Context) string {
	blockDeviceMappingAMI, err := i.GetMetadata(ctx, "block-device-mapping/ami")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch blockDeviceMappingAMI: %v", err))
	}
	return blockDeviceMappingAMI
}

func (i Client) MustGetBlockDeviceMappingAMI() string {
	ctx := context.Background()
	blockDeviceMappingAMI, err := i.GetMetadata(ctx, "block-device-mapping/ami")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch blockDeviceMappingAMI: %v", err))
	}
	return blockDeviceMappingAMI
}

func (i Client) GetBlockDeviceMappingAMIWithContext(ctx context.Context) (string, error) {
	blockDeviceMappingAMI, err := i.GetMetadata(ctx, "block-device-mapping/ami")
	if err != nil {
		return "", err
	}
	return blockDeviceMappingAMI, nil
}

func (i Client) GetBlockDeviceMappingAMI() (string, error) {
	ctx := context.Background()
	blockDeviceMappingAMI, err := i.GetMetadata(ctx, "block-device-mapping/ami")
	if err != nil {
		return "", err
	}
	return blockDeviceMappingAMI, nil
}

func (i Client) MustGetEventsMaintenanceHistoryWithContext(ctx context.Context) string {
	eventsMaintenanceHistory, err := i.GetMetadata(ctx, "events/maintenance/history")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsMaintenanceHistory: %v", err))
	}
	return eventsMaintenanceHistory
}

func (i Client) MustGetEventsMaintenanceHistory() string {
	ctx := context.Background()
	eventsMaintenanceHistory, err := i.GetMetadata(ctx, "events/maintenance/history")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsMaintenanceHistory: %v", err))
	}
	return eventsMaintenanceHistory
}

func (i Client) GetEventsMaintenanceHistoryWithContext(ctx context.Context) (string, error) {
	eventsMaintenanceHistory, err := i.GetMetadata(ctx, "events/maintenance/history")
	if err != nil {
		return "", err
	}
	return eventsMaintenanceHistory, nil
}

func (i Client) GetEventsMaintenanceHistory() (string, error) {
	ctx := context.Background()
	eventsMaintenanceHistory, err := i.GetMetadata(ctx, "events/maintenance/history")
	if err != nil {
		return "", err
	}
	return eventsMaintenanceHistory, nil
}

func (i Client) MustGetEventsMaintenanceScheduledWithContext(ctx context.Context) string {
	eventsMaintenanceScheduled, err := i.GetMetadata(ctx, "events/maintenance/scheduled")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsMaintenanceScheduled: %v", err))
	}
	return eventsMaintenanceScheduled
}

func (i Client) MustGetEventsMaintenanceScheduled() string {
	ctx := context.Background()
	eventsMaintenanceScheduled, err := i.GetMetadata(ctx, "events/maintenance/scheduled")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsMaintenanceScheduled: %v", err))
	}
	return eventsMaintenanceScheduled
}

func (i Client) GetEventsMaintenanceScheduledWithContext(ctx context.Context) (string, error) {
	eventsMaintenanceScheduled, err := i.GetMetadata(ctx, "events/maintenance/scheduled")
	if err != nil {
		return "", err
	}
	return eventsMaintenanceScheduled, nil
}

func (i Client) GetEventsMaintenanceScheduled() (string, error) {
	ctx := context.Background()
	eventsMaintenanceScheduled, err := i.GetMetadata(ctx, "events/maintenance/scheduled")
	if err != nil {
		return "", err
	}
	return eventsMaintenanceScheduled, nil
}

func (i Client) MustGetEventsRecommendationsRebalanceWithContext(ctx context.Context) string {
	eventsRecommendationsRebalance, err := i.GetMetadata(ctx, "events/recommendations/rebalance")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsRecommendationsRebalance: %v", err))
	}
	return eventsRecommendationsRebalance
}

func (i Client) MustGetEventsRecommendationsRebalance() string {
	ctx := context.Background()
	eventsRecommendationsRebalance, err := i.GetMetadata(ctx, "events/recommendations/rebalance")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsRecommendationsRebalance: %v", err))
	}
	return eventsRecommendationsRebalance
}

func (i Client) GetEventsRecommendationsRebalanceWithContext(ctx context.Context) (string, error) {
	eventsRecommendationsRebalance, err := i.GetMetadata(ctx, "events/recommendations/rebalance")
	if err != nil {
		return "", err
	}
	return eventsRecommendationsRebalance, nil
}

func (i Client) GetEventsRecommendationsRebalance() (string, error) {
	ctx := context.Background()
	eventsRecommendationsRebalance, err := i.GetMetadata(ctx, "events/recommendations/rebalance")
	if err != nil {
		return "", err
	}
	return eventsRecommendationsRebalance, nil
}

func (i Client) MustGetIamInfoWithContext(ctx context.Context) string {
	iamInfo, err := i.GetMetadata(ctx, "iam/info")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch iamInfo: %v", err))
	}
	return iamInfo
}

func (i Client) MustGetIamInfo() string {
	ctx := context.Background()
	iamInfo, err := i.GetMetadata(ctx, "iam/info")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch iamInfo: %v", err))
	}
	return iamInfo
}

func (i Client) GetIamInfoWithContext(ctx context.Context) (string, error) {
	iamInfo, err := i.GetMetadata(ctx, "iam/info")
	if err != nil {
		return "", err
	}
	return iamInfo, nil
}

func (i Client) GetIamInfo() (string, error) {
	ctx := context.Background()
	iamInfo, err := i.GetMetadata(ctx, "iam/info")
	if err != nil {
		return "", err
	}
	return iamInfo, nil
}
