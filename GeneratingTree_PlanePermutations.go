// Plane permutations are those which avoid the pattern 213'54
// Avoiding 213'54 is equivalent to avoiding 2-14-3
// This generating tree based enumeration is based on the following paper:
// https://arxiv.org/pdf/1702.04529.pdf

package main

import (
    "fmt"
    "runtime"
)

// struct to mimick set in Python
// code from https://play.golang.org/p/_FvECoFvhq
type ArraySet struct {
	set map[[20]int]bool
}

func NewArraySet() *ArraySet {
	return &ArraySet{make(map[[20]int]bool)}
}

func (set *ArraySet) Add(p [20]int) bool {
	_, found := set.set[p]
	set.set[p] = true
	return !found	//False if it existed already
}


func (set *ArraySet) Get(i [20]int) bool {
	_, found := set.set[i]
	return found	//true if it existed already
}

func (set *ArraySet) Remove(i [20]int) {
	delete(set.set, i)
}

// -----

// func to mimick min in Python
func min (a, b int) (int) {
	if a<=b {return a}
	return b
}

func localExp (perm []int, a int, level int) ([]int){

	// Local expansion as described in the paper
	newPerm := make([]int, level+1)
	for i, k := range perm {
		if k < a {
			newPerm[i] = k
		} else {
			newPerm[i] = k+1
		}
	}
	newPerm[level] = a
	return newPerm
}

func isPlane (perm []int) (bool) {
	n := len(perm)
	steps := make([]int, 0, n)
	for k:=0; k<n-1; k++ {
		if perm[k] < perm[k+1] - 1 {steps = append(steps, k)}
	}

	for _,s := range steps {
		m, M := perm[s], perm[s+1]
		two := 1000
		prefix, suffix := perm[:s], perm[s+2:]
		for _,k := range prefix {
			if (k > m) && (k < M - 1) {
				two = min(k, two)
			}
		}

		for _,k := range suffix {
			if (k > two) && (k < M) {
				return false
			}
		}
	}

	return true
}


func main() {
	procs := 2
	runtime.GOMAXPROCS(procs)

	curLevel := NewArraySet()
	var arr [20]int
	copy(arr[:], []int{1,2,3})
	curLevel.Add(arr)
	copy(arr[:], []int{1,3,2})
	curLevel.Add(arr)
	copy(arr[:], []int{2,1,3})
	curLevel.Add(arr)
	copy(arr[:], []int{3,1,2})
	curLevel.Add(arr)
	copy(arr[:], []int{2,3,1})
	curLevel.Add(arr)
	copy(arr[:], []int{3,2,1})
	curLevel.Add(arr)
	level := 3

	for level < 20 {
		newLevel := NewArraySet()
		for perm := range curLevel.set {
			for a:=1; a<=level+1; a++ {
				newPerm:= localExp(perm[:], a, level)
				if isPlane(newPerm) {
					var finPerm [20]int
					copy(finPerm[:], newPerm)
					newLevel.Add(finPerm)
				}
			}
		}

		fmt.Println(len(newLevel.set))
		curLevel = newLevel
		level += 1
	}
}


