create or replace function is_intersect(reserveStartTime timestamp, 
					reserveEndTime timestamp,
					choosenStartTime timestamp,
					choosenEndTime timestamp) returns boolean language plpgsql as $$
    declare
        ans boolean;
    begin
    ans = false;
    if ((choosenStartTime >= reserveStartTime and choosenStartTime < reserveEndTime) or
		(choosenEndTime <= reserveEndTime and choosenEndTime > reserveStartTime) or
		(choosenStartTime <= reserveStartTime and choosenEndTime >= reserveEndTime)) then
		ans = true;
		end if;
    return ans;
	end;
    $$;
