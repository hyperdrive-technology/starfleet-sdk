# Starfleet SDK - TODO

**Repository Status**: 🟡 PARTIALLY COMPLETE - Core implementation exists, missing publishing & validation

## 📋 Completed ✅

- ✅ Core TypeScript interfaces and types (SceneNode, SceneEdge, SceneFile)
- ✅ Plugin interfaces (Importer, Provider, AnimationHook)
- ✅ Utility functions (createTransform, validateScene, etc.)
- ✅ TypeScript package structure with proper exports
- ✅ Go bindings and structs in `go/models.go`
- ✅ JSON Schema in `schema/scenefile.schema.json`
- ✅ Basic test structure in `ts/src/index.test.ts`
- ✅ Build configuration with tsup
- ✅ Examples directory with basic usage
- ✅ README documentation

## 🔄 In Progress / Pending

### Critical - Publishing & CI/CD 🚨
- [ ] **GitHub Actions workflow for automated publishing**
  - [ ] TypeScript package publish to NPM on tag
  - [ ] Go module publishing (git tags)
  - [ ] Automated testing on PR/push
  - [ ] Lint checking and type validation
- [ ] **Version management strategy**
  - [ ] Semantic versioning policy
  - [ ] CHANGELOG.md generation
  - [ ] Release notes automation

### Testing & Validation 🧪
- [ ] **Comprehensive test suite**
  - [ ] Unit tests for utility functions
  - [ ] Validation schema tests
  - [ ] JSON Schema compliance tests
  - [ ] Go structs validation tests
- [ ] **Test coverage reporting**
  - [ ] TypeScript coverage with vitest
  - [ ] Go coverage with go test
  - [ ] Coverage badges and reporting

### Documentation 📚
- [ ] **API documentation generation**
  - [ ] TypeScript API docs with TypeDoc
  - [ ] Go documentation with godoc
  - [ ] Hosted documentation site
- [ ] **Usage examples**
  - [ ] More comprehensive examples beyond basic-usage
  - [ ] Integration examples with importers
  - [ ] Provider implementation examples

### Package Management 📦
- [ ] **NPM package publishing**
  - [ ] First v0.1.0 release to NPM
  - [ ] Package verification and testing
  - [ ] Public access configuration
- [ ] **Go module setup**
  - [ ] Module proxy registration
  - [ ] Version tagging strategy
  - [ ] Dependencies management

### JSON Schema Enhancements 🔧
- [ ] **Schema validation**
  - [ ] Runtime schema validation
  - [ ] Schema version compatibility
  - [ ] Migration utilities for schema updates
- [ ] **Code generation from schema**
  - [ ] Automated TypeScript types from JSON Schema
  - [ ] Automated Go structs from JSON Schema
  - [ ] Schema update synchronization

### Protocol Buffers (Optional) 🔌
- [ ] **Proto definitions** (if gRPC support needed)
  - [ ] Scene proto definitions
  - [ ] Code generation for multiple languages
  - [ ] Backward compatibility management

## 🎯 Next Actions (Priority Order)

1. **🚨 HIGH: Set up GitHub Actions for publishing**
   ```yaml
   # .github/workflows/publish.yml
   - TypeScript build & test & publish on tag
   - Go module tag creation and validation
   ```

2. **🚨 HIGH: Publish v0.1.0 to NPM**
   ```bash
   cd ts
   pnpm build
   npm publish --access public
   ```

3. **🔧 MEDIUM: Expand test coverage**
   - Add comprehensive unit tests for all utility functions
   - Add integration tests for common usage patterns

4. **📚 MEDIUM: Generate API documentation**
   - Set up TypeDoc for TypeScript API docs
   - Create hosted documentation site

5. **🔌 LOW: Protocol Buffer support** (only if needed for gRPC)

## 🔗 Dependencies & Integration

### Downstream Dependencies
- `starfleet-importer-brainboard` ✅ (already using @starfleet/sdk)
- `@starfleet/cli` 🔄 (needs @starfleet/sdk dependency)
- `@starfleet/viewer-2d-react` 🔄 (needs @starfleet/sdk dependency)

### Blocking Issues
- **NPM Publishing**: Required for other packages to use published version
- **Version Management**: Needed for stable dependency management across ecosystem

## 📊 Success Criteria

- [ ] Published @starfleet/sdk@0.1.0 on NPM
- [ ] CI/CD pipeline working (tests pass, auto-publish on tags)
- [ ] Test coverage >80% for critical functions
- [ ] All downstream packages using published SDK version
- [ ] Documentation site live and accessible
- [ ] Go module properly tagged and accessible

## 🔧 Development Commands

```bash
# Build and test
cd ts
pnpm install
pnpm build
pnpm test

# Publish (once ready)
pnpm version patch
pnpm publish --access public

# Go module testing
cd go
go mod tidy
go test ./...
```

## 📝 Notes

- The SDK is the foundation for the entire Starfleet ecosystem
- Must maintain strict semantic versioning due to multiple dependents
- TypeScript types should be the source of truth, with Go/JSON Schema generated
- Keep minimal dependencies to avoid conflicts in consuming packages
- Focus on stability and backward compatibility
