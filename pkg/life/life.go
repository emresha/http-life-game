package life

import (
	"math/rand"
	"time"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

// Используйте код из предыдущего урока по игре «Жизнь»
func NewWorld(height, width int) (*World, error) {
	// создаём тип World с количеством слайсов hight (количество строк)
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) Next(x, y int) bool {
	n := w.Neighbors(x, y)       // получим количество живых соседей
	alive := w.Cells[y][x]       // текущее состояние клетки
	if n < 4 && n > 1 && alive { // если соседей двое или трое, а клетка жива
		return true // то следующее состояние — жива
	}
	if n == 3 && !alive { // если клетка мертва, но у неё трое соседей
		return true // клетка оживает
	}

	return false // в любых других случаях — клетка мертва
}

func (w World) Neighbors(x, y int) int {
	deltas := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0

	for _, delta := range deltas {
		neighborX := x + delta[0]
		neighborY := y + delta[1]

		if neighborX >= 0 && neighborX < w.Width && neighborY >= 0 && neighborY < w.Height {
			if w.Cells[neighborY][neighborX] {
				count++
			}
		}
	}

	return count
}

func NextState(oldWorld, newWorld *World) {
	// переберём все клетки, чтобы понять, в каком они состоянии
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

// RandInit заполняет поля на указанное число процентов
func (w *World) RandInit(percentage int) {
	// Количество живых клеток
	numAlive := percentage * w.Height * w.Width / 100
	// Заполним живыми первые клетки
	w.fillAlive(numAlive)
	// Получаем рандомные числа
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Рандомно меняем местами
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
