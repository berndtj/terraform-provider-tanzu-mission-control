---
Title: "Cluster Group Resource"
Description: |-
    Creating the cluster group resource to a cluster.
---

# Cluster Group

Group clusters together into a cluster group for more efficient management using this Terraform module.

A cluster group is a logical grouping of clusters in the Tanzu Mission Control object hierarchy that provides for better management of your Kubernetes clusters.

To create a cluster group, you must have `organization.edit` permissions in Tanzu Mission Control.
For more information, please refer [create a Cluster Group.][cluster-group]

[cluster-group]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-E36D5205-5291-4201-809D-B8E26277CCF4.html

## Example Usage

```terraform
# Create Tanzu Mission Control cluster group
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "tf-cluster-group"
  meta {
    description = "Create cluster group through terraform"
    labels = {
      "key1" : "value1",
      "key2" : "value2"
    }
  }
}

# Create cluster group with minimal information
resource "tanzu-mission-control_cluster_group" "create_cluster_group_min_info" {
  name = "tf-cluster-group-min-info"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String)

### Optional

- **id** (String) The ID of this resource.
- **meta** (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- **description** (String) Description of the resource
- **labels** (Map of String) Labels for the resource

Read-Only:

- **annotations** (Map of String) Annotations for the resource
- **resource_version** (String) Resource version of the resource
- **uid** (String) UID of the resource