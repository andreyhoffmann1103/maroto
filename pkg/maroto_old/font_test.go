package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewFont(t *testing.T) {
	font := maroto.NewFont(&mocks.Pdf{}, 10, maroto.Arial, maroto.Bold)

	assert.NotNil(t, font)
	assert.Equal(t, fmt.Sprintf("%T", font), "*maroto.font")
	assert.Equal(t, font.GetFamily(), maroto.Arial)
	assert.Equal(t, font.GetStyle(), maroto.Bold)
	assert.Equal(t, font.GetSize(), 10.0)
}

func TestFont_GetSetFamily(t *testing.T) {
	cases := []struct {
		name        string
		family      maroto.Family
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertFont  func(t *testing.T, family maroto.Family)
	}{
		{
			"PdfMaroto.Arial",
			maroto.Arial,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "arial", "B", 10.0)
			},
			func(t *testing.T, family maroto.Family) {
				assert.Equal(t, family, maroto.Arial)
			},
		},
		{
			"PdfMaroto.Helvetica",
			maroto.Helvetica,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "helvetica", "B", 10.0)
			},
			func(t *testing.T, family maroto.Family) {
				assert.Equal(t, family, maroto.Helvetica)
			},
		},
		{
			"PdfMaroto.Symbol",
			maroto.Symbol,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "symbol", "B", 10.0)
			},
			func(t *testing.T, family maroto.Family) {
				assert.Equal(t, family, maroto.Symbol)
			},
		},
		{
			"PdfMaroto.ZapBats",
			maroto.ZapBats,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "zapfdingbats", "B", 10.0)
			},
			func(t *testing.T, family maroto.Family) {
				assert.Equal(t, family, maroto.ZapBats)
			},
		},
		{
			"PdfMaroto.Courier",
			maroto.Courier,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "courier", "B", 10.0)
			},
			func(t *testing.T, family maroto.Family) {
				assert.Equal(t, family, maroto.Courier)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := maroto.NewFont(pdf, 10, maroto.Arial, maroto.Bold)

		// Act
		font.SetFamily(c.family)

		// Assert
		c.assertCalls(t, pdf)
		c.assertFont(t, font.GetFamily())
	}
}

func TestFont_GetSetStyle(t *testing.T) {
	cases := []struct {
		name        string
		style       maroto.Style
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertStyle func(t *testing.T, style maroto.Style)
	}{
		{
			"PdfMaroto.Normal",
			maroto.Normal,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "")
			},
			func(t *testing.T, style maroto.Style) {
				assert.Equal(t, style, maroto.Normal)
			},
		},
		{
			"PdfMaroto.Bold",
			maroto.Bold,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "B")
			},
			func(t *testing.T, style maroto.Style) {
				assert.Equal(t, style, maroto.Bold)
			},
		},
		{
			"PdfMaroto.Italic",
			maroto.Italic,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "I")
			},
			func(t *testing.T, style maroto.Style) {
				assert.Equal(t, style, maroto.Italic)
			},
		},
		{
			"PdfMaroto.BoldItalic",
			maroto.BoldItalic,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "BI")
			},
			func(t *testing.T, style maroto.Style) {
				assert.Equal(t, style, maroto.BoldItalic)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := maroto.NewFont(pdf, 10, maroto.Arial, maroto.Bold)

		// Act
		font.SetStyle(c.style)

		// Assert
		c.assertCalls(t, pdf)
		c.assertStyle(t, font.GetStyle())
	}
}

func TestFont_GetSetSize(t *testing.T) {
	// Arrange
	pdf := &mocks.Pdf{}
	pdf.On("SetFontSize", mock.Anything)
	font := maroto.NewFont(pdf, 10, maroto.Arial, maroto.Bold)

	// Act
	font.SetSize(16)

	// Assert
	pdf.AssertNumberOfCalls(t, "SetFontSize", 1)
	pdf.MethodCalled("SetFontSize", 16)
	assert.Equal(t, font.GetSize(), 16.0)
}

func TestFont_GetSetFont(t *testing.T) {
	cases := []struct {
		name        string
		family      maroto.Family
		style       maroto.Style
		size        float64
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertFont  func(t *testing.T, family maroto.Family, style maroto.Style, size float64)
	}{
		{
			"PdfMaroto.Arial, PdfMaroto.Normal, 16",
			maroto.Arial,
			maroto.Normal,
			16.0,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "arial", "", 16.0)
			},
			func(t *testing.T, family maroto.Family, style maroto.Style, size float64) {
				assert.Equal(t, family, maroto.Arial)
				assert.Equal(t, style, maroto.Normal)
				assert.Equal(t, 16, int(size))
			},
		},
		{
			"PdfMaroto.Helvetica, PdfMaroto.Bold, 13",
			maroto.Helvetica,
			maroto.Bold,
			13,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "helvetica", "B", 13.0)
			},
			func(t *testing.T, family maroto.Family, style maroto.Style, size float64) {
				assert.Equal(t, family, maroto.Helvetica)
				assert.Equal(t, style, maroto.Bold)
				assert.Equal(t, 13, int(size))
			},
		},
		{
			"PdfMaroto.Symbol, PdfMaroto.Italic, 10",
			maroto.Symbol,
			maroto.Italic,
			10,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "symbol", "I", 10.0)
			},
			func(t *testing.T, family maroto.Family, style maroto.Style, size float64) {
				assert.Equal(t, family, maroto.Symbol)
				assert.Equal(t, style, maroto.Italic)
				assert.Equal(t, 10, int(size))
			},
		},
		{
			"PdfMaroto.ZapBats, PdfMaroto.BoldItalic, 5",
			maroto.ZapBats,
			maroto.BoldItalic,
			5,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "zapfdingbats", "BI", 5.0)
			},
			func(t *testing.T, family maroto.Family, style maroto.Style, size float64) {
				assert.Equal(t, family, maroto.ZapBats)
				assert.Equal(t, style, maroto.BoldItalic)
				assert.Equal(t, 5, int(size))
			},
		},
		{
			"PdfMaroto.Courier, PdfMaroto.Normal, 12",
			maroto.Courier,
			maroto.Normal,
			12,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "courier", "", 12.0)
			},
			func(t *testing.T, family maroto.Family, style maroto.Style, size float64) {
				assert.Equal(t, family, maroto.Courier)
				assert.Equal(t, style, maroto.Normal)
				assert.Equal(t, 12, int(size))
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := maroto.NewFont(pdf, 10, maroto.Arial, maroto.Bold)

		// Act
		font.SetFont(c.family, c.style, c.size)
		family, style, size := font.GetFont()

		// Assert
		c.assertCalls(t, pdf)
		c.assertFont(t, family, style, size)
	}
}
