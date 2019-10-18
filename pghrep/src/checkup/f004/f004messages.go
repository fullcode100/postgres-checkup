package f004

const MSG_NO_RECOMMENDATIONS string = "All good 👍\n"
const MSG_TOTAL_BLOAT_EXCESS_CONCLUSION string = "[P1] Total table (heap) bloat estimation is ~%s, it is %.2f%% of the overall size of all tables and %.2f%% of the DB size. " +
	"Removing the table bloat will reduce the total DB size down to ~%s. Free disk space will be increased by ~%s. " +
	"Total size of tables is %.2f times bigger than it could be. " +
	"Notice that this is only an estimation, sometimes it may be significantly off.\n"
const MSG_TOTAL_BLOAT_LOW_CONCLUSION string = "The estimated table (heap) bloat in this DB is low, just ~%.2f%% (~%s). No action is needed now. Keep watching it though.\n"
const MSG_BLOAT_CRITICAL_RECOMMENDATION string = "[P1] Reduce and prevent the high level of table bloat:\n" +
	"    - to prevent a high level of bloat in the future, tune autovacuum: consider more aggressive autovacuum settings (see F001);\n" +
	"    - eliminate or reduce the current table bloat using one of the approaches listed below.\n"
const MSG_BLOAT_WARNING_RECOMMENDATION_TABLES string = "The following tables have size > 1 MiB and table bloat estimate > %.2f%%. Use this list to reduce the bloat applying one of the approaches described below. Here are these tables: %s."
const MSG_BLOAT_WARNING_RECOMMENDATION string = "[P2] Consider the following:\n" +
	"    - to prevent a high level of bloat in the future, tune autovacuum: consider more aggressive autovacuum settings (see F001);\n" +
	"    - eliminate or reduce the current table bloat using one of the approaches listed below.\n"
const MSG_BLOAT_GENERAL_RECOMMENDATION_1 string = "If you want to get exact bloat numbers, clone the database, get table sizes, then apply " +
	"database-wide `VACUUM FULL` (it eliminate all the bloat), and get new table sizes. Then compare old and new numbers.\n"
const MSG_BLOAT_GENERAL_RECOMMENDATION_2 string = "To reduce the table bloat, consider one of the following approaches:\n" +
	"    - [`VACUUM FULL`](https://www.postgresql.org/docs/current/sql-vacuum.html) (:warning:  requires downtime / maintenance window),\n" +
	"    - one of the tools reducing the bloat online, without interrupting the operations:\n" +
	"        - [pg_repack](https://github.com/reorg/pg_repack),\n" +
	"        - [pg_squeeze](https://github.com/cybertec-postgresql/pg_squeeze),\n" +
	"        - [pgcompacttable](https://github.com/dataegret/pgcompacttable).\n"
const MSG_BLOAT_PX_RECOMMENDATION string = "Read more on this topic:\n" +
	"    - [Bloat estimation for tables](http://blog.ioguix.net/postgresql/2014/09/10/Bloat-estimation-for-tables.html) (2014, ioguix)\n" +
	"    - [Show database bloat](https://wiki.postgresql.org/wiki/Show_database_bloat) (PostgreSQL wiki)\n" +
	"    - [PostgreSQL Bloat: origins, monitoring and managing](https://www.compose.com/articles/postgresql-bloat-origins-monitoring-and-managing/) (2016, Compose)\n" +
	"    - [Dealing with significant Postgres database bloat — what are your options?](https://medium.com/compass-true-north/dealing-with-significant-postgres-database-bloat-what-are-your-options-a6c1814a03a5) (2018, Compass)\n" +
	"    - [Postgres database bloat analysis](https://about.gitlab.com/handbook/engineering/infrastructure/blueprint/201901-postgres-bloat/) (2019, GitLab)\n"
const MSG_BLOAT_WARNING_CONCLUSION_1 string = "[P2] There is %d table with size > 1 MiB and table bloat estimate >= %.2f%% and < %.2f%%:  \n%s\n"
const MSG_BLOAT_CRITICAL_CONCLUSION_1 string = "[P1] The following %d table has significant size (>1 MiB) and bloat estimate > %.2f%%:  \n%s  \n"
const MSG_BLOAT_WARNING_CONCLUSION_N string = "[P2] There are %d tables with size > 1 MiB and table bloat estimate >= %.2f%% and < %.2f%%:  \n%s  \n"
const MSG_BLOAT_CRITICAL_CONCLUSION_N string = "[P1] The following %d tables have significant size (>1 MiB) and bloat estimate > %.2f%%:  \n%s  \n"
const TABLE_DETAILS string = "    - `%s`: size %s, can be reduced %.2f times, by ~%s (~%.2f%%)  \n"
