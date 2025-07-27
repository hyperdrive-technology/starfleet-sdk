package starfleet

import (
	"encoding/json"
	"testing"
)

// TestVector3_JSON tests Vector3 JSON marshaling/unmarshaling
func TestVector3_JSON(t *testing.T) {
	original := Vector3{X: 1.5, Y: 2.5, Z: 3.5}
	
	// Test marshaling
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal Vector3: %v", err)
	}
	
	// Test unmarshaling
	var result Vector3
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Vector3: %v", err)
	}
	
	// Verify values
	if result.X != original.X || result.Y != original.Y || result.Z != original.Z {
		t.Errorf("Vector3 mismatch: got %+v, want %+v", result, original)
	}
}

// TestEuler3_JSON tests Euler3 JSON marshaling/unmarshaling
func TestEuler3_JSON(t *testing.T) {
	original := Euler3{X: 0.1, Y: 0.2, Z: 0.3}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal Euler3: %v", err)
	}
	
	var result Euler3
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Euler3: %v", err)
	}
	
	if result.X != original.X || result.Y != original.Y || result.Z != original.Z {
		t.Errorf("Euler3 mismatch: got %+v, want %+v", result, original)
	}
}

// TestTransform_JSON tests Transform JSON marshaling/unmarshaling
func TestTransform_JSON(t *testing.T) {
	original := Transform{
		Position: Vector3{X: 1, Y: 2, Z: 3},
		Rotation: Euler3{X: 0.1, Y: 0.2, Z: 0.3},
		Scale:    Scale3{X: 2, Y: 2, Z: 2},
	}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal Transform: %v", err)
	}
	
	var result Transform
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Transform: %v", err)
	}
	
	if result.Position != original.Position || result.Rotation != original.Rotation || result.Scale != original.Scale {
		t.Errorf("Transform mismatch: got %+v, want %+v", result, original)
	}
}

// TestColor_JSON tests Color JSON marshaling/unmarshaling
func TestColor_JSON(t *testing.T) {
	original := Color{R: 0.8, G: 0.4, B: 0.2, A: 0.9}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal Color: %v", err)
	}
	
	var result Color
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Color: %v", err)
	}
	
	if result.R != original.R || result.G != original.G || result.B != original.B || result.A != original.A {
		t.Errorf("Color mismatch: got %+v, want %+v", result, original)
	}
}

// TestMaterial_JSON tests Material JSON marshaling/unmarshaling
func TestMaterial_JSON(t *testing.T) {
	color := Color{R: 1, G: 0, B: 0, A: 1}
	
	original := Material{
		Color:       &color,
		Emissive:    &Color{R: 0.1, G: 0.1, B: 0.1, A: 1},
		Metalness:   0.8,
		Roughness:   0.2,
		Opacity:     0.9,
		Transparent: true,
		Wireframe:   false,
		Texture:     "texture.jpg",
	}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal Material: %v", err)
	}
	
	var result Material
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Material: %v", err)
	}
	
	// Check all fields
	if result.Color == nil || *result.Color != *original.Color {
		t.Errorf("Material color mismatch")
	}
	if result.Metalness != original.Metalness {
		t.Errorf("Material metalness mismatch: got %f, want %f", result.Metalness, original.Metalness)
	}
	if result.Texture != original.Texture {
		t.Errorf("Material texture mismatch: got %s, want %s", result.Texture, original.Texture)
	}
}

// TestSceneNode_JSON tests SceneNode JSON marshaling/unmarshaling
func TestSceneNode_JSON(t *testing.T) {
	original := SceneNode{
		ID:   "test-node",
		Type: "server",
		Name: "Test Server",
		Transform: Transform{
			Position: Vector3{X: 0, Y: 0, Z: 0},
			Rotation: Euler3{X: 0, Y: 0, Z: 0},
			Scale:    Scale3{X: 1, Y: 1, Z: 1},
		},
		Geometry: &Geometry{
			Type:       GeometryBox,
			Parameters: map[string]interface{}{"width": 2, "height": 1, "depth": 1},
		},
		Material: &Material{
			Color: &Color{R: 0, G: 1, B: 0, A: 1},
		},
		Visible:    true,
		Metadata:   map[string]interface{}{"cpu": "2.4GHz", "ram": "16GB"},
		Tags:       []string{"production", "web"},
		Metrics:    map[string]interface{}{"cpu_usage": 85.5},
		Status:     NodeStatusHealthy,
		Parent:     "parent-node",
		Children:   []string{"child1", "child2"},
		Extensions: map[string]interface{}{"custom": "value"},
	}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal SceneNode: %v", err)
	}
	
	var result SceneNode
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal SceneNode: %v", err)
	}
	
	// Check basic fields
	if result.ID != original.ID || result.Type != original.Type || result.Name != original.Name {
		t.Errorf("SceneNode basic fields mismatch")
	}
	
	// Check transform
	if result.Transform.Position != original.Transform.Position {
		t.Errorf("SceneNode transform position mismatch")
	}
	
	// Check visible
	if result.Visible != original.Visible {
		t.Errorf("SceneNode visible mismatch: got %t, want %t", result.Visible, original.Visible)
	}
	
	// Check children
	if len(result.Children) != len(original.Children) {
		t.Errorf("SceneNode children count mismatch")
	}
}

