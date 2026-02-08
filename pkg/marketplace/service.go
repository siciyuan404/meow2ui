package marketplace

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Template struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Tags       []string  `json:"tags"`
	Version    string    `json:"version"`
	Schema     string    `json:"schema"`
	Theme      string    `json:"theme"`
	Status     string    `json:"status"`
	Owner      string    `json:"owner"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ReviewNote string    `json:"reviewNote,omitempty"`
}

type Rating struct {
	TemplateID string    `json:"templateId"`
	UserID     string    `json:"userId"`
	Score      int       `json:"score"`
	Comment    string    `json:"comment,omitempty"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}

type SearchInput struct {
	Category string
	Tag      string
	Query    string
}

type SaveFromSessionInput struct {
	Name       string
	Category   string
	Tags       []string
	Schema     string
	Theme      string
	Owner      string
	SessionID  string
	VersionID  string
	ChangeNote string
}

type Service struct {
	mu      sync.RWMutex
	tpls    map[string]Template
	ratings map[string][]Rating
	history map[string][]string
}

func NewService() *Service {
	return &Service{
		tpls:    map[string]Template{},
		ratings: map[string][]Rating{},
		history: map[string][]string{},
	}
}

func (s *Service) SaveFromSession(_ context.Context, in SaveFromSessionInput) (Template, error) {
	if strings.TrimSpace(in.Name) == "" {
		return Template{}, fmt.Errorf("name is required")
	}
	if strings.TrimSpace(in.Schema) == "" {
		return Template{}, fmt.Errorf("schema is required")
	}
	now := time.Now()
	t := Template{
		ID:        util.NewID("tpl"),
		Name:      in.Name,
		Category:  in.Category,
		Tags:      append([]string(nil), in.Tags...),
		Version:   "v1",
		Schema:    in.Schema,
		Theme:     in.Theme,
		Status:    "draft",
		Owner:     in.Owner,
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tpls[t.ID] = t
	s.history[t.ID] = []string{fmt.Sprintf("v1: %s", strings.TrimSpace(in.ChangeNote))}
	return t, nil
}

func (s *Service) SubmitReview(_ context.Context, templateID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tpls[templateID]
	if !ok {
		return fmt.Errorf("template not found")
	}
	t.Status = "submitted"
	t.UpdatedAt = time.Now()
	s.tpls[templateID] = t
	return nil
}

func (s *Service) Review(_ context.Context, templateID, decision, note string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tpls[templateID]
	if !ok {
		return fmt.Errorf("template not found")
	}
	switch decision {
	case "publish":
		t.Status = "published"
	case "block":
		t.Status = "blocked"
	default:
		return fmt.Errorf("invalid decision")
	}
	t.ReviewNote = note
	t.UpdatedAt = time.Now()
	s.tpls[templateID] = t
	return nil
}

func (s *Service) AddVersion(_ context.Context, templateID, schema, changeNote string) (Template, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tpls[templateID]
	if !ok {
		return Template{}, fmt.Errorf("template not found")
	}
	n := len(s.history[templateID]) + 1
	t.Version = fmt.Sprintf("v%d", n)
	t.Schema = schema
	t.UpdatedAt = time.Now()
	s.tpls[templateID] = t
	s.history[templateID] = append(s.history[templateID], fmt.Sprintf("v%d: %s", n, changeNote))
	return t, nil
}

func (s *Service) Search(_ context.Context, in SearchInput) []Template {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Template, 0, len(s.tpls))
	q := strings.ToLower(strings.TrimSpace(in.Query))
	for _, t := range s.tpls {
		if in.Category != "" && t.Category != in.Category {
			continue
		}
		if in.Tag != "" && !contains(t.Tags, in.Tag) {
			continue
		}
		if q != "" && !strings.Contains(strings.ToLower(t.Name), q) {
			continue
		}
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.After(out[j].UpdatedAt) })
	return out
}

func (s *Service) AddRating(_ context.Context, templateID, userID string, score int, comment string) (Rating, error) {
	if score < 1 || score > 5 {
		return Rating{}, fmt.Errorf("score must be 1-5")
	}
	r := Rating{TemplateID: templateID, UserID: userID, Score: score, Comment: comment, Status: "visible", CreatedAt: time.Now()}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tpls[templateID]; !ok {
		return Rating{}, fmt.Errorf("template not found")
	}
	s.ratings[templateID] = append(s.ratings[templateID], r)
	return r, nil
}

func (s *Service) FlagComment(_ context.Context, templateID, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	items := s.ratings[templateID]
	for i := range items {
		if items[i].UserID == userID {
			items[i].Status = "flagged"
			s.ratings[templateID] = items
			return nil
		}
	}
	return fmt.Errorf("comment not found")
}

func (s *Service) ApplyTemplate(_ context.Context, templateID, sessionID string) (map[string]any, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tpls[templateID]
	if !ok {
		return nil, fmt.Errorf("template not found")
	}
	if t.Theme == "" {
		return nil, fmt.Errorf("missing theme dependency")
	}
	return map[string]any{"sessionId": sessionID, "templateId": templateID, "version": t.Version, "status": "applied"}, nil
}

func contains(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
