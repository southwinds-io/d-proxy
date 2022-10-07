module southwinds.dev/d-proxy

go 1.19

replace (
	southwinds.dev/http => ../http
	southwinds.dev/source_client => ../source_client
	southwinds.dev/types => ../types
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.4.0
	southwinds.dev/http v0.0.0-00010101000000-000000000000
	southwinds.dev/source_client v0.0.0-00010101000000-000000000000
	southwinds.dev/types v0.0.0-00010101000000-000000000000
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0 // indirect
	github.com/invopop/jsonschema v0.6.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	github.com/swaggo/http-swagger v1.3.3 // indirect
	github.com/swaggo/swag v1.8.1 // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/tools v0.1.10 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
