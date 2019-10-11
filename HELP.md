# A. General  / Infrastructural

This group determines the available resources such as hardware characteristics or OS version and settings, as well as current resource usage and PostgreSQL settings. This information helps to identify suboptimal settings and existing or potential issues related to environment where Postgres works. It may be used for creating recommendations for improving database performance. A-group also serves as a basis for some reports in other groups.

### A001 System, CPU, RAM, Virtualization

General information about operational systems where the observed Postgres master and its replicas operate.

> Insights:
>
> - Hardware and software differences (OS versions, Linux kernel versions, CPU, Memory). If the observed master and its replicas run on different platforms, it might cause issues with binary replication.
>
> - Memory settings tuning. (Examples: is swap enabled? Are huge pages used?) Observing state of memory about memory consumption by database may lead to recommendations of changes to improve system performance.
>
> - Information about virtualization type.


### A002 Postgres Version Information

This report answers the following questions:
- Do all nodes have the same Postgres version?
- Is the minor version being used up-to-date? Keeping the minor version of the database up-to-date is recommended to decrease chances to encounter with bugs, performance and security issues?
- Is the major version currently supported by the community?
- Will the major version be supported by the community during the next 12 months?
- If the minor version is not the most recent, are any critical bugfixes released that need to be applied ASAP?


### A003 Postgres Settings

Shows all PostgreSQL settings and their values grouped into categories.
May be used as a quick reference.


### A004 Cluster Information

A quick overview of "what is happening with the observed Postgres nodes?".

The following is included:

- The uptime. Sometimes low uptime may indicate an unplanned, accidental restart of the database.
- General information: how many databases are on one instance, what is their size, replication mode, age of statistics.
- Information about replicas, replication modes, replication delays.
- Ratio of forced checkpoints among all checkpoints registered since statistics reset time.
> Insights: Frequent checkpoints in most cases create an excessive load on the disk subsystem. Identifying this fact will allow the more optimal disk utilization.
- How big is the observed database (the cluster may have multiple databases)?
> Insight: if the database is smaller than RAM, there are good chances to avoid intensive disk IO in most operations
- Cache Effectiveness: percentage of buffer pool hits.
> Insight: if it is not more than 95% on all nodes, it might be a good sign that the buffer pool size needs to be increased.
- Successful Commits: percentage of successfully committed transactions.
> Insight: if the value is not more than 99%, it might be a sign of logic issues with application code leading to high rates of ROLLBACK events.
- Temp Files per day: how many temporary files were generated per day in average, since last statistics reset time.
> Insight: if this value is high (thousands), it is a signal that work_mem should be increased.
- Deadlocks per day.
> Insight: significant (dozens) daily number of deadlocks is a sign of issues with application logic that needs redesign.

### A005 Extensions

Provides a list of all available and installed (in the current observed database) extensions, with versions. Insight: if there is a newer version of an installed extension, the report will highlight it, meaning that update is needed.

### A006 Postgres Setting Deviations

Helps to check that there are no differences in Postgres configuration on the observed nodes (except `transaction_read_only` and pg_stat_kcache’s `linux_hz`).

> Insights:
> - In general, any differences in configuration on master and its replicas might lead to issues in case of failover. An example: the master is tuned, while replicas are not tuned at all or tuned poorly, in the event of failover, a new master cannot operate properly due to poor tuning.


### A007 Altered Settings

There are multiple ways to change database settings globally:

- explicitly, in the configuration file postgresql.conf, and
- implicitly, using 'ALTER SYSTEM' commands.

This report checks if there are settings which were set by implicit (ALTER SYSTEM) way.

Possible sources of configuration settings (presented in the first column of the report’s table):

* `postgresql.auto.conf`: changed via 'ALTER SYSTEM' command.
* `%any other file pattern%`: changed in additional config included to the main one.
* `postgresql.conf`: non-default values are set in postgresql.conf.

### A008 Disk Usage and File System Type

Shows detailed file systems information related to Postgres database.
 
