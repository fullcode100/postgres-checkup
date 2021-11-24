package h004

const MSG_REDUNDANT_INDEXES_FOUND_P2_CONCLUSION string = "[P2] %d redundant index(es) have been found and their total size " +
	"exceeds %.2f%% of the database size."
const MSG_REDUNDANT_INDEXES_FOUND_P3_CONCLUSION string = "[P3] %d redundant index(es) have been found."

const MSG_REDUNDANT_INDEXES_FOUND_R1 string = "Use the database migration provided below to drop the redundant indexes. " +
	"Keep in mind, that under load, it is recommended to use `DROP INDEX CONCURRENTLY` (and `CREATE INDEX CONCURRENTLY` " +
	"if reverting is needed) to avoid blocking issues."
const MSG_REDUNDANT_INDEXES_FOUND_R2 string = "Be careful dropping the indexes. If you have multiple setups of your " +
	"software, the analysis of just a single setup might be not enough. Some indexes might be used (and therefore, needed) " +
	"only on a limited number of setups. Also, in some cases, developers prepare indexes for new features in advance – " +
	"in such cases, dropping those indexes is not a good idea."
const MSG_REDUNDANT_INDEXES_FOUND_R3 string = "If there are some doubts, consider a more careful approach. Before actual" +
	"dropping, indexes disable listed in this report. For this, use queries like `UPDATE pg_index SET indisvalid = false " +
	"WHERE indexrelid::regclass = (select oid from pg_class where relname = 'u_users_email');. Indexes will " +
	"continue to get updates. In case of some performance degradations, re-enable the corresponding indexes, " +
	"setting indisvalid to true. If everything looks fine, after a significant period of observations, " +
	"proceed with DROP INDEX CONCURRENTLY."

const MSG_REDUNDANT_INDEXES_FOUND_DO string = "\"DO\" database migrations  \n%s"
const MSG_REDUNDANT_INDEXES_FOUND_UNDO string = "\"UNDO\" database migrations  \n%s"
