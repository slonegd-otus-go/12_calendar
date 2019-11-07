SELECT id, description, start_time, duration 
FROM events
WHERE '2006-01-02 15:04:17' BETWEEN start_time::timestamp AND start_time::timestamp + duration::interval
