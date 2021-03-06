package role

import "github.com/kirk-enterprise/openstack-golang-sdk/lib/ifaces"

type ProjectGroupRole struct {
	*AbstractRole

	_ bool
}

func NewProjectGroupRole(client ifaces.Openstacker) *ProjectGroupRole {
	return &ProjectGroupRole{
		AbstractRole: NewAbstractRole(client, ProjectUrl, GroupUrl),
	}
}
