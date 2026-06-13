# Monitoring Dashboard

ByteDance has a unique mointoring tool called Metrics.
It was very easy to trace down the issue, no matter is upstream or downstream.

But grafana gives a great advantages of overall monitoring.

I created a overall general component monitoring dashboard.

1. service overview, including CPU, MEM, panic, tce aborting, qps/error(upstream/server/downstream), single pod issue
2. DB overview, cpu, mem, disk, row lock, qps db command, read/qps, disk latency
3. DB proxy overview, cpu, db return error
4. MQ overview, in/out(qps,network), message stack(by queue, imbalance consumption speed)
5. service overall, basic for all IDCs
6. migration dashboard.
7. redis overview
