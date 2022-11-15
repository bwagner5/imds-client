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

func (i IMDS) MustGetRamDiskIDWithContext(ctx context.Context) string {
	ramDiskID, err := i.GetMetadata(ctx, "meta-data/ramdisk-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch ramDiskID: %v", err))
	}
	return ramDiskID
}

func (i IMDS) MustGetRamDiskID() string {
	ctx := context.Background()
	ramDiskID, err := i.GetMetadata(ctx, "meta-data/ramdisk-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch ramDiskID: %v", err))
	}
	return ramDiskID
}

func (i IMDS) GetRamDiskIDWithContext(ctx context.Context) (string, error) {
	ramDiskID, err := i.GetMetadata(ctx, "meta-data/ramdisk-id")
	if err != nil {
		return "", err
	}
	return ramDiskID, nil
}

func (i IMDS) GetRamDiskID() (string, error) {
	ctx := context.Background()
	ramDiskID, err := i.GetMetadata(ctx, "meta-data/ramdisk-id")
	if err != nil {
		return "", err
	}
	return ramDiskID, nil
}

func (i IMDS) MustGetReservationIDWithContext(ctx context.Context) string {
	reservationID, err := i.GetMetadata(ctx, "meta-data/reservation-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch reservationID: %v", err))
	}
	return reservationID
}

func (i IMDS) MustGetReservationID() string {
	ctx := context.Background()
	reservationID, err := i.GetMetadata(ctx, "meta-data/reservation-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch reservationID: %v", err))
	}
	return reservationID
}

func (i IMDS) GetReservationIDWithContext(ctx context.Context) (string, error) {
	reservationID, err := i.GetMetadata(ctx, "meta-data/reservation-id")
	if err != nil {
		return "", err
	}
	return reservationID, nil
}

func (i IMDS) GetReservationID() (string, error) {
	ctx := context.Background()
	reservationID, err := i.GetMetadata(ctx, "meta-data/reservation-id")
	if err != nil {
		return "", err
	}
	return reservationID, nil
}

func (i IMDS) MustGetAvailabilityZoneWithContext(ctx context.Context) string {
	availabilityZone, err := i.GetMetadata(ctx, "meta-data/placement/availability-zone")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch availabilityZone: %v", err))
	}
	return availabilityZone
}

func (i IMDS) MustGetAvailabilityZone() string {
	ctx := context.Background()
	availabilityZone, err := i.GetMetadata(ctx, "meta-data/placement/availability-zone")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch availabilityZone: %v", err))
	}
	return availabilityZone
}

func (i IMDS) GetAvailabilityZoneWithContext(ctx context.Context) (string, error) {
	availabilityZone, err := i.GetMetadata(ctx, "meta-data/placement/availability-zone")
	if err != nil {
		return "", err
	}
	return availabilityZone, nil
}

func (i IMDS) GetAvailabilityZone() (string, error) {
	ctx := context.Background()
	availabilityZone, err := i.GetMetadata(ctx, "meta-data/placement/availability-zone")
	if err != nil {
		return "", err
	}
	return availabilityZone, nil
}

func (i IMDS) MustGetAvailabilityZoneIDWithContext(ctx context.Context) string {
	availabilityZoneID, err := i.GetMetadata(ctx, "meta-data/placement/availability-zone-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch availabilityZoneID: %v", err))
	}
	return availabilityZoneID
}

func (i IMDS) MustGetAvailabilityZoneID() string {
	ctx := context.Background()
	availabilityZoneID, err := i.GetMetadata(ctx, "meta-data/placement/availability-zone-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch availabilityZoneID: %v", err))
	}
	return availabilityZoneID
}

func (i IMDS) GetAvailabilityZoneIDWithContext(ctx context.Context) (string, error) {
	availabilityZoneID, err := i.GetMetadata(ctx, "meta-data/placement/availability-zone-id")
	if err != nil {
		return "", err
	}
	return availabilityZoneID, nil
}

func (i IMDS) GetAvailabilityZoneID() (string, error) {
	ctx := context.Background()
	availabilityZoneID, err := i.GetMetadata(ctx, "meta-data/placement/availability-zone-id")
	if err != nil {
		return "", err
	}
	return availabilityZoneID, nil
}

