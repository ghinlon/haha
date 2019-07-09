# context switch

# Links

* [performance monitoring - Find out which task is generating a lot of context switches on linux - Server Fault](https://serverfault.com/questions/190049/find-out-which-task-is-generating-a-lot-of-context-switches-on-linux)

# Check

```
// sar [xxx interval count]
sar -w 1 3
       -w     Report task creation and system switching activity.
              proc/s
                     Total number of tasks created per second.
              cswch/s
                     Total number of context switches per second.


vmstat 1 3
   System
       in: The number of interrupts per second, including the clock.
       cs: The number of context switches per second.
```

pid level

```
pidstat -wt 1 3
       -w     Report  task  switching  activity  (kernels 2.6.23 and later
              only).  The following values are displayed:

              PID
                     The identification number of  the  task  being  moni-
                     tored.

              cswch/s
                     Total  number  of voluntary context switches the task
                     made per second.  A voluntary context  switch  occurs
                     when  a  task  blocks  because it requires a resource
                     that is unavailable.

              nvcswch/s
                     Total number of non voluntary  context  switches  the
                     task  made  per second.  A involuntary context switch
                     takes place when a task executes for the duration  of
                     its  time  slice and then is forced to relinquish the
                     processor.

              Command
                     The command name of the task.

       -t     Also display statistics for threads associated with selected tasks.

              This option adds the following values to the reports:

              TGID
                     The identification number of the thread group leader.

              TID
                     The identification number of the thread being monitored.
```



