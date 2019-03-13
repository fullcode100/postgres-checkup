# {{ .checkId }} Table sizes #

## Observations ##
Data collected: {{ DtFormat .timestamptz }}  
Current database: {{ .database }}  
{{ if .hosts.master }}
{{ if (index .results .hosts.master) }}
{{ if (index (index .results .hosts.master) "data") }}
### Master (`{{.hosts.master}}`) ###
Table | Rows | &#9660;&nbsp;Total size | Table size | Index(es) Size | TOAST Size
------|------|------------|------------|----------------|------------
{{ range $i, $key := (index (index (index .results .hosts.master) "data") "_keys") }}
{{- $value := (index (index (index $.results $.hosts.master) "data") $key) -}}
{{ index $value "Table"}} | {{ index $value "Rows"}} | {{ index $value "Total Size"}} | {{ index $value "Table Size"}} | {{ index $value "Index(es) Size"}} | {{ index $value "TOAST Size"}}
{{ end }}
{{- else -}}{{/*Master data*/}}
No data
{{- end }}{{/*Master data*/}}
{{- else -}}{{/*Master results*/}}
No data
{{- end }}{{/*Master results*/}}
{{- else -}}{{/*Master*/}}
No data
{{ end }}{{/*Master*/}}

## Conclusions ##


## Recommendations ##

