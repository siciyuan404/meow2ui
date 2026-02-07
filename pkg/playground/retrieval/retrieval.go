package retrieval

import (
	"context"
	"sort"
	"strings"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
)

type RetrievalQuery struct {
	Text       string
	ThemeID    string
	CategoryID string
	Tags       []string
	Limit      int
	MinScore   float64
}

type RetrievalHit struct {
	ItemID  string
	Title   string
	Tags    []string
	Summary string
	Score   float64
}

type Retriever interface {
	Retrieve(context.Context, RetrievalQuery) ([]RetrievalHit, error)
}

type KeywordRetriever struct {
	repo store.PlaygroundRepository
}

func NewKeywordRetriever(repo store.PlaygroundRepository) *KeywordRetriever {
	return &KeywordRetriever{repo: repo}
}

func (r *KeywordRetriever) Retrieve(ctx context.Context, q RetrievalQuery) ([]RetrievalHit, error) {
	if q.Limit <= 0 {
		q.Limit = 3
	}
	items, err := r.repo.SearchItems(ctx, q.CategoryID, q.Tags, q.Text)
	if err != nil {
		return nil, err
	}
	hits := make([]RetrievalHit, 0, len(items))
	for _, it := range items {
		score := scoreItem(it, q)
		if score < q.MinScore {
			continue
		}
		hits = append(hits, RetrievalHit{
			ItemID:  it.ID,
			Title:   it.Title,
			Tags:    append([]string(nil), it.Tags...),
			Summary: summarize(it.SchemaSnapshot),
			Score:   score,
		})
	}
	sort.Slice(hits, func(i, j int) bool { return hits[i].Score > hits[j].Score })
	if len(hits) > q.Limit {
		hits = hits[:q.Limit]
	}
	return hits, nil
}

func scoreItem(it domain.PlaygroundItem, q RetrievalQuery) float64 {
	score := 0.0
	text := strings.ToLower(strings.TrimSpace(q.Text))
	if text != "" && strings.Contains(strings.ToLower(it.Title), text) {
		score += 3
	}
	if q.ThemeID != "" && it.ThemeID == q.ThemeID {
		score += 2
	}
	if q.CategoryID != "" && it.CategoryID == q.CategoryID {
		score += 1
	}
	qTags := map[string]struct{}{}
	for _, t := range q.Tags {
		qTags[t] = struct{}{}
	}
	for _, t := range it.Tags {
		if _, ok := qTags[t]; ok {
			score += 2
		}
	}
	return score
}

func summarize(schema string) string {
	s := strings.TrimSpace(schema)
	if len(s) <= 120 {
		return s
	}
	return s[:120]
}
