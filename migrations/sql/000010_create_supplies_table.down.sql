DROP INDEX IF EXISTS idx_supply_usages_batch_id;
DROP INDEX IF EXISTS idx_supply_usages_student_id;
DROP INDEX IF EXISTS idx_supply_usages_class_session_id;
DROP INDEX IF EXISTS idx_supply_usages_used_at;
DROP INDEX IF EXISTS idx_supply_batches_supply_id;
DROP INDEX IF EXISTS idx_supply_batches_purchase_date;

DROP TABLE IF EXISTS supply_usages;
DROP TABLE IF EXISTS supply_batches;
DROP TABLE IF EXISTS supplies;
