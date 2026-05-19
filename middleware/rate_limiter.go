package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type client struct {
	requests int
	lastSeen time.Time
}

var clients = make(map[string]*client)

var mutex sync.Mutex

func RateLimiter() gin.HandlerFunc {

	return func(c *gin.Context) {

		ip := c.ClientIP()

		mutex.Lock()

		if _, found := clients[ip]; !found {

			clients[ip] = &client{
				requests: 1,
				lastSeen: time.Now(),
			}

		} else {

			if time.Since(clients[ip].lastSeen) > time.Minute {

				clients[ip].requests = 1
				clients[ip].lastSeen = time.Now()

			} else {

				clients[ip].requests++
			}
		}

		if clients[ip].requests > 20 {

			mutex.Unlock()

			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})

			c.Abort()

			return
		}

		mutex.Unlock()

		c.Next()
	}
}
