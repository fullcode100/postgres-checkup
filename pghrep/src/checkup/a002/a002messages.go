package a002

// Analysis of Postgres major version
const MSG_UNKNOWN_VERSION_CONCLUSION string = "[P1] PostgreSQL major version `%s` is unknown (on `%s`). " +
    "It means that PostgreSQL Global Development Group does not support your PostgreSQL version. " +
    "In case of bugs and security issues you are on your own.  \n" // NikolayS: very rare; most likely it's a devel version.
const MSG_UNKNOWN_VERSION_RECOMMENDATION string = "[P1] On `%s`, consider using one of supported major versions.  \n"
const MSG_NOT_SUPPORTED_VERSION_CONCLUSION string = "[P1] Postgres major version being used is `%s` and it is " +
	"NOT supported by PostgreSQL Global Development Group. The support has ended: `%s`. This is a major issue. New bugs and security " +
	"problems will not be fixed by the Postgres community. You are on your own! Read more: " +
	"[Versioning Policy](https://www.postgresql.org/support/versioning/).  \n"
const MSG_NOT_SUPPORTED_VERSION_RECOMMENDATION string = "[P1] Consider upgrading Postgres version `%s` to one of the " +
	"versions supported by the community. To minimize downtime, consider using [pg_upgrade](https://www.postgresql.org/docs/current/pgupgrade.html) " +
	"or one of the solutions based on logical replication.  \n"
const MSG_LAST_YEAR_SUPPORTED_VERSION_CONCLUSION string = "[P2] PostgreSQL Global Development Group will stop supporting version `%s`" +
	" within the next 12 months. End of life is scheduled `%s`. After that, you will be on your own!" +
	" Read more: [Versioning Policy](https://www.postgresql.org/support/versioning/).  \n"
const MSG_SUPPORTED_VERSION_CONCLUSION string = "Postgres major version being used is `%s` and it is " +
	"currently supported by PostgreSQL Global Development Group. End of life is scheduled %s. It means that in case " +
	"of bugs and security issues, updates (new minor versions) with fixes will be released and available for use." +
	" Read more: [Versioning Policy](https://www.postgresql.org/support/versioning/).  \n"
const MSG_NOT_LAST_MAJOR_VERSION_CONCLUSION string = "[P3] Consider upgrading to the newest major version: %.0f. " +
    "It has a lot of new features and improvements.  \n"

// Analysis of minor version
const MSG_LAST_MINOR_VERSION_CONCLUSION string = "`%s` is the most up-to-date Postgres minor version in the branch `%s`.  \n"
const MSG_NOT_LAST_MINOR_VERSION_CONCLUSION_1 string = "[P2] The minor version being used (`%s`) is not up-to-date " +
    "(the newest version: `%s`). See [the full list of changes between %s and %s](https://why-upgrade.depesz.com/show?from=%s&to=%s).  \n"
const MSG_NOT_LAST_MINOR_VERSION_CONCLUSION_N string = "[P2] The minor versions being used (`%s`) are not up-to-date " +
    "(the newest version: `%s`). See [the full list of changes between %s and %s](https://why-upgrade.depesz.com/show?from=%s&to=%s).  \n"
const MSG_NOT_ALL_VERSIONS_SAME_CONCLUSION_1 string = "[P2] Postgres minor version on the master and replica(s) differ. Node `%s` uses Postgres `%s`.  \n"
const MSG_NOT_ALL_VERSIONS_SAME_CONCLUSION_N string = "[P2] Postgres minor version on the master and replica(s) differ. Nodes `%s` use Postgres `%s`.  \n"
const MSG_NOT_ALL_VERSIONS_SAME_RECOMMENDATION string = "[P2] Please upgrade Postgres so its versions on all nodes match.  \n"
const MSG_ALL_VERSIONS_SAME_CONCLUSION string = "All nodes have the same Postgres version (`%s`).  \n"

const MSG_NOT_LAST_MINOR_VERSION_RECOMMENDATION string = "[P2] Consider performing upgrade to the newest minor version: `%s`.  \n"
const MSG_GENERAL_RECOMMENDATION_1 string = "For more information about minor and major upgrades see:  \n" +
	"    - Official documentation: https://www.postgresql.org/docs/current/upgrading.html  \n"

const MSG_GENERAL_RECOMMENDATION_2 string = "    - [Major-version upgrading with minimal downtime](https://www.depesz.com/2016/11/08/major-version-upgrading-with-minimal-downtime/) (depesz.com)  \n" +
	"    - [Upgrading PostgreSQL on AWS RDS with minimum or zero downtime](https://medium.com/preply-engineering/postgres-multimaster-34f2446d5e14)  \n" +
	"    - [Near-Zero Downtime Automated Upgrades of PostgreSQL Clusters in Cloud](https://www.2ndquadrant.com/en/blog/near-zero-downtime-automated-upgrades-postgresql-clusters-cloud/) (2ndQuadrant.com)  \n" +
	"    - [Updating a 50 terabyte PostgreSQL database](https://medium.com/adyen/updating-a-50-terabyte-postgresql-database-f64384b799e7)  \n"
