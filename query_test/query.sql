SELECT u.id AS id, u.name AS name, p.name AS parent_name
FROM mst_users AS u
LEFT JOIN mst_users AS p
ON u.parent_id = p.id;