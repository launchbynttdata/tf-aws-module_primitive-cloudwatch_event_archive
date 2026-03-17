resource_names_map = {
  "event_bus" = {
    name       = "eventbus1"
    max_length = 64
  }
  "event_archive" = {
    name       = "eventarchive1"
    max_length = 48
  }
}

logical_product_family  = "launch"
logical_product_service = "eventbridge"
class_env               = "dev"
instance_env            = 1
instance_resource       = 1

description    = "Example EventBridge archive for testing"
retention_days = 7
