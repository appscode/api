package v1alpha2

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var certificateDeployRequestSchema *gojsonschema.Schema
var certificateDescribeRequestSchema *gojsonschema.Schema
var certificateLoadRequestSchema *gojsonschema.Schema
var certificateListRequestSchema *gojsonschema.Schema
var certificateDeleteRequestSchema *gojsonschema.Schema

func init() {
	var err error
	certificateDeployRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    },
    "owner": {
      "type": "string"
    },
    "secret_name": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	certificateDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "owner": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	certificateLoadRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cert_data": {
      "type": "string"
    },
    "key_data": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "owner": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	certificateListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "owner": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	certificateDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "owner": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *CertificateDeployRequest) Valid() (*gojsonschema.Result, error) {
	return certificateDeployRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDeployRequest) IsRequest() {}

func (m *CertificateDescribeRequest) Valid() (*gojsonschema.Result, error) {
	return certificateDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDescribeRequest) IsRequest() {}

func (m *CertificateLoadRequest) Valid() (*gojsonschema.Result, error) {
	return certificateLoadRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateLoadRequest) IsRequest() {}

func (m *CertificateListRequest) Valid() (*gojsonschema.Result, error) {
	return certificateListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateListRequest) IsRequest() {}

func (m *CertificateDeleteRequest) Valid() (*gojsonschema.Result, error) {
	return certificateDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDeleteRequest) IsRequest() {}

