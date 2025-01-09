package main

// Config defines format of configuration file.
type Config struct {
	Source struct {
		Type     SourceType `json:"type"`
		URL      string     `json:"url"`      // type: http
		Auth     AuthType   `json:"auth"`     // type: http
		User     string     `json:"user"`     // type: http
		Password string     `json:"password"` // type: http
		Path     string     `json:"path"`     // type: file
	} `json:"src"`
	Convert struct {
		Width  int        `json:"width"`
		Height int        `json:"height"`
		Format FormatType `json:"format"`
	} `json:"convert"`
	Destination struct {
		Type             DestinationType `json:"type"`
		PathLayout       string          `json:"path_layout"`        // type: s3, file
		LayoutMode       LayoutMode      `json:"layout_mode"`        // type: s3, file
		CreateLatestFile bool            `json:"create_latest_file"` // type: s3, file
		LatestFilePath   string          `json:"latest_file_path"`   // type: s3, file
		Endpoint         string          `json:"endpoint"`           // type: s3
		Bucket           string          `json:"bucket"`             // type: s3
		AccessKeyID      string          `json:"access_key_id"`      // type: s3
		SecretAccessKey  string          `json:"secret_access_key"`  // type: s3
		Region           string          `json:"region"`             // type: s3
	} `json:"dest"`
}

// SourceType defines type of sources.
type SourceType string

// SourceType constants are value of each source type.
const (
	SourceTypeHTTP SourceType = "http"
	SourceTypeFile SourceType = "file"
)

type DestinationType string

// DestinationType constants are value of each destination type.
const (
	DestinationTypeS3   DestinationType = "s3"
	DestinationTypeFile DestinationType = "file"
)

type AuthType string

const (
	AuthTypeBasic  AuthType = "basic"
	AuthTypeDigest AuthType = "digest"
)

// LayoutMode defines type of layout modes.
type LayoutMode int

// LayoutMode constants are value of each layout mode.
const (
	LayoutModeAll      LayoutMode = 0
	LayoutModeDisable  LayoutMode = 1
	LayoutModeFileName LayoutMode = 2
)

// FormatType defines type of formats.
type FormatType string
