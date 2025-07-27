// Package starfleet provides core types and interfaces for the Starfleet ecosystem
package starfleet

import (
	"time"
)

// =============================================================================
// CORE SCENE TYPES
// =============================================================================

// Vector3 represents a 3D position in scene space
type Vector3 struct {
	X float64 `json:"x" validate:"required"`
	Y float64 `json:"y" validate:"required"`
	Z float64 `json:"z" validate:"required"`
}

// Euler3 represents 3D rotation (Euler angles in radians)
type Euler3 struct {
	X float64 `json:"x" validate:"required"`
	Y float64 `json:"y" validate:"required"`
	Z float64 `json:"z" validate:"required"`
}

// Scale3 represents 3D scale factors
type Scale3 struct {
	X float64 `json:"x" validate:"required"`
	Y float64 `json:"y" validate:"required"`
	Z float64 `json:"z" validate:"required"`
}

// Transform represents a transform matrix for 3D positioning
type Transform struct {
	Position Vector3 `json:"position" validate:"required"`
	Rotation Euler3  `json:"rotation" validate:"required"`
	Scale    Scale3  `json:"scale" validate:"required"`
}

// Color represents RGBA color representation
type Color struct {
	R float64 `json:"r" validate:"required,min=0,max=1"`
	G float64 `json:"g" validate:"required,min=0,max=1"`
	B float64 `json:"b" validate:"required,min=0,max=1"`
	A float64 `json:"a,omitempty" validate:"omitempty,min=0,max=1"`
}

// Material represents material properties for 3D rendering
type Material struct {
	Color       *Color  `json:"color,omitempty"`
	Emissive    *Color  `json:"emissive,omitempty"`
	Metalness   float64 `json:"metalness,omitempty" validate:"omitempty,min=0,max=1"`
	Roughness   float64 `json:"roughness,omitempty" validate:"omitempty,min=0,max=1"`
	Opacity     float64 `json:"opacity,omitempty" validate:"omitempty,min=0,max=1"`
	Transparent bool    `json:"transparent,omitempty"`
	Wireframe   bool    `json:"wireframe,omitempty"`
	Texture     string  `json:"texture,omitempty"`
}

// GeometryType represents the type of geometry
type GeometryType string

const (
	GeometryBox      GeometryType = "box"
	GeometrySphere   GeometryType = "sphere"
	GeometryCylinder GeometryType = "cylinder"
	GeometryPlane    GeometryType = "plane"
	GeometryCustom   GeometryType = "custom"
)

