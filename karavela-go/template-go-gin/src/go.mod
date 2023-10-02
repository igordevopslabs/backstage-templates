module ${{ values.apiName | lower }}

go ${{ values.goVersion | lower }}

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/sirupsen/logrus v1.9.3
)
