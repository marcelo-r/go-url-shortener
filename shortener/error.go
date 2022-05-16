package shortener

import "errors"

// ErrEmptyRepository signals it tried to create a repository with an empty
// repository
var ErrEmptyRepository = errors.New("unable to use empty repository")
