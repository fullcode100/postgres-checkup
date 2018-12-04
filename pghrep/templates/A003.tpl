Current values
===

Postgres settings

Master DB server is {{.hosts.master}}
{{ range $key, $value := (index (index .results .hosts.master) "data") }}
  {{ $key }} = {{ $value.setting}} {{ if $value.unit }}{{ $value.unit }} {{ end }}
{{ end }}

{{ if gt (len .hosts.replicas) 0 }}
Slave DB servers:
  {{ range $skey, $host := .hosts.replicas }}
  DB slave server: {{ $host }}
    {{ if (index $.results $host) }}
      {{ range $key, $value := (index (index $.results $host) "data") }}
        {{ $key }} = {{ $value.setting}} {{ if $value.unit }}{{ $value.unit }} {{ end }}
      {{ end }}
    {{ else }}
      No data
    {{ end}}
  {{ end }}
{{ end }}

Conclusions
===
{{.Conclusion}}

Recommendations
===
{{.Recommended}}
