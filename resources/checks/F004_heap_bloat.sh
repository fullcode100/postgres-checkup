sql=$(curl -s -L https://raw.githubusercontent.com/NikolayS/postgres_dba/f1effb54dcfcc7075960a3a51a412d4d2796064a/sql/b1_table_estimation.sql | awk '{gsub("; *$", "", $0); print $0}')

${CHECK_HOST_CMD} "${_PSQL} -f -" <<SQL
with data as (
  $sql
), limited_data as (
  select * from data limit 100
)
select json_object_agg(ld."Table", ld) as json from limited_data ld;

SQL
