// Copyright 2020 cetc-30. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.



package sm3

//go:noescape

func block(dig *digest, p []byte)