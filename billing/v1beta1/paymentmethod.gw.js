// Code generated by protoc-gen-grpc-js-client
// source: paymentmethod.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function paymentMethodsCheck(p, conf) {
    path = '/billing/v1beta1/paymentmethods'
    return xhr(path, 'GET', conf, p);
}

var services = {
    paymentMethods: {
        check: paymentMethodsCheck
    }
};

module.exports = {billing: {v1beta1: services}};
