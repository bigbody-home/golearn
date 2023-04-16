package main

type IntSlice []string

func (i1 IntSlice) Len() int {
	//TODO implement me
	//panic("implement me")
	return len(i1)
}

func (i1 IntSlice) Less(i, j int) bool {
	//TODO implement me
	//panic("implement me")
	return i1[i] < i1[j]
}

func (i1 IntSlice) Swap(i, j int) {
	//TODO implement me
	//panic("implement me")
	i1[i], i1[j] = i1[j], i1[i]
}
