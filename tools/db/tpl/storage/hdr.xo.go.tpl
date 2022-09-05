{{ define "header" }}
// Package {{ pkg }} contains generated code from xo.
package {{ pkg }}

// Code generated by xo. DO NOT EDIT.

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

{{ range protonames }} 
	{{ printf "%q" . }}
{{- end }}
)
{{ end }}
