package reader

import "time"

type Dir []Item

type Item struct {
	Name         string
	LastModified time.Time
	IsDir        bool
}
