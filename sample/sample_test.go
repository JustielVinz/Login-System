package sample

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
)

func TestCreateSession(t *testing.T) {

	// Set up a request and response recorder for testing
	req, err := http.NewRequest("GET", "/createsession", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	// Call the CreateNewSession function with the recorder and request
	CreateNewSession(recorder, req)

	// Check the HTTP status code in the response
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected HTTP status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func CreateNewSession(recorder *httptest.ResponseRecorder, req *http.Request) {
	//geenrate the session key
	// Initialize the session store
	store := sessions.NewCookieStore([]byte("your-secret-key"))

	// Retrieve the session from the request
	session, err := store.Get(req, "my-session")

	if err != nil {
		http.Error(recorder, err.Error(), http.StatusInternalServerError)
		return
	}

	// Simulate creating a new session or setting some session data
	session.Values["data"] = "example data"
	session.Values["authenticated"] = true // Simulate an authenticated user

	// Save the session
	err = session.Save(req, recorder)
	if err != nil {
		http.Error(recorder, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestStoreDataInSession(t *testing.T) {
	// Initialize the session store
	store := sessions.NewCookieStore([]byte("your-secret-key"))

	// Create a request with a session
	req, err := http.NewRequest("GET", "/example", nil)
	if err != nil {
		t.Fatal(err)
	}
	session, _ := store.Get(req, "my-session")

	// Simulate storing data in the session
	dataToStore := "example data"
	session.Values["data"] = dataToStore

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Save the session
	err = session.Save(req, rr)
	if err != nil {
		t.Fatal(err)
	}

	// Check the session to see if the data was stored correctly
	session, err = store.Get(req, "my-session")
	if err != nil {
		t.Fatal(err)
	}

	storedData, ok := session.Values["data"].(string)
	if !ok {
		t.Errorf("Data type in session is not as expected")
	}

	if storedData != dataToStore {
		t.Errorf("Data in session is not as expected: got %v, want %v", storedData, dataToStore)
	}
}

func TestAuthentication(t *testing.T) {
	// Initialize the session store
	store := sessions.NewCookieStore([]byte("your-secret-key"))

	// Create a request with a session
	req, err := http.NewRequest("GET", "/example", nil)
	if err != nil {
		t.Fatal(err)
	}
	session, _ := store.Get(req, "my-session")

	// Simulate authenticating a user
	session.Values["authenticated"] = true
	session.Values["username"] = "authenticated_user"

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Save the session
	err = session.Save(req, rr)
	if err != nil {
		t.Fatal(err)
	}

	// Check the authentication status using the AuthenticateUser function
	isAuthenticated := AuthenticateUser(req)

	if !isAuthenticated {
		t.Errorf("Authentication failed for user: %s", "authenticated_user")
	}
}

func AuthenticateUser(req *http.Request) bool {
	// Retrieve the session from the request
	store := sessions.NewCookieStore([]byte("your-secret-key"))
	session, _ := store.Get(req, "my-session")

	// Check if the "authenticated" key exists in the session
	if val, ok := session.Values["authenticated"].(bool); ok && val {
		// The "authenticated" key exists and is true, indicating the user is authenticated.
		return true
	}

	// If the "authenticated" key doesn't exist or is false, the user is not authenticated.
	return false
}
