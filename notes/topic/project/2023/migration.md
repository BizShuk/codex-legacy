# Migration

## Vision/OKR

### Build comprehensive migration system

[2023Q3]

- Cross-IDC RPC for async[shuk]
- Collect all Dao into async/migrate_sdk
- Async Read-Send MQ flow

[2023Q4]

- Data IDC design refactor(read/write/delete)
- Data IDC mark
- DataMigratedVerification
- Online comparison and delete flow
- New modules (ES/Offline DB/redis)
- Migration Monitoring Dashboard
- Cross-IDC RPC for async Monitoring Dashboard
- Performance

### EU-TTP deployment/data migration

- Service Deployment(10/01, 10/15~)
- Data Migration(11/20~)

### Service Responsibility refactor

- Access level design(active/inactive/migrating/suspend)
  - User read/write
  - Group read/write, admin, migration, wallet
  - All read only
  - No read/write
- Wallet Account
- Income+/C2C Payout/C&S

### E2E Payment Traffic Route Topology

- Analyze the E2E traffic route
- Systematically generate the report/platform
