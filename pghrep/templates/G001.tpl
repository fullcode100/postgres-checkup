# {{ .checkId }} Memory-related settings #

## Observations ##

### Master (`{{.hosts.master}}`) ###

Setting name | Value | Unit
-------------|-------|------
{{ range $i, $key := (index (index (index .results .hosts.master) "data") "_keys") }}
    {{- $value := (index (index (index $.results $.hosts.master) "data") $key) -}}
    [{{ $key }}](https://postgresqlco.nf/en/doc/param/{{ $key }})|{{ $value.setting}}|{{ $value.unit }}
{{ end -}}

{{ if gt (len .hosts.replicas) 0 }}
### Replica servers: ###
    {{ range $skey, $host := .hosts.replicas }}
#### Replica (`{{ $host }}`) ####
        {{ if (index $.results $host) }}
Setting name | Value | Unit
-------------|-------|------
{{ range $i, $key := (index (index (index $.results $host) "data") "_keys") }}
    {{- $value := (index (index (index $.results $host) "data") $key) -}}
    [{{ $key }}](https://postgresqlco.nf/en/doc/param/{{ $key }})|{{ $value.setting}}|{{ $value.unit }}
{{ end }}
        {{- else -}}
No data
        {{- end -}}
    {{- end -}}
{{- end }}

## Conclusions ##


## Recommendations ##