func (i IMDS) MustGetGroupNameWithContext(ctx context.Context) string {
	groupName, err := i.GetMetadata(ctx, "meta-data/placement/group-name")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch groupName: %v", err))
	}
	return groupName
}

func (i IMDS) MustGetGroupName() string {
	ctx := context.Background()
	groupName, err := i.GetMetadata(ctx, "meta-data/placement/group-name")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch groupName: %v", err))
	}
	return groupName
}

func (i IMDS) GetGroupNameWithContext(ctx context.Context) (string, error) {
	groupName, err := i.GetMetadata(ctx, "meta-data/placement/group-name")
	if err != nil {
		return "", err
	}
	return groupName, nil
}

func (i IMDS) GetGroupName() (string, error) {
	ctx := context.Background()
	groupName, err := i.GetMetadata(ctx, "meta-data/placement/group-name")
	if err != nil {
		return "", err
	}
	return groupName, nil
}

func (i IMDS) MustGetHostIDWithContext(ctx context.Context) string {
	hostID, err := i.GetMetadata(ctx, "meta-data/placement/host-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch hostID: %v", err))
	}
	return hostID
}

func (i IMDS) MustGetHostID() string {
	ctx := context.Background()
	hostID, err := i.GetMetadata(ctx, "meta-data/placement/host-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch hostID: %v", err))
	}
	return hostID
}

func (i IMDS) GetHostIDWithContext(ctx context.Context) (string, error) {
	hostID, err := i.GetMetadata(ctx, "meta-data/placement/host-id")
	if err != nil {
		return "", err
	}
	return hostID, nil
}

func (i IMDS) GetHostID() (string, error) {
	ctx := context.Background()
	hostID, err := i.GetMetadata(ctx, "meta-data/placement/host-id")
	if err != nil {
		return "", err
	}
	return hostID, nil
}

func (i IMDS) MustGetPartitionNumberWithContext(ctx context.Context) int {
	partitionNumber, err := i.GetMetadata(ctx, "meta-data/placement/partition-number")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch partitionNumber: %v", err))
	}
	partitionNumberNum, err := strconv.Atoi(partitionNumber)
	if err != nil {
		panic(fmt.Sprintf("unable to convert meta-data/placement/partition-number of %s to integer: %v", partitionNumber, err))
	}
	return partitionNumberNum
}

func (i IMDS) MustGetPartitionNumber() int {
	ctx := context.Background()
	partitionNumber, err := i.GetMetadata(ctx, "meta-data/placement/partition-number")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch partitionNumber: %v", err))
	}
	partitionNumberNum, err := strconv.Atoi(partitionNumber)
	if err != nil {
		panic(fmt.Sprintf("unable to convert meta-data/placement/partition-number of %s to integer: %v", partitionNumber, err))
	}
	return partitionNumberNum
}

func (i IMDS) GetPartitionNumberWithContext(ctx context.Context) (int, error) {
	partitionNumber, err := i.GetMetadata(ctx, "meta-data/placement/partition-number")
	if err != nil {
		return 0, err
	}
	partitionNumberNum, err := strconv.Atoi(partitionNumber)
	if err != nil {
		return 0, fmt.Errorf("unable to convert meta-data/placement/partition-number of %s to integer: %w", partitionNumber, err)
	}
	return partitionNumberNum, nil
}

func (i IMDS) GetPartitionNumber() (int, error) {
	ctx := context.Background()
	partitionNumber, err := i.GetMetadata(ctx, "meta-data/placement/partition-number")
	if err != nil {
		return 0, err
	}
	partitionNumberNum, err := strconv.Atoi(partitionNumber)
	if err != nil {
		return 0, fmt.Errorf("unable to convert meta-data/placement/partition-number of %s to integer: %w", partitionNumber, err)
	}
	return partitionNumberNum, nil
}

func (i IMDS) MustGetRegionWithContext(ctx context.Context) string {
	region, err := i.GetMetadata(ctx, "meta-data/placement/region")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch region: %v", err))
	}
	return region
}

func (i IMDS) MustGetRegion() string {
	ctx := context.Background()
	region, err := i.GetMetadata(ctx, "meta-data/placement/region")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch region: %v", err))
	}
	return region
}

func (i IMDS) GetRegionWithContext(ctx context.Context) (string, error) {
	region, err := i.GetMetadata(ctx, "meta-data/placement/region")
	if err != nil {
		return "", err
	}
	return region, nil
}

