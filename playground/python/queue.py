from Queue import Queue
from LifoQueue import Queue
from PriorityQueue import Queue

a = Queue()
print "IsEmpty:", a.empty()
print "IsFull:", a.full()
print "Queue size:", a.qsize()

# a.put(item[, block[, timeout]])
# block is true and timeout is None (the default),
print "Put:", a.put("a")

# a.get([block[, timeout]])
# block is true and timeout is None (the default),
print "Get:", a.get()

# join
# Blocks until all items in the queue have been gotten and processed.
# a.join()

# task_down
# task count increase by put or get and decrese by task_down called by put or get subsequently
# a.task_down()

print a
