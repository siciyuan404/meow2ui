package a2ui

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
)

type ValidationResult struct {
	Valid  bool
	Errors []ValidationError
}

type ValidationError struct {
	Code    string `json:"code"`
	Path    string `json:"path"`
	Message string `json:"message"`
}

type ComponentRule struct {
	Name      string
	PropTypes map[string]string
}

type Registry struct {
	rules map[string]ComponentRule
}

func NewRegistry() *Registry {
	return &Registry{rules: map[string]ComponentRule{}}
}

func (r *Registry) Register(rule ComponentRule) {
	r.rules[rule.Name] = rule
}

func (r *Registry) Rule(name string) (ComponentRule, bool) {
	rule, ok := r.rules[name]
	return rule, ok
}

type Service struct {
	registry *Registry
}

func NewService() *Service {
	r := NewRegistry()
	r.Register(ComponentRule{Name: "Container", PropTypes: map[string]string{"className": "string", "title": "string"}})
	r.Register(ComponentRule{Name: "Text", PropTypes: map[string]string{"content": "string", "className": "string"}})
	r.Register(ComponentRule{Name: "Button", PropTypes: map[string]string{"label": "string", "disabled": "boolean"}})
	r.Register(ComponentRule{Name: "Card", PropTypes: map[string]string{"title": "string", "className": "string"}})
	return &Service{registry: r}
}

func (s *Service) ValidateSchema(schema domain.UISchema) ValidationResult {
	errs := []ValidationError{}
	if strings.TrimSpace(schema.Version) == "" {
		errs = append(errs, ValidationError{Code: "A2UI_SCHEMA_REQUIRED_FIELD_MISSING", Path: "schema.version", Message: "schema.version is required"})
	}
	if strings.TrimSpace(schema.Root.ID) == "" {
		errs = append(errs, ValidationError{Code: "A2UI_SCHEMA_REQUIRED_FIELD_MISSING", Path: "root.id", Message: "root.id is required"})
	}
	if strings.TrimSpace(schema.Root.Type) == "" {
		errs = append(errs, ValidationError{Code: "A2UI_SCHEMA_REQUIRED_FIELD_MISSING", Path: "root.type", Message: "root.type is required"})
	}
	s.validateComponent(schema.Root, "root", &errs)
	return ValidationResult{Valid: len(errs) == 0, Errors: errs}
}

func (s *Service) validateComponent(c domain.UIComponent, path string, errs *[]ValidationError) {
	if strings.TrimSpace(c.ID) == "" {
		*errs = append(*errs, ValidationError{Code: "A2UI_SCHEMA_REQUIRED_FIELD_MISSING", Path: fmt.Sprintf("%s.id", path), Message: fmt.Sprintf("%s.id is required", path)})
	}
	if strings.TrimSpace(c.Type) == "" {
		*errs = append(*errs, ValidationError{Code: "A2UI_SCHEMA_REQUIRED_FIELD_MISSING", Path: fmt.Sprintf("%s.type", path), Message: fmt.Sprintf("%s.type is required", path)})
	} else {
		rule, ok := s.registry.Rule(c.Type)
		if !ok {
			*errs = append(*errs, ValidationError{Code: "A2UI_COMPONENT_NOT_ALLOWED", Path: fmt.Sprintf("%s.type", path), Message: fmt.Sprintf("component type not allowed: %s", c.Type)})
		} else {
			s.validateProps(c.Props, rule, path, errs)
		}
	}
	for i := range c.Children {
		s.validateComponent(c.Children[i], fmt.Sprintf("%s.children[%d]", path, i), errs)
	}
}

func (s *Service) validateProps(props map[string]any, rule ComponentRule, path string, errs *[]ValidationError) {
	if props == nil {
		return
	}
	for key, value := range props {
		expected, ok := rule.PropTypes[key]
		if !ok {
			continue
		}
		if !matchesType(value, expected) {
			*errs = append(*errs, ValidationError{Code: "A2UI_PROP_TYPE_INVALID", Path: fmt.Sprintf("%s.props.%s", path, key), Message: fmt.Sprintf("prop %s expects %s", key, expected)})
		}
	}
}

func matchesType(v any, expected string) bool {
	switch expected {
	case "string":
		_, ok := v.(string)
		return ok
	case "number":
		switch v.(type) {
		case int, int32, int64, float32, float64:
			return true
		default:
			return false
		}
	case "boolean":
		_, ok := v.(bool)
		return ok
	case "object":
		_, ok := v.(map[string]any)
		return ok
	case "array":
		_, ok := v.([]any)
		return ok
	default:
		return true
	}
}

func (s *Service) ParseSchema(jsonText string) (domain.UISchema, error) {
	var schema domain.UISchema
	if err := json.Unmarshal([]byte(jsonText), &schema); err != nil {
		return domain.UISchema{}, err
	}
	return schema, nil
}

func (s *Service) ApplyPatch(base domain.UISchema, patch domain.PatchDocument) (domain.UISchema, error) {
	out := base
	for _, op := range patch.Operations {
		switch strings.ToLower(op.Op) {
		case "update_root_props":
			if out.Root.Props == nil {
				out.Root.Props = map[string]any{}
			}
			for k, v := range op.Value {
				out.Root.Props[k] = v
			}
		case "add_root_child":
			child := domain.UIComponent{}
			b, _ := json.Marshal(op.Value)
			if err := json.Unmarshal(b, &child); err != nil {
				return domain.UISchema{}, err
			}
			out.Root.Children = append(out.Root.Children, child)
		default:
			return domain.UISchema{}, fmt.Errorf("unsupported patch op: %s", op.Op)
		}
	}
	return out, nil
}
