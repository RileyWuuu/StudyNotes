# Find the prime numbers from the given integer

## solutions:
> Ofcourse we can use a simple for loop to run through the given number. <br> But by testing the code, we'll find out that it takes a lot of time. <br>
Therefore we use goroutines to make it faster.

> We can use time.Sleep to make sure our main func stays awake, but a more convenient way is to use waitgroup. (Since by using time.Sleep, we're won't know how long we need)