// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package freetype

import (
	"errors"
)

var (
	ErrCannotOpenResource          = errors.New("cannot open resource")
	ErrUnknownFileFormat           = errors.New("unknown file format")
	ErrInvalidFileFormat           = errors.New("broken file")
	ErrInvalidVersion              = errors.New("invalid FreeType version")
	ErrLowerModuleVersion          = errors.New("module version is too low")
	ErrInvalidArgument             = errors.New("invalid argument")
	ErrUnimplementedFeature        = errors.New("unimplemented feature")
	ErrInvalidTable                = errors.New("broken table")
	ErrInvalidOffset               = errors.New("broken offset within table")
	ErrArrayTooLarge               = errors.New("array allocation size too large")
	ErrMissingModule               = errors.New("missing module")
	ErrMissingProperty             = errors.New("missing property")
	ErrInvalidGlyphIndex           = errors.New("invalid glyph index")
	ErrInvalidCharacterCode        = errors.New("invalid character code")
	ErrInvalidGlyphFormat          = errors.New("unsupported glyph image format")
	ErrCannotRenderGlyph           = errors.New("cannot render this glyph format")
	ErrInvalidOutline              = errors.New("invalid outline")
	ErrInvalidComposite            = errors.New("invalid composite glyph")
	ErrTooManyHints                = errors.New("too many hints")
	ErrInvalidPixelSize            = errors.New("invalid pixel size")
	ErrInvalidHandle               = errors.New("invalid object handle")
	ErrInvalidLibraryHandle        = errors.New("invalid library handle")
	ErrInvalidDriverHandle         = errors.New("invalid module handle")
	ErrInvalidFaceHandle           = errors.New("invalid face handle")
	ErrInvalidSizeHandle           = errors.New("invalid size handle")
	ErrInvalidSlotHandle           = errors.New("invalid glyph slot handle")
	ErrInvalidCharMapHandle        = errors.New("invalid charmap handle")
	ErrInvalidCacheHandle          = errors.New("invalid cache manager handle")
	ErrInvalidStreamHandle         = errors.New("invalid stream handle")
	ErrTooManyDrivers              = errors.New("too many modules")
	ErrTooManyExtensions           = errors.New("too many extensions")
	ErrOutOfMemory                 = errors.New("out of memory")
	ErrUnlistedObject              = errors.New("unlisted object")
	ErrCannotOpenStream            = errors.New("cannot open stream")
	ErrInvalidStreamSeek           = errors.New("invalid stream seek")
	ErrInvalidStreamSkip           = errors.New("invalid stream skip")
	ErrInvalidStreamRead           = errors.New("invalid stream read")
	ErrInvalidStreamOperation      = errors.New("invalid stream operation")
	ErrInvalidFrameOperation       = errors.New("invalid frame operation")
	ErrNestedFrameAccess           = errors.New("nested frame access")
	ErrInvalidFrameRead            = errors.New("invalid frame read")
	ErrRasterUninitialized         = errors.New("raster uninitialized")
	ErrRasterCorrupted             = errors.New("raster corrupted")
	ErrRasterOverflow              = errors.New("raster overflow")
	ErrRasterNegativeHeight        = errors.New("negative height while rastering")
	ErrTooManyCaches               = errors.New("too many registered caches")
	ErrInvalidOpcode               = errors.New("invalid opcode")
	ErrTooFewArguments             = errors.New("too few arguments")
	ErrStackOverflow               = errors.New("stack overflow")
	ErrCodeOverflow                = errors.New("code overflow")
	ErrBadArgument                 = errors.New("bad argument")
	ErrDivideByZero                = errors.New("division by zero")
	ErrInvalidReference            = errors.New("invalid reference")
	ErrDebugOpCode                 = errors.New("found debug opcode")
	ErrENDFInExecStream            = errors.New("found ENDF opcode in execution stream")
	ErrNestedDEFS                  = errors.New("nested DEFS")
	ErrInvalidCodeRange            = errors.New("invalid code range")
	ErrExecutionTooLong            = errors.New("execution context too long")
	ErrTooManyFunctionDefs         = errors.New("too many function definitions")
	ErrTooManyInstructionDefs      = errors.New("too many instruction definitions")
	ErrTableMissing                = errors.New("SFNT font table missing")
	ErrHorizHeaderMissing          = errors.New("horizontal header (hhea) table missing")
	ErrLocationsMissing            = errors.New("locations (loca) table missing")
	ErrNameTableMissing            = errors.New("name table missing")
	ErrCMapTableMissing            = errors.New("character map (cmap) table missing")
	ErrHmtxTableMissing            = errors.New("horizontal metrics (hmtx) table missing")
	ErrPostTableMissing            = errors.New("PostScript (post) table missing")
	ErrInvalidHorizMetrics         = errors.New("invalid horizontal metrics")
	ErrInvalidCharMapFormat        = errors.New("invalid character map (cmap) format")
	ErrInvalidPPem                 = errors.New("invalid ppem value")
	ErrInvalidVertMetrics          = errors.New("invalid vertical metrics")
	ErrCouldNotFindContext         = errors.New("could not find context")
	ErrInvalidPostTableFormat      = errors.New("invalid PostScript (post) table format")
	ErrInvalidPostTable            = errors.New("invalid PostScript (post) table")
	ErrSyntaxError                 = errors.New("opcode syntax error")
	ErrStackUnderflow              = errors.New("argument stack underflow")
	ErrIgnore                      = errors.New("ignore")
	ErrNoUnicodeGlyphName          = errors.New("no Unicode glyph name found")
	ErrGlyphTooBig                 = errors.New("glyph to big for hinting")
	ErrMissingStartfontField       = errors.New("`STARTFONT' field missing")
	ErrMissingFontField            = errors.New("`FONT' field missing")
	ErrMissingSizeField            = errors.New("`SIZE' field missing")
	ErrMissingFontboundingboxField = errors.New("`FONTBOUNDINGBOX' field missing")
	ErrMissingCharsField           = errors.New("`CHARS' field missing")
	ErrMissingStartcharField       = errors.New("`STARTCHAR' field missing")
	ErrMissingEncodingField        = errors.New("`ENCODING' field missing")
	ErrMissingBbxField             = errors.New("`BBX' field missing")
	ErrBbxTooBig                   = errors.New("`BBX' too big")
	ErrCorruptedFontHeader         = errors.New("Font header corrupted or missing fields")
	ErrCorruptedFontGlyphs         = errors.New("Font glyphs corrupted or missing fields")

	lookupErr = map[int]error{
		0x01: ErrCannotOpenResource,
		0x02: ErrUnknownFileFormat,
		0x03: ErrInvalidFileFormat,
		0x04: ErrInvalidVersion,
		0x05: ErrLowerModuleVersion,
		0x06: ErrInvalidArgument,
		0x07: ErrUnimplementedFeature,
		0x08: ErrInvalidTable,
		0x09: ErrInvalidOffset,
		0x0A: ErrArrayTooLarge,
		0x0B: ErrMissingModule,
		0x0C: ErrMissingProperty,
		0x10: ErrInvalidGlyphIndex,
		0x11: ErrInvalidCharacterCode,
		0x12: ErrInvalidGlyphFormat,
		0x13: ErrCannotRenderGlyph,
		0x14: ErrInvalidOutline,
		0x15: ErrInvalidComposite,
		0x16: ErrTooManyHints,
		0x17: ErrInvalidPixelSize,
		0x20: ErrInvalidHandle,
		0x21: ErrInvalidLibraryHandle,
		0x22: ErrInvalidDriverHandle,
		0x23: ErrInvalidFaceHandle,
		0x24: ErrInvalidSizeHandle,
		0x25: ErrInvalidSlotHandle,
		0x26: ErrInvalidCharMapHandle,
		0x27: ErrInvalidCacheHandle,
		0x28: ErrInvalidStreamHandle,
		0x30: ErrTooManyDrivers,
		0x31: ErrTooManyExtensions,
		0x40: ErrOutOfMemory,
		0x41: ErrUnlistedObject,
		0x51: ErrCannotOpenStream,
		0x52: ErrInvalidStreamSeek,
		0x53: ErrInvalidStreamSkip,
		0x54: ErrInvalidStreamRead,
		0x55: ErrInvalidStreamOperation,
		0x56: ErrInvalidFrameOperation,
		0x57: ErrNestedFrameAccess,
		0x58: ErrInvalidFrameRead,
		0x60: ErrRasterUninitialized,
		0x61: ErrRasterCorrupted,
		0x62: ErrRasterOverflow,
		0x63: ErrRasterNegativeHeight,
		0x70: ErrTooManyCaches,
		0x80: ErrInvalidOpcode,
		0x81: ErrTooFewArguments,
		0x82: ErrStackOverflow,
		0x83: ErrCodeOverflow,
		0x84: ErrBadArgument,
		0x85: ErrDivideByZero,
		0x86: ErrInvalidReference,
		0x87: ErrDebugOpCode,
		0x88: ErrENDFInExecStream,
		0x89: ErrNestedDEFS,
		0x8A: ErrInvalidCodeRange,
		0x8B: ErrExecutionTooLong,
		0x8C: ErrTooManyFunctionDefs,
		0x8D: ErrTooManyInstructionDefs,
		0x8E: ErrTableMissing,
		0x8F: ErrHorizHeaderMissing,
		0x90: ErrLocationsMissing,
		0x91: ErrNameTableMissing,
		0x92: ErrCMapTableMissing,
		0x93: ErrHmtxTableMissing,
		0x94: ErrPostTableMissing,
		0x95: ErrInvalidHorizMetrics,
		0x96: ErrInvalidCharMapFormat,
		0x97: ErrInvalidPPem,
		0x98: ErrInvalidVertMetrics,
		0x99: ErrCouldNotFindContext,
		0x9A: ErrInvalidPostTableFormat,
		0x9B: ErrInvalidPostTable,
		0xA0: ErrSyntaxError,
		0xA1: ErrStackUnderflow,
		0xA2: ErrIgnore,
		0xA3: ErrNoUnicodeGlyphName,
		0xA4: ErrGlyphTooBig,
		0xB0: ErrMissingStartfontField,
		0xB1: ErrMissingFontField,
		0xB2: ErrMissingSizeField,
		0xB3: ErrMissingFontboundingboxField,
		0xB4: ErrMissingCharsField,
		0xB5: ErrMissingStartcharField,
		0xB6: ErrMissingEncodingField,
		0xB7: ErrMissingBbxField,
		0xB8: ErrBbxTooBig,
		0xB9: ErrCorruptedFontHeader,
		0xBA: ErrCorruptedFontGlyphs,
	}
)
