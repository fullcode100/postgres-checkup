# {{ .checkId }} Globally aggregated query metrics

## Observations ##
{{ if .hosts.master }}
### Master (`{{.hosts.master}}`) ###
Start: {{ (index (index (index .results .hosts.master) "data") "start_timestamptz") }}  
End: {{ (index (index (index .results .hosts.master) "data") "end_timestamptz") }}  
Period seconds: {{ (index (index (index .results .hosts.master) "data") "period_seconds") }}  
Period age: {{ (index (index (index .results .hosts.master) "data") "period_age") }}  

Error (calls): {{ NumFormat (index (index (index .results .hosts.master) "data") "absolute_error_calls") 2 }} ({{ NumFormat (index (index (index .results .hosts.master) "data") "relative_error_calls") 2 }}%)  
Error (total time): {{ NumFormat (index (index (index .results .hosts.master) "data") "absolute_error_total_time") 2 }} ({{ NumFormat (index (index (index .results .hosts.master) "data") "relative_error_total_time") 2 }}%)

Calls | Total&nbsp;time | Rows | shared_blks_hit | shared_blks_read | shared_blks_dirtied | shared_blks_written | blk_read_time | blk_write_time | kcache_reads | kcache_writes | kcache_user_time_ms | kcache_system_time 
-------|------------|------|-----------------|------------------|---------------------|---------------------|---------------|----------------|--------------|---------------|---------------------|--------------------
{{ range $i, $key := (index (index (index (index .results .hosts.master) "data") "aggregated") "_keys") }}
{{- $value := (index (index (index (index $.results $.hosts.master) "data") "aggregated") $key) -}}
{{- NumFormat $value.diff_calls 2 }}<br/>{{ NumFormat $value.per_sec_calls 2 }}/sec<br/>{{ NumFormat $value.per_call_calls 2 }}/call<br/>{{ NumFormat $value.ratio_calls 2 }}% |
{{- MsFormat $value.diff_total_time }}<br/>{{ MsFormat $value.per_sec_total_time }}/sec<br/>{{ MsFormat $value.per_call_total_time }}/call<br/>{{ NumFormat $value.ratio_total_time 2 }}% |
{{- NumFormat $value.diff_rows 2 }}<br/>{{ NumFormat $value.per_sec_rows 2 }}/sec<br/>{{ NumFormat $value.per_call_rows 2 }}/call<br/>{{ NumFormat $value.ratio_rows 2 }}% |
{{- NumFormat $value.diff_shared_blks_hit 2 }}&nbsp;blks<br/>{{ NumFormat $value.per_sec_shared_blks_hit 2 }}&nbsp;blks/sec<br/>{{ NumFormat $value.per_call_shared_blks_hit 2 }}&nbsp;blks/call<br/>{{ NumFormat $value.ratio_shared_blks_hit 2 }}% |
{{- NumFormat $value.diff_shared_blks_read 2 }}&nbsp;blks<br/>{{ NumFormat $value.per_sec_shared_blks_read 2 }}&nbsp;blks/sec<br/>{{ NumFormat $value.per_call_shared_blks_read 2 }}&nbsp;blks/call<br/>{{ NumFormat $value.ratio_shared_blks_read 2 }}% |
{{- NumFormat $value.diff_shared_blks_dirtied 2 }}&nbsp;blks<br/>{{ NumFormat $value.per_sec_shared_blks_dirtied 2 }}&nbsp;blks/sec<br/>{{ NumFormat $value.per_call_shared_blks_dirtied 2 }}&nbsp;blks/call<br/>{{ NumFormat $value.ratio_shared_blks_dirtied 2 }}% |
{{- NumFormat $value.diff_shared_blks_written 2 }}&nbsp;blks<br/>{{ NumFormat $value.per_sec_shared_blks_written 2 }}&nbsp;blks/sec<br/>{{ NumFormat $value.per_call_shared_blks_written 2 }}&nbsp;blks/call<br/>{{ NumFormat $value.ratio_shared_blks_written 2 }}% |
{{- MsFormat $value.diff_blk_read_time }}<br/>{{ MsFormat $value.per_sec_blk_read_time }}/sec<br/>{{ MsFormat $value.per_call_blk_read_time }}/call<br/>{{ NumFormat $value.ratio_blk_read_time 2 }}% |
{{- MsFormat $value.diff_blk_write_time }}<br/>{{ MsFormat $value.per_sec_blk_write_time }}/sec<br/>{{ MsFormat $value.per_call_blk_write_time }}/call<br/>{{ MsFormat $value.per_call_blk_write_time }}/call<br/>{{ NumFormat $value.ratio_blk_write_time 2 }}% |
{{- NumFormat $value.diff_kcache_reads 2 }}&nbsp;bytes<br/>{{ NumFormat $value.per_sec_kcache_reads 2 }}&nbsp;bytes/sec<br/>{{ NumFormat $value.per_call_kcache_reads 2 }}&nbsp;bytes/call<br/>{{ NumFormat $value.ratio_kcache_reads 2 }}% |
{{- NumFormat $value.diff_kcache_writes 2 }}&nbsp;bytes<br/>{{ NumFormat $value.per_sec_kcache_writes 2 }}&nbsp;bytes/sec<br/>{{ NumFormat $value.per_call_kcache_writes 2 }}&nbsp;bytes/call<br/>{{ NumFormat $value.ratio_kcache_writes 2 }}% |
{{- MsFormat $value.diff_kcache_user_time_ms }}<br/>{{ MsFormat $value.per_sec_kcache_user_time_ms }}/sec<br/>{{ MsFormat $value.per_call_kcache_user_time_ms }}/call<br/>{{ NumFormat $value.ratio_kcache_user_time_ms 2 }}% |
{{- MsFormat $value.diff_kcache_system_time_ms }}<br/>{{ MsFormat $value.per_sec_kcache_system_time_ms }}/sec<br/>{{ MsFormat $value.per_call_kcache_system_time_ms }}/call<br/>{{ NumFormat $value.ratio_kcache_system_time_ms 2 }}%
{{ end }}{{/* range */}}
{{ else }}{{/* if .host.master*/}}
No data
{{ end }}{{/* if .host.master*/}}

