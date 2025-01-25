SELECT
    code.sys_code_text as menu_type_name,
    p.menu_name as parent_menu_name,
    a.*
FROM sys_menu a
left join sys_code code on code.sys_code_type = 'menu_type'  and a.menu_type = code.sys_code
left join sys_menu p on a.parent_id = p.sys_menu_id
where a.belong_project= {{.belong_project}}
{{ if .menu_code }}
and a.menu_code = {{.menu_code }}
{{ end }}
{{ if .sys_menu_id}}
and a.sys_menu_id = {{.sys_menu_id }}
{{ end }}



{{ if .menu_type }}
and a.menu_type = {{.menu_type }}
{{ end }}

{{ if .exclude }}
and a.sys_menu_id != {{.exclude}}
{{ end }}

{{ if .search }}
and ( a.menu_name like {{.search}} or a.url like {{.search}} or a.menu_code like {{.search}})
{{ end }}
order by a.order_index
