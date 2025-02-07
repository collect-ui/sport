select a.*
from tb_config a
where 1=1
{{ if .name }}
 and a.name = {{.name}}
{{ end }}
{{ if .exclude}}
and a.tb_config_id != {{ .exclude }}
{{ end }}
{{ if .search }}
and  (a.name like {{.search}} or a.value like {{.search}} or a.description like {{.search}})
{{ end}}
{{ if .project }}
and a.project = {{.project}}
{{ end}}
order by a.order_index