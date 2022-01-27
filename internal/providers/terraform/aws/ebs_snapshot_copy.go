package aws

import (
	"github.com/infracost/infracost/internal/resources/aws"
	"github.com/infracost/infracost/internal/schema"
)

func getEBSSnapshotCopyRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_ebs_snapshot_copy",
		RFunc: NewEbsSnapshotCopy,
		ReferenceAttributes: []string{
			"volume_id",
			"source_snapshot_id",
		},
	}
}
func NewEbsSnapshotCopy(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.EbsSnapshotCopy{Address: d.Address, Region: d.Get("region").String()}
	sourceSnapshotRefs := d.References("source_snapshot_id")
	if len(sourceSnapshotRefs) > 0 {
		volumeRefs := sourceSnapshotRefs[0].References("volume_id")
		if len(volumeRefs) > 0 {
			if volumeRefs[0].Get("size").Exists() {
				r.VolumeRefSize = volumeRefs[0].Get("size").Float()
			}
		}
	}
	r.PopulateUsage(u)
	return r.BuildResource()
}
