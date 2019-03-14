# {{ .checkId }} Memory-related settings #

## Observations ##
Data collected: {{ DtFormat .timestamptz }}  
{{ if .hosts.master }}
{{ if (index .results .hosts.master) }}
{{ if (index (index .results .hosts.master) "data") }}
### Master (`{{.hosts.master}}`) ###

Setting name | Value | Unit | Pretty value
-------------|-------|------|--------------
{{ range $i, $key := (index (index (index .results .hosts.master) "data") "_keys") }}
    {{- $value := (index (index (index $.results $.hosts.master) "data") $key) -}}
    [{{ $key }}](https://postgresqlco.nf/en/doc/param/{{ $key }}) | {{ $value.setting }}| {{ $value.unit }} | {{ UnitValue $value.setting $value.unit}}
{{ end -}}
{{ end }}{{/* master data */}}
{{ end }}{{/* master results */}}
{{ end }}{{/* master */}}

{{ if gt (len .hosts.replicas) 0 }}
### Replica servers: ###
    {{ range $skey, $host := .hosts.replicas }}
#### Replica (`{{ $host }}`) ####
        {{ if (index $.results $host) }}
Setting name | Value
-------------|-------
{{ range $i, $key := (index (index (index $.results $host) "data") "_keys") }}
    {{- $value := (index (index (index $.results $host) "data") $key) -}}
    [{{ $key }}](https://postgresqlco.nf/en/doc/param/{{ $key }})|{{ UnitValue $value.setting $value.unit}}
{{ end }}
        {{- else -}}
No data
        {{- end -}}
    {{- end -}}
{{- end }}

## Conclusions ##


## Recommendations ##

