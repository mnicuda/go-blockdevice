// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package blockdevice

// Options is the functional options struct.
type Options struct {
	CreateGPT     bool
	ExclusiveLock bool
	Mode          int
}

// Option is the functional option func.
type Option func(*Options)

// WithNewGPT opens the blockdevice with a new GPT.
func WithNewGPT(o bool) Option {
	return func(args *Options) {
		args.CreateGPT = o
	}
}

// WithExclusiveLock locks the blockdevice for exclusive access using flock().
func WithExclusiveLock(o bool) Option {
	return func(args *Options) {
		args.ExclusiveLock = o
	}
}

// WithMode opens blockdevice in a specific mode.
func WithMode(value int) Option {
	return func(args *Options) {
		args.Mode = value
	}
}

// NewDefaultOptions initializes a Options struct with default values.
func NewDefaultOptions(setters ...Option) *Options {
	opts := &Options{
		CreateGPT: false,
		Mode:      DefaultMode,
	}

	for _, setter := range setters {
		setter(opts)
	}

	return opts
}
