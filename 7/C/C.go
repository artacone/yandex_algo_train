package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func readInput() (nStudents, d int, students [][2]int) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	nStudents, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	d, _ = strconv.Atoi(sc.Text())
	students = make([][2]int, nStudents)
	i := 0
	for sc.Scan() {
		s, _ := strconv.Atoi(sc.Text())
		students[i] = [2]int{s, i}
		i++
	}
	sort.Slice(students, func(i, j int) bool {
		return students[i][0] < students[j][0]
	})
	return
}

func countVariants(nStudents, d int, students [][2]int) (int, []int) {
	variants := make([]int, nStudents)
	nVariants := 1

	for left, right := 0, 1; right < nStudents; left, right = left+1, right+1 {
		next := students[left][0] + d
		for right < nStudents && students[right][0] <= next {
			right++
		}
		if right-left > nVariants {
			nVariants = right - left
		}
	}
	for i := 0; i < nStudents; {
		for j := 1; j <= nVariants && i < nStudents; j++ {
			variants[students[i][1]] = j
			i++
		}
	}
	return nVariants, variants
}

func writeOutput(nVariants int, variants []int) error {
	wr := bufio.NewWriter(os.Stdout)
	_, err := wr.WriteString(strconv.Itoa(nVariants))
	if err != nil {
		return err
	}
	err = wr.WriteByte('\n')
	if err != nil {
		return err
	}

	for _, v := range variants {
		_, err := wr.WriteString(strconv.Itoa(v))
		if err != nil {
			return err
		}
		err = wr.WriteByte(' ')
		if err != nil {
			return err
		}
	}
	err = wr.WriteByte('\n')
	if err != nil {
		return err
	}
	return wr.Flush()
}

func main() {
	nStudents, d, students := readInput()

	nVariants, variants := countVariants(nStudents, d, students)

	err := writeOutput(nVariants, variants)
	if err != nil {
		return
	}
}
