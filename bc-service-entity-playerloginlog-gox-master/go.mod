module github.com/PGITAb/bc-service-entity-playerloginlog-go

go 1.13

replace (
	github.com/PGITAb/bc-proto-auth-go => ../bc-proto-lib/bc-proto-auth-go
	github.com/PGITAb/bc-proto-backoffice-go => ../bc-proto-lib/bc-proto-backoffice-go
	github.com/PGITAb/bc-proto-commonstruct-go => ../bc-proto-lib/bc-proto-commonstruct-go
	github.com/PGITAb/bc-proto-diagnostic-go => ../bc-proto-lib/bc-proto-diagnostic-go
	github.com/PGITAb/bc-proto-entity-accesstoken-go => ../bc-proto-lib/bc-proto-entity-accesstoken-go
	github.com/PGITAb/bc-proto-entity-bannerinfo-go => ../bc-proto-lib/bc-proto-entity-bannerinfo-go
	github.com/PGITAb/bc-proto-entity-bet-go => ../bc-proto-lib/bc-proto-entity-bet-go
	github.com/PGITAb/bc-proto-entity-betlimit-go => ../bc-proto-lib/bc-proto-entity-betlimit-go
	github.com/PGITAb/bc-proto-entity-card-go => ../bc-proto-lib/bc-proto-entity-card-go
	github.com/PGITAb/bc-proto-entity-common-go => ../bc-proto-lib/bc-proto-entity-common-go
	github.com/PGITAb/bc-proto-entity-fundtran-go => ../bc-proto-lib/bc-proto-entity-fundtran-go
	github.com/PGITAb/bc-proto-entity-gameround-go => ../bc-proto-lib/bc-proto-entity-gameround-go
	github.com/PGITAb/bc-proto-entity-gametable-go => ../bc-proto-lib/bc-proto-entity-gametable-go
	github.com/PGITAb/bc-proto-entity-host-go => ../bc-proto-lib/bc-proto-entity-host-go
	github.com/PGITAb/bc-proto-entity-massagepoint-go => ../bc-proto-lib/bc-proto-entity-massagepoint-go
	github.com/PGITAb/bc-proto-entity-operatorprofile-go => ../bc-proto-lib/bc-proto-entity-operatorprofile-go
	github.com/PGITAb/bc-proto-entity-playerloginlog-go => ../bc-proto-lib/bc-proto-entity-playerloginlog-go
	github.com/PGITAb/bc-proto-entity-playerprofile-go => ../bc-proto-lib/bc-proto-entity-playerprofile-go
	github.com/PGITAb/bc-proto-entity-playerwallet-go => ../bc-proto-lib/bc-proto-entity-playerwallet-go
	github.com/PGITAb/bc-proto-entity-reader-go => ../bc-proto-lib/bc-proto-entity-reader-go
	github.com/PGITAb/bc-proto-entity-sitenewsinfo-go => ../bc-proto-lib/bc-proto-entity-sitenewsinfo-go
	github.com/PGITAb/bc-proto-entity-staffloginlog-go => ../bc-proto-lib/bc-proto-entity-staffloginlog-go
	github.com/PGITAb/bc-proto-entity-staffprofile-go => ../bc-proto-lib/bc-proto-entity-staffprofile-go
	github.com/PGITAb/bc-proto-entity-tip-go => ../bc-proto-lib/bc-proto-entity-tip-go
	github.com/PGITAb/bc-proto-entity-uisettingsinfo-go => ../bc-proto-lib/bc-proto-entity-uisettingsinfo-go
	github.com/PGITAb/bc-proto-gamehistory-go => ../bc-proto-lib/bc-proto-gamehistory-go
	github.com/PGITAb/bc-proto-gamelobby-go => ../bc-proto-lib/bc-proto-gamelobby-go
	github.com/PGITAb/bc-proto-gomw-go => ../bc-proto-lib/bc-proto-gomw-go
	github.com/PGITAb/bc-proto-gosw-go => ../bc-proto-lib/bc-proto-gosw-go
	github.com/PGITAb/bc-proto-livegamehost-ba-go => ../bc-proto-lib/bc-proto-livegamehost-ba-go
	github.com/PGITAb/bc-proto-livegamehost-go => ../bc-proto-lib/bc-proto-livegamehost-go
	github.com/PGITAb/bc-proto-livegamehost-pitboss-go => ../bc-proto-lib/bc-proto-livegamehost-pitboss-go
	github.com/PGITAb/bc-proto-livegametable-ba-go => ../bc-proto-lib/bc-proto-livegametable-ba-go
	github.com/PGITAb/bc-proto-livegametable-go => ../bc-proto-lib/bc-proto-livegametable-go
	github.com/PGITAb/bc-proto-wallet-go => ../bc-proto-lib/bc-proto-wallet-go
	github.com/PGITAb/bc-proto-walletstruct-go => ../bc-proto-lib/bc-proto-walletstruct-go

)

require (
	github.com/PGITAb/bc-proto-entity-common-go v0.0.0-00010101000000-000000000000
	github.com/PGITAb/bc-proto-entity-playerloginlog-go v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.7
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.9.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.4.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.1
	xorm.io/core v0.7.0
)
