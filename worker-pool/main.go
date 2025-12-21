package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct{
	ID int
}


func worker(id int, jobs <- chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("woker #%d percessing job #%d\n", id, job.ID)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	numWorkers := 3
	numJobs := 10
	var wg sync.WaitGroup

	jobs := make(chan Job, numJobs)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)

	}
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j} 
	}
	close(jobs)
	
	wg.Wait()
}