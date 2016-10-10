// Code generated by protoc-gen-grpc-js-client
// source: cluster.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function ClustersList(p, conf) {
	path = '/kubernetes/v1beta1/clusters'
	return xhr(path, 'GET', conf, p);
}

function ClustersDescribe(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['uid']
	delete p['uid']
	return xhr(path, 'GET', conf, p);
}

function ClustersCreate(p, conf) {
	path = '/kubernetes/v1beta1/clusters'
	return xhr(path, 'POST', conf, null, p);
}

function ClustersUpdate(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['name']
	delete p['name']
	return xhr(path, 'PUT', conf, null, p);
}

function ClustersScale(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['name'] + '/actions/scale'
	delete p['name']
	return xhr(path, 'PUT', conf, null, p);
}

function ClustersUpgrade(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['name'] + '/actions/upgrade'
	delete p['name']
	return xhr(path, 'PUT', conf, null, p);
}

function ClustersDelete(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['name']
	delete p['name']
	return xhr(path, 'DELETE', conf, p);
}

function ClustersClientConfig(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['name'] + '/client-config'
	delete p['name']
	return xhr(path, 'GET', conf, p);
}

function ClustersInstances(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster_name'] + '/instances'
	delete p['cluster_name']
	return xhr(path, 'GET', conf, p);
}

function ClustersStartupScript(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['uid'] + '/startup-script/' + p['role']
	delete p['uid']
	delete p['role']
	return xhr(path, 'GET', conf, p);
}

function ClustersInstanceByIP(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['phid'] + '/instance-by-ip/' + p['external_ip']
	delete p['phid']
	delete p['external_ip']
	return xhr(path, 'GET', conf, p);
}

module.exports = {
  Clusters: {
      List: ClustersList,
      Describe: ClustersDescribe,
      Create: ClustersCreate,
      Update: ClustersUpdate,
      Scale: ClustersScale,
      Upgrade: ClustersUpgrade,
      Delete: ClustersDelete,
      ClientConfig: ClustersClientConfig,
      Instances: ClustersInstances,
      StartupScript: ClustersStartupScript,
      InstanceByIP: ClustersInstanceByIP
  }
};