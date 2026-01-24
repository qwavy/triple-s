package services

import "regexp"

type Bucket struct {
	Name string
}

var nameRegex = regexp.MustCompile(`^[a-z0-9.-]{3,63}$`)

func (b *Bucket) CreateBucket(name string) error {
	if nameRegex.MatchString(name)
}
