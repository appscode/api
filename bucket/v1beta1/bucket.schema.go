package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var bucketListRequestSchema *gojsonschema.Schema

func init() {
	var err error
	bucketListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cloud_credential": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *BucketListRequest) IsValid() (*gojsonschema.Result, error) {
	return bucketListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BucketListRequest) IsRequest() {}

func (m *BucketListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
