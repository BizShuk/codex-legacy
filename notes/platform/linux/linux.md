# linux

# TODO: clean up

### cmd

sar -P ALL 1
sar -q
sar -q 1 1
sar -r
strace -o hello.log ./hello
free

taskset -c 0,4 ./sched num num num
    -c 0,4 => cpu 0, cpu 4 no share cache

time taskset -c 0 ./sched 1 10000 1000

ps -eo pid,comm,etime,time

nice -n 5 echo hello

readelf => output virtual memory
cat /proc<pid>/maps => output virtual memory

swapon --show
thrashing
sar -W
sar -S

ulimit -v

dd if=/dev/zero of=testfile oflag=direct bs=1M count=1K

sar -B1
sar -d-p

##### latency

一個程式完成的時間 可能被插隊延長

##### throuput

### 記憶體

OOM => sysctl vm.panic_on_om => default 0 -> 1
記憶體碎片化 存取到其他用途記憶體 多行程處理困難

虛擬記憶體 分頁表
=> 分頁無效
linux => demand paging
SIGSEGV

### copy on write

fork => copy virtual mem of parent (lock for write of parent of child process)

### SIZE

X86_64 128TB, 4KB/page 8byte/page item

大型分頁 /sys/kernel/mm/tranparent_hugepage/enabled (always/never)

# fs

### copy from device to device

dd if=/dev/sdx of=/dev/sdx bs=64K conv=noerror,sync status=progress

### partition

sudo parted
    mkpart [primary|extended|logical] start end

### list block

lsblk -f

###

df -h

sudo umount /path/to/dir
sudo mount /dev/sdx /path/to/dir
sudo mkfs -t ext4 /dev/sdx
