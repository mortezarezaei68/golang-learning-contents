Q1) Why we use wait and close channel in go func?


wg.Wait() یعنی صبر کن همه goroutine ها تموم بشن
جلوگیری از:
خروج زودهنگام main
goroutine leak
Close channel:
دیگه داده نمی‌فرستم
consumer می‌فهمه range کی تموم بشه
اگر نبندی → برنامه ممکنه hang کنه


مثال کوتاه:
go func() {
    wg.Wait()  
    close(ch)  
}()


Q2) What is blocking and deadlock in goroutine?

Blocking
یعنی goroutine منتظر send/receive می‌مونه


مثال:
ch := make(chan int)

go func() {
    ch <- 1 
}()

fmt.Println(<-ch)


Deadlock
یعنی همه goroutine ها بلوک می‌شن


هیچ‌کس آزاد نمی‌کنه → برنامه کرش می‌کنه


مثال:
ch := make(chan int)
ch <- 1     
fmt.Println("never reached")
// fatal error: all goroutines are asleep - deadlock!


Q3) Why this fanIn implementation goes wrong?
مشکل‌های اصلی:
out unbuffered است → اگر consumer آماده نباشه، goroutine داخل fanIn block می‌شود


وقتی ورودی‌ها بسته می‌شن، ارسال به out ممکنه قبل از خروج باعث hang شود


همچنین اگر یکی از کانال‌ها خیلی کند باشد، select ممکنه مدت طولانی منتظر بماند


 رفتار درست این است که:
کانال‌ها بعد از بسته شدن با nil غیرفعال می‌شن (کار درسته)


ولی باید مطمئن باشی receiver همیشه هست


راه امن‌تر:
out := make(chan int, 1) // buffer to prevent blocking

یا استفاده از wg:
func fanIn(ch1, ch2 <-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    forward := func(ch <-chan int) {
        defer wg.Done()
        for v := range ch {
            out <- v
        }
    }

    wg.Add(2)
    go forward(ch1)
    go forward(ch2)

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}