// TestSceneEdge_JSON tests SceneEdge JSON marshaling/unmarshaling
func TestSceneEdge_JSON(t *testing.T) {
	original := SceneEdge{
		ID:      "edge-1",
		Source:  "node-1",
		Target:  "node-2",
		Type:    "connection",
		Color:   &Color{R: 1, G: 1, B: 0, A: 0.8},
		Width:   2.0,
		Style:   EdgeStyleSolid,
		Opacity: 0.8,
		Metadata: map[string]interface{}{"port": 8080},
		Metrics:  map[string]interface{}{"latency": 50},
		Extensions: map[string]interface{}{"custom": "edge-value"},
	}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal SceneEdge: %v", err)
	}
	
	var result SceneEdge
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal SceneEdge: %v", err)
	}
	
	// Check basic fields
	if result.ID != original.ID || result.Source != original.Source || result.Target != original.Target {
		t.Errorf("SceneEdge basic fields mismatch")
	}
	
	// Check type
	if result.Type != original.Type {
		t.Errorf("SceneEdge type mismatch: got %s, want %s", result.Type, original.Type)
	}
	
	// Check width
	if result.Width != original.Width {
		t.Errorf("SceneEdge width mismatch: got %f, want %f", result.Width, original.Width)
	}
}

// TestSceneGraph_JSON tests SceneGraph JSON marshaling/unmarshaling
func TestSceneGraph_JSON(t *testing.T) {
	original := SceneGraph{
		Nodes: []SceneNode{
			{
				ID:   "node-1",
				Type: "server",
				Name: "Server 1",
				Transform: Transform{
					Position: Vector3{X: 0, Y: 0, Z: 0},
					Rotation: Euler3{X: 0, Y: 0, Z: 0},
					Scale:    Scale3{X: 1, Y: 1, Z: 1},
				},
			},
		},
		Edges: []SceneEdge{
			{
				ID:     "edge-1",
				Source: "node-1",
				Target: "node-2",
			},
		},
		Bounds: &Bounds{
			Min: Vector3{X: -10, Y: -10, Z: -10},
			Max: Vector3{X: 10, Y: 10, Z: 10},
		},
		Camera: &Camera{
			Position: Vector3{X: 0, Y: 10, Z: 10},
			Target:   Vector3{X: 0, Y: 0, Z: 0},
			FOV:      75,
		},
	}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal SceneGraph: %v", err)
	}
	
	var result SceneGraph
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal SceneGraph: %v", err)
	}
	
	// Check nodes count
	if len(result.Nodes) != len(original.Nodes) {
		t.Errorf("SceneGraph nodes count mismatch: got %d, want %d", len(result.Nodes), len(original.Nodes))
	}
	
	// Check edges count
	if len(result.Edges) != len(original.Edges) {
		t.Errorf("SceneGraph edges count mismatch: got %d, want %d", len(result.Edges), len(original.Edges))
	}
	
	// Check bounds
	if result.Bounds == nil || original.Bounds == nil {
		t.Errorf("SceneGraph bounds nil mismatch")
	} else if result.Bounds.Min != original.Bounds.Min || result.Bounds.Max != original.Bounds.Max {
		t.Errorf("SceneGraph bounds mismatch")
	}
}

