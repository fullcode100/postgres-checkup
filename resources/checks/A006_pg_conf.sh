# Collect pg_settings and pg_config values data to compare with same values from other domains
# Require comparation in plugin
pg_settings=$(ssh ${HOST} "${_PSQL} -c \"select json_object_agg(s.name, s) from pg_settings s;\"")
pg_config=$(ssh ${HOST} "${_PSQL} -c \"select json_object_agg(c.name, c) from pg_config c;\"")
echo '{ "pg_settings": '$pg_settings', "pg_config":'$pg_config'}'

