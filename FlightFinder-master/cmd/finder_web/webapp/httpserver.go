package webapp

import (
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	apiserver "github.com/mateuszmidor/FlightFinder/cmd/finder_web/webapp/apiserver/go"
	"github.com/mateuszmidor/FlightFinder/pkg/application"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/aws"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv"
	"github.com/oroshnivskyy/go-gin-aws-x-ray/xray"
)

func Run(http_port, flights_data_dir, web_data_dir string, metrics application.MetricsClient, tracing application.TracingClient, cache infrastructure.CacheClient) {
	router := gin.Default()
	router.Use(tracing)
	router.LoadHTMLGlob(path.Join(web_data_dir, "*.html"))
	router.StaticFile("favicon.ico", path.Join(web_data_dir, "favicon.ico"))
	router.Use(allowLocalSwaggerPreviewCORS)
	router.Use(finder(flights_data_dir))
	router.Use(airports(flights_data_dir))
	for _, r := range apiserver.GetRoutes() {
		// raw handler
		handler := r.HandlerFunc

		// add caching
		handler = handlerWithCache(cache, r.Method, handler)

		// add metrics
		handler = handlerWithMetrics(metrics, r.Method, r.Pattern, handler)

		// register
		router.Handle(r.Method, r.Pattern, handler)
	}
	log.Fatal(router.Run(":" + http_port))
}

func MakeMetricsClient(aws_region string) application.MetricsClient {
	log.Print("Initializing AWS CloudWatch as metrics client...")
	var client application.MetricsClient
	client, err := aws.NewMetricsClient(aws_region)
	if err != nil {
		log.Printf("AWS CloudWatch not available - will not push metrics. Error: %v", err)
		return &application.NullMericsClient{}
	}
	log.Print("Done.")
	return client
}

func MakeTracingClient(enabled bool) application.TracingClient {
	if enabled {
		return xray.Middleware(nil)
	} else {
		return application.NullTracingClient
	}
}

func MakeCacheClient(addr, pass string) infrastructure.CacheClient {
	log.Print("Initializing Redis as cache client...")
	var client infrastructure.CacheClient
	client, err := aws.NewCacheClient(addr, pass, time.Duration(0))
	if err != nil {
		log.Printf("Redis not available - will not cache search results. Error: %v", err)
		return &infrastructure.NullCacheClient{}
	}
	log.Print("Done.")
	return client
}

func handlerWithMetrics(metrics application.MetricsClient, method string, pattern string, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		handler(c)
		duration := time.Since(start)
		metrics.PutRequestMetrics(pattern, method, duration)
	}
}

func handlerWithTracing(tracing application.TracingClient, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
		tracing(c)
	}
}

func handlerWithCache(cache infrastructure.CacheClient, method string, handler gin.HandlerFunc) gin.HandlerFunc {
	if method == http.MethodGet {
		return cache.CachePage(handler)
	}
	return handler
}

func allowLocalSwaggerPreviewCORS(c *gin.Context) {
	const SwaggerViewerLocalhost = "http://localhost:18512"
	c.Header("Access-Control-Allow-Origin", SwaggerViewerLocalhost)
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
}

func finder(csv_dir string) func(*gin.Context) {
	repo := csv.NewFlightsDataRepoCSV(csv_dir)
	finder := application.NewConnectionFinder(repo)

	return func(c *gin.Context) {
		c.Set("finder", finder)
	}
}

func airports(csv_dir string) func(*gin.Context) {
	repo := csv.NewFlightsDataRepoCSV(csv_dir)
	airports := application.NewAirportFinder(repo)

	return func(c *gin.Context) {
		c.Set("airports", airports)
	}

}
