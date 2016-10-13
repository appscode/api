// Code generated by protoc-gen-grpc-js-client
// source: event.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function eventsConstructive(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster_name'] + '/actions/construct-event'
    delete p['cluster_name']
    return xhr(path, 'PUT', conf, null, p);
}

function eventsDestructive(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster_name'] + '/actions/destruct-event'
    delete p['cluster_name']
    return xhr(path, 'PUT', conf, null, p);
}

var services = {
    events: {
        constructive: eventsConstructive,
        destructive: eventsDestructive
    }
};

module.exports = {kubernetes: {v1beta1: services}};
