package queue

import "errors"

type PriorityQueue struct {
	Data []int
	Size int
}

func (p *PriorityQueue) Init() {
	p.Data = make([]int, 1024)
	p.Size = 0
}

func (p *PriorityQueue) Length() int {
	return p.Size
}

func (p *PriorityQueue) Insert(num int) {
	p.Size++
	p.Data[p.Size] = num
	p.up(p.Size)
}

func (p *PriorityQueue) Pop() (int, error) {
	if p.Size < 1 {
		return -1, errors.New("without num")
	}
	res := p.Data[1]
	p.Data[1], p.Data[p.Size] = p.Data[p.Size], p.Data[1]
	p.Size--
	p.down(1)
	return res, nil
}

func (p *PriorityQueue) Max() int {
	return p.Data[1]
}

func (p *PriorityQueue) down(i int) {
	largest := i
	if 2*i < p.Size && p.Data[2*i] > p.Data[largest] {
		largest = 2 * i
	}
	if 2*i+1 < p.Size && p.Data[2*i+1] > p.Data[largest] {
		largest = 2*i + 1
	}
	if largest != i {
		p.Data[i], p.Data[largest] = p.Data[largest], p.Data[i]
		p.down(largest)
	}
}

func (p *PriorityQueue) up(i int) {
	if i <= 1 {
		return
	}
	for parent := i / 2; parent >= 1; parent /= 2 {
		if p.Data[i] > p.Data[parent] {
			p.Data[i], p.Data[parent] = p.Data[parent], p.Data[i]
			i = parent
		} else {
			break
		}
	}
}
