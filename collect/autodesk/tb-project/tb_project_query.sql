select a.*
from tb_project a
where 1=1
{{ if .name }}
 and a.name = {{.name}}
{{ end }}
{{ if .exclude}}
and a.tb_project_id != {{ .exclude }}
{{ end }}
{{ if .search }}
and  a.name like {{.search}}
{{ end}}
