# Disaster Recovery

### Backup and Recovery
Backup your system periodically.

##### Strategy

- RTO, Recovery Time Objective
    
    Backup frequency

- RPO, Recovery Point Objective
    
    Recovery time from failure

### Pilot Light Approach
經常性上傳backup 然後用backup恢復環境, few mins.

### Warm Standby Approach, Active-Passive Failover
一台即時備份 當災難發生時 轉移EIP到目標伺服器 並啟動Auto-Scaling機制

### Multi-Site Approach, Active-Active Failover
創造多個環境伺服器 並分擔流量, 在災難發生在其中一個區域時 繼續使用其他區域分擔流量