func (i IMDS) GetRegion() (string, error) {
	ctx := context.Background()
	region, err := i.GetMetadata(ctx, "meta-data/placement/region")
	if err != nil {
		return "", err
	}
	return region, nil
}

func (i IMDS) MustGetPublicHostnameWithContext(ctx context.Context) string {
	publicHostname, err := i.GetMetadata(ctx, "meta-data/public-hostname")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch publicHostname: %v", err))
	}
	return publicHostname
}

func (i IMDS) MustGetPublicHostname() string {
	ctx := context.Background()
	publicHostname, err := i.GetMetadata(ctx, "meta-data/public-hostname")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch publicHostname: %v", err))
	}
	return publicHostname
}

func (i IMDS) GetPublicHostnameWithContext(ctx context.Context) (string, error) {
	publicHostname, err := i.GetMetadata(ctx, "meta-data/public-hostname")
	if err != nil {
		return "", err
	}
	return publicHostname, nil
}

func (i IMDS) GetPublicHostname() (string, error) {
	ctx := context.Background()
	publicHostname, err := i.GetMetadata(ctx, "meta-data/public-hostname")
	if err != nil {
		return "", err
	}
	return publicHostname, nil
}

func (i IMDS) MustGetPublicIPv4WithContext(ctx context.Context) string {
	publicIPv4, err := i.GetMetadata(ctx, "meta-data/public-ipv4")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch publicIPv4: %v", err))
	}
	return publicIPv4
}

func (i IMDS) MustGetPublicIPv4() string {
	ctx := context.Background()
	publicIPv4, err := i.GetMetadata(ctx, "meta-data/public-ipv4")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch publicIPv4: %v", err))
	}
	return publicIPv4
}

func (i IMDS) GetPublicIPv4WithContext(ctx context.Context) (string, error) {
	publicIPv4, err := i.GetMetadata(ctx, "meta-data/public-ipv4")
	if err != nil {
		return "", err
	}
	return publicIPv4, nil
}

func (i IMDS) GetPublicIPv4() (string, error) {
	ctx := context.Background()
	publicIPv4, err := i.GetMetadata(ctx, "meta-data/public-ipv4")
	if err != nil {
		return "", err
	}
	return publicIPv4, nil
}

func (i IMDS) MustGetLocalHostnameWithContext(ctx context.Context) string {
	localHostname, err := i.GetMetadata(ctx, "meta-data/local-hostname")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch localHostname: %v", err))
	}
	return localHostname
}

func (i IMDS) MustGetLocalHostname() string {
	ctx := context.Background()
	localHostname, err := i.GetMetadata(ctx, "meta-data/local-hostname")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch localHostname: %v", err))
	}
	return localHostname
}

func (i IMDS) GetLocalHostnameWithContext(ctx context.Context) (string, error) {
	localHostname, err := i.GetMetadata(ctx, "meta-data/local-hostname")
	if err != nil {
		return "", err
	}
	return localHostname, nil
}

func (i IMDS) GetLocalHostname() (string, error) {
	ctx := context.Background()
	localHostname, err := i.GetMetadata(ctx, "meta-data/local-hostname")
	if err != nil {
		return "", err
	}
	return localHostname, nil
}

func (i IMDS) MustGetLocalIPv4WithContext(ctx context.Context) string {
	localIPv4, err := i.GetMetadata(ctx, "meta-data/local-ipv4")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch localIPv4: %v", err))
	}
	return localIPv4
}

func (i IMDS) MustGetLocalIPv4() string {
	ctx := context.Background()
	localIPv4, err := i.GetMetadata(ctx, "meta-data/local-ipv4")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch localIPv4: %v", err))
	}
	return localIPv4
}

func (i IMDS) GetLocalIPv4WithContext(ctx context.Context) (string, error) {
	localIPv4, err := i.GetMetadata(ctx, "meta-data/local-ipv4")
	if err != nil {
		return "", err
	}
	return localIPv4, nil
}

func (i IMDS) GetLocalIPv4() (string, error) {
	ctx := context.Background()
	localIPv4, err := i.GetMetadata(ctx, "meta-data/local-ipv4")
	if err != nil {
		return "", err
	}
	return localIPv4, nil
}

