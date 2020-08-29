package interceptor

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Session(SessionName string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(SessionName, store)
}
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600 //一小时
	sessionSecret := "Damon"
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
		Path:   "/",
	})
	return store
}