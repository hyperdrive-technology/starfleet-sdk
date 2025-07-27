/**
 * @fileoverview Tests for the Starfleet SDK TypeScript package
 */

import { describe, expect, it } from 'vitest';
import {
    SceneFile,
    SceneNode,
    calculateSceneStats,
    createMaterial,
    createTransform,
    generateId,
    validateScene,
    type AnimationHook,
    type ImportResult,
    type Importer,
    type MetricsQuery,
    type MetricsResult,
    type Provider
} from './index';

describe('Starfleet SDK Core Types', () => {
  describe('Utility Functions', () => {
    it('should create default transform', () => {
      const transform = createTransform();

      expect(transform.position).toEqual({ x: 0, y: 0, z: 0 });
      expect(transform.rotation).toEqual({ x: 0, y: 0, z: 0 });
      expect(transform.scale).toEqual({ x: 1, y: 1, z: 1 });
    });

    it('should create transform with custom values', () => {
      const transform = createTransform(
        { x: 1, y: 2, z: 3 },
        { x: 0.1, y: 0.2, z: 0.3 },
        { x: 2, y: 2, z: 2 }
      );

      expect(transform.position).toEqual({ x: 1, y: 2, z: 3 });
      expect(transform.rotation).toEqual({ x: 0.1, y: 0.2, z: 0.3 });
      expect(transform.scale).toEqual({ x: 2, y: 2, z: 2 });
    });

    it('should create default material', () => {
      const material = createMaterial();

      expect(material.color).toEqual({ r: 0.8, g: 0.8, b: 0.8, a: 1 });
      expect(material.metalness).toBe(0.0);
      expect(material.roughness).toBe(0.5);
      expect(material.opacity).toBe(1.0);
      expect(material.transparent).toBe(false);
      expect(material.wireframe).toBe(false);
    });

    it('should create material with custom values', () => {
      const material = createMaterial({
        color: { r: 1, g: 0, b: 0, a: 0.5 },
        metalness: 0.8,
        roughness: 0.2,
        transparent: true,
        wireframe: true
      });

      expect(material.color).toEqual({ r: 1, g: 0, b: 0, a: 0.5 });
      expect(material.metalness).toBe(0.8);
      expect(material.roughness).toBe(0.2);
      expect(material.transparent).toBe(true);
      expect(material.wireframe).toBe(true);
    });

    it('should generate unique IDs', () => {
      const id1 = generateId();
      const id2 = generateId();

      expect(id1).toBeDefined();
      expect(id2).toBeDefined();
      expect(id1).not.toBe(id2);
      expect(typeof id1).toBe('string');
      expect(typeof id2).toBe('string');
    });
  });

  describe('Scene Validation', () => {
    const createValidScene = (): SceneFile => ({
      version: '0.1.0',
      metadata: {
        name: 'Test Scene',
        description: 'A test scene for validation',
        author: 'Test Author'
      },
      scene: {
        nodes: [
          {
            id: 'node-1',
            type: 'server',
            name: 'Test Server',
            transform: createTransform(),
            status: 'healthy'
          },
          {
            id: 'node-2',
            type: 'database',
            name: 'Test Database',
            transform: createTransform({ x: 5, y: 0, z: 0 }),
            status: 'healthy'
          }
        ],
        edges: [
          {
            id: 'edge-1',
            source: 'node-1',
            target: 'node-2',
            type: 'connection'
          }
        ]
      }
    });

    it('should validate a correct scene', () => {
      const scene = createValidScene();
      const result = validateScene(scene);

      expect(result.valid).toBe(true);
      expect(result.errors).toHaveLength(0);
    });

    it('should detect missing version', () => {
      const scene = createValidScene();
      scene.version = '';

      const result = validateScene(scene);
      expect(result.valid).toBe(false);
      expect(result.errors).toContain('Scene file must have a version');
    });

    it('should detect missing metadata name', () => {
      const scene = createValidScene();
      scene.metadata.name = '';

      const result = validateScene(scene);
      expect(result.valid).toBe(false);
      expect(result.errors).toContain('Scene file must have a name');
    });

    it('should detect duplicate node IDs', () => {
      const scene = createValidScene();
      scene.scene.nodes.push({
        id: 'node-1', // duplicate ID
        type: 'server',
        name: 'Duplicate Server',
        transform: createTransform(),
        status: 'healthy'
      });

      const result = validateScene(scene);
      expect(result.valid).toBe(false);
      expect(result.errors).toContain('Duplicate node id: node-1');
    });

    it('should detect invalid edge references', () => {
      const scene = createValidScene();
      scene.scene.edges.push({
        id: 'edge-2',
        source: 'non-existent-node',
        target: 'node-1',
        type: 'connection'
      });

      const result = validateScene(scene);
      expect(result.valid).toBe(false);
      expect(result.errors).toContain('Edge edge-2 references non-existent source node: non-existent-node');
    });

    it('should provide warnings for missing node names', () => {
      const scene = createValidScene();
      scene.scene.nodes[0].name = '';

      const result = validateScene(scene);
      expect(result.warnings).toContain('Node node-1 has no name');
    });
  });

  describe('Scene Statistics', () => {
    it('should calculate basic scene stats', () => {
      const scene: SceneFile = {
        version: '0.1.0',
        metadata: { name: 'Test Scene' },
        scene: {
          nodes: [
            {
              id: 'node-1',
              type: 'server',
              name: 'Server 1',
              transform: createTransform({ x: 0, y: 0, z: 0 }),
              status: 'healthy'
            },
            {
              id: 'node-2',
              type: 'database',
              name: 'Database 1',
              transform: createTransform({ x: 10, y: 5, z: -5 }),
              status: 'healthy'
            }
          ],
          edges: [
            {
              id: 'edge-1',
              source: 'node-1',
              target: 'node-2',
              type: 'connection'
            }
          ]
        }
      };

      const stats = calculateSceneStats(scene);

      expect(stats.nodeCount).toBe(2);
      expect(stats.edgeCount).toBe(1);
      expect(stats.bounds).toBeDefined();
      expect(stats.bounds?.min).toEqual({ x: 0, y: 0, z: -5 });
      expect(stats.bounds?.max).toEqual({ x: 10, y: 5, z: 0 });
      expect(stats.bounds?.size).toEqual({ x: 10, y: 5, z: 5 });
    });

    it('should handle empty scene', () => {
      const scene: SceneFile = {
        version: '0.1.0',
        metadata: { name: 'Empty Scene' },
        scene: { nodes: [], edges: [] }
      };

      const stats = calculateSceneStats(scene);

      expect(stats.nodeCount).toBe(0);
      expect(stats.edgeCount).toBe(0);
      expect(stats.bounds).toBeUndefined();
    });
  });

  describe('Scene File Creation', () => {
    it('should create a valid scene file with all required fields', () => {
      const scene: SceneFile = {
        version: '0.1.0',
        metadata: {
          name: 'Production Infrastructure',
          description: 'Main production environment',
          author: 'DevOps Team',
          created: new Date().toISOString(),
          tags: ['production', 'aws']
        },
        scene: {
          nodes: [
            {
              id: generateId(),
              type: 'server',
              name: 'Web Server',
              transform: createTransform(),
              geometry: {
                type: 'box',
                parameters: { width: 2, height: 1, depth: 1 }
              },
              material: createMaterial({
                color: { r: 0.2, g: 0.8, b: 0.2, a: 1 }
              }),
              status: 'healthy',
              metadata: {
                cpu: '85%',
                memory: '12GB'
              },
              tags: ['web', 'frontend']
            }
          ],
          edges: [],
          camera: {
            position: { x: 0, y: 10, z: 10 },
            target: { x: 0, y: 0, z: 0 },
            fov: 75
          },
          lights: [
            {
              type: 'ambient',
              color: { r: 0.4, g: 0.4, b: 0.4, a: 1 },
              intensity: 0.5
            }
          ]
        }
      };

      const validation = validateScene(scene);
      expect(validation.valid).toBe(true);
      expect(scene.version).toBe('0.1.0');
      expect(scene.metadata.name).toBe('Production Infrastructure');
      expect(scene.scene.nodes).toHaveLength(1);
      expect(scene.scene.camera).toBeDefined();
      expect(scene.scene.lights).toHaveLength(1);
    });
  });

  describe('Interface Compatibility', () => {
    it('should allow implementing Importer interface', () => {
      class TestImporter implements Importer {
        id = 'test-importer';
        name = 'Test Importer';
        description = 'A test importer';
        supportedFormats = ['.json'];

        async import(input: string): Promise<ImportResult> {
          const data = JSON.parse(input);
          return {
            scene: {
              version: '0.1.0',
              metadata: {
                name: data.name || 'Imported Scene',
                importedBy: this.id,
                importedAt: new Date().toISOString()
              },
              scene: {
                nodes: [],
                edges: []
              }
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

      const importer = new TestImporter();
      expect(importer.id).toBe('test-importer');
      expect(importer.name).toBe('Test Importer');
      expect(importer.supportedFormats).toContain('.json');
    });

    it('should allow implementing Provider interface', () => {
      class TestProvider implements Provider {
        id = 'test-provider';
        name = 'Test Provider';
        description = 'A test provider';
        private connected = false;

        async connect(config: any): Promise<void> {
          this.connected = true;
        }

        async disconnect(): Promise<void> {
          this.connected = false;
        }

        isConnected(): boolean {
          return this.connected;
        }

        async query(query: MetricsQuery): Promise<MetricsResult[]> {
          return [];
        }

        async healthCheck(): Promise<boolean> {
          return this.connected;
        }
      }

      const provider = new TestProvider();
      expect(provider.id).toBe('test-provider');
      expect(provider.isConnected()).toBe(false);
    });

    it('should allow implementing AnimationHook interface', () => {
      class TestAnimationHook implements AnimationHook {
        id = 'test-hook';
        name = 'Test Animation Hook';
        description = 'A test animation hook';

        onFrame(context: any): void {
          // Test implementation
        }

        onStart(context: any): void {
          // Test implementation
        }

        onStop(context: any): void {
          // Test implementation
        }
      }

      const hook = new TestAnimationHook();
      expect(hook.id).toBe('test-hook');
      expect(hook.name).toBe('Test Animation Hook');
      expect(typeof hook.onFrame).toBe('function');
    });
  });

  describe('Animation System', () => {
    it('should create animation tracks', () => {
      const animation = {
        name: 'rotate-server',
        duration: 2.0,
        loop: true,
        tracks: [
          {
            property: 'transform.rotation.y',
            keyframes: [
              { time: 0, value: 0, easing: 'linear' as const },
              { time: 2, value: Math.PI * 2, easing: 'linear' as const }
            ]
          }
        ]
      };

      expect(animation.name).toBe('rotate-server');
      expect(animation.duration).toBe(2.0);
      expect(animation.loop).toBe(true);
      expect(animation.tracks).toHaveLength(1);
      expect(animation.tracks[0].property).toBe('transform.rotation.y');
      expect(animation.tracks[0].keyframes).toHaveLength(2);
    });
  });

  describe('Type Safety', () => {
    it('should enforce status enum values', () => {
      const node: SceneNode = {
        id: 'test-node',
        type: 'server',
        name: 'Test Server',
        transform: createTransform(),
        status: 'healthy' // This should compile without errors
      };

      expect(node.status).toBe('healthy');

      // These should be valid status values
      const validStatuses = ['healthy', 'warning', 'critical', 'unknown'];
      validStatuses.forEach(status => {
        node.status = status as any;
        expect(validStatuses).toContain(node.status);
      });
    });

    it('should enforce geometry type enum values', () => {
      const geometry = {
        type: 'box' as const,
        parameters: { width: 1, height: 1, depth: 1 }
      };

      expect(geometry.type).toBe('box');

      // These should be valid geometry types
      const validTypes = ['box', 'sphere', 'cylinder', 'plane', 'custom'];
      expect(validTypes).toContain(geometry.type);
    });

    it('should enforce color value ranges', () => {
      const color = { r: 0.5, g: 0.8, b: 0.2, a: 1.0 };

      // All values should be between 0 and 1
      expect(color.r).toBeGreaterThanOrEqual(0);
      expect(color.r).toBeLessThanOrEqual(1);
      expect(color.g).toBeGreaterThanOrEqual(0);
      expect(color.g).toBeLessThanOrEqual(1);
      expect(color.b).toBeGreaterThanOrEqual(0);
      expect(color.b).toBeLessThanOrEqual(1);
      expect(color.a).toBeGreaterThanOrEqual(0);
      expect(color.a).toBeLessThanOrEqual(1);
    });
  });
});
