package main

func isSubsequence(subSeq string, str string) bool {

	seqIdx := 0

	for i := 0; i < len(str); i++ {
		if len(subSeq) == seqIdx {
			return true
		}
		if str[i] == subSeq[seqIdx] {
			seqIdx++
		}
	}

	return len(subSeq) == seqIdx

}
