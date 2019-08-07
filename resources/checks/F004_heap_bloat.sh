if [[ ! -z ${IS_LARGE_DB+x} ]] && [[ ${IS_LARGE_DB} == "1" ]]; then
  MIN_RELPAGES=100
else
  MIN_RELPAGES=0
fi

${CHECK_HOST_CMD} "${_PSQL} -f -" <<SQL
with data as (
  with overrided_tables as (
    select
      pc.oid as table_id,
      pn.nspname as scheme_name,
      pc.relname as table_name,
      pc.reloptions as options
    from pg_class pc
    join pg_namespace pn on pn.oid = pc.relnamespace
    where reloptions::text ~ 'autovacuum'
  ), step1 as (
    select
      tbl.oid tblid,
      ns.nspname as schema_name,
      tbl.relname as table_name,
      tbl.reltuples,
      tbl.relpages as heappages,
      coalesce(toast.relpages, 0) as toastpages,
      coalesce(toast.reltuples, 0) as toasttuples,
      coalesce(substring(array_to_string(tbl.reloptions, ' ') from '%fillfactor=#"__#"%' for '#')::int2, 100) as fillfactor,
      current_setting('block_size')::numeric as bs,
      case when version() ~ 'mingw32|64-bit|x86_64|ppc64|ia64|amd64' then 8 else 4 end as ma, -- NS: TODO: check it
      24 as page_hdr,
      23 + case when max(coalesce(null_frac, 0)) > 0 then (7 + count(*)) / 8 else 0::int end
        + case when tbl.relhasoids then 4 else 0 end as tpl_hdr_size,
      sum((1 - coalesce(s.null_frac, 0)) * coalesce(s.avg_width, 1024)) as tpl_data_size,
      bool_or(att.atttypid = 'pg_catalog.name'::regtype) or count(att.attname) <> count(s.attname) as is_na
    from pg_attribute as att
    join pg_class as tbl on att.attrelid = tbl.oid and tbl.relkind = 'r'
    join pg_namespace as ns on ns.oid = tbl.relnamespace
    join pg_stats as s on s.schemaname = ns.nspname and s.tablename = tbl.relname and not s.inherited and s.attname = att.attname
    left join pg_class as toast on tbl.reltoastrelid = toast.oid
    where att.attnum > 0 and not att.attisdropped and s.schemaname not in ('pg_catalog', 'information_schema') and tbl.relpages > ${MIN_RELPAGES}
    group by 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, tbl.relhasoids
    order by 2, 3
  ), step2 as (
    select
      *,
      (
        4 + tpl_hdr_size + tpl_data_size + (2 * ma)
        - case when tpl_hdr_size % ma = 0 then ma else tpl_hdr_size % ma end
        - case when ceil(tpl_data_size)::int % ma = 0 then ma else ceil(tpl_data_size)::int % ma end
      ) as tpl_size,
      bs - page_hdr as size_per_block,
      (heappages + toastpages) as tblpages
    from step1
  ), step3 as (
    select
      *,
      ceil(reltuples / ((bs - page_hdr) / tpl_size)) + ceil(toasttuples / 4) as est_tblpages,
      ceil(reltuples / ((bs - page_hdr) * fillfactor / (tpl_size * 100))) + ceil(toasttuples / 4) as est_tblpages_ff
      -- , stattuple.pgstattuple(tblid) as pst
    from step2
  ), step4 as (
    select
      *,
      tblpages * bs as real_size,
      (tblpages - est_tblpages) * bs as extra_size,
      case when tblpages - est_tblpages > 0 then 100 * (tblpages - est_tblpages) / tblpages::float else 0 end as extra_ratio,
      (tblpages - est_tblpages_ff) * bs as bloat_size,
      case when tblpages - est_tblpages_ff > 0 then 100 * (tblpages - est_tblpages_ff) / tblpages::float else 0 end as bloat_ratio
      -- , (pst).free_percent + (pst).dead_tuple_percent as real_frag
    from step3
    left join pg_stat_user_tables su on su.relid = tblid
    -- WHERE NOT is_na
    --   AND tblpages*((pst).free_percent + (pst).dead_tuple_percent)::float4/100 >= 1
  )
  select
    case is_na when true then 'TRUE' else '' end as "is_na",
    coalesce(nullif(step4.schema_name, 'public') || '.', '') || step4.table_name as "table_name",
    case
      when extra_size::numeric >= 0
        then extra_size::numeric
      else null
    end as "extra_size_bytes",
    extra_ratio as "extra_ratio_percent",
    case
      when bloat_size::numeric >= 0
        then bloat_size::numeric
      else null
    end as "bloat_size_bytes",
    bloat_ratio as "bloat_ratio_percent",
    real_size as "real_size_bytes",
    case
      when (real_size - bloat_size)::numeric >=0
        then (real_size - bloat_size)::numeric
        else null
      end as "live_data_size_bytes",
    greatest(last_autovacuum, last_vacuum)::timestamp(0)::text
      || case greatest(last_autovacuum, last_vacuum)
        when last_autovacuum then ' (auto)'
      else '' end as "last_vaccuum",
    (
      select
        coalesce(substring(array_to_string(reloptions, ' ') from 'fillfactor=([0-9]+)')::smallint, 100)
      from pg_class
      where oid = tblid
    ) as "fillfactor",
    case when ot.table_id is not null then true else false end as overrided_settings,
    case
      when bloat_size::numeric >= 0 and (real_size - bloat_size)::numeric >=0
        then real_size::numeric / (real_size - bloat_size)::numeric
        else null
      end as "bloat_ratio_factor"
  from step4
  left join overrided_tables ot on ot.table_id = step4.tblid
  order by bloat_size desc nulls last
), limited_data as (
  select * from data
), num_limited_data as (
  select
    row_number() over () num,
    limited_data.*
  from limited_data
), limited_json_data as (
  select json_object_agg(ld."table_name", ld) as json from num_limited_data ld
), total_data as (
  select
    sum(1) as count,
    sum("extra_size_bytes") as "extra_size_bytes_sum",
    sum("real_size_bytes") as "real_size_bytes_sum",
    sum("bloat_size_bytes") as "bloat_size_bytes_sum",
    sum("live_data_size_bytes") as "live_data_size_bytes_sum",
    (sum("bloat_size_bytes")::numeric/sum("real_size_bytes")::numeric * 100) as "bloat_ratio_percent_avg",
    (sum("real_size_bytes")::numeric/sum("live_data_size_bytes")::numeric) as "bloat_ratio_factor_avg",
    sum("extra_size_bytes") as "extra_size_bytes_sum"
  from data
)
select
  json_build_object(
    'heap_bloat',
    (select * from limited_json_data),
    'heap_bloat_total',
    (select row_to_json(total_data) from total_data),
    'overrided_settings_count',
    (select count(1) from limited_data where overrided_settings = true),
    'database_size_bytes',
    (select pg_database_size(current_database())),
    'min_table_size_bytes',
    (select ${MIN_RELPAGES} * 8192)
  )
SQL

# This SQL is derived from https://github.com/ioguix/pgsql-bloat-estimation/blob/master/table/table_bloat.sql

#Copyright (c) 2015, Jehan-Guillaume (ioguix) de Rorthais
#All rights reserved.
#
#Redistribution and use in source and binary forms, with or without
#modification, are permitted provided that the following conditions are met:
#
#* Redistributions of source code must retain the above copyright notice, this
#  list of conditions and the following disclaimer.
#
#* Redistributions in binary form must reproduce the above copyright notice,
#  this list of conditions and the following disclaimer in the documentation
#  and/or other materials provided with the distribution.
#
#THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
#AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
#IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
#DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
#FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
#DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
#SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
#CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
#OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
#OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.