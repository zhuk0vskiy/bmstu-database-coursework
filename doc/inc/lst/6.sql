select equipment.id, 
               equipment.name,
               equipment.type,
               equipment.studio_id,
             to_char(reserve.start_time, 'YYYY-MM-DD HH24:MI:SS'),
             to_char(reserve.end_time, 'YYYY-MM-DD HH24:MI:SS')
        from equipment, reserve where equipment.studio_id = $1 and equipment.type = $2 and exists 
            (select * from reserved_equipments where equipment.id = reserved_equipments.equipment_id)
