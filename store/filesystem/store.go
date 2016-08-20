// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package filesystem provides an upspin.StoreServer
// that serves files from a local file system.
// It must be used in conjuction with the upspin.DirServer
// implementation in package upspin.io/dir/filesystem.
package filesystem

import (
	"upspin.io/access"
	"upspin.io/errors"
	"upspin.io/log"
	"upspin.io/path"
	"upspin.io/upspin"
)

// New creates a new StoreServer instance with the
// provided server context and configuration options.
// The only valid configuration option is "root", which
// specifies a path to the file system root.
func New(ctx upspin.Context, options ...string) (upspin.StoreServer, error) {
	const op = "store/filesystem.New"

	s := &server{server: ctx}

	var err error
	s.root, s.defaultAccess, err = newRoot(ctx, options)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return s, nil
}

type server struct {
	// Set by New.
	root          string
	server        upspin.Context
	defaultAccess *access.Access

	// Set by Dial.
	user upspin.Context
}

var errNotDialed = errors.E(errors.Internal, errors.Str("must Dial before making request"))

func (s *server) Get(ref upspin.Reference) ([]byte, []upspin.Location, error) {
	const op = "store/filesystem.Get"
	log.Debug.Println(op, ref)

	if s.user == nil {
		return nil, nil, errors.E(op, errNotDialed)
	}

	pathName := upspin.PathName(s.server.UserName()) + "/" + upspin.PathName(ref)
	parsed, err := path.Parse(pathName)
	if err != nil {
		return nil, nil, errors.E(op, err)
	}

	// Verify that the requesting user can access this file.
	if ok, err := can(s.root, s.defaultAccess, s.user.UserName(), access.Read, parsed); err != nil {
		return nil, nil, errors.E(op, err)
	} else if !ok {
		return nil, nil, errors.E(op, parsed.Path(), access.ErrPermissionDenied)
	}

	data, err := readFile(s.root, pathName)
	if err != nil {
		return nil, nil, errors.E(op, err)
	}
	return data, nil, nil
}

func (s *server) Dial(ctx upspin.Context, e upspin.Endpoint) (upspin.Service, error) {
	const op = "store/filesystem.Dial"

	dialed := *s
	dialed.user = ctx.Copy()
	return &dialed, nil
}

func (s *server) Ping() bool {
	return true
}

func (s *server) Close() {
}

// Methods that are not implemented.

var errNotImplemented = errors.Str("not implemented")

func (s *server) Put(ciphertext []byte) (upspin.Reference, error) {
	return "", errors.E("store/filesystem.Put", errNotImplemented)
}

func (s *server) Delete(ref upspin.Reference) error {
	return errors.E("store/filesystem.Delete", errNotImplemented)
}

// Methods that do not apply to this server.

func (s *server) Configure(options ...string) error {
	return errors.Str("store/filesystem: Configure should not be called")
}

func (s *server) Authenticate(upspin.Context) error {
	return errors.Str("store/filesystem: Authenticate should not be called")
}

func (s *server) Endpoint() upspin.Endpoint {
	return upspin.Endpoint{} // No endpoint.
}