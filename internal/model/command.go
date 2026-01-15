package model

import (
	"time"
)

type Command struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	Cmd       string     `db:"cmd"`
	Cwd       *string    `db:"cwd"`
	Tags      string     `db:"tags"`
	Confirm   bool       `db:"confirm"`
	Frequency int        `db:"frequency"`
	LastUsed  *time.Time `db:"last_used"`
	CreatedAt time.Time  `db:"created_at"`
}

func (c *Command) HasTag(tag string) bool {
	if c.Tags == "" {
		return false
	}
	tags := splitTags(c.Tags)
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func (c *Command) IsDangerous() bool {
	return c.HasTag("danger")
}

func splitTags(tags string) []string {
	if tags == "" {
		return nil
	}
	var result []string
	current := ""
	for _, ch := range tags {
		if ch == ',' {
			if current != "" {
				result = append(result, current)
			}
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}
