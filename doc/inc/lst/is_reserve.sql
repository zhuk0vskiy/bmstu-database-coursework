create or replace function is_reserve(userId integer,
                                      roomId integer,
                                      producerId integer,
                                      instrumentalistId integer,
                                      startTime timestamp,
                                      endTime timestamp) returns integer language plpgsql as $$
    declare
        count integer;
    begin
        select count(*) into count
        from reserve
        where (user_id = userId or
               room_id = roomId or
               producer_id = producerId or
               instrumentalist_id = instrumentalistId)
          and is_intersect(reserve.start_time,
                           reserve.end_time,
                           startTime,
                           endTime);

        return count;
    end;
    $$;