{{ if gt (len .hosts.replicas) 0 }}
### Replica servers: ###
{{ range $key, $host := .hosts.replicas }}
#### Replica (`{{ $host }}`) ####
{{ if (index $.results $host) }}
Start: {{ (index (index (index $.results $host) "data") "start_timestamptz") }}  
End: {{ (index (index (index $.results $host) "data") "end_timestamptz") }}  
Period seconds: {{ (index (index (index $.results $host) "data") "period_seconds") }}  
Period age: {{ (index (index (index $.results $host) "data") "period_age") }}  

Calls | Total&nbsp;time | Rows | shared_blks_hit | shared_blks_read | shared_blks_dirtied | shared_blks_written | blk_read_time | blk_write_time | kcache_reads | kcache_writes | kcache_user_time_ms | kcache_system_time 
-------|------------|------|-----------------|------------------|---------------------|---------------------|---------------|----------------|--------------|---------------|---------------------|--------------------
{{ range $i, $key := (index (index (index (index $.results $host) "data") "aggregated") "_keys") }}
{{- $value := (index (index (index (index $.results $host) "data") "aggregated") $key) -}}
{{- NumFormat $value.diff_calls 2 }}<br/>{{ NumFormat $value.per_sec_calls 2 }}/sec<br/>{{ NumFormat $value.per_call_calls 2 }}/call<br/>{{ NumFormat $value.ratio_calls 2 }}% |
{{- MsFormat $value.diff_total_time }}<br/>{{ MsFormat $value.per_sec_total_time }}/sec<br/>{{ MsFormat $value.per_call_total_time }}/call<br/>{{ NumFormat $value.ratio_total_time 2 }}% |
{{- NumFormat $value.diff_rows 2 }}<br/>{{ NumFormat $value.per_sec_rows 2 }}/sec<br/>{{ NumFormat $value.per_call_rows 2 }}/call<br/>{{ NumFormat $value.ratio_rows 2 }}% |
{{- NumFormat $value.diff_shared_blks_hit 2 }}&nbsp;blks<br/>{{ NumFormat $value.per_sec_shared_blks_hit 2 }}&nbsp;blks/sec<br/>{{ NumFormat $value.per_call_shared_blks_hit 2 }}&nbsp;blks/call<br/>{{ NumFormat $value.ratio_shared_blks_hit 2 }}% |
{{- NumFormat $value.diff_shared_blks_read 2 }}&nbsp;blks<br/>{{ NumFormat $value.per_sec_shared_blks_read 2 }}&nbsp;blks/sec<br/>{{ NumFormat $value.per_call_shared_blks_read 2 }}&nbsp;blks/call<br/>{{ NumFormat $value.ratio_shared_blks_read 2 }}% |
{{- NumFormat $value.diff_shared_blks_dirtied 2 }}&nbsp;blks<br/>{{ NumFormat $value.per_sec_shared_blks_dirtied 2 }}&nbsp;blks/sec<br/>{{ NumFormat $value.per_call_shared_blks_dirtied 2 }}&nbsp;blks/call<br/>{{ NumFormat $value.ratio_shared_blks_dirtied 2 }}% |
{{- NumFormat $value.diff_shared_blks_written 2 }}&nbsp;blks<br/>{{ NumFormat $value.per_sec_shared_blks_written 2 }}&nbsp;blks/sec<br/>{{ NumFormat $value.per_call_shared_blks_written 2 }}&nbsp;blks/call<br/>{{ NumFormat $value.ratio_shared_blks_written 2 }}% |
{{- MsFormat $value.diff_blk_read_time }}<br/>{{ MsFormat $value.per_sec_blk_read_time }}/sec<br/>{{ MsFormat $value.per_call_blk_read_time }}/call<br/>{{ NumFormat $value.ratio_blk_read_time 2 }}% |
{{- MsFormat $value.diff_blk_write_time }}<br/>{{ MsFormat $value.per_sec_blk_write_time }}/sec<br/>{{ MsFormat $value.per_call_blk_write_time }}/call<br/>{{ MsFormat $value.per_call_blk_write_time }}/call<br/>{{ NumFormat $value.ratio_blk_write_time 2 }}% |
{{- NumFormat $value.diff_kcache_reads 2 }}&nbsp;bytes<br/>{{ NumFormat $value.per_sec_kcache_reads 2 }}&nbsp;bytes/sec<br/>{{ NumFormat $value.per_call_kcache_reads 2 }}&nbsp;bytes/call<br/>{{ NumFormat $value.ratio_kcache_reads 2 }}% |
{{- NumFormat $value.diff_kcache_writes 2 }}&nbsp;bytes<br/>{{ NumFormat $value.per_sec_kcache_writes 2 }}&nbsp;bytes/sec<br/>{{ NumFormat $value.per_call_kcache_writes 2 }}&nbsp;bytes/call<br/>{{ NumFormat $value.ratio_kcache_writes 2 }}% |
{{- MsFormat $value.diff_kcache_user_time_ms }}<br/>{{ MsFormat $value.per_sec_kcache_user_time_ms }}/sec<br/>{{ MsFormat $value.per_call_kcache_user_time_ms }}/call<br/>{{ NumFormat $value.ratio_kcache_user_time_ms 2 }}% |
{{- MsFormat $value.diff_kcache_system_time_ms }}<br/>{{ MsFormat $value.per_sec_kcache_system_time_ms }}/sec<br/>{{ MsFormat $value.per_call_kcache_system_time_ms }}/call<br/>{{ NumFormat $value.ratio_kcache_system_time_ms 2 }}%
{{ end }}{{/* range */}}
{{- else -}}{{/* if host data */}}
No data
{{- end -}}{{/* if host data */}}
{{- end -}}{{/* hosts range */}}
{{- end -}}{{/* if replicas */}}

## Conclusions ##


## Recommendations ##