func (i IMDS) MustGetMacWithContext(ctx context.Context) string {
	mac, err := i.GetMetadata(ctx, "meta-data/mac")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch mac: %v", err))
	}
	return mac
}

func (i IMDS) MustGetMac() string {
	ctx := context.Background()
	mac, err := i.GetMetadata(ctx, "meta-data/mac")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch mac: %v", err))
	}
	return mac
}

func (i IMDS) GetMacWithContext(ctx context.Context) (string, error) {
	mac, err := i.GetMetadata(ctx, "meta-data/mac")
	if err != nil {
		return "", err
	}
	return mac, nil
}

func (i IMDS) GetMac() (string, error) {
	ctx := context.Background()
	mac, err := i.GetMetadata(ctx, "meta-data/mac")
	if err != nil {
		return "", err
	}
	return mac, nil
}

func (i IMDS) MustGetInstanceActionWithContext(ctx context.Context) string {
	instanceAction, err := i.GetMetadata(ctx, "meta-data/instance-action")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceAction: %v", err))
	}
	return instanceAction
}

func (i IMDS) MustGetInstanceAction() string {
	ctx := context.Background()
	instanceAction, err := i.GetMetadata(ctx, "meta-data/instance-action")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceAction: %v", err))
	}
	return instanceAction
}

func (i IMDS) GetInstanceActionWithContext(ctx context.Context) (string, error) {
	instanceAction, err := i.GetMetadata(ctx, "meta-data/instance-action")
	if err != nil {
		return "", err
	}
	return instanceAction, nil
}

func (i IMDS) GetInstanceAction() (string, error) {
	ctx := context.Background()
	instanceAction, err := i.GetMetadata(ctx, "meta-data/instance-action")
	if err != nil {
		return "", err
	}
	return instanceAction, nil
}

func (i IMDS) MustGetInstanceIDWithContext(ctx context.Context) string {
	instanceID, err := i.GetMetadata(ctx, "meta-data/instance-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceID: %v", err))
	}
	return instanceID
}

func (i IMDS) MustGetInstanceID() string {
	ctx := context.Background()
	instanceID, err := i.GetMetadata(ctx, "meta-data/instance-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceID: %v", err))
	}
	return instanceID
}

func (i IMDS) GetInstanceIDWithContext(ctx context.Context) (string, error) {
	instanceID, err := i.GetMetadata(ctx, "meta-data/instance-id")
	if err != nil {
		return "", err
	}
	return instanceID, nil
}

func (i IMDS) GetInstanceID() (string, error) {
	ctx := context.Background()
	instanceID, err := i.GetMetadata(ctx, "meta-data/instance-id")
	if err != nil {
		return "", err
	}
	return instanceID, nil
}

func (i IMDS) MustGetInstanceLifecycleWithContext(ctx context.Context) string {
	instanceLifecycle, err := i.GetMetadata(ctx, "meta-data/instance-life-cycle")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceLifecycle: %v", err))
	}
	return instanceLifecycle
}

func (i IMDS) MustGetInstanceLifecycle() string {
	ctx := context.Background()
	instanceLifecycle, err := i.GetMetadata(ctx, "meta-data/instance-life-cycle")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceLifecycle: %v", err))
	}
	return instanceLifecycle
}

func (i IMDS) GetInstanceLifecycleWithContext(ctx context.Context) (string, error) {
	instanceLifecycle, err := i.GetMetadata(ctx, "meta-data/instance-life-cycle")
	if err != nil {
		return "", err
	}
	return instanceLifecycle, nil
}

func (i IMDS) GetInstanceLifecycle() (string, error) {
	ctx := context.Background()
	instanceLifecycle, err := i.GetMetadata(ctx, "meta-data/instance-life-cycle")
	if err != nil {
		return "", err
	}
	return instanceLifecycle, nil
}

func (i IMDS) MustGetInstanceTypeWithContext(ctx context.Context) string {
	instanceType, err := i.GetMetadata(ctx, "meta-data/instance-type")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceType: %v", err))
	}
	return instanceType
}

func (i IMDS) MustGetInstanceType() string {
	ctx := context.Background()
	instanceType, err := i.GetMetadata(ctx, "meta-data/instance-type")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch instanceType: %v", err))
	}
	return instanceType
}

