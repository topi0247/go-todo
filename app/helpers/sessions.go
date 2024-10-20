package helpers

import (
	"net/http"
	"udemy-todo-app/app/models"
	"udemy-todo-app/infrastructure/db"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

const (
	SessionKey  = "session-key"
	SessionName = "session-name"
)

func AppendFlash(w http.ResponseWriter, r *http.Request, key string, value string) {
	session, _ := store.Get(r, SessionName)
	session.AddFlash(value, key)
	session.Save(r, w)
}

func getFlashMessages(w http.ResponseWriter, r *http.Request, session *sessions.Session, key string) []string {
	flashes := session.Flashes(key)
	session.Save(r, w)
	messages := []string{}
	for _, v := range flashes {
		messages = append(messages, v.(string))
	}
	return messages
}

func GetFlashes(w http.ResponseWriter, r *http.Request) struct {
	FlashSuccess []string
	FlashError   []string
	FlashNotice  []string
} {
	session, _ := store.Get(r, SessionName)

	flash := struct {
		FlashSuccess []string
		FlashError   []string
		FlashNotice  []string
	}{
		FlashSuccess: getFlashMessages(w, r, session, FlashSuccess),
		FlashError:   getFlashMessages(w, r, session, FlashError),
		FlashNotice:  getFlashMessages(w, r, session, FlashNotice),
	}

	return flash
}

func ClearFlashes(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SessionName)
	session.Flashes(FlashError)
	session.Flashes(FlashSuccess)
	session.Flashes(FlashNotice)
	session.Save(r, w)
}

func CreateSession(w http.ResponseWriter, r *http.Request, userUUID string) {
	session, _ := store.Get(r, SessionName)
	session.Values["userUUID"] = userUUID
	session.Save(r, w)
}

func GetSession(r *http.Request) string {
	session, _ := store.Get(r, SessionName)
	userUUID := session.Values["userUUID"]
	if userUUID == nil {
		return ""
	}
	return userUUID.(string)
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SessionName)
	delete(session.Values, "userUUID")
	session.Save(r, w)
}

func Authenticate(w http.ResponseWriter, r *http.Request) bool {
	userUUID := GetSession(r)
	if userUUID == "" {
		AppendFlash(w, r, FlashError, "ログインしてください")
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	} else {
		_, err := models.Users(
			models.UserWhere.UUID.EQ(userUUID),
		).One(r.Context(), db.DB)
		if err != nil {
			AppendFlash(w, r, FlashError, "予期せぬエラーが発生しました。再ログインしてください")
			http.Redirect(w, r, "/login", http.StatusFound)
			return false
		}
	}
	return true
}

func CurrentUser(r *http.Request) (user *models.User) {
	userUUID := GetSession(r)
	if userUUID == "" {
		user = &models.User{}
	} else {
		user, _ = models.Users(
			models.UserWhere.UUID.EQ(userUUID),
		).One(r.Context(), db.DB)
	}
	return user
}
