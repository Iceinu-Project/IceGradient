package IceGradient

import (
	"fmt"
	"strings"
)

// FontColor Convert a string (RGB code), slice (RGB array), or individual RGB values to ANSI font color
func FontColor(input ...interface{}) string {
	var rgb []int

	if len(input) == 1 {
		switch v := input[0].(type) {
		case string:
			rgb, _ = CodeToRGB(v)
		case []int:
			rgb = v
			if len(rgb) != 3 {
				return ""
			}
		case [3]int:
			rgb = v[:]
		default:
			return ""
		}
	} else if len(input) == 3 {
		for _, v := range input {
			if value, ok := v.(int); ok {
				rgb = append(rgb, value)
			} else {
				return ""
			}
		}
	} else {
		return ""
	}

	return RGBToANSI(rgb...)
}

// BackgroundColor Convert a string (RGB code), slice (RGB array), or individual RGB values to ANSI background color
func BackgroundColor(input ...interface{}) string {
	var rgb []int
	var err error

	if len(input) == 1 {
		switch v := input[0].(type) {
		case string:
			if strings.HasPrefix(v, "\033[") {
				rgb, err = ANSIToRGB(v)
				if err != nil {
					return ""
				}
			} else {
				rgb, _ = CodeToRGB(v)
			}
		case []int:
			rgb = v
			if len(rgb) != 3 {
				return ""
			}
		case [3]int:
			rgb = v[:]
		default:
			return ""
		}
	} else if len(input) == 3 {
		for _, v := range input {
			if value, ok := v.(int); ok {
				rgb = append(rgb, value)
			} else {
				return ""
			}
		}
	} else {
		return ""
	}

	return RGBToBgANSI(rgb...)
}

// GradientText Generate text with a gradient font color
func GradientText(text string, colors ...interface{}) string {
	if len(colors) < 2 {
		return text
	}

	// Convert all colors to RGB
	var rgbColors [][]int
	for _, color := range colors {
		rgb, err := ConvertToRGB(color)
		if err != nil {
			return text
		}
		rgbColors = append(rgbColors, rgb)
	}

	// Split text into lines
	lines := strings.Split(text, "\n")
	totalChars := 0
	for _, line := range lines {
		totalChars += len(line)
	}

	// Calculate gradient steps
	steps := totalChars - 1
	segments := len(rgbColors) - 1
	segmentLength := steps / segments

	var result strings.Builder
	charIndex := 0
	for _, line := range lines {
		for _, char := range line {
			segment := charIndex / (segmentLength + 1)
			if segment >= segments {
				segment = segments - 1
			}
			startColor := rgbColors[segment]
			endColor := rgbColors[segment+1]

			// Calculate the interpolated color
			ratio := float64(charIndex%(segmentLength+1)) / float64(segmentLength+1)
			r := int(float64(startColor[0])*(1-ratio) + float64(endColor[0])*ratio)
			g := int(float64(startColor[1])*(1-ratio) + float64(endColor[1])*ratio)
			b := int(float64(startColor[2])*(1-ratio) + float64(endColor[2])*ratio)

			// Append the colored character to the result
			result.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm%c", r, g, b, char))
			charIndex++
		}
		result.WriteString("\n")
	}

	// Reset color at the end
	result.WriteString("\033[0m")
	return result.String()
}

// GradientBackgroundText Generate text with a gradient background color
func GradientBackgroundText(text string, colors ...interface{}) string {
	if len(colors) < 2 {
		return text
	}

	// Convert all colors to RGB
	var rgbColors [][]int
	for _, color := range colors {
		rgb, err := ConvertToRGB(color)
		if err != nil {
			return text
		}
		rgbColors = append(rgbColors, rgb)
	}

	// Split text into lines
	lines := strings.Split(text, "\n")
	totalChars := 0
	for _, line := range lines {
		totalChars += len(line)
	}

	// Calculate gradient steps
	steps := totalChars - 1
	segments := len(rgbColors) - 1
	segmentLength := steps / segments

	var result strings.Builder
	charIndex := 0
	for _, line := range lines {
		for _, char := range line {
			segment := charIndex / (segmentLength + 1)
			if segment >= segments {
				segment = segments - 1
			}
			startColor := rgbColors[segment]
			endColor := rgbColors[segment+1]

			// Calculate the interpolated color
			ratio := float64(charIndex%(segmentLength+1)) / float64(segmentLength+1)
			r := int(float64(startColor[0])*(1-ratio) + float64(endColor[0])*ratio)
			g := int(float64(startColor[1])*(1-ratio) + float64(endColor[1])*ratio)
			b := int(float64(startColor[2])*(1-ratio) + float64(endColor[2])*ratio)

			// Append the colored character to the result
			result.WriteString(fmt.Sprintf("\033[48;2;%d;%d;%dm%c", r, g, b, char))
			charIndex++
		}
		// Reset color at the end of each line
		result.WriteString("\033[0m\n")
	}

	return result.String()
}

// GradientTextWithBackground Generate text with both gradient font and background
func GradientTextWithBackground(text string, fontColors []interface{}, bgColors []interface{}) string {
	if len(fontColors) < 2 || len(bgColors) < 2 {
		return text
	}

	// Generate gradient font text
	fontText := GradientText(text, fontColors...)

	// Generate gradient background text
	bgText := GradientBackgroundText(text, bgColors...)

	// Combine font and background gradients
	var result strings.Builder
	fontLines := strings.Split(fontText, "\n")
	bgLines := strings.Split(bgText, "\n")

	for i := range fontLines {
		if i < len(bgLines) {
			result.WriteString(bgLines[i])
			result.WriteString(fontLines[i])
			result.WriteString("\033[0m\n")
		}
	}

	return result.String()
}

// BlendColors Blends multiple colors into one average color
func BlendColors(colors ...interface{}) ([]int, error) {
	if len(colors) == 0 {
		return nil, fmt.Errorf("at least one color is required")
	}

	var totalR, totalG, totalB int
	for _, color := range colors {
		rgb, err := ConvertToRGB(color)
		if err != nil {
			return nil, err
		}
		totalR += rgb[0]
		totalG += rgb[1]
		totalB += rgb[2]
	}

	numColors := len(colors)
	averageColor := []int{
		totalR / numColors,
		totalG / numColors,
		totalB / numColors,
	}

	return averageColor, nil
}

// Convert a slice of strings to a slice of interfaces
func convertStringsToInterfaces(colors []string) []interface{} {
	var result []interface{}
	for _, color := range colors {
		result = append(result, color)
	}
	return result
}

// GradientTextWithTemplate Generate text with a gradient font color using a predefined template
func GradientTextWithTemplate(text string, templateName string) string {
	colors, ok := GradientTemplates[templateName]
	if !ok {
		return text
	}
	interfaceColors := convertStringsToInterfaces(colors)
	return GradientText(text, interfaceColors...)
}

// GradientBackgroundTextWithTemplate Generate text with a gradient background color using a predefined template
func GradientBackgroundTextWithTemplate(text string, templateName string) string {
	colors, ok := GradientTemplates[templateName]
	if !ok {
		return text
	}
	interfaceColors := convertStringsToInterfaces(colors)
	return GradientBackgroundText(text, interfaceColors...)
}
