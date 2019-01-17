# {{ .checkId }} Unused/Rarely Used Indexes #

## Observations ##

{{ if .resultData }}

{{- if .resultData.unused_indexes -}}
### Never Used Indexes ###
Index | {{.hosts.master}} {{ range $skey, $host := .hosts.replicas }}| {{ $host }} {{ end }}| Usage
--------|-------{{ range $skey, $host := .hosts.replicas }}|--------{{ end }}|-----
{{ range $key, $value := (index .resultData "unused_indexes") }}
{{- if ne $key "_keys" -}}
{{- if eq $value.master.reason "Never Used Indexes" -}}
{{- if $value.usage -}}
{{- else -}}
{{ $key }} | {{ $value.master.idx_scan }}{{ range $skey, $host := $.hosts.replicas }}|{{ (index $value $host).idx_scan }}{{- end -}} | {{ if $value.usage }} Used{{ else }}Not used {{ end }}
{{/* new line */}}
{{- end -}}{{/* value.usage */}}
{{- end -}}{{/**/}}
{{- end }}{{/* in ! _keys */}}
{{- end }}{{/* range unused_indexes */}}

### Other unused indexes ###
Index | Reason |{{.hosts.master}} {{ range $skey, $host := .hosts.replicas }}| {{ $host }} {{ end }}| Usage
------|--------|-------{{ range $skey, $host := .hosts.replicas }}|--------{{ end }}|-----
{{ range $key, $value := (index .resultData "unused_indexes") }}
{{- if ne $key "_keys" -}}
{{- if ne $value.master.reason "Never Used Indexes" -}}
{{ $key }} | {{ $value.master.reason }} | Index&nbsp;size:{{ Nobr $value.master.index_size }} Table&nbsp;size:{{ Nobr $value.master.table_size }} {{ range $skey, $host := $.hosts.replicas }} |Index&nbsp;size:{{ Nobr (index $value $host).index_size }} Table&nbsp;size:{{ Nobr (index $value $host).table_size }}{{- end -}} | {{ if $value.usage }} Used{{ else }}Not used {{ end }}
{{/* new line */}}
{{- end -}}
{{- end }}{{/* ! "_keys" */}}
{{- end }}{{/* range unused_indexes */}}
{{ end }}{{/* if unused_indexes */}}

{{- if .resultData.redundant_indexes -}}

### Redundant indexes ###

Index | {{.hosts.master}} {{ range $skey, $host := .hosts.replicas }}| {{ $host }} {{ end }}| Usage | Index size
--------|-------{{ range $skey, $host := .hosts.replicas }}|--------{{ end }}|-----|-----
{{ range $key, $value := (index .resultData "redundant_indexes") }}
{{- if ne $key "_keys" -}}
{{- if $value.usage -}}
{{- else -}}
{{ $key }} | {{ $value.master.index_usage }}{{ range $skey, $host := $.hosts.replicas }}|{{ (index $value $host).index_usage }}{{- end -}} | {{ if $value.usage }} Used{{ else }}Not used {{ end }} | {{ $value.master.index_size }}
{{/* new line */}}
{{- end -}}{{/* value.usage */}}
{{- end }}{{/* in ! _keys */}}
{{- end }}{{/* range redundant_indexes */}}
{{end}}{{/* if redundant_indexes */}}

## Conclusions ##


## Recommendations ##
{{ if .resultData.dropCode }}
#### Drop code ####
```
{{ range $i, $drop_code := (index .resultData  "dropCode") }}{{ $drop_code }}
{{ end }}
```
{{ end }}

{{ if .resultData.revertCode }}
#### Revert code ####
```
{{ range $i, $revert_code := (index .resultData  "revertCode") }}{{ $revert_code }}
{{ end }}
```
{{ end }}


{{- else -}}
No data
{{- end }}