func (i IMDS) GetInstanceTypeWithContext(ctx context.Context) (string, error) {
	instanceType, err := i.GetMetadata(ctx, "meta-data/instance-type")
	if err != nil {
		return "", err
	}
	return instanceType, nil
}

func (i IMDS) GetInstanceType() (string, error) {
	ctx := context.Background()
	instanceType, err := i.GetMetadata(ctx, "meta-data/instance-type")
	if err != nil {
		return "", err
	}
	return instanceType, nil
}

func (i IMDS) MustGetKernelIDWithContext(ctx context.Context) string {
	kernelID, err := i.GetMetadata(ctx, "meta-data/kernel-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch kernelID: %v", err))
	}
	return kernelID
}

func (i IMDS) MustGetKernelID() string {
	ctx := context.Background()
	kernelID, err := i.GetMetadata(ctx, "meta-data/kernel-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch kernelID: %v", err))
	}
	return kernelID
}

func (i IMDS) GetKernelIDWithContext(ctx context.Context) (string, error) {
	kernelID, err := i.GetMetadata(ctx, "meta-data/kernel-id")
	if err != nil {
		return "", err
	}
	return kernelID, nil
}

func (i IMDS) GetKernelID() (string, error) {
	ctx := context.Background()
	kernelID, err := i.GetMetadata(ctx, "meta-data/kernel-id")
	if err != nil {
		return "", err
	}
	return kernelID, nil
}

func (i IMDS) MustGetAmiIDWithContext(ctx context.Context) string {
	amiID, err := i.GetMetadata(ctx, "meta-data/ami-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiID: %v", err))
	}
	return amiID
}

func (i IMDS) MustGetAmiID() string {
	ctx := context.Background()
	amiID, err := i.GetMetadata(ctx, "meta-data/ami-id")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiID: %v", err))
	}
	return amiID
}

func (i IMDS) GetAmiIDWithContext(ctx context.Context) (string, error) {
	amiID, err := i.GetMetadata(ctx, "meta-data/ami-id")
	if err != nil {
		return "", err
	}
	return amiID, nil
}

func (i IMDS) GetAmiID() (string, error) {
	ctx := context.Background()
	amiID, err := i.GetMetadata(ctx, "meta-data/ami-id")
	if err != nil {
		return "", err
	}
	return amiID, nil
}

func (i IMDS) MustGetAmiLaunchIndexWithContext(ctx context.Context) int {
	amiLaunchIndex, err := i.GetMetadata(ctx, "meta-data/ami-launch-index")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiLaunchIndex: %v", err))
	}
	amiLaunchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		panic(fmt.Sprintf("unable to convert meta-data/ami-launch-index of %s to integer: %v", amiLaunchIndex, err))
	}
	return amiLaunchIndexNum
}

func (i IMDS) MustGetAmiLaunchIndex() int {
	ctx := context.Background()
	amiLaunchIndex, err := i.GetMetadata(ctx, "meta-data/ami-launch-index")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiLaunchIndex: %v", err))
	}
	amiLaunchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		panic(fmt.Sprintf("unable to convert meta-data/ami-launch-index of %s to integer: %v", amiLaunchIndex, err))
	}
	return amiLaunchIndexNum
}

func (i IMDS) GetAmiLaunchIndexWithContext(ctx context.Context) (int, error) {
	amiLaunchIndex, err := i.GetMetadata(ctx, "meta-data/ami-launch-index")
	if err != nil {
		return 0, err
	}
	amiLaunchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		return 0, fmt.Errorf("unable to convert meta-data/ami-launch-index of %s to integer: %w", amiLaunchIndex, err)
	}
	return amiLaunchIndexNum, nil
}

func (i IMDS) GetAmiLaunchIndex() (int, error) {
	ctx := context.Background()
	amiLaunchIndex, err := i.GetMetadata(ctx, "meta-data/ami-launch-index")
	if err != nil {
		return 0, err
	}
	amiLaunchIndexNum, err := strconv.Atoi(amiLaunchIndex)
	if err != nil {
		return 0, fmt.Errorf("unable to convert meta-data/ami-launch-index of %s to integer: %w", amiLaunchIndex, err)
	}
	return amiLaunchIndexNum, nil
}

