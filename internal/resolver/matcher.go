package resolver

import (
	"math"
	"sort"
	"strings"
	"time"

	"github.com/sahilm/fuzzy"
	"github.com/verse/apotheke/internal/model"
)

type Match struct {
	Command *model.Command
	Score   int
	Matched string
}

type Resolver struct{}

func New() *Resolver {
	return &Resolver{}
}

type commandSource []*model.Command

func (c commandSource) String(i int) string {
	return c[i].Name
}

func (c commandSource) Len() int {
	return len(c)
}

func (r *Resolver) Resolve(query string, commands []*model.Command) []Match {
	if len(commands) == 0 {
		return nil
	}

	var matches []Match

	for _, cmd := range commands {
		if cmd.Name == query {
			matches = append(matches, Match{
				Command: cmd,
				Score:   1000000,
				Matched: cmd.Name,
			})
			return matches
		}
	}

	var prefixMatches []Match
	for _, cmd := range commands {
		if strings.HasPrefix(cmd.Name, query) {
			score := 100000 - len(cmd.Name) + frecencyScore(cmd)
			prefixMatches = append(prefixMatches, Match{
				Command: cmd,
				Score:   score,
				Matched: cmd.Name,
			})
		}
	}

	if len(prefixMatches) > 0 {
		sort.Slice(prefixMatches, func(i, j int) bool {
			return prefixMatches[i].Score > prefixMatches[j].Score
		})
		return prefixMatches
	}

	source := commandSource(commands)
	results := fuzzy.FindFrom(query, source)

	for _, result := range results {
		cmd := commands[result.Index]
		score := result.Score + frecencyScore(cmd)
		matches = append(matches, Match{
			Command: cmd,
			Score:   score,
			Matched: cmd.Name,
		})
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Score > matches[j].Score
	})

	return matches
}

func frecencyScore(cmd *model.Command) int {
	if cmd.LastUsed == nil {
		return cmd.Frequency
	}

	hoursSinceUse := time.Since(*cmd.LastUsed).Hours()
	recency := 1.0 / (1.0 + hoursSinceUse/24.0)
	score := float64(cmd.Frequency) * recency * 10

	return int(math.Min(score, 10000))
}