> Insights:
> - Is there enough free disk space?
> - Are there no network file systems for Postgres such as NFS? Use of a network file system reduces the performance of the database and might lead to data corruption.
> - Is stats_temp_directory located in RAM (tmpfs)? The default location of the statistics collector directory inside PGDATA, so it might create excessive load on the disk subsystem.
> - Are file systems equal across all observed nodes (comparing master with replicas)?

### A010 Data Checksums, wal_log_hints

This report gives understanding, how securely the database stores data.

> Insights:
> - When the checksums are enabled, the pg_verify_checksums utility can be used (added in version 11; renamed to pg_checksums in version 12).
> - pg_rewind utility requires wal_log_hints to be turned on (default value is off).
# D. Monitoring / Troubleshooting

Reports of this group help to understand if the database configuration settings for collecting statistics are correct. Without statistics, monitoring systems will not be able to function well.


### D002 Useful Linux Tools

Checks if some common diagnostics Linux tools are installed on the system.

> Insights:
> It is worth having at least one or two utilities in every category (memory, CPU, network, I/O, etc.).
> Such tools should be installed in advance to diagnose incidents in a timed fashion.


### D004 pg_stat_statements, Tuning opts, pg_stat_kcache

Checks if the extensions for SQL query analysis are available and installed in the
current database: pg_stat_statements and pg_stat_kcache.
If extenstions are available, their settings will be shown.

> Insights:
> - Query analysis reports (section K) require at least pg_stat_statements to be installed.
> - In addition, pg_stat_statements are used by most monitoring systems.
> - Inspection of pg_stat_statements is essential for correct query analysis.
> - For example, `pg_stat_statements.track = all` might double some SQL queries if stored procedures are used (to avoid duplicates and understand workload correctly, `top` should be used; however, it will exclude in-function SQL statements).


# F. Autovacuum, Bloat

### F001 Autovacuum: Current Settings

Shows global and per-table (if any) autovacuum-related Postgres settings.

> Insights:
> - Is any tuning applied (values are not default)?
> - Are there any custom table autovacuum settings? There are cases when the tables have a custom autovacuum configuration. Tracking such tables will allow you to understand the nature of the functioning of autovacuum workers. Such tables are marked with asterisk (\*) in the following reports.

### F002 Autovacuum: Transaction ID Wraparound Check

Shows a distance in % to transaction ID wraparound disaster for every database.

> Insights:
> If % is higher than 50%, autovacuum tuning should be considered as soon as possible.
> By identifying objects that are older than the specified threshold, settings for adjusting the autovacuum settings may be suggested.


### F003 Autovacuum: Dead Tuples

Current levels of dead tuples in tables in the observed database.

> Insights:
> The report's key metric, "Dead Tuples Ratio, %", shows the percentage of dead tuples in the tables.
> The column "Since the last autovacuum" gives understanding about the effectiveness of the autovacuum tuning.


### F004, F005 Autovacuum: Heap Bloat, Index Bloat (Estimated)

Estimated table and index bloat is presented in this report.

> Insights:
> - Objects with a high percentage of bloat lead to wasted disk space, degradation in query performance, additional CPU costs, and excessive read load on the disk.
> This report is based on estimations. The errors in bloat estimates may be significant (in some cases, up to 15% and event more). Use it only as an indicator of potential issues.
> - Checks the following things:
>     - Extreme (>90%) level of heap or index bloat estimated.
>     - Significant (>40%) level of heap or index bloat estimated.

### F008 Autovacuum: Resource Usage

Shows a table with Postgres settings related to autovacuum resource usage.

> Insights:
> - Is  `autovacuum_max_workers`  not default? (When CPU cores or vCPUs >= 10).
> - Is `autovacuum_vacuum_cost_limit` / `vacuum_cost_limit` not default?
> - Isn't `maintenance_work_mem` / `autovacuum_work_mem` too low compared to table sizes and RAM?


# G. Performance / Connections / Memory-related Settings

