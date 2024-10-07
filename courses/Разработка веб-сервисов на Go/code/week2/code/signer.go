package main

import "strconv"

// сюда писать код
func ExecutePipeline(jobs ...job) {

	for _, j := range jobs {

	}
}

func SingleHash(in, out chan interface{}) {
	data := (<-in).(string)
	out <- DataSignerCrc32(data) + "~" + DataSignerCrc32(DataSignerMd5(data))
}

func MultiHash(in, out chan interface{}) {
	data := (<-in).(string)
	var res string
	for i := 0; i < 6; i++ {
		res += DataSignerCrc32(strconv.Itoa(i) + data)
	}
	out <- res
}

func CombineResults(in, out chan interface{}) {
	var res string
	for v := range in {
		res += v.(string)
	}
	out <- res
}
