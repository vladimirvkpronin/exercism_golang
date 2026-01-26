package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	initial1 := initial
	for _, i := range s {
		initial1 = fn(initial1, i)
	}
	return initial1
	//panic("Please implement the Foldl function")
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	initial1 := initial
	for i := len(s) - 1; i >= 0; i-- {
		initial1 = fn(s[i], initial1)
	}
	return initial1
	//panic("Please implement the Foldr function")
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0, len(s))
    
    for _, item := range s {
        if fn(item) {
            result = append(result, item)
        }
    }
    
    return result
	//panic("Please implement the Filter function")
}

func (s IntList) Length() int {
	return len(s)
	//panic("Please implement the Length function")
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, len(s))
	for i, item := range s {
		result[i] = fn(item)
	}
	return result
	//panic("Please implement the Map function")
}

func (s IntList) Reverse() IntList {
	if len(s) == 0 {
		return IntList{}
	}
	result := make(IntList, len(s))
	for i, item := range s {
		result[len(s)-i-1] = item
	}
	return result
	//panic("Please implement the Reverse function")
}

func (s IntList) Append(lst IntList) IntList {
	if len(s) == 0 {
		return lst
	}
	if len(lst) == 0 {
		return s
	}
	  result := make(IntList, len(s)+len(lst))
    
    // Копируем элементы из текущего списка
    copy(result, s)
    
    // Копируем элементы из добавляемого списка
    copy(result[len(s):], lst)
    
    return result
	//panic("Please implement the Append function")
}

func (s IntList) Concat(lists []IntList) IntList {
	total := len(s)
	for _, lst := range lists {
		total += len(lst)
	}
	result := make(IntList, total)
	copy(result, s)
	pos := len(s)
	for _, lst := range lists {
		copy(result[pos:], lst)
		pos += len(lst)
	}
	return result
	//panic("Please implement the Concat function")
}
