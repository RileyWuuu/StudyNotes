# Gouroutine Study Note

## Channel 
![Channel](https://user-images.githubusercontent.com/71340325/183274342-f7c82641-6e28-4927-9142-36f6c55fe2ac.jpg)

*  a means through which different goroutines communicate.


## Goroutine
* Goroutines are actions or methods that run concurrently with other functions or methods.
* We use channel for goroutines' communication in order to synchronize execution and pass values.
* Concurrency :
    * the capability to deal with lots of things at once. (But not at the same time)
    * ![Goroutine](https://user-images.githubusercontent.com/71340325/183274211-2d283d1e-5f24-4734-8328-9f9756d5803f.jpg)

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
![WaitGroup](https://user-images.githubusercontent.com/71340325/183274069-897977a7-71b5-40d7-9f37-f2c740204b27.jpg)



* WaitGroup is mainly used to wait for a group of goroutines to exit. We can specify the number of goroutines we need to wait for to exit by adding, and then decrease by Done. If it is 0, we can exit

## Mutex
![mutex](https://user-images.githubusercontent.com/71340325/183273747-83c4bc7c-6ba5-4f39-a0db-8a88b225fbe4.jpg)


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
   
   > However, the result turns out to be 995000 (which we're expecting 1000*1000) <br>
   > This is called "race conditions", a situation that occurs when a device attempts to perform more than one operations at the same time.
* When situations like this happens, we use mutex to constraint the device to run one goroutine at a time.
```
type bank struct{
   balance int
   mux sync.RWMutex
}
func main(){
   var wg sync.WaitGroup
   b :=&bank{}
   n:=1000
   wg.Add(n)
   for i:=0;i<n;i++{
      go func(){
         b.Deposit(1000)
         log.Printf("(Write) Deposit amount: %v",1000)
         wg.Done()
      }()
   }
   wg.Add(n)
   for i:=0;i<n;i++{
      go func(){
         log.Printf("(Read) Balance: %v",b.balance())
         wg.Done()
      }()
   }
   wg.wait
}
func (b *bank)deposit int{
   b.mux.RWLock()
   time.Sleep(time.Second)
   b.balance += amount
   b.mux.RWUnlock()
}
func (b *bank)balance()(balance int){
   b.mux.RWLock()
   time.Sleep(time.Second)
   balance = b.balance
   b.mux.RWUnlock()
   return
}
```
> In the code above, we use sync.RWMutex, meaning a lock that allows multiple readers but only one writer, so that reading and writing will be processed simultaneously.


## sync/atomic
* when we need to implement synchronization algorithms, we can use mutex or package atomic
* atomic works through CPU command
* Commands:
   *  Add
   *  CompareAndSwap
   *  Load
   *  Store
   *  Swap

### Differences between atomic and mutex
* same in **behavior** but different in **performance**
![mutexVSatomic](https://user-images.githubusercontent.com/71340325/183278863-4cab83e9-504c-4856-a799-5616b4ad6aa7.jpg)
