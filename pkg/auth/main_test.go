package auth

import (
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

var testRedisMock redismock.ClientMock
var testRedisClient *redis.Client

func TestMain(m *testing.M) {
	client, mock := redismock.NewClientMock()
	testRedisClient = client
	testRedisMock = mock

	status := m.Run()
	if err := testRedisMock.ExpectationsWereMet(); err != nil {
		panic(err)
	}

	os.Exit(status)
}
