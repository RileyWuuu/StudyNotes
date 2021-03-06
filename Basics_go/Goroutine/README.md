# Gouroutine Study Note

## Channel
![Goroutine1](https://user-images.githubusercontent.com/71340325/177242997-128fdc54-7eae-46f5-806d-06ef87dab0a2.jpg)

*  a means through which different goroutines communicate.


## Goroutine
* Goroutines are actions or methods that run concurrently with other functions or methods.
* We use channel for goroutines' communication in order to synchronize execution and pass values.
* Concurrency :
    * the capability to deal with lots of things at once. (But not at the same time)
    * ![Goroutine](https://user-images.githubusercontent.com/71340325/177244854-119aa66f-fc93-4997-88a1-6dadeecdc6d4.jpg)
 
## WaitGroup
* Needs to be done through pointer (so that it all points directly to the same struct, avoids call by value situation)
* waits for a collection of goroutines to finish.
* main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished.
![Sync waitgroup](https://user-images.githubusercontent.com/71340325/177271353-d6dc38ad-b639-41a3-99bf-cc995e534730.jpg)

* WaitGroup is mainly used to wait for a group of goroutines to exit. We can specify the number of goroutines we need to wait for to exit by adding, and then decrease by Done. If it is 0, we can exit
