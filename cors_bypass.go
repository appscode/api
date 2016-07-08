package api

import (
	alert "github.com/appscode/api/alert/v0.1"
	artifactory "github.com/appscode/api/artifactory/v0.1"
	billing "github.com/appscode/api/billing/v0.1"
	bucket "github.com/appscode/api/bucket/v0.1"
	certificate "github.com/appscode/api/certificate/v0.1"
	ci "github.com/appscode/api/ci/v0.1"
	credential "github.com/appscode/api/credential/v0.1"
	db "github.com/appscode/api/db/v0.1"
	glusterfs "github.com/appscode/api/glusterfs/v0.1"
	kubernetes "github.com/appscode/api/kubernetes/v0.1"
	loadbalancer "github.com/appscode/api/loadbalancer/v0.1"
	mailinglist "github.com/appscode/api/mailinglist/v0.1"
	namespace "github.com/appscode/api/namespace/v0.1"
	pv "github.com/appscode/api/pv/v0.1"
	"github.com/gengo/grpc-gateway/runtime"
	backup "github.com/appscode/api/backup/v0.1"
)

// This is a hackish method to add support javascript
// CORS request to the grpc-gateway.
// TODO: Make those auto generated.
func Patterens() []runtime.Pattern {
	ps := make([]runtime.Pattern, 0)

	ps = append(ps, alert.Patterns()...)
	ps = append(ps, credential.Patterns()...)
	ps = append(ps, namespace.Patterns()...)
	ps = append(ps, kubernetes.Patterns()...)
	ps = append(ps, db.Patterns()...)
	ps = append(ps, ci.Patterns()...)
	ps = append(ps, pv.Patterns()...)
	ps = append(ps, glusterfs.Patterns()...)
	ps = append(ps, bucket.Patterns()...)
	ps = append(ps, certificate.Patterns()...)
	ps = append(ps, mailinglist.Patterns()...)
	ps = append(ps, loadbalancer.Patterns()...)
	ps = append(ps, artifactory.Patterns()...)
	ps = append(ps, billing.Patterns()...)
	ps = append(ps, backup.Patterns()...)
	return ps
}
