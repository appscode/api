// Code generated by protoc-gen-grpc-js-client
// source: cluster.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function clustersList(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['kube_cluster'] + '/glusterfs'
    delete p['kube_cluster']
    return xhr(path, 'GET', conf, p);
}

function clustersDescribe(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['kube_cluster'] + '/namespaces/' + p['kube_namespace'] + '/glusterfs/' + p['name']
    delete p['kube_cluster']
    delete p['kube_namespace']
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

function clustersCreate(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['kube_cluster'] + '/namespaces/' + p['kube_namespace'] + '/glusterfs'
    delete p['kube_cluster']
    delete p['kube_namespace']
    return xhr(path, 'POST', conf, null, p);
}

function clustersDelete(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['kube_cluster'] + '/namespaces/' + p['kube_namespace'] + '/glusterfs/' + p['name']
    delete p['kube_cluster']
    delete p['kube_namespace']
    delete p['name']
    return xhr(path, 'DELETE', conf, p);
}

var services = {
    clusters: {
        list: clustersList,
        describe: clustersDescribe,
        create: clustersCreate,
        delete: clustersDelete
    }
};

module.exports = {glusterfs: {v1beta1: services}};
