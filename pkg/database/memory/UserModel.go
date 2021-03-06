package memory

import (
	"fmt"
	"strings"

	"github.com/quan-to/chevron/pkg/models"
	"github.com/quan-to/chevron/pkg/uuid"
)

func (h *DbDriver) AddUser(um models.User) (string, error) {
	h.lock.Lock()
	defer h.lock.Unlock()

	um.ID = uuid.EnsureUUID(h.log)

	h.users = append(h.users, um)

	return um.ID, nil
}

func (h *DbDriver) GetUser(username string) (um *models.User, err error) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	for _, v := range h.users {
		if strings.EqualFold(username, v.Username) {
			n := v // Copy
			return &n, nil
		}
	}

	return um, fmt.Errorf("not found")
}

func (h *DbDriver) UpdateUser(um models.User) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	if um.ID == "" {
		return fmt.Errorf("not found")
	}

	for i, v := range h.users {
		if v.ID == um.ID {
			h.users[i] = um
			return nil
		}
	}

	return fmt.Errorf("not found")
}
