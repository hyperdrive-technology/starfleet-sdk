# Starfleet SDK

> **Foundation interfaces and types for the Starfleet ecosystem** - A MIT-licensed contract for building importers, providers, and plugins.

[![npm version](https://badge.fury.io/js/@starfleet%2Fsdk.svg)](https://badge.fury.io/js/@starfleet%2Fsdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/hyperdrive-technology/starfleet-sdk-go.svg)](https://pkg.go.dev/github.com/hyperdrive-technology/starfleet-sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

The Starfleet SDK provides the core TypeScript interfaces, Go bindings, and JSON Schema definitions that enable interoperability across the entire Starfleet ecosystem. Whether you're building importers for infrastructure diagrams, providers for live metrics, or custom plugins for 3D visualization, this SDK defines the contracts you need.

Starfleet is a comprehensive platform for creating interactive 3D visualizations of infrastructure and systems. Think of it as "Infrastructure as Diagrams" with real-time data integration.

## Features

- 🎯 **Type-safe contracts** - Comprehensive TypeScript interfaces and Go structs
- 🔌 **Plugin architecture** - Standardized interfaces for importers, providers, and animation hooks
- 📐 **3D scene modeling** - Complete scene graph with nodes, edges, transforms, and materials
- 🎨 **Material system** - PBR-ready materials with textures and animations
- 📊 **Live data integration** - Real-time metrics and monitoring data support
- ✅ **JSON Schema validation** - Validate scene files in any language
- 🌐 **Multi-language support** - TypeScript and Go bindings included

## Packages

| Language | Package | Description |
|----------|---------|-------------|
| TypeScript | [`@starfleet/sdk`](https://npmjs.com/package/@starfleet/sdk) | Core Scene types, plugin interfaces, utility helpers |
| Go | [`github.com/hyperdrive-technology/starfleet-sdk-go`](https://pkg.go.dev/github.com/hyperdrive-technology/starfleet-sdk-go) | Generated structs and helper functions |
| JSON Schema | [`/schema/scenefile.schema.json`](./schema/scenefile.schema.json) | Validation schema for any toolchain |

## Core Concepts

### Scene Structure
- **SceneFile**: Top-level container with metadata and scene graph
- **SceneNode**: Individual elements (servers, databases, containers, etc.)
- **SceneEdge**: Connections and relationships between nodes
- **Transform**: 3D positioning, rotation, and scale
- **Material**: PBR materials with colors, textures, and properties

### Plugin Architecture
- **Importer**: Transform external data sources (Terraform, K8s, etc.) into SceneFiles
- **Provider**: Supply live data and metrics to scene nodes (Prometheus, Datadog, etc.)
- **AnimationHook**: Custom animation behaviors for scene elements

## Quick Start

### TypeScript

```bash
npm install @starfleet/sdk
```

```typescript
import {
  SceneFile,
  createTransform,
  createMaterial,
  generateId,
  validateScene,
  Importer,
  Provider
} from '@starfleet/sdk';

// Create a simple infrastructure scene
const scene: SceneFile = {
  version: '0.1.0',
  metadata: {
    name: 'My Infrastructure',
    description: 'A sample infrastructure scene',
    author: 'Your Name'
  },
  scene: {
    nodes: [
      {
        id: generateId(),
        type: 'server',
        name: 'Web Server',
        transform: createTransform(
          { x: 0, y: 0, z: 0 }  // position
        ),
        material: createMaterial({
          color: { r: 0.2, g: 0.8, b: 0.2, a: 1 }
        }),
        status: 'healthy',
        metadata: {
          cpu: '85%',
          memory: '12GB'
        }
      }
    ],
    edges: []
  }
};

// Validate the scene
const validation = validateScene(scene);
if (!validation.valid) {
  console.error('Scene validation failed:', validation.errors);
}
```

### Define a Custom Importer

```typescript
import { Importer, ImportResult } from '@starfleet/sdk';

class TerraformImporter implements Importer {
  id = 'terraform-importer';
  name = 'Terraform Infrastructure Importer';
  description = 'Imports Terraform state into 3D scenes';
  supportedFormats = ['.tfstate', '.tf'];

  async import(input: string): Promise<ImportResult> {
    const tfData = JSON.parse(input);

    // Transform Terraform resources into SceneNodes
    const nodes = tfData.resources.map(resource => ({
      id: resource.name,
      type: resource.type,
      name: resource.name,
      transform: createTransform({
        x: Math.random() * 10,
        y: 0,
        z: Math.random() * 10
      }),
      material: createMaterial({
        color: getColorForResourceType(resource.type)
      }),
      metadata: resource.instances[0]?.attributes || {}
    }));

    return {
      scene: {
        version: '0.1.0',
        metadata: {
          name: 'Terraform Infrastructure',
          importedBy: this.id,
          importedAt: new Date().toISOString()
        },
        scene: { nodes, edges: [] }
      },
      warnings: [],
      errors: []
    };
  }

  async validate(input: string): Promise<boolean> {
    try {
      JSON.parse(input);
      return true;
    } catch {
      return false;
    }
  }
}
```

### Define a Custom Provider

```typescript
import { Provider, MetricsQuery, MetricsResult } from '@starfleet/sdk';

class PrometheusProvider implements Provider {
  id = 'prometheus-provider';
  name = 'Prometheus Metrics Provider';
  description = 'Provides metrics from Prometheus';

  private baseUrl: string = '';
  private connected = false;

  async connect(config: any): Promise<void> {
    this.baseUrl = config.url;
    // Test connection
    const response = await fetch(`${this.baseUrl}/-/healthy`);
    if (!response.ok) {
      throw new Error('Failed to connect to Prometheus');
    }
    this.connected = true;
  }

  async disconnect(): Promise<void> {
    this.connected = false;
  }

  isConnected(): boolean {
    return this.connected;
  }

  async query(query: MetricsQuery): Promise<MetricsResult[]> {
    const results: MetricsResult[] = [];

    for (const nodeId of query.nodeIds || []) {
      for (const metricName of query.metricNames || []) {
        // Query Prometheus for this metric
        const promQuery = `${metricName}{instance="${nodeId}"}`;
        const response = await fetch(
          `${this.baseUrl}/api/v1/query?query=${encodeURIComponent(promQuery)}`
        );

        const data = await response.json();

        results.push({
          nodeId,
          metricName,
          dataPoints: data.data.result.map((r: any) => ({
            timestamp: new Date(r.value[0] * 1000),
            value: parseFloat(r.value[1]),
            tags: r.metric
          })),
          unit: getUnitForMetric(metricName)
        });
      }
    }

    return results;
  }

  async healthCheck(): Promise<boolean> {
    if (!this.connected) return false;

    try {
      const response = await fetch(`${this.baseUrl}/-/healthy`);
      return response.ok;
    } catch {
      return false;
    }
  }
}
```

### Go

```bash
go mod init your-project
go get github.com/hyperdrive-technology/starfleet-sdk-go
```

```go
package main

import (
    "fmt"
    "time"

    "github.com/hyperdrive-technology/starfleet-sdk-go"
)

func main() {
    // Create a scene with some infrastructure nodes
    scene := starfleet.NewSceneFile("My Infrastructure")

    // Add a server node
    server := starfleet.SceneNode{
        ID:   "server-1",
        Type: "server",
        Name: "Web Server",
        Transform: starfleet.NewTransform(),
        Material: starfleet.NewMaterial(),
        Status: "healthy",
        Metadata: map[string]interface{}{
            "cpu": "85%",
            "memory": "12GB",
        },
    }

    scene.AddNode(server)

    // Add a database node
    database := starfleet.SceneNode{
        ID:   "db-1",
        Type: "database",
        Name: "Primary Database",
        Transform: starfleet.NewTransformWithPosition(5, 0, 0),
        Material: starfleet.NewMaterial(),
        Status: "healthy",
    }

    scene.AddNode(database)

    // Connect them with an edge
    edge := starfleet.SceneEdge{
        ID:     "conn-1",
        Source: "server-1",
        Target: "db-1",
        Type:   "data-connection",
        Width:  0.1,
        Style:  "solid",
    }

    scene.AddEdge(edge)

    fmt.Printf("Created scene with %d nodes and %d edges\n",
        scene.GetNodeCount(), scene.GetEdgeCount())
}
```

## Repository Structure

```
starfleet-sdk/
├── schema/
│   └── scenefile.schema.json    # JSON Schema for validation
├── ts/
│   ├── src/
│   │   └── index.ts            # TypeScript interfaces and utilities
│   ├── dist/                   # Built packages (generated)
│   ├── package.json
│   └── tsconfig.json
├── go/
│   ├── models.go               # Go struct definitions
│   ├── models_test.go          # Go tests
│   └── go.mod
├── examples/
│   ├── basic-usage/            # Basic usage examples
│   └── README.md              # Examples documentation
├── .github/
│   └── workflows/
│       └── publish.yml         # CI/CD for publishing packages
├── CHANGELOG.md
├── LICENSE
└── README.md
```

## Scene File Format

A Starfleet scene file is a JSON document that represents a 3D infrastructure visualization:

```json
{
  "version": "0.1.0",
  "metadata": {
    "name": "Production Infrastructure",
    "description": "Our main production environment",
    "author": "DevOps Team",
    "created": "2024-01-15T10:00:00Z",
    "tags": ["production", "aws", "kubernetes"]
  },
  "scene": {
    "nodes": [
      {
        "id": "web-server-1",
        "type": "server",
        "name": "Web Server",
        "transform": {
          "position": { "x": 0, "y": 0, "z": 0 },
          "rotation": { "x": 0, "y": 0, "z": 0 },
          "scale": { "x": 1, "y": 1, "z": 1 }
        },
        "geometry": {
          "type": "box",
          "parameters": { "width": 2, "height": 1, "depth": 1 }
        },
        "material": {
          "color": { "r": 0.2, "g": 0.8, "b": 0.2, "a": 1 },
          "metalness": 0.1,
          "roughness": 0.7
        },
        "status": "healthy",
        "metadata": {
          "cpu": "85%",
          "memory": "12GB",
          "instance_type": "t3.large"
        }
      }
    ],
    "edges": [
      {
        "id": "connection-1",
        "source": "web-server-1",
        "target": "database-1",
        "type": "data-connection",
        "color": { "r": 0.5, "g": 0.5, "b": 0.5, "a": 0.8 },
        "width": 0.1,
        "style": "solid"
      }
    ]
  }
}
```

## Integration with Starfleet Ecosystem

The SDK is designed to work seamlessly with other Starfleet components:

- **[starfleet](https://github.com/hyperdrive-technology/starfleet)** - Main monorepo with CLI and React components
- **[starfleet-gateway](https://github.com/hyperdrive-technology/starfleet-gateway)** - Go service for metrics aggregation
- **[starfleet-importer-tf](https://github.com/hyperdrive-technology/starfleet-importer-tf)** - Terraform infrastructure import
- **[starfleet-provider-otel](https://github.com/hyperdrive-technology/starfleet-provider-otel)** - OpenTelemetry metrics provider
- **[starfleet-provider-hyperdrive](https://github.com/hyperdrive-technology/starfleet-provider-hyperdrive)** - Real-time data bridge

## Plugin Development

### Creating an Importer

Importers transform external data sources into Starfleet scene files:

1. Implement the `Importer` interface
2. Define supported file formats
3. Parse input data and create SceneNodes/SceneEdges
4. Return a complete SceneFile with metadata

### Creating a Provider

Providers supply live data to scene nodes:

1. Implement the `Provider` interface
2. Handle connection lifecycle (connect/disconnect)
3. Query external systems for metrics
4. Return standardized MetricsResult data

### Creating Animation Hooks

Animation hooks add custom behaviors to scenes:

1. Implement the `AnimationHook` interface
2. Handle frame updates and node/edge changes
3. Update scene properties based on data or time

## Building from Source

### TypeScript Package

```bash
cd ts
npm install
npm run build
npm test
```

### Go Module

```bash
cd go
go mod tidy
go build
go test
```

## Versioning

This project follows [Semantic Versioning](https://semver.org/).

Breaking changes to the core interfaces will result in a major version bump. All ecosystem packages should pin to the same major version of the SDK to ensure compatibility.

## Contributing

We welcome contributions! Please see our [contributing guidelines](CONTRIBUTING.md) for details.

### Development Workflow

1. Fork the repository
2. Create a feature branch
3. Make your changes (ensure tests pass)
4. Update documentation if needed
5. Submit a pull request

## License

MIT - See [LICENSE](LICENSE) for details.

## Support

- 📖 [Documentation](https://github.com/hyperdrive-technology/starfleet-sdk#readme)
- 🐛 [Issues](https://github.com/hyperdrive-technology/starfleet-sdk/issues)
- 💬 [Discussions](https://github.com/hyperdrive-technology/starfleet-sdk/discussions)
- 🌟 [Starfleet Organization](https://github.com/hyperdrive-technology)

---

Built with ❤️ by the [Hyperdrive Technology](https://github.com/hyperdrive-technology) team.
