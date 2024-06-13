select equipment.id,
               equipment.name,
               equipment.type,
               equipment.studio_id
        from equipment where equipment.studio_id = $1
