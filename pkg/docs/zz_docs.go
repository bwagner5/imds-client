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
package doc

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
type InstanceMetadataCategory struct {
	Category    string
	Description string
	Version     string
}

type DynamicCategory struct {
	Category    string
	Description string
	Version     string
}

var InstanceMetadataCategoryEntries = []InstanceMetadataCategory{
	{
		Category:    "ami-id",
		Description: "The AMI ID used to launch the instance.",
		Version:     "1.0",
	},
	{
		Category:    "ami-launch-index",
		Description: "If you started more than one instance at the same time, this value indicates the order in which the instance was launched. The value of the first instance launched is 0.",
		Version:     "1.0",
	},
	{
		Category:    "ami-manifest-path",
		Description: "The path to the AMI manifest file in Amazon S3. If you used an Amazon EBS-backed AMI to launch the instance, the returned result is unknown.",
		Version:     "1.0",
	},
	{
		Category:    "ancestor-ami-ids",
		Description: "The AMI IDs of any instances that were rebundled to create this AMI. This value will only exist if the AMI manifest file contained an ancestor-amis key.",
		Version:     "2007-10-10",
	},
	{
		Category:    "autoscaling/target-lifecycle-state",
		Description: "Value showing the target Auto Scaling lifecycle state that an Auto Scaling instance is transitioning to. Present when the instance transitions to one of the target lifecycle states after March 10, 2022. Possible values: `Detached` | `InService` | `Standby` | `Terminated` | `Warmed:Hibernated` | `Warmed:Running` | `Warmed:Stopped` | `Warmed:Terminated`. See [Retrieve the target lifecycle state through instance metadata](https://docs.aws.amazon.com/autoscaling/ec2/userguide/retrieving-target-lifecycle-state-through-imds.html) in the *Amazon EC2 Auto Scaling User Guide*.",
		Version:     "2021-07-15",
	},
	{
		Category:    "block-device-mapping/ami",
		Description: "The virtual device that contains the root/boot file system.",
		Version:     "2007-12-15",
	},
	{
		Category:    "block-device-mapping/ebsN",
		Description: "The virtual devices associated with any Amazon EBS volumes. Amazon EBS volumes are only available in metadata if they were present at launch time or when the instance was last started. The N indicates the index of the Amazon EBS volume (such as ebs1 or ebs2).",
		Version:     "2007-12-15",
	},
	{
		Category:    "block-device-mapping/ephemeralN",
		Description: "The virtual devices for any non-NVMe instance store volumes. The N indicates the index of each volume. The number of instance store volumes in the block device mapping might not match the actual number of instance store volumes for the instance. The instance type determines the number of instance store volumes that are available to an instance. If the number of instance store volumes in a block device mapping exceeds the number available to an instance, the additional instance store volumes are ignored.",
		Version:     "2007-12-15",
	},
	{
		Category:    "block-device-mapping/root",
		Description: "The virtual devices or partitions associated with the root devices or partitions on the virtual device, where the root (/ or C:) file system is associated with the given instance.",
		Version:     "2007-12-15",
	},
	{
		Category:    "block-device-mapping/swap",
		Description: "The virtual devices associated with swap. Not always present.",
		Version:     "2007-12-15",
	},
	{
		Category:    "elastic-gpus/associations/elastic-gpu-id",
		Description: "If there is an Elastic GPU attached to the instance, contains a JSON string with information about the Elastic GPU, including its ID and connection information.",
		Version:     "2016-11-30",
	},
	{
		Category:    "elastic-inference/associations/eia-id",
		Description: "If there is an Elastic Inference accelerator attached to the instance, contains a JSON string with information about the Elastic Inference accelerator, including its ID and type.",
		Version:     "2018-11-29",
	},
	{
		Category:    "events/maintenance/history",
		Description: "If there are completed or canceled maintenance events for the instance, contains a JSON string with information about the events. For more information, see [To view event history about completed or canceled events](monitoring-instances-status-check_sched.md#viewing-event-history).",
		Version:     "2018-08-17",
	},
	{
		Category:    "events/maintenance/scheduled",
		Description: "If there are active maintenance events for the instance, contains a JSON string with information about the events. For more information, see [View scheduled events](monitoring-instances-status-check_sched.md#viewing_scheduled_events).",
		Version:     "2018-08-17",
	},
	{
		Category:    "events/recommendations/rebalance",
		Description: "The approximate time, in UTC, when the EC2 instance rebalance recommendation notification is emitted for the instance. The following is an example of the metadata for this category: {\"noticeTime\": \"2020-11-05T08:22:00Z\"}. This category is available only after the notification is emitted. For more information, see [EC2 instance rebalance recommendations](rebalance-recommendations.md).",
		Version:     "2020-10-27",
	},
	{
		Category:    "hostname",
		Description: "If the EC2 instance is using IP-based naming (IPBN), this is the private IPv4 DNS hostname of the instance. If the EC2 instance is using Resource-based naming (RBN), this is the RBN. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0). For more information about IPBN and RBN, see [Amazon EC2 instance hostname types](ec2-instance-naming.md).",
		Version:     "1.0",
	},
	{
		Category:    "iam/info",
		Description: "If there is an IAM role associated with the instance, contains information about the last time the instance profile was updated, including the instance's LastUpdated date, InstanceProfileArn, and InstanceProfileId. Otherwise, not present.",
		Version:     "2012-01-12",
	},
	{
		Category:    "iam/security-credentials/role-name",
		Description: "If there is an IAM role associated with the instance, role-name is the name of the role, and role-name contains the temporary security credentials associated with the role (for more information, see [Retrieve security credentials from instance metadata](iam-roles-for-amazon-ec2.md#instance-metadata-security-credentials)). Otherwise, not present.",
		Version:     "2012-01-12",
	},
	{
		Category:    "identity-credentials/ec2/info",
		Description: "[Internal use only] Information about the credentials in identity-credentials/ec2/security-credentials/ec2-instance. These credentials are used by AWS features such as EC2 Instance Connect, and do not have any additional AWS API permissions or privileges beyond identifying the instance.",
		Version:     "2018-05-23",
	},
	{
		Category:    "identity-credentials/ec2/security-credentials/ec2-instance",
		Description: "[Internal use only] Credentials that allow on-instance software to identify itself to AWS to support features such as EC2 Instance Connect. These credentials do not have any additional AWS API permissions or privileges.",
		Version:     "2018-05-23",
	},
	{
		Category:    "instance-action",
		Description: "Notifies the instance that it should reboot in preparation for bundling. Valid values: none | shutdown | bundle-pending.",
		Version:     "2008-09-01",
	},
	{
		Category:    "instance-id",
		Description: "The ID of this instance.",
		Version:     "1.0",
	},
	{
		Category:    "instance-life-cycle",
		Description: "The purchasing option of this instance. For more information, see [Instance purchasing options](instance-purchasing-options.md).",
		Version:     "2019-10-01",
	},
	{
		Category:    "instance-type",
		Description: "The type of instance. For more information, see [Instance types](instance-types.md).",
		Version:     "2007-08-29",
	},
	{
		Category:    "ipv6",
		Description: "The IPv6 address of the instance. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0) network interface and the first IPv6 address assigned. If no IPv6 address exists on network interface[0], this item is not set and results in an HTTP 404 response.",
		Version:     "2021-01-03",
	},
	{
		Category:    "kernel-id",
		Description: "The ID of the kernel launched with this instance, if applicable.",
		Version:     "2008-02-01",
	},
	{
		Category:    "local-hostname",
		Description: "In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0). If the EC2 instance is using IP-based naming (IPBN), this is the private IPv4 DNS hostname of the instance. If the EC2 instance is using Resource-based naming (RBN), this is the RBN. For more information about IPBN, RBN, and EC2 instance naming, see [Amazon EC2 instance hostname types](ec2-instance-naming.md).",
		Version:     "2007-01-19",
	},
	{
		Category:    "local-ipv4",
		Description: "The private IPv4 address of the instance. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0). If this is an IPv6-only instance, this item is not set and results in an HTTP 404 response.",
		Version:     "1.0",
	},
	{
		Category:    "mac",
		Description: "The instance's media access control (MAC) address. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0).",
		Version:     "2011-01-01",
	},
	{
		Category:    "metrics/vhostmd",
		Description: "No longer available.",
		Version:     "2011-05-01",
	},
	{
		Category:    "network/interfaces/macs/mac/device-number",
		Description: "The unique device number associated with that interface. The device number corresponds to the device name; for example, a device-number of 2 is for the eth2 device. This category corresponds to the DeviceIndex and device-index fields that are used by the Amazon EC2 API and the EC2 commands for the AWS CLI.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/interface-id",
		Description: "The ID of the network interface.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/ipv4-associations/public-ip",
		Description: "The private IPv4 addresses that are associated with each public IP address and assigned to that interface.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/ipv6s",
		Description: "The IPv6 addresses associated with the interface. Returned only for instances launched into a VPC.",
		Version:     "2016-06-30",
	},
	{
		Category:    "network/interfaces/macs/mac/local-hostname",
		Description: "The private IPv4 DNS hostname of the instance. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0). If this is a IPv6-only instance, this is the resource-based name. For more information about IPBN and RBN, see [Amazon EC2 instance hostname types](ec2-instance-naming.md).",
		Version:     "2007-01-19",
	},
	{
		Category:    "network/interfaces/macs/mac/local-ipv4s",
		Description: "The private IPv4 addresses associated with the interface. If this is an IPv6-only network interface, this item is not set and results in an HTTP 404 response.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/mac",
		Description: "The instance's MAC address.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/network-card-index",
		Description: "The index of the network card. Some instance types support multiple network cards.",
		Version:     "2020-11-01",
	},
	{
		Category:    "network/interfaces/macs/mac/owner-id",
		Description: "The ID of the owner of the network interface. In multiple-interface environments, an interface can be attached by a third party, such as Elastic Load Balancing. Traffic on an interface is always billed to the interface owner.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/public-hostname",
		Description: "The interface's public DNS (IPv4). This category is only returned if the enableDnsHostnames attribute is set to true. For more information, see [Using DNS with Your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html) in the Amazon VPC User Guide. If the instance only has a public-IPv6 address and no public-IPv4 address, this item is not set and results in an HTTP 404 response.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/public-ipv4s",
		Description: "The public IP address or Elastic IP addresses associated with the interface. There may be multiple IPv4 addresses on an instance.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/security-groups",
		Description: "Security groups to which the network interface belongs.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/security-group-ids",
		Description: "The IDs of the security groups to which the network interface belongs.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/subnet-id",
		Description: "The ID of the subnet in which the interface resides.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/subnet-ipv4-cidr-block",
		Description: "The IPv4 CIDR block of the subnet in which the interface resides.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/subnet-ipv6-cidr-blocks",
		Description: "The IPv6 CIDR block of the subnet in which the interface resides.",
		Version:     "2016-06-30",
	},
	{
		Category:    "network/interfaces/macs/mac/vpc-id",
		Description: "The ID of the VPC in which the interface resides.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/vpc-ipv4-cidr-block",
		Description: "The primary IPv4 CIDR block of the VPC.",
		Version:     "2011-01-01",
	},
	{
		Category:    "network/interfaces/macs/mac/vpc-ipv4-cidr-blocks",
		Description: "The IPv4 CIDR blocks for the VPC.",
		Version:     "2016-06-30",
	},
	{
		Category:    "network/interfaces/macs/mac/vpc-ipv6-cidr-blocks",
		Description: "The IPv6 CIDR block of the VPC in which the interface resides.",
		Version:     "2016-06-30",
	},
	{
		Category:    "placement/availability-zone",
		Description: "The Availability Zone in which the instance launched.",
		Version:     "2008-02-01",
	},
	{
		Category:    "placement/availability-zone-id",
		Description: "The static Availability Zone ID in which the instance is launched. The Availability Zone ID is consistent across accounts. However, it might be different from the Availability Zone, which can vary by account.",
		Version:     "2019-10-01",
	},
	{
		Category:    "placement/group-name",
		Description: "The name of the placement group in which the instance is launched.",
		Version:     "2020-08-24",
	},
	{
		Category:    "placement/host-id",
		Description: "The ID of the host on which the instance is launched. Applicable only to Dedicated Hosts.",
		Version:     "2020-08-24",
	},
	{
		Category:    "placement/partition-number",
		Description: "The number of the partition in which the instance is launched.",
		Version:     "2020-08-24",
	},
	{
		Category:    "placement/region",
		Description: "The AWS Region in which the instance is launched.",
		Version:     "2020-08-24",
	},
	{
		Category:    "product-codes",
		Description: "AWS Marketplace product codes associated with the instance, if any.",
		Version:     "2007-03-01",
	},
	{
		Category:    "public-hostname",
		Description: "The instance's public DNS (IPv4). This category is only returned if the enableDnsHostnames attribute is set to true. For more information, see [Using DNS with Your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html) in the Amazon VPC User Guide. If the instance only has a public-IPv6 address and no public-IPv4 address, this item is not set and results in an HTTP 404 response.",
		Version:     "2007-01-19",
	},
	{
		Category:    "public-ipv4",
		Description: "The public IPv4 address. If an Elastic IP address is associated with the instance, the value returned is the Elastic IP address.",
		Version:     "2007-01-19",
	},
	{
		Category:    "public-keys/0/openssh-key",
		Description: "Public key. Only available if supplied at instance launch time.",
		Version:     "1.0",
	},
	{
		Category:    "ramdisk-id",
		Description: "The ID of the RAM disk specified at launch time, if applicable.",
		Version:     "2007-10-10",
	},
	{
		Category:    "reservation-id",
		Description: "The ID of the reservation.",
		Version:     "1.0",
	},
	{
		Category:    "security-groups",
		Description: "The names of the security groups applied to the instance. After launch, you can change the security groups of the instances. Such changes are reflected here and in network/interfaces/macs/**mac**/security-groups.",
		Version:     "1.0",
	},
	{
		Category:    "services/domain",
		Description: "The domain for AWS resources for the Region.",
		Version:     "2014-02-25",
	},
	{
		Category:    "services/partition",
		Description: "The partition that the resource is in. For standard AWS Regions, the partition is `aws`. If you have resources in other partitions, the partition is `aws-partitionname`. For example, the partition for resources in the China (Beijing) Region is `aws-cn`.",
		Version:     "2015-10-20",
	},
	{
		Category:    "spot/instance-action",
		Description: "The action (hibernate, stop, or terminate) and the approximate time, in UTC, when the action will occur. This item is present only if the Spot Instance has been marked for hibernate, stop, or terminate. For more information, see [instance-action](spot-instance-termination-notices.md#instance-action-metadata).",
		Version:     "2016-11-15",
	},
	{
		Category:    "spot/termination-time",
		Description: "The approximate time, in UTC, that the operating system for your Spot Instance will receive the shutdown signal. This item is present and contains a time value (for example, 2015-01-05T18:02:00Z) only if the Spot Instance has been marked for termination by Amazon EC2. The termination-time item is not set to a time if you terminated the Spot Instance yourself. For more information, see [termination-time](spot-instance-termination-notices.md#termination-time-metadata).",
		Version:     "2014-11-05",
	},
	{
		Category:    "tags/instance",
		Description: "The instance tags associated with the instance. Only available if you explicitly allow access to tags in instance metadata. For more information, see [Allow access to tags in instance metadata](Using_Tags.md#allow-access-to-tags-in-IMDS).",
		Version:     "2021-03-23",
	},
}

var DynamicCategoryEntries = []DynamicCategory{
	{
		Category:    "fws/instance-monitoring",
		Description: "Value showing whether the customer has enabled detailed one-minute monitoring in CloudWatch. Valid values: enabled | disabled",
		Version:     "2009-04-04",
	},
	{
		Category:    "instance-identity/document",
		Description: "JSON containing instance attributes, such as instance-id, private IP address, etc. See [Instance identity documents](instance-identity-documents.md).",
		Version:     "2009-04-04",
	},
	{
		Category:    "instance-identity/pkcs7",
		Description: "Used to verify the document's authenticity and content against the signature. See [Instance identity documents](instance-identity-documents.md).",
		Version:     "2009-04-04",
	},
	{
		Category:    "instance-identity/signature",
		Description: "Data that can be used by other parties to verify its origin and authenticity. See [Instance identity documents](instance-identity-documents.md).",
		Version:     "2009-04-04",
	},
}

