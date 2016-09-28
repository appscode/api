package v1beta2

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var secretUpdateRequestSchema *gojsonschema.Schema
var updateResourceRequestSchema *gojsonschema.Schema
var configMapUpdateRequestSchema *gojsonschema.Schema
var copyResourceRequestSchema *gojsonschema.Schema
var deleteResourceRequestSchema *gojsonschema.Schema
var listResourceRequestSchema *gojsonschema.Schema
var describeResourceRequestSchema *gojsonschema.Schema

func init() {
	var err error
	secretUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "add": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    },
    "cluster": {
      "type": "string"
    },
    "deleted": {
      "items": {
        "type": "string"
      },
      "type": "array"
    },
    "name": {
      "maxLength": 63,
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "update": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	updateResourceRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1beta2Raw": {
      "properties": {
        "data": {
          "type": "string"
        },
        "format": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "raw": {
      "$ref": "#/definitions/v1beta2Raw"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	configMapUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "add": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    },
    "cluster": {
      "type": "string"
    },
    "deleted": {
      "items": {
        "type": "string"
      },
      "type": "array"
    },
    "name": {
      "maxLength": 63,
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "update": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	copyResourceRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1beta2KubeObject": {
      "properties": {
        "cluster": {
          "type": "string"
        },
        "name": {
          "maxLength": 63,
          "type": "string"
        },
        "namespace": {
          "maxLength": 63,
          "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "destination": {
      "$ref": "#/definitions/v1beta2KubeObject"
    },
    "source": {
      "$ref": "#/definitions/v1beta2KubeObject"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	deleteResourceRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "type": "string"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	listResourceRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	describeResourceRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "raw": {
      "type": "string"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *SecretUpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return secretUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SecretUpdateRequest) IsRequest() {}

func (m *UpdateResourceRequest) IsValid() (*gojsonschema.Result, error) {
	return updateResourceRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UpdateResourceRequest) IsRequest() {}

func (m *ConfigMapUpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return configMapUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ConfigMapUpdateRequest) IsRequest() {}

func (m *CopyResourceRequest) IsValid() (*gojsonschema.Result, error) {
	return copyResourceRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CopyResourceRequest) IsRequest() {}

func (m *DeleteResourceRequest) IsValid() (*gojsonschema.Result, error) {
	return deleteResourceRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DeleteResourceRequest) IsRequest() {}

func (m *ListResourceRequest) IsValid() (*gojsonschema.Result, error) {
	return listResourceRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ListResourceRequest) IsRequest() {}

func (m *DescribeResourceRequest) IsValid() (*gojsonschema.Result, error) {
	return describeResourceRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DescribeResourceRequest) IsRequest() {}

func (m *DescribeResourceResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *ClientActionResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *ListResourceResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
