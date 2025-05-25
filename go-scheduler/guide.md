┌───────────── minute (0 - 59)
│ ┌───────────── hour (0 - 23)
│ │ ┌───────────── day of month (1 - 31)
│ │ │ ┌───────────── month (1 - 12)
│ │ │ │ ┌───────────── day of week (0 - 6) (Sunday=0)
│ │ │ │ │
│ │ │ │ │
* * * * * 


Task | Cron Expression | Notes
Product list sync | "*/5 * * * *" | Every 5 mins
Product details update | "0 2 * * *" | Daily at 2 AM
Backup | "30 3 * * 1" | Mondays at 3:30 AM
