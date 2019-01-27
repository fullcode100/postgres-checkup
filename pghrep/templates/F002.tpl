# {{ .checkId }} Autovacuum: Transaction wraparound check #

## Observations ##
{{ if .hosts.master }}
### Master (`{{.hosts.master}}`) ###
{{ if index (index (index .results .hosts.master) "data") "per_instance" }}
#### Databases ####
 Database | &#9660;&nbsp;Age | Capacity used, % | Warning | datfrozenxid
----------|-----|------------------|---------|--------------
{{ range $i, $key := (index (index (index (index .results .hosts.master) "data") "per_instance") "_keys") }}
{{- $value := (index (index (index (index $.results $.hosts.master) "data") "per_instance") $key) -}}
{{ index $value "datname"}} | 
{{- NumFormat (index $value "age") -1 }} |
{{- index $value "capacity_used"}} |
{{- if (index $value "warning") }} &#9888; {{ else }} {{ end }} |
{{- NumFormat (index $value "datfrozenxid") -1}}
{{ end }}{{/* range */}}
{{- end -}}{{/* if per_instance exists */}}

{{/* if index (index (index .results .hosts.master) "data") "per_database" */}}
#### Tables in the observed database ####
 Relation | Age | &#9660;&nbsp;Capacity used, % | Warning |rel_relfrozenxid | toast_relfrozenxid 
----------|-----|------------------|---------|-----------------|--------------------
{{ range $i, $key := (index (index (index (index .results .hosts.master) "data") "per_database") "_keys") }}
{{- $value := (index (index (index (index $.results $.hosts.master) "data") "per_database") $key) -}}
{{ index $value "relation"}} | 
{{- NumFormat (index $value "age") -1 }} |
{{- index $value "capacity_used"}} |
{{- if (index $value "warning") }} &#9888; {{ else }} {{ end }} |
{{- NumFormat (index $value "rel_relfrozenxid") -1}} |
{{- NumFormat (index $value "toast_relfrozenxid") -1}} |
{{ end }}{{/* range */}}
{{/*- end -*/}}{{/* if per_instance exists */}}

{{- else }}
No data
{{- end }}

## Conclusions ##


## Recommendations ##

