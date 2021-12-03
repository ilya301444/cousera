package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
	"fmt"
	"time"
)


func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})
	for _, singleJob := range jobs {
		wg.Add(1)
		out := make(chan interface{})
		go jobMaker(singleJob, in, out, wg)
		in = out
	}
	wg.Wait()
}

func jobMaker(job job, in, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	job(in, out)
}

func crc32signer(data string, out chan string) {
	out <- DataSignerCrc32(data)
}

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for v := range in {
		wg.Add(1)
		go singleHashJob(v.(int), out, wg, mu)
	}
	wg.Wait()
}

func singleHashJob(value int, out chan interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	data := strconv.Itoa(value)
	
	mu.Lock()
	md5 := DataSignerMd5(data)
	mu.Unlock()

	crc32Chan := make(chan string)
	go crc32signer(data, crc32Chan)

	crc32 := <-crc32Chan
	crc32md5 := DataSignerCrc32(md5)

	//result := crc32 + "~" + crc32md5

	/*fmt.Printf("%s SingleHash data %s\n", data, data)
	fmt.Printf("%s SingleHash md5(data) %s\n", data, md5)
	fmt.Printf("%s SingleHash crc32(md5(data)) %s\n", data, crc32md5)
	fmt.Printf("%s SingleHash crc32(data) %s\n", data, crc32)
	fmt.Printf("%s SingleHash result %s\n\n", data, result)*/
	out <- crc32 + "~" + crc32md5
}

func MultiHash(in, out chan interface{}) {
	maxTh := 6
	wg := &sync.WaitGroup{}
	for v := range in {
		wg.Add(1)
		go multiHashJob(v.(string), out, wg, maxTh)
	}
	wg.Wait()
}

func multiHashJob(value string, out chan interface{}, wg *sync.WaitGroup, maxTh int) {
	defer wg.Done()
	results := make([]string, maxTh)
	wgJob := &sync.WaitGroup{}
	muJob := &sync.Mutex{}
	for th := 0; th < maxTh; th++ {
		wgJob.Add(1)
		data := strconv.Itoa(th) + value
		go func(data string, th int, results []string, wgJob *sync.WaitGroup, muJob *sync.Mutex) {
			result := DataSignerCrc32(data)
			muJob.Lock()
			results[th] = result
			muJob.Unlock()
			//fmt.Printf("%s MultiHash: crc32(th+step1)) %v %s\n", value, th, results[th])
			wgJob.Done()
		}(data, th, results, wgJob, muJob)
	}
	wgJob.Wait()
	//resultStr :=
	//fmt.Printf("%s MultiHash result: %s\n\n", value, resultStr)
	out <- strings.Join(results[:], "")
}

func CombineResults(in, out chan interface{}) {
	resultsSlice := make([]string, 0)
	for v := range in {
		resultsSlice = append(resultsSlice, v.(string))
	}
	sort.Strings(resultsSlice)
	//result :=
	//fmt.Printf("CombineResults %s\n", result)
	out <- strings.Join(resultsSlice[:], "_")
}



// сюда писать код
/*
import(
	"strconv"
	"strings"
	"sort"
	"fmt"
	"sync"
	"time"
)

type dataInChanal struct { //identificator for Multihash
	data string
}

var wg sync.WaitGroup
//
//var chNext = make(chan struct{}, 1)
//var inp interface{}
var strmas = []string{}

func ExecutePipeline(jobs ...job){
	in, out := make(chan interface{}), make(chan interface{})
	go jobs[0](in, out)
	//for j := 0;j < 7;j++ {
		for i := 1;i < len(jobs) - 1;i++ {
			wg.Add(1)
			go jobs[i](in, out)
		}
	//}
	wg.Wait()
		sort.Strings(strmas)
		

	go jobs[len(jobs)-1](in, out)
	time.Sleep(time.Millisecond)
}

func SingleHash(in, out chan interface{}) {
	for i := 0;i < 7;i++ {
		dat := <-out
		dataint := dat.(int)
		data := strconv.Itoa(dataint)
		fmt.Print(data, " ")
		
		chandataCrc32 := make(chan string) 
		go func(data string){
			chandataCrc32 <- DataSignerCrc32(DataSignerMd5(data))
		}(data)
		data = DataSignerCrc32(data) + "~"
		dataCrc := <-chandataCrc32
		data += dataCrc
		fmt.Println(data)
		
		in <-dataInChanal{data}
		//inp = dataInChanal{data}
	}
	wg.Done()
}

func MultiHash(in, out chan interface{}){
	for i := 0;i < 7;i++ {
		dat := <-in
		//dat := inp
		//fmt.Println("Multi hash ", i)
		switch dat.(type) {
			case string:
				fmt.Println("Other data ", i) 
				in <-dat
				i--
				continue
		}
		dataInCh := dat.(dataInChanal)
		data := dataInCh.data
		resdata := [7]string{}
		waitGr := sync.WaitGroup{}
		var mu sync.Mutex
		for j := 0;j < 6;j++ {
			waitGr.Add(1)
			go func(j int){
				temp := DataSignerCrc32(strconv.Itoa(j) + data)
				mu.Lock()
				resdata[j] = temp
				mu.Unlock()
				waitGr.Done()
			}(j)
		}
		waitGr.Wait()
		in <- strings.Join(resdata[:], "")
		//inp = strings.Join(resdata[:], "")
	}
	wg.Done()
}

func CombineResults(in, out chan interface{}){
	
	strmas := []string{}
	for i := 0;i < 7;i++ {
		dat := <-in
		//dat := inp
		var data string
		switch dat.(type) {
			case dataInChanal:
				in <- dat
				i--
				continue						
		}
		data = dat.(string)	
		strmas = append(strmas, data)
	}
	sort.Strings(strmas)
	wg.Done()
	in <- strings.Join(strmas, "_")
	
	//in strings.Join(strmas, "_")
}
*/

func main(){
	inputData := []int{0, 1, 1, 2, 3, 5, 8}
	hashSignJobs := []job{
		job(func(in, out chan interface{}) {
			for _, fibNum := range inputData {
				out <- fibNum
			}
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			data, ok := dataRaw.(string)
			if !ok {
				fmt.Println("cant convert result data to string")
				fmt.Println(data)
			}
			
			testResult := data
			fmt.Println(testResult)
			
		}),
	}
	
	ExecutePipeline(hashSignJobs...)
	time.Sleep(time.Millisecond)
	//var s string
	//fmt.Scan(&s)
}


/*
какие были ошибки 
блоблемма с локом горутин
было реализованна передача информации через и оут 
но тут проблемма в том что может лочится передача так как горутины нельзя делать 
буферезированными, ну и не все горутины умеют переключаться 

Решение использовать только 1 канал для горутин с переключением */

/* вариант с лоченьем горутин ни к чему хорошему не привел */

 /* надо либо создавать массив каналов и их передавать */
 
 /* либо лочить ввод и вывод отдельно 
	черзе буферезированый канал
	
 */



