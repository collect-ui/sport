select a.*,
(
   select GROUP_CONCAT(name )
   from ldap_group lg
   join role_ldap_group rlg on rlg.ldap_group_id =lg.ldap_group_id
   where rlg.role_id =a.role_id
) as ldap_names,
(
    select GROUP_CONCAT(lg.ldap_group_id)
    from ldap_group lg
    join role_ldap_group rlg   on rlg.ldap_group_id =lg.ldap_group_id
    where rlg.role_id =a.role_id
) as ldap_group_ids,
(
    select count(1)
    from user_role_id_list b
    join user_account c on c.user_id = b.user_id and c.is_delete= '0'
    where b.role_id = a.role_id
) as user_count
from (require(base.sql)) as a
order by a.order_index