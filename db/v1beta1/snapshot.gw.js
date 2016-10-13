// Code generated by protoc-gen-grpc-js-client
// source: snapshot.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function snapshotsList(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/snapshots'
    delete p['cluster']
    delete p['uid']
    return xhr(path, 'GET', conf, p);
}

function snapshotsBackupSchedule(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/actions/schedule-backup'
    delete p['cluster']
    delete p['uid']
    return xhr(path, 'POST', conf, null, p);
}

function snapshotsBackupUnschedule(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/actions/unschedule-backup'
    delete p['cluster']
    delete p['uid']
    return xhr(path, 'PUT', conf, p);
}

function snapshotsRestore(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/actions/restore'
    delete p['cluster']
    delete p['uid']
    return xhr(path, 'POST', conf, null, p);
}

var services = {
    snapshots: {
        list: snapshotsList,
        backupSchedule: snapshotsBackupSchedule,
        backupUnschedule: snapshotsBackupUnschedule,
        restore: snapshotsRestore
    }
};

module.exports = {db: {v1beta1: services}};
