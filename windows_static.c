// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This stub is to fix a compilation error with Go 1.2
//
// It is likely that we can remove this once the external linker is used on
// Windows in Go 1.3, right now this is just kind of a hack to get it to
// compile OK with Go 1.2.
int _get_output_format( void ){ return 0; }
