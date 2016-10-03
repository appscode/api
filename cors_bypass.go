package api

import (
	alert "github.com/appscode/api/alert/v1beta1"
	artifactory "github.com/appscode/api/artifactory/v1beta1"
	backup "github.com/appscode/api/backup/v1beta1"
	billing "github.com/appscode/api/billing/v1beta1"
	bucket "github.com/appscode/api/bucket/v1beta1"
	certificate "github.com/appscode/api/certificate/v1beta1"
	ci "github.com/appscode/api/ci/v1beta1"
	credential "github.com/appscode/api/credential/v1beta1"
	db "github.com/appscode/api/db/v0.1"
	glusterfs "github.com/appscode/api/glusterfs/v1beta1"
	kubernetes "github.com/appscode/api/kubernetes/v1beta1"
	kubernetes_v1beta2 "github.com/appscode/api/kubernetes/v1beta2"
	loadbalancer "github.com/appscode/api/loadbalancer/v1beta1"
	mailinglist "github.com/appscode/api/mailinglist/v1beta1"
	namespace "github.com/appscode/api/namespace/v1beta1"
	pv "github.com/appscode/api/pv/v1beta1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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
	ps = append(ps, kubernetes_v1beta2.Patterns()...)
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
