// Code generated by protoc-gen-grpc-js-client
// source: cluster.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function clustersList(p, conf) {
    path = '/kubernetes/v1beta1/clusters'
    return xhr(path, 'GET', conf, p);
}

function clustersDescribe(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['uid']
    delete p['uid']
    return xhr(path, 'GET', conf, p);
}

function clustersCreate(p, conf) {
    path = '/kubernetes/v1beta1/clusters'
    return xhr(path, 'POST', conf, null, p);
}

function clustersUpdate(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['name']
    delete p['name']
    return xhr(path, 'PUT', conf, null, p);
}

function clustersScale(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['name'] + '/actions/scale'
    delete p['name']
    return xhr(path, 'PUT', conf, null, p);
}

function clustersUpgrade(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['name'] + '/actions/upgrade'
    delete p['name']
    return xhr(path, 'PUT', conf, null, p);
}

function clustersDelete(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['name']
    delete p['name']
    return xhr(path, 'DELETE', conf, p);
}

function clustersClientConfig(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['name'] + '/client-config'
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

function clustersInstances(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster_name'] + '/instances'
    delete p['cluster_name']
    return xhr(path, 'GET', conf, p);
}

function clustersStartupScript(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['uid'] + '/startup-script/' + p['role']
    delete p['uid']
    delete p['role']
    return xhr(path, 'GET', conf, p);
}

function clustersInstanceByIP(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['phid'] + '/instance-by-ip/' + p['external_ip']
    delete p['phid']
    delete p['external_ip']
    return xhr(path, 'GET', conf, p);
}

var services = {
    clusters: {
        list: clustersList,
        describe: clustersDescribe,
        create: clustersCreate,
        update: clustersUpdate,
        scale: clustersScale,
        upgrade: clustersUpgrade,
        delete: clustersDelete,
        clientConfig: clustersClientConfig,
        instances: clustersInstances,
        startupScript: clustersStartupScript,
        instanceByIP: clustersInstanceByIP
    }
};

module.exports = {kubernetes: {v1beta1: services}};
