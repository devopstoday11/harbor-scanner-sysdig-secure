package secure

import "time"

type ScanResponse struct {
	AnalysisStatus string         `json:"analysis_status"`
	CreatedAt      time.Time      `json:"created_at"`
	ImageDigest    string         `json:"imageDigest"`
	ImageStatus    string         `json:"image_status"`
	LastUpdated    time.Time      `json:"last_updated"`
	ParentDigest   string         `json:"parentDigest"`
	UserID         string         `json:"userId"`
	ImageContent   *ImageContent  `json:"image_content"`
	ImageDetail    []*ImageDetail `json:"image_detail"`
}

type ImageContent struct {
	Metadata *Metadata `json:"metadata"`
}
type Metadata struct {
	Arch           string `json:"arch"`
	Distro         string `json:"distro"`
	DistroVersion  string `json:"distro_version"`
	DockerfileMode string `json:"dockerfile_mode"`
	ImageSize      int    `json:"image_size"`
	LayerCount     int    `json:"layer_count"`
}

type ImageDetail struct {
	CreatedAt   time.Time `json:"created_at"`
	Dockerfile  string    `json:"dockerfile"`
	FullDigest  string    `json:"full_digest"`
	FullTag     string    `json:"full_tag"`
	ImageDigest string    `json:"image_digest"`
	ImageID     string    `json:"imageId"`
	LastUpdated time.Time `json:"last_updated"`
	Registry    string    `json:"registry"`
	Repo        string    `json:"repo"`
	UserID      string    `json:"userId"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type VulnerabilityReport struct {
	ImageDigest       string           `json:"imageDigest"`
	VulnerabilityType string           `json:"vulnerability_type"`
	Vulnerabilities   []*Vulnerability `json:"vulnerabilities"`
}

type Vulnerability struct {
	Feed           string     `json:"feed"`
	FeedGroup      string     `json:"feed_group"`
	Fix            string     `json:"fix"`
	NVDData        []*NVDData `json:"nvd_data"`
	Package        string     `json:"package"`
	PackageCPE     string     `json:"package_cpe"`
	PackageName    string     `json:"package_name"`
	PackagePath    string     `json:"package_path"`
	PackageType    string     `json:"package_type"`
	PackageVersion string     `json:"package_version"`
	Severity       string     `json:"severity"`
	URL            string     `json:"url"`
	Vuln           string     `json:"vuln"`
}

type NVDData struct {
	ID    string `json:"id"`
	CSSV2 *CSS   `json:"css_v2"`
	CSSV3 *CSS   `json:"css_v3"`
}

type CSS struct {
	BaseScore           float32 `json:"base_score"`
	ExploitabilityScore float32 `json:"base_score"`
	ImpactScore         float32 `json:"impact_score"`
}
