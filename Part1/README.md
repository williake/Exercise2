# Mutex and Channel basics

### What is an atomic operation?
> An atomic operation is an operation that can is indivisible, either the operation gets done or it doesn't; it's not possible to stop the operation halfway.

### What is a semaphore?
> a semaphore is a variable that keeps track of a resource, if it is available or not. 0 usually signifies that it's locked or unavailable. 1 or higher means that it is unlocked/available.

### What is a mutex?
> Mutex stands for mutual exclusion. multiple processes can access the object, but not at the same time.

### What is the difference between a mutex and a binary semaphore?
> a binary semaphore is a type of mutex.

### What is a critical section?
> The critical sections are parts of the program where access to to shared resources are protected. that it cannot accesed by more than one process.

### What is the difference between race conditions and data races?
 > a data race is when 2 instrunctions from different threads access the same memory location at the same time, with atleast one of them a write, without synchronizing the order of the instructions. -> leads to different results.
 a race condition are errors where timing and and order leads to wrong program behaviour. data race and race conditions are not mutually exlusive, and one does not lead to the other or vice versa.

### List some advantages of using message passing over lock-based synchronization primitives.
> it's easier to scale. It's better for bigger systems, like for example the internet.  

### List some advantages of using lock-based synchronization primitives over message passing.
> it's faster over message passing. easier to implement.
