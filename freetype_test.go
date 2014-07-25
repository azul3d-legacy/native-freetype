// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package freetype

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestFreetype(t *testing.T) {
	ctx, err := Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("FreeType context initialized")

	fontFile, err := os.Open("vera/Vera.ttf")
	if err != nil {
		t.Fatal(err)
	}

	fontFileData, err := ioutil.ReadAll(fontFile)
	if err != nil {
		t.Fatal(err)
	}

	font, err := ctx.Load(fontFileData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Loaded the font.")

	size := 500 * 64 // 50pt (expressed in 26.6 font units)
	dpi := 100       // 72 pixels per inch
	font.SetSize(size, 0, dpi, 0)
	//font.SetSizePixels(512, 512)

	glyphIndex := font.Index('A')
	t.Log("Glyph index for 'A':", glyphIndex)

	glyph, err := font.Load(glyphIndex)
	if err != nil {
		t.Fatal("Failed to load glyph.")
	}
	t.Log("Loaded the glyph.")

	m := glyph.HMetrics
	m = glyph.VMetrics
	xKern, yKern, err := font.Kerning('A', 'V')
	if err != nil {
		xKern = -1
		yKern = -1
	}

	t.Log("")
	t.Log("BBox:", font.BBox)
	t.Log("UnitsPerEm:", font.UnitsPerEm)
	t.Log("Ascender:", font.Ascender)
	t.Log("Descender:", font.Descender)
	t.Log("LineHeight:", font.LineHeight)
	t.Log("MaxAdvanceWidth:", font.MaxAdvanceWidth)
	t.Log("MaxAdvanceHeight:", font.MaxAdvanceHeight)
	t.Log("UnderlinePosition:", font.UnderlinePosition)
	t.Log("UnderlineThickness:", font.UnderlineThickness)
	t.Log("")
	t.Log("Glyph:")
	t.Log("")
	t.Log("    Width:", glyph.Width)
	t.Log("    Height:", glyph.Height)
	t.Log("")
	t.Log("    HMetrics:")
	t.Log("        BearingX:", m.BearingX)
	t.Log("        BearingY:", m.BearingY)
	t.Log("        Advance:", m.Advance)
	t.Log("        UnhintedAdvance:", m.UnhintedAdvance)
	t.Log("")
	t.Log("    VMetrics:")
	t.Log("        BearingX:", m.BearingX)
	t.Log("        BearingY:", m.BearingY)
	t.Log("        Advance:", m.Advance)
	t.Log("        UnhintedAdvance:", m.UnhintedAdvance)
	t.Log("")
	t.Logf("Kerning('A', 'V'): [%v, %v]\n", xKern, yKern)

	outFile, err := os.Create("test_freetype_out.png")
	if err != nil {
		t.Fatal(err)
	}

	glyphImage, err := glyph.Image()
	if err != nil {
		t.Fatal(err)
	}

	err = png.Encode(outFile, glyphImage)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Wrote test_freetype_out.png file.")
}