func (i IMDS) MustGetAmiManifestPathWithContext(ctx context.Context) string {
	amiManifestPath, err := i.GetMetadata(ctx, "meta-data/ami-manifest-path")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiManifestPath: %v", err))
	}
	return amiManifestPath
}

func (i IMDS) MustGetAmiManifestPath() string {
	ctx := context.Background()
	amiManifestPath, err := i.GetMetadata(ctx, "meta-data/ami-manifest-path")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch amiManifestPath: %v", err))
	}
	return amiManifestPath
}

func (i IMDS) GetAmiManifestPathWithContext(ctx context.Context) (string, error) {
	amiManifestPath, err := i.GetMetadata(ctx, "meta-data/ami-manifest-path")
	if err != nil {
		return "", err
	}
	return amiManifestPath, nil
}

func (i IMDS) GetAmiManifestPath() (string, error) {
	ctx := context.Background()
	amiManifestPath, err := i.GetMetadata(ctx, "meta-data/ami-manifest-path")
	if err != nil {
		return "", err
	}
	return amiManifestPath, nil
}

func (i IMDS) MustGetAutoscalingTargetLifecycleStateWithContext(ctx context.Context) string {
	autoscalingTargetLifecycleState, err := i.GetMetadata(ctx, "meta-data/autoscaling/target-lifecycle-state")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch autoscalingTargetLifecycleState: %v", err))
	}
	return autoscalingTargetLifecycleState
}

func (i IMDS) MustGetAutoscalingTargetLifecycleState() string {
	ctx := context.Background()
	autoscalingTargetLifecycleState, err := i.GetMetadata(ctx, "meta-data/autoscaling/target-lifecycle-state")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch autoscalingTargetLifecycleState: %v", err))
	}
	return autoscalingTargetLifecycleState
}

func (i IMDS) GetAutoscalingTargetLifecycleStateWithContext(ctx context.Context) (string, error) {
	autoscalingTargetLifecycleState, err := i.GetMetadata(ctx, "meta-data/autoscaling/target-lifecycle-state")
	if err != nil {
		return "", err
	}
	return autoscalingTargetLifecycleState, nil
}

func (i IMDS) GetAutoscalingTargetLifecycleState() (string, error) {
	ctx := context.Background()
	autoscalingTargetLifecycleState, err := i.GetMetadata(ctx, "meta-data/autoscaling/target-lifecycle-state")
	if err != nil {
		return "", err
	}
	return autoscalingTargetLifecycleState, nil
}

func (i IMDS) MustGetBlockDeviceMappingAMIWithContext(ctx context.Context) string {
	blockDeviceMappingAMI, err := i.GetMetadata(ctx, "meta-data/block-device-mapping/ami")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch blockDeviceMappingAMI: %v", err))
	}
	return blockDeviceMappingAMI
}

func (i IMDS) MustGetBlockDeviceMappingAMI() string {
	ctx := context.Background()
	blockDeviceMappingAMI, err := i.GetMetadata(ctx, "meta-data/block-device-mapping/ami")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch blockDeviceMappingAMI: %v", err))
	}
	return blockDeviceMappingAMI
}

func (i IMDS) GetBlockDeviceMappingAMIWithContext(ctx context.Context) (string, error) {
	blockDeviceMappingAMI, err := i.GetMetadata(ctx, "meta-data/block-device-mapping/ami")
	if err != nil {
		return "", err
	}
	return blockDeviceMappingAMI, nil
}

func (i IMDS) GetBlockDeviceMappingAMI() (string, error) {
	ctx := context.Background()
	blockDeviceMappingAMI, err := i.GetMetadata(ctx, "meta-data/block-device-mapping/ami")
	if err != nil {
		return "", err
	}
	return blockDeviceMappingAMI, nil
}

func (i IMDS) MustGetEventsMaintenanceHistoryWithContext(ctx context.Context) string {
	eventsMaintenanceHistory, err := i.GetMetadata(ctx, "meta-data/events/maintenance/history")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsMaintenanceHistory: %v", err))
	}
	return eventsMaintenanceHistory
}

func (i IMDS) MustGetEventsMaintenanceHistory() string {
	ctx := context.Background()
	eventsMaintenanceHistory, err := i.GetMetadata(ctx, "meta-data/events/maintenance/history")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsMaintenanceHistory: %v", err))
	}
	return eventsMaintenanceHistory
}

