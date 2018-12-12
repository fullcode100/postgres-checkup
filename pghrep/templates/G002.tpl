# {{ .checkId }} Connections and current activity #

## Observations ##

### Master (`{{.hosts.master}}`) ###
```
{{(index (index .results .hosts.master) "data").raw}}
```
{{/* newline */}}
{{/* newline */}}

{{- if gt (len .hosts.replicas) 0 -}}
### Replica servers: ###
{{/* newline */}}
{{/* newline */}}
  {{- range $skey, $host := .hosts.replicas -}}
#### Replica (`{{ $host }}`) ####
    {{- if (index $.results $host) -}}
{{/* newline */}}
{{/* newline */}}
```
{{ (index (index $.results $host) "data").raw }}
```
    {{- else -}}
```
No data
```
    {{- end -}}
  {{- end -}}
{{- end -}}

{{/* newline */}}
{{/* newline */}}
## Conclusions ##


## Recommendations ##

