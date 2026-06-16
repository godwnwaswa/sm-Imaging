package imaging

import (
	"image"
)

// EnhanceHD applies mild sharpening and contrast adjustments to make an image appear more "HD".
// It is useful for enhancing low-quality images to give them a richer, crisper look.
func EnhanceHD(img image.Image) *image.NRGBA {
	// 1. Mild sharpening to bring out details
	sharpened := Sharpen(img, 0.8)

	// 2. Adjust contrast slightly to make colors pop (Sigmoid curve preserves highlights/shadows)
	// midpoint 0.5, factor 1.5 gives a nice gentle contrast boost
	enhanced := AdjustSigmoid(sharpened, 0.5, 1.5)

	return enhanced
}

// UpscaleHD resizes the image to the specified width and height using the Lanczos filter (which provides 
// high-quality upscaling) and then applies HD enhancements (sharpening and contrast) to maintain high quality.
func UpscaleHD(img image.Image, width, height int) *image.NRGBA {
	// High-quality resize using Lanczos filter
	resized := Resize(img, width, height, Lanczos)
	
	// Apply HD enhancements
	return EnhanceHD(resized)
}

// FitHD scales down the image using the Lanczos filter to fit the specified
// maximum width and height, and then applies HD enhancements.
func FitHD(img image.Image, width, height int) *image.NRGBA {
	fitted := Fit(img, width, height, Lanczos)
	return EnhanceHD(fitted)
}
