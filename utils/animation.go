package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Spinner() {
	frames := []string{"-", "\\", "|", "/"}
	delay := 100 * time.Millisecond

	for i := 0; i < 50; i++ {
		fmt.Printf("\r%s", frames[i%len(frames)])
		time.Sleep(delay)
	}
	fmt.Printf("\r%s\r", "          ")
	fmt.Printf("\n")
}

func LoadingBar() {
	width := 50
	delay := 100 * time.Millisecond

	for i := 0; i <= width; i++ {
		progress := i * 100 / width
		bar := "[" + strings.Repeat("=", i) + strings.Repeat(" ", width-i) + "]"
		fmt.Printf("\r%s %d%%", bar, progress)
		time.Sleep(delay)
	}
	fmt.Println() // Print a new line when done
}

func Scrolling() {
	message := "This is a scrolling text message"
	delay := 100 * time.Millisecond

	for i := 0; i <= len(message); i++ {
		fmt.Printf("\r%s", message[i:]+message[:i])
		time.Sleep(delay)
	}
	fmt.Println() // Print a new line when done
}

func BouncingBall() {
	width, height := 40, 10
	ballX, ballY := width/2, 1
	directionX, directionY := 1, 1
	delay := 100 * time.Millisecond

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Update ball position
		ballX += directionX
		ballY += directionY

		// Bounce off the walls
		if ballX <= 0 || ballX >= width-1 {
			directionX *= -1
		}
		if ballY <= 0 || ballY >= height-1 {
			directionY *= -1
		}

		// Draw the ball
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if x == ballX && y == ballY {
					fmt.Print("O")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

		// Pause for a moment
		time.Sleep(delay)
	}
}

func Matrix() {
	rand.Seed(time.Now().UnixNano())
	width, height := 80, 25
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	matrix := make([][]rune, height)
	for i := range matrix {
		matrix[i] = make([]rune, width)
	}

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Generate random falling characters
		for x := 0; x < width; x++ {
			if rand.Intn(3) == 0 {
				matrix[0][x] = chars[rand.Intn(len(chars))]
			}
		}

		// Move characters down
		for y := 1; y < height; y++ {
			for x := 0; x < width; x++ {
				matrix[y][x] = matrix[y-1][x]
			}
		}

		// Print the matrix
		for _, row := range matrix {
			fmt.Println(string(row))
		}

		// Pause for a moment
		time.Sleep(50 * time.Millisecond)
	}

}

func CountDownTimer() {
	duration := 10 * time.Second // Change the countdown duration here
	ticker := time.NewTicker(1 * time.Second)

	for remaining := duration; remaining >= 0; {
		fmt.Printf("\rTime remaining: %s", remaining.Round(time.Second))
		remaining -= 1 * time.Second
		<-ticker.C
	}

	fmt.Println("\nCountdown complete!")

}

func TextMorphing() {
	startText := "Hello, World!"
	endText := "byeee, World!"
	duration := 5 * time.Second // Total animation duration
	frameCount := 50            // Number of frames
	delay := duration / time.Duration(frameCount)

	// Ensure both start and end texts have the same length
	if len(startText) != len(endText) {
		fmt.Println("Texts must have the same length")
		return
	}

	for i := 0; i <= frameCount; i++ {
		interpolatedText := interpolateText(startText, endText, float64(i)/float64(frameCount))
		fmt.Print("\r" + interpolatedText)
		time.Sleep(delay)
	}

	fmt.Println() // Print a new line when done
}

func interpolateText(start, end string, progress float64) string {
	if len(start) != len(end) {
		return ""
	}

	interpolated := make([]rune, len(start))
	for i := 0; i < len(start); i++ {
		startChar := rune(start[i])
		endChar := rune(end[i])

		// Interpolate the character based on progress
		interpolatedChar := startChar + rune(float64(endChar-startChar)*progress)

		interpolated[i] = interpolatedChar
	}

	return string(interpolated)

}
func FireWork() {
	rand.Seed(time.Now().UnixNano())
	width, height := 80, 25
	screen := make([][]rune, height)
	for i := range screen {
		screen[i] = make([]rune, width)
	}

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Simulate fireworks explosions
		for i := 0; i < 5; i++ {
			x := rand.Intn(width)
			y := rand.Intn(height-5) + 5 // Ensure fireworks appear above the bottom
			color := randColor()
			explodeFirework(screen, x, y, color)
		}

		// Update and display the screen
		drawScreen(screen)

		// Pause for a moment
		time.Sleep(500 * time.Millisecond)
	}
}

func explodeFirework(screen [][]rune, x, y int, color rune) {
	for i := 0; i < 5; i++ {
		explodeStar(screen, x, y, color)
		time.Sleep(50 * time.Millisecond)
	}
}

func explodeStar(screen [][]rune, x, y int, color rune) {
	for i := 0; i < 8; i++ {
		dx, dy := direction(i)
		for j := 1; j <= 5; j++ {
			newX, newY := x+j*dx, y+j*dy
			if newX >= 0 && newX < len(screen[0]) && newY >= 0 && newY < len(screen) {
				screen[newY][newX] = color
			}
		}
	}
}

func direction(i int) (int, int) {
	switch i {
	case 0:
		return 1, 0
	case 1:
		return -1, 0
	case 2:
		return 0, 1
	case 3:
		return 0, -1
	case 4:
		return 1, 1
	case 5:
		return -1, 1
	case 6:
		return 1, -1
	case 7:
		return -1, -1
	default:
		return 0, 0
	}
}

func drawScreen(screen [][]rune) {
	for _, row := range screen {
		fmt.Println(string(row))
	}
}

func randColor() rune {
	colors := []rune{'R', 'G', 'B', 'Y', 'M', 'C', 'W'}
	return colors[rand.Intn(len(colors))]
}

func AsciiArt() {
	frames := []string{
		"  ___  ",
		" ( o ) ",
		"  \\ /  ",
		"  _|_  ",
	}

	delay := 250 * time.Millisecond

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Display each frame
		for _, frame := range frames {
			fmt.Println(frame)
		}

		// Pause for a moment
		time.Sleep(delay)
	}

}

func RainFall() {
	rand.Seed(time.Now().UnixNano())
	width, height := 80, 25
	screen := make([][]rune, height)
	for i := range screen {
		screen[i] = make([]rune, width)
	}

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Simulate raindrops
		for x := 0; x < width; x++ {
			if rand.Intn(5) == 0 {
				screen[0][x] = ' '
			}
		}

		for y := height - 1; y > 0; y-- {
			for x := 0; x < width; x++ {
				screen[y][x] = screen[y-1][x]
			}
		}

		// Generate new raindrops
		for x := 0; x < width; x++ {
			if rand.Intn(5) == 0 {
				screen[0][x] = '|'
			}
		}

		// Print the screen
		for _, row := range screen {
			fmt.Println(string(row))
		}

		// Pause for a moment
		time.Sleep(50 * time.Millisecond)
	}

}
