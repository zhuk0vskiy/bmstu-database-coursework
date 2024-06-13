select equipment.id,
               equipment.name,
               equipment.type,
               equipment.studio_id
        from equipment where equipment.studio_id = $1 and equipment.type = $2 and not exists
            (select * from reserved_equipments where equipment.id = reserved_equipments.equipment_id)
