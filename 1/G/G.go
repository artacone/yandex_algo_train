package main

import "fmt"

func makeBlanks(mMetal *int, mBlank int) int {
	nBlanks := *mMetal / mBlank
	*mMetal = *mMetal % mBlank
	return nBlanks
}

func makeDetails(mMetal *int, nBlanks, nDetailsPerBlank, mExcessivePerBlank int) int {
	nDetails := nDetailsPerBlank * nBlanks
	*mMetal += mExcessivePerBlank * nBlanks
	return nDetails
}

func countDetails(mMetal, mBlank, mDetail int) int{
	var nDetails, nBlanks int
	nDetailsPerBlank := mBlank / mDetail
	mExcessivePerBlank := mBlank % mDetail
	if nDetailsPerBlank == 0 {
		return 0
	}
	for ; mMetal >= mBlank; {
		nBlanks = makeBlanks(&mMetal, mBlank)
		nDetails += makeDetails(&mMetal, nBlanks, nDetailsPerBlank, mExcessivePerBlank)
	}
	return nDetails
}

func main() {
	var mMetal, mBlank, mDetail int

	fmt.Scan(&mMetal, &mBlank, &mDetail)
	fmt.Println(countDetails(mMetal, mBlank, mDetail))
}