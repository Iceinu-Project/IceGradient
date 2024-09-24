package gradient

import (
	"fmt"
	"strconv"
	"strings"
)

// CodeToRGB Transform an RGB code to an RGB slice
func CodeToRGB(code string) ([]int, error) {
	// 去掉可能存在的 '#' 字符
	code = strings.TrimPrefix(strings.ToLower(code), "#")

	// 确保 RGBToCode 代码有效（6 个十六进制数字）
	if len(code) != 6 {
		return nil, fmt.Errorf("invalid RGBToCode code: %s", code)
	}

	// 解析 RGBToCode 组件
	r, err := strconv.ParseInt(code[0:2], 16, 0)
	if err != nil {
		return nil, err
	}
	g, err := strconv.ParseInt(code[2:4], 16, 0)
	if err != nil {
		return nil, err
	}
	b, err := strconv.ParseInt(code[4:6], 16, 0)
	if err != nil {
		return nil, err
	}

	return []int{int(r), int(g), int(b)}, nil
}

// RGBToANSI Generate an ANSI color string from an RGB slice
func RGBToANSI(rgb ...int) string {
	// 确保 RGBToCode 组件有效
	if len(rgb) != 3 {
		return ""
	}

	// 计算 ANSI 颜色代码
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
}

// RGB Convert an RGB code to an ANSI color string
func RGB(code string) (string, error) {
	rgb, err := CodeToRGB(code)
	if err != nil {
		return "", err
	}
	return RGBToANSI(rgb...), nil
}

// ANSIToRGB Convert an ANSI color string to an RGB slice
func ANSIToRGB(ansi string) ([]int, error) {
	// 基础颜色映射
	baseColors := map[string][]int{
		"\033[30m": {0, 0, 0},       // Black
		"\033[31m": {255, 0, 0},     // Red
		"\033[32m": {0, 255, 0},     // Green
		"\033[33m": {255, 255, 0},   // Yellow
		"\033[34m": {0, 0, 255},     // Blue
		"\033[35m": {255, 0, 255},   // Magenta
		"\033[36m": {0, 255, 255},   // Cyan
		"\033[37m": {255, 255, 255}, // White
	}

	// 检查是否为基础颜色
	if rgb, exists := baseColors[ansi]; exists {
		return rgb, nil
	}

	// 确保 ANSI 字符串有效
	if !strings.HasPrefix(ansi, "\033[38;2;") || !strings.HasSuffix(ansi, "m") {
		return nil, fmt.Errorf("invalid ANSI color string: %s", ansi)
	}

	// 去掉前缀和后缀
	ansi = strings.TrimPrefix(ansi, "\033[38;2;")
	ansi = strings.TrimSuffix(ansi, "m")

	// 分割 RGB 组件
	components := strings.Split(ansi, ";")
	if len(components) != 3 {
		return nil, fmt.Errorf("invalid ANSI color string: %s", ansi)
	}

	// 解析 RGB 组件
	r, err := strconv.Atoi(components[0])
	if err != nil {
		return nil, err
	}
	g, err := strconv.Atoi(components[1])
	if err != nil {
		return nil, err
	}
	b, err := strconv.Atoi(components[2])
	if err != nil {
		return nil, err
	}

	return []int{r, g, b}, nil
}

// RGBToCode Generate an RGB code from an RGB slice
func RGBToCode(rgb ...int) string {
	// 确保 RGBToCode 组件有效
	if len(rgb) != 3 {
		return ""
	}

	// 生成 RGBToCode 代码
	return fmt.Sprintf("#%02x%02x%02x", rgb[0], rgb[1], rgb[2])
}

// ANSI Convert an ANSI color string to an RGB code
func ANSI(ansi string) (string, error) {
	rgb, err := ANSIToRGB(ansi)
	if err != nil {
		return "", err
	}
	return RGBToCode(rgb...), nil
}

// RGBToBgANSI Generate an ANSI background color string from an RGB slice
func RGBToBgANSI(rgb ...int) string {
	// 确保 RGB 组件有效
	if len(rgb) != 3 {
		return ""
	}

	// 计算 ANSI 背景颜色代码
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
}

// ConvertToRGB Convert various color formats to RGB slice
func ConvertToRGB(color interface{}) ([]int, error) {
	switch v := color.(type) {
	case string:
		if strings.HasPrefix(v, "\033[") {
			return ANSIToRGB(v)
		} else if strings.HasPrefix(v, "#") {
			return CodeToRGB(v[1:])
		} else {
			return CodeToRGB(v)
		}
	case []int:
		if len(v) == 3 {
			return v, nil
		}
	case [3]int:
		return v[:], nil
	}
	return nil, fmt.Errorf("invalid color format")
}