// Geometry represents geometry definition for 3D objects
type Geometry struct {
	Type       GeometryType           `json:"type" validate:"required"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Asset      string                 `json:"asset,omitempty"`
}

// EasingType represents animation easing types
type EasingType string

const (
	EasingLinear    EasingType = "linear"
	EasingEaseIn    EasingType = "ease-in"
	EasingEaseOut   EasingType = "ease-out"
	EasingEaseInOut EasingType = "ease-in-out"
)

// Keyframe represents an animation keyframe
type Keyframe struct {
	Time   float64     `json:"time" validate:"required"`
	Value  interface{} `json:"value" validate:"required"`
	Easing EasingType  `json:"easing,omitempty"`
}

// AnimationTrack represents an animation track definition
type AnimationTrack struct {
	Property  string     `json:"property" validate:"required"`
	Keyframes []Keyframe `json:"keyframes" validate:"required"`
}

// Animation represents an animation definition
type Animation struct {
	Name     string           `json:"name" validate:"required"`
	Duration float64          `json:"duration" validate:"required"`
	Loop     bool             `json:"loop,omitempty"`
	Tracks   []AnimationTrack `json:"tracks" validate:"required"`
}

// NodeStatus represents the status of a scene node
type NodeStatus string

const (
	NodeStatusHealthy  NodeStatus = "healthy"
	NodeStatusWarning  NodeStatus = "warning"
	NodeStatusCritical NodeStatus = "critical"
	NodeStatusUnknown  NodeStatus = "unknown"
)

// SceneNode represents an individual node in the scene graph
type SceneNode struct {
	ID         string                 `json:"id" validate:"required"`
	Type       string                 `json:"type" validate:"required"`
	Name       string                 `json:"name" validate:"required"`
	Transform  Transform              `json:"transform" validate:"required"`
	Geometry   *Geometry              `json:"geometry,omitempty"`
	Material   *Material              `json:"material,omitempty"`
	Visible    bool                   `json:"visible,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Tags       []string               `json:"tags,omitempty"`
	Metrics    map[string]interface{} `json:"metrics,omitempty"`
	Status     NodeStatus             `json:"status,omitempty"`
	Animations []Animation            `json:"animations,omitempty"`
	Parent     string                 `json:"parent,omitempty"`
	Children   []string               `json:"children,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// EdgeStyle represents the style of an edge
type EdgeStyle string

const (
	EdgeStyleSolid  EdgeStyle = "solid"
	EdgeStyleDashed EdgeStyle = "dashed"
	EdgeStyleDotted EdgeStyle = "dotted"
)

// SceneEdge represents a connection between two nodes
type SceneEdge struct {
	ID         string                 `json:"id" validate:"required"`
	Source     string                 `json:"source" validate:"required"`
	Target     string                 `json:"target" validate:"required"`
	Type       string                 `json:"type,omitempty"`
	Color      *Color                 `json:"color,omitempty"`
	Width      float64                `json:"width,omitempty"`
	Style      EdgeStyle              `json:"style,omitempty"`
	Opacity    float64                `json:"opacity,omitempty" validate:"omitempty,min=0,max=1"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Metrics    map[string]interface{} `json:"metrics,omitempty"`
	Animations []Animation            `json:"animations,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// LightType represents the type of light
type LightType string

const (
	LightAmbient     LightType = "ambient"
	LightDirectional LightType = "directional"
	LightPoint       LightType = "point"
	LightSpot        LightType = "spot"
)

// Light represents a light definition
type Light struct {
	Type      LightType `json:"type" validate:"required"`
	Color     *Color    `json:"color,omitempty"`
	Intensity float64   `json:"intensity,omitempty"`
	Position  *Vector3  `json:"position,omitempty"`
	Direction *Vector3  `json:"direction,omitempty"`
}

// Fog represents fog definition
type Fog struct {
	Color Color   `json:"color" validate:"required"`
	Near  float64 `json:"near" validate:"required"`
	Far   float64 `json:"far" validate:"required"`
}

// Environment represents environment settings
type Environment struct {
	Background interface{} `json:"background,omitempty"` // Can be Color or string
	Fog        *Fog        `json:"fog,omitempty"`
}

// Camera represents camera settings
type Camera struct {
	Position Vector3 `json:"position" validate:"required"`
	Target   Vector3 `json:"target" validate:"required"`
	FOV      float64 `json:"fov,omitempty"`
	Near     float64 `json:"near,omitempty"`
	Far      float64 `json:"far,omitempty"`
}

// Bounds represents scene bounds
type Bounds struct {
	Min Vector3 `json:"min" validate:"required"`
	Max Vector3 `json:"max" validate:"required"`
}

// SceneGraph represents a scene graph containing all nodes and edges
type SceneGraph struct {
	Nodes       []SceneNode  `json:"nodes" validate:"required"`
	Edges       []SceneEdge  `json:"edges" validate:"required"`
	Bounds      *Bounds      `json:"bounds,omitempty"`
	Camera      *Camera      `json:"camera,omitempty"`
	Lights      []Light      `json:"lights,omitempty"`
	Environment *Environment `json:"environment,omitempty"`
}

// SceneMetadata represents scene metadata
type SceneMetadata struct {
	Name         string                 `json:"name" validate:"required"`
	Description  string                 `json:"description,omitempty"`
	Author       string                 `json:"author,omitempty"`
	Version      string                 `json:"version,omitempty"`
	Created      *time.Time             `json:"created,omitempty"`
	Updated      *time.Time             `json:"updated,omitempty"`
	Tags         []string               `json:"tags,omitempty"`
	ImportSource string                 `json:"importSource,omitempty"`
	ImportedAt   *time.Time             `json:"importedAt,omitempty"`
	ImportedBy   string                 `json:"importedBy,omitempty"`
	Extensions   map[string]interface{} `json:"extensions,omitempty"`
}

// SceneFile represents a complete scene file
type SceneFile struct {
	Version    string                 `json:"version" validate:"required"`
	Metadata   SceneMetadata          `json:"metadata" validate:"required"`
	Scene      SceneGraph             `json:"scene" validate:"required"`
	Assets     map[string]string      `json:"assets,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// =============================================================================
// PLUGIN INTERFACES
// =============================================================================

// ImporterConfig represents configuration for importers
type ImporterConfig map[string]interface{}

// ImportResult represents the result of an import operation
type ImportResult struct {
	Scene    SceneFile `json:"scene" validate:"required"`
	Warnings []string  `json:"warnings,omitempty"`
	Errors   []string  `json:"errors,omitempty"`
}

// ProviderConfig represents configuration for providers
type ProviderConfig map[string]interface{}

// MetricsQuery represents a query for metrics data
type MetricsQuery struct {
	NodeIDs     []string               `json:"nodeIds,omitempty"`
	MetricNames []string               `json:"metricNames,omitempty"`
	From        *time.Time             `json:"from,omitempty"`
	To          *time.Time             `json:"to,omitempty"`
	Resolution  int                    `json:"resolution,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"`
}

// MetricsDataPoint represents a single metrics data point
type MetricsDataPoint struct {
	Timestamp time.Time         `json:"timestamp" validate:"required"`
	Value     interface{}       `json:"value" validate:"required"`
	Tags      map[string]string `json:"tags,omitempty"`
}

// MetricsResult represents the result of a metrics query
type MetricsResult struct {
	NodeID     string                 `json:"nodeId" validate:"required"`
	MetricName string                 `json:"metricName" validate:"required"`
	DataPoints []MetricsDataPoint     `json:"dataPoints" validate:"required"`
	Unit       string                 `json:"unit,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

// =============================================================================
// UTILITY TYPES
// =============================================================================

// ValidationResult represents the result of scene validation
type ValidationResult struct {
	Valid    bool     `json:"valid"`
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
}

// SceneStatsSize represents the size of scene bounds
type SceneStatsSize struct {
	Min  Vector3 `json:"min"`
	Max  Vector3 `json:"max"`
	Size Vector3 `json:"size"`
}

// SceneStats represents statistics about a scene
type SceneStats struct {
	NodeCount      int             `json:"nodeCount"`
	EdgeCount      int             `json:"edgeCount"`
	TotalVertices  int             `json:"totalVertices,omitempty"`
	TotalTriangles int             `json:"totalTriangles,omitempty"`
	MemoryUsage    int64           `json:"memoryUsage,omitempty"`
	Bounds         *SceneStatsSize `json:"bounds,omitempty"`
}

// =============================================================================
// UTILITY FUNCTIONS
// =============================================================================

// NewTransform creates a new transform with default values
func NewTransform() Transform {
	return Transform{
		Position: Vector3{X: 0, Y: 0, Z: 0},
		Rotation: Euler3{X: 0, Y: 0, Z: 0},
		Scale:    Scale3{X: 1, Y: 1, Z: 1},
	}
}

// NewTransformWithPosition creates a new transform with specified position
func NewTransformWithPosition(x, y, z float64) Transform {
	return Transform{
		Position: Vector3{X: x, Y: y, Z: z},
		Rotation: Euler3{X: 0, Y: 0, Z: 0},
		Scale:    Scale3{X: 1, Y: 1, Z: 1},
	}
}

// NewMaterial creates a new material with default values
func NewMaterial() Material {
	return Material{
		Color:       &Color{R: 0.8, G: 0.8, B: 0.8, A: 1.0},
		Metalness:   0.0,
		Roughness:   0.5,
		Opacity:     1.0,
		Transparent: false,
		Wireframe:   false,
	}
}

// NewColor creates a new color with RGB values
func NewColor(r, g, b float64) Color {
	return Color{R: r, G: g, B: b, A: 1.0}
}

// NewColorWithAlpha creates a new color with RGBA values
func NewColorWithAlpha(r, g, b, a float64) Color {
	return Color{R: r, G: g, B: b, A: a}
}

// NewSceneFile creates a new scene file with minimal required fields
func NewSceneFile(name string) SceneFile {
	now := time.Now()
	return SceneFile{
		Version: "0.1.0",
		Metadata: SceneMetadata{
			Name:    name,
			Created: &now,
			Updated: &now,
		},
		Scene: SceneGraph{
			Nodes: []SceneNode{},
			Edges: []SceneEdge{},
		},
		Assets: make(map[string]string),
	}
}

// AddNode adds a node to the scene graph
func (sf *SceneFile) AddNode(node SceneNode) {
	sf.Scene.Nodes = append(sf.Scene.Nodes, node)
}

// AddEdge adds an edge to the scene graph
func (sf *SceneFile) AddEdge(edge SceneEdge) {
	sf.Scene.Edges = append(sf.Scene.Edges, edge)
}

// FindNode finds a node by ID
func (sf *SceneFile) FindNode(id string) *SceneNode {
	for i := range sf.Scene.Nodes {
		if sf.Scene.Nodes[i].ID == id {
			return &sf.Scene.Nodes[i]
		}
	}
	return nil
}

// FindEdge finds an edge by ID
func (sf *SceneFile) FindEdge(id string) *SceneEdge {
	for i := range sf.Scene.Edges {
		if sf.Scene.Edges[i].ID == id {
			return &sf.Scene.Edges[i]
		}
	}
	return nil
}

// GetNodeCount returns the number of nodes in the scene
func (sf *SceneFile) GetNodeCount() int {
	return len(sf.Scene.Nodes)
}

// GetEdgeCount returns the number of edges in the scene
func (sf *SceneFile) GetEdgeCount() int {
	return len(sf.Scene.Edges)
}

// UpdateTimestamp updates the scene's updated timestamp
func (sf *SceneFile) UpdateTimestamp() {
	now := time.Now()
	sf.Metadata.Updated = &now
}
