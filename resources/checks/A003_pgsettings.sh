# Collect pg_settings artifact
${CHECK_HOST_CMD} "${_PSQL_NO_TIMEOUT} -c \"select json_object_agg(s.name, s) from (select * from pg_settings s order by category, name) s;\""
