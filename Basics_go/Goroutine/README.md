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
* Examples:
```
package main
import(
   "fmt"
   "sync"
)
type bank struct{
   balance int
}
func main(){
   var wg sync.WaitGroup
   b:=&bank{}
   wg.Add(3)
   go func(){
      deposit(1000)
      wg.Done()
   }()
      go func(){
      deposit(1050)
      wg.Done()
   }()
      go func(){
      deposit(2000)
      wg.Done()
   }()
   wg.Wait()
   fmt.Println(b.balance)
}
func (b *bank) deposit(amount int){
   b.balance += amount
}
```
 
## WaitGroup
* Needs to be done through pointer (so that it all points directly to the same struct, avoids call by value situation)
* waits for a collection of goroutines to finish.
* main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished.
![Sync waitgroup](https://user-images.githubusercontent.com/71340325/177271353-d6dc38ad-b639-41a3-99bf-cc995e534730.jpg)

* WaitGroup is mainly used to wait for a group of goroutines to exit. We can specify the number of goroutines we need to wait for to exit by adding, and then decrease by Done. If it is 0, we can exit

## Mutex
![image](https://user-images.githubusercontent.com/71340325/183039270-224caeb0-120a-4c08-b4ab-375fa03bdba4.png)
> Illustration created for “A Journey With Go”, made from the original Go Gopher, created by Renee French.

* A method used as a lock to ensure that only one Goroutine is accessing the critical section of code at any point of time.
   * Two methods:
      * Lock
      * Unlock
   * Code between a call to lock and unlock will be executed by only one Goroutine.
   ```
   mutex.Lock()
   x += 1
   mutex.Unlock()
   ```
 * Examples:
   * We have a bank where we can deposit and check our balance. But we can't provide services to clients one by one
   ```
   //the solutions might be: (following the code above)
   //suppose we have 1000 actions(each action deposits 1000 into account) a time
   func main(){
      var wg sync.WaitGroup
      b :=&bank{}
      
      n:=1000
      wg.Add(n)
      for i:=0;i<n;i++{
         go func(){
            b.Deposit(1000)
            wg.Done()
         }()
      }
      fmt.Println(b.balance)
   }
   ```
   
   > However, the result turns out to be 946000 (which we're expecting 1000*1000) <br>
   > This is called "race conditions", a situation that occurs when a device attempts to perform more than one operations at the same time.
* When situations like this happens, we use mutex to constraint the device to run one goroutine at a time.
```
type bank struct{
   balance int
   mux sync.Mutex
}
func main(){
   var wg sync.WaitGroup
      b :=&bank{}
      
      n:=1000
      wg.Add(n)
      for i:=0;i<n;i++{
         go func(){
            b.Deposit(1000)
            wg.Done()
         }()
      }
      fmt.Println(b.balance)
}
func (b *bank)deposit int{
   b.mux.Lock()
   b.balance += amount
   b.mux.Unlock()
}
```
