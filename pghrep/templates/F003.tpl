# {{ .checkId }} Autovacuum: Dead tuples #

## Observations ##
{{ if .hosts.master }}
### Master (`{{.hosts.master}}`) ###
Stats reset: {{ DtFormat (index (index (index .results .hosts.master) "data") "database_stat").stats_reset }}  
Report created: {{ DtFormat .timestamptz }}

 Relation | Type | Since last autovacuum | Since last vacuum | Autovacuum Count | Vacuum Count | n_tup_ins | n_tup_upd | n_tup_del | pg_class.reltuples | n_live_tup | n_dead_tup | &#9660;Dead Tuples Ratio, %
----------|------|-----------------------|-------------------|----------|---------|-----------|-----------|-----------|--------------------|------------|------------|-----------
{{ range $i, $key := (index (index (index (index .results .hosts.master) "data") "dead_tuples") "_keys") }}
{{- $value := (index (index (index (index $.results $.hosts.master) "data") "dead_tuples") $key) -}}
{{ index $value "relation"}} | 
{{- index $value "relkind"}} | 
{{- index $value "since_last_autovacuum"}} |
{{- index $value "since_last_vacuum"}} |
{{- NumFormat (index $value "av_count") -1 }} |
{{- NumFormat (index $value "v_count") -1 }} |
{{- NumFormat (index $value "n_tup_ins") -1 }} |
{{- NumFormat (index $value "n_tup_upd") -1 }} |
{{- NumFormat (index $value "n_tup_del") -1 }} |
{{- NumFormat (index $value "pg_class_reltuples") -1 }} |
{{- NumFormat (index $value "n_live_tup") -1 }} |
{{- NumFormat (index $value "n_dead_tup") -1 }} |
{{- index $value "dead_ratio"}}
{{ end }}
{{- else }}
No data
{{- end }}

## Conclusions ##


## Recommendations ##