// TestSceneFile_JSON tests complete SceneFile JSON marshaling/unmarshaling
func TestSceneFile_JSON(t *testing.T) {
	original := SceneFile{
		Version: "0.1.0",
		Metadata: SceneMetadata{
			Name:        "Test Scene",
			Description: "A test scene",
			Author:      "Test Author",
			Version:     "1.0.0",
			Tags:        []string{"test", "example"},
		},
		Scene: SceneGraph{
			Nodes: []SceneNode{
				{
					ID:   "node-1",
					Type: "server",
					Name: "Test Server",
					Transform: Transform{
						Position: Vector3{X: 0, Y: 0, Z: 0},
						Rotation: Euler3{X: 0, Y: 0, Z: 0},
						Scale:    Scale3{X: 1, Y: 1, Z: 1},
					},
				},
			},
			Edges: []SceneEdge{},
		},
		Assets: map[string]string{
			"texture1": "https://example.com/texture.jpg",
		},
		Extensions: map[string]interface{}{
			"custom": "scene-value",
		},
	}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal SceneFile: %v", err)
	}
	
	var result SceneFile
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal SceneFile: %v", err)
	}
	
	// Check version
	if result.Version != original.Version {
		t.Errorf("SceneFile version mismatch: got %s, want %s", result.Version, original.Version)
	}
	
	// Check metadata
	if result.Metadata.Name != original.Metadata.Name {
		t.Errorf("SceneFile metadata name mismatch")
	}
	
	// Check scene nodes
	if len(result.Scene.Nodes) != len(original.Scene.Nodes) {
		t.Errorf("SceneFile scene nodes count mismatch")
	}
	
	// Check assets
	if len(result.Assets) != len(original.Assets) {
		t.Errorf("SceneFile assets count mismatch")
	}
}

// TestKeyframe_JSON tests Keyframe JSON marshaling/unmarshaling
func TestKeyframe_JSON(t *testing.T) {
	original := Keyframe{
		Time:   1.5,
		Value:  42.0,
		Easing: EasingEaseInOut,
	}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal Keyframe: %v", err)
	}
	
	var result Keyframe
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Keyframe: %v", err)
	}
	
	if result.Time != original.Time {
		t.Errorf("Keyframe time mismatch: got %f, want %f", result.Time, original.Time)
	}
	
	if result.Easing != original.Easing {
		t.Errorf("Keyframe easing mismatch: got %s, want %s", result.Easing, original.Easing)
	}
}

// TestAnimation_JSON tests Animation JSON marshaling/unmarshaling
func TestAnimation_JSON(t *testing.T) {
	original := Animation{
		Name:     "test-animation",
		Duration: 2.0,
		Loop:     true,
		Tracks: []AnimationTrack{
			{
				Property: "transform.position.x",
				Keyframes: []Keyframe{
					{Time: 0, Value: 0},
					{Time: 2, Value: 10},
				},
			},
		},
	}
	
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal Animation: %v", err)
	}
	
	var result Animation
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Animation: %v", err)
	}
	
	if result.Name != original.Name || result.Duration != original.Duration {
		t.Errorf("Animation basic fields mismatch")
	}
	
	if len(result.Tracks) != len(original.Tracks) {
		t.Errorf("Animation tracks count mismatch")
	}
	
	if result.Loop != original.Loop {
		t.Errorf("Animation loop mismatch: got %t, want %t", result.Loop, original.Loop)
	}
}

// BenchmarkSceneFile_Marshal benchmarks SceneFile JSON marshaling
func BenchmarkSceneFile_Marshal(b *testing.B) {
	sceneFile := SceneFile{
		Version: "0.1.0",
		Metadata: SceneMetadata{
			Name: "Benchmark Scene",
		},
		Scene: SceneGraph{
			Nodes: make([]SceneNode, 100),
			Edges: make([]SceneEdge, 50),
		},
	}
	
	// Initialize nodes
	for i := range sceneFile.Scene.Nodes {
		sceneFile.Scene.Nodes[i] = SceneNode{
			ID:   "node-" + string(rune(i)),
			Type: "server",
			Name: "Server " + string(rune(i)),
			Transform: Transform{
				Position: Vector3{X: float64(i), Y: float64(i), Z: float64(i)},
				Rotation: Euler3{X: 0, Y: 0, Z: 0},
				Scale:    Scale3{X: 1, Y: 1, Z: 1},
			},
		}
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(sceneFile)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSceneFile_Unmarshal benchmarks SceneFile JSON unmarshaling
func BenchmarkSceneFile_Unmarshal(b *testing.B) {
	sceneFile := SceneFile{
		Version: "0.1.0",
		Metadata: SceneMetadata{
			Name: "Benchmark Scene",
		},
		Scene: SceneGraph{
			Nodes: make([]SceneNode, 100),
			Edges: make([]SceneEdge, 50),
		},
	}
	
	data, err := json.Marshal(sceneFile)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result SceneFile
		err := json.Unmarshal(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}