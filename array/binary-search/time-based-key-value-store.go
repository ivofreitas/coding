package binarysearch

type TimeValue struct {
	Value     string
	Timestamp int
}

type TimeMap struct {
	store map[string][]TimeValue
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]TimeValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if len(this.store[key]) == 0 {
		this.store[key] = make([]TimeValue, 0)
	}
	this.store[key] = append(this.store[key], TimeValue{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.store[key]) == 0 {
		return ""
	}

	left, right, res := 0, len(this.store[key])-1, ""
	for left <= right {
		middle := (right + left) / 2

		if this.store[key][middle].Timestamp <= timestamp {
			res = this.store[key][middle].Value
			left = middle + 1
		} else if this.store[key][middle].Timestamp > timestamp {
			right = middle - 1
		} else {

		}
	}

	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