func (i IMDS) GetEventsMaintenanceHistoryWithContext(ctx context.Context) (string, error) {
	eventsMaintenanceHistory, err := i.GetMetadata(ctx, "meta-data/events/maintenance/history")
	if err != nil {
		return "", err
	}
	return eventsMaintenanceHistory, nil
}

func (i IMDS) GetEventsMaintenanceHistory() (string, error) {
	ctx := context.Background()
	eventsMaintenanceHistory, err := i.GetMetadata(ctx, "meta-data/events/maintenance/history")
	if err != nil {
		return "", err
	}
	return eventsMaintenanceHistory, nil
}

func (i IMDS) MustGetEventsMaintenanceScheduledWithContext(ctx context.Context) string {
	eventsMaintenanceScheduled, err := i.GetMetadata(ctx, "meta-data/events/maintenance/scheduled")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsMaintenanceScheduled: %v", err))
	}
	return eventsMaintenanceScheduled
}

func (i IMDS) MustGetEventsMaintenanceScheduled() string {
	ctx := context.Background()
	eventsMaintenanceScheduled, err := i.GetMetadata(ctx, "meta-data/events/maintenance/scheduled")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsMaintenanceScheduled: %v", err))
	}
	return eventsMaintenanceScheduled
}

func (i IMDS) GetEventsMaintenanceScheduledWithContext(ctx context.Context) (string, error) {
	eventsMaintenanceScheduled, err := i.GetMetadata(ctx, "meta-data/events/maintenance/scheduled")
	if err != nil {
		return "", err
	}
	return eventsMaintenanceScheduled, nil
}

func (i IMDS) GetEventsMaintenanceScheduled() (string, error) {
	ctx := context.Background()
	eventsMaintenanceScheduled, err := i.GetMetadata(ctx, "meta-data/events/maintenance/scheduled")
	if err != nil {
		return "", err
	}
	return eventsMaintenanceScheduled, nil
}

func (i IMDS) MustGetEventsRecommendationsRebalanceWithContext(ctx context.Context) string {
	eventsRecommendationsRebalance, err := i.GetMetadata(ctx, "meta-data/events/recommendations/rebalance")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsRecommendationsRebalance: %v", err))
	}
	return eventsRecommendationsRebalance
}

func (i IMDS) MustGetEventsRecommendationsRebalance() string {
	ctx := context.Background()
	eventsRecommendationsRebalance, err := i.GetMetadata(ctx, "meta-data/events/recommendations/rebalance")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch eventsRecommendationsRebalance: %v", err))
	}
	return eventsRecommendationsRebalance
}

func (i IMDS) GetEventsRecommendationsRebalanceWithContext(ctx context.Context) (string, error) {
	eventsRecommendationsRebalance, err := i.GetMetadata(ctx, "meta-data/events/recommendations/rebalance")
	if err != nil {
		return "", err
	}
	return eventsRecommendationsRebalance, nil
}

func (i IMDS) GetEventsRecommendationsRebalance() (string, error) {
	ctx := context.Background()
	eventsRecommendationsRebalance, err := i.GetMetadata(ctx, "meta-data/events/recommendations/rebalance")
	if err != nil {
		return "", err
	}
	return eventsRecommendationsRebalance, nil
}

func (i IMDS) MustGetIamInfoWithContext(ctx context.Context) string {
	iamInfo, err := i.GetMetadata(ctx, "meta-data/iam/info")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch iamInfo: %v", err))
	}
	return iamInfo
}

func (i IMDS) MustGetIamInfo() string {
	ctx := context.Background()
	iamInfo, err := i.GetMetadata(ctx, "meta-data/iam/info")
	if err != nil {
		panic(fmt.Sprintf("unable to fetch iamInfo: %v", err))
	}
	return iamInfo
}

func (i IMDS) GetIamInfoWithContext(ctx context.Context) (string, error) {
	iamInfo, err := i.GetMetadata(ctx, "meta-data/iam/info")
	if err != nil {
		return "", err
	}
	return iamInfo, nil
}

func (i IMDS) GetIamInfo() (string, error) {
	ctx := context.Background()
	iamInfo, err := i.GetMetadata(ctx, "meta-data/iam/info")
	if err != nil {
		return "", err
	}
	return iamInfo, nil
}

