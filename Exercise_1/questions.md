Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?

>concurrency: multiple tasks are making progress, but not necessarily at the same instant. parallelism: multiple tasks are making progress at the same time.
What is the difference between a *race condition* and a *data race*?
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> the scheduler is responsible for deciding which of the runnable threads in the system should be executed by the CPU at any given time.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> threads can be used to solve problems that can be divided into smaller tasks that can be executed concurrently.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> fibers are lightweight threads that are managed by the application instead of the OS. They are faster to create and destroy than OS threads, and they are also faster to switch between.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> it can make the programmer's life harder because of the complexity of managing shared resources and creating more bugs that are hard to reproduce.

What do you think is best - *shared variables* or *message passing*?
> shared variables are easier to understand and use, but message passing is more reliable and less error-prone.


