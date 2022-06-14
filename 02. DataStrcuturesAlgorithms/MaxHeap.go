package main

type MaxHeap struct {
	array []int
}

//Insert element to the Heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {

}

//Get parent index
func parent (i int) int {
	return (i-1)/2
}

func left(i int) int {
	return 2*i + 1
}

func right (i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap (i1, i2 int) {
	h.array[i1], h.array[i2] = h.array
}

func main() {

}