### G001 Memory-related Settings

Shows Postgres settings related to memory usage.

> Insights:
> - Memory management in PostgreSQL is important for good database performance.
> - Questions worth answering:
>     - Are the Resource Consumption parameters tuned?
>     - Is the level of OOM risks low or high?
>     - Can any signs of sub-optimal behaviour be observed from memory usage?

### G002 Connections and Current Activity

A detailed snapshot report of all connections, grouped by users, databases and state type.

> Insights:
> - The report helps to detect the count of potentially bad conditions like:
>     - Many "idle in the transaction" sessions. Are there any "idle in transaction" connections with state changed more than 1 minute / 1 hour ago, or with transaction age more than 1 minute / hour?
>     - Long active connections. Are there any  "active" connections with state changed more than 1 minute / 1 hour ago?
>     - How close is `max_connections`?

### G003 Timeouts, Locks, Deadlocks

Provides information about how "timeout" and locking-related settings are tuned, shows deadlocks counter for every database since statistics reset.

> Insights:
> - Questions worth answering:
>     - Is `statement_timeout` > 0 and <= 30 seconds (good choice for an OLTP system)?
>     - Is `idle_in_transaction_session_timeout` >0 and < 20 minutes (preventing autovacuum and locking issues)?
>     - Is `max_locks_per_transaction` not default (for example, low value may interrupt pg_dump)?

# H. Index Analysis

### H001 Invalid Indexes

The list of broken indexes (invalid state) to be removed or reindexed.

### H002 Unused and Redundant Indexes

Shows the list of never used and rarely used indexes.
Helps to understand how much space they occupy.

> Insights:
> - Questions worth answering:
>     - Is the total size of unused indexes less than 10% of the DB size (only if statistics is older than 1 week)?
>     - Is statistics saved across restarts?
> - If statistics age is low, the report should be used with caution.

### H003 Non-indexed Foreign Keys

Checks if all foreign keys have indexes in referencing tables.

### H004 Redundant Indexes

Shows the list of redundant indexes.
Helps to understand how much space they occupy.


# K. SQL Query Analysis

This is important group of reports providing deep SQL query analysis, based on pg_stat_statements and (optional) pg_stat_kcache.

### K001 Globally Aggregated Query Metrics

Aggregated statistics about all queries performed during the observation period.

> Insights:
> - One of the most interesting metrics is `s/sec` ("seconds per second") in the "Total time" column. It signals roughly how many CPU cores should be present to execute queries. `3s/sec` would mean "3 CPUs are needed to process these queries" (w/o considering context switches). This might be very helpful for capacity planning.

### K002 Workload type ("First-word" Analysis)

Helps to understand which type of workload is the most frequent (selects, inserts, updates, etc.) during the observation period.
The grouping is based on the first word of every query.

### K003 Top Queries by total_time

One of the most comprehensive and deep reports. Shows Top query groups
ordered by total execution time during the observation period (`total_time` in
pg_stat_statements). Good start for query optimization.

> Insights:
> - The first question to answer: Are there any query groups with `total_time` ratio >50% of overall `total_time`?  If we have this type of query, it is definitely worth optimizing it.
> - Full query text is available by the link below each query group.


# L. DB Schema Analysis

The reports of this group are designed for architectural checks, which are
crucial for making decisions on changing the structure of the database in the
face of a growing amount of data.

### L001 Table Sizes

Displays the size of tables and their components (indexes, TOAST, the table itself).

> - Questions worth answering:
>     - Does the size of indexes for each table not exceed heap (with toast) size?
>     - Are there any non-indexes tables which size is > 10 MiB?
>     - Are there any non-partitioned tables of size > 100 GiB?

### L003 Integer (int2, int4) Out-of-range Risks in PKs

Shows primary keys with risks of integer capacity overflow (reached above 10%).
If the capacity of the primary key is exhausted, this will most likely lead to
the shutdown of the service.

This report helps to protect the database from disaster on integer overflow.
