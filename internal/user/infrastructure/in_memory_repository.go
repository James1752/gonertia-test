package user_infrastructure

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	user_domain "github.com/James1752/gonertia-test/internal/user/domain"
	"github.com/google/uuid"
)

// UserInMemoryRepository will store users in memory
type UserInMemoryRepository struct {
	users map[uuid.UUID]*user_domain.User
	mu    sync.RWMutex // To handle concurrency safely
}

// NewUserInMemoryRepository creates and returns a new UserInMemoryRepository
func NewUserInMemoryRepository() *UserInMemoryRepository {
	return &UserInMemoryRepository{
		users: make(map[uuid.UUID]*user_domain.User),
	}
}

func logUsersAsJSON(users map[uuid.UUID]*user_domain.User) {
	// Marshal the map into JSON for structured logging
	usersJson, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Printf("Error marshaling users map: %v", err)
		return
	}

	// Log the JSON string
	log.Printf("Users: %s", usersJson)
}

// GetUserById retrieves a user by ID from the in-memory map
func (r *UserInMemoryRepository) GetUserById(ID uuid.UUID) (*user_domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	log.Printf("Getting User with id: %s\n", ID)

	user, exists := r.users[ID]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// CreateUser creates a new user and stores it in the in-memory map
func (r *UserInMemoryRepository) CreateUser(user *user_domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user already exists by ID (can also check by email, depending on your logic)
	if _, exists := r.users[user.UserID]; exists {
		return errors.New("user already exists")
	}

	log.Printf("Creating User with id: %s\n", user.UserID)

	// Store the user in memory
	r.users[user.UserID] = user

	logUsersAsJSON(r.users)

	return nil
}

// UpdateUser updates an existing user in the in-memory map
func (r *UserInMemoryRepository) UpdateUser(user *user_domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user exists
	if _, exists := r.users[user.UserID]; !exists {
		return errors.New("user not found")
	}

	log.Printf("Updating User with id: %s\n", user.UserID)

	// Update the user data
	r.users[user.UserID] = user

	return nil
}

// DeleteUser removes a user from the in-memory map
func (r *UserInMemoryRepository) DeleteUser(ID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user exists
	if _, exists := r.users[ID]; !exists {
		return errors.New("user not found")
	}

	log.Printf("Deleting User with id: %s\n", ID)

	// Remove the user from memory
	delete(r.users, ID)

	return nil
